package umanagement

// reset password route and logic. When the user isn't logged in.
// changing password when logged in is in the changesensitiveinfo.go

import (
	"bibleapp/services"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"
	"time"
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)



func AskRstPassword(w http.ResponseWriter, r *http.Request) {
    // user gets a form with an email field (if email doesn't exists, don't tell him)
    // if request is a POST, try to send the mail with the reset link
    var req struct  {
        Emailrequested string `json:"email"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    fmt.Println(req.Emailrequested)
    query := `SELECT id FROM Users WHERE email=?;`
    var res struct {
        Id int `db:"id"`
    }
    err := services.DB.Get(&res, query, req.Emailrequested) // replace 1 with userid from r.context (created by authmiddleware)
    if err != nil {
        fmt.Println("error asking for email rst passwrd: ", err)
    }
    if res.Id == -1 {
        // user not found respond as if normal
        fmt.Println("User not found")
    } else {
        // send email
        err := SendRstPasswordEmail(req.Emailrequested)
        if err != nil {
            fmt.Println("Error sending email: ", err)
        }
    }


    w.WriteHeader(http.StatusOK);
    w.Write([]byte("email sent"));
}


func RstPasswordForm(w http.ResponseWriter, r *http.Request) {
    // returns a simple form with the url param inside.
    // do not check the parameters now but later when the form is submitted
    vars := mux.Vars(r);
    _= vars["token"]
    _= vars["email"]
    _= vars["seed"]
    _, err := strconv.ParseInt(vars["datelim"], 10, 64)
    if err != nil {
        fmt.Println("invalid parameter", err);
        http.Error(w, "invalid parameter", http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK);
    w.Write([]byte("return an html page with form for pass reset"));

}


func RstPassword(w http.ResponseWriter, r *http.Request) {
    // endpoint that get the token from the link the user clicked on and verify that it is valid (then changes the database user email status)
    type resetpassform struct {
        Password string `json:"password"`
        Token string `json:"token"`
        Email string `json:"email"`
        Seed string `json:"seed"`
        DateLim int64 `json:"dl"`
    }
    var req resetpassform
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        fmt.Println("invalid parameter", err);
        http.Error(w, "invalid parameter", http.StatusBadRequest)
        return
    }

    valid, err := verifypasswordtoken(req.Token, req.Seed, req.Email, req.DateLim)
    if err != nil {
        fmt.Println("error with verify password token: ", err);
        http.Error(w, "error with token", http.StatusBadRequest)
        return
    } 
    if valid == false {
        fmt.Println("invalid password token: ", err);
        http.Error(w, "invalid token", http.StatusBadRequest)
        return
    } else {
        // change database password 
        pass, err := HashPassword(req.Password)
        if err != nil {
            fmt.Println("error hashing register")
            w.WriteHeader(http.StatusBadRequest);
            w.Write([]byte("error with your password"));
            return
        }
        dbquery := `UPDATE Users SET password=? WHERE email=?;`
        _, err = services.DB.Exec(dbquery, pass, req.Email)
        if err != nil {
            fmt.Println("error changing password status: ", err)
            http.Error(w, "error changing your password, try again later", http.StatusInternalServerError)
            return
        } else {
            w.WriteHeader(http.StatusOK);
            w.Write([]byte("password changed"));
        }
    }
}


// =======================================
//           NON endpoint func
// =======================================


func HashPassword(password string) ([]byte, error) {
    // note: bcrypt accept password that are less than 72 bytes long
    pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return pass, err

}



func SendRstPasswordEmail (email string) (error) {
    // this isn't a handler, just a function that send an email with the link to verify the email.
    _, err := mail.ParseAddress(email)
    if err != nil {
        // return err
    }
    
    // generate seed
    randomBytes := make([]byte, 20)
    if _, err := rand.Read(randomBytes); err != nil {
        fmt.Println("Error generating seed for password reset")  
        return errors.New("error generating seed")
    }
    seed := base64.RawURLEncoding.EncodeToString(randomBytes)

    // compute token
    datelimit := time.Now().Add(time.Hour * 1).Unix() // link is only valid for 1 hour
    token := generatepasswordtoken(seed, email, datelimit)
    fmt.Printf("generated link: localhost:8080/users/account/reset-password-form/%v/%v/%v/%v\n", token, seed, email, datelimit) // later send a real email

    return nil
}

func generatepasswordtoken(seed string, email string, limitdate int64) string {
    privatekey := "3qc1DaODDEFFWFdMSzZgrSIMNdaF43avz5isxXE9z7mWIH0lfHzxza7QllhwZo3F"
    h := sha256.New()
    s := strconv.FormatInt(limitdate, 10)
    h.Write([]byte(privatekey +seed + email + s))
    outputtoken := hex.EncodeToString(h.Sum(nil))
    return outputtoken
}

func verifypasswordtoken(token string, seed string, email string, limitdate int64) (bool, error) {
    currentdate := time.Now().Unix()
    if currentdate > limitdate {
        return false, errors.New("token is expired")
    }
    outputtoken := generatepasswordtoken(seed, email, limitdate) 

    if outputtoken != token {
        // token don't match = modified token or value
        return false, errors.New("invalid token")
    }


    return true, nil
}
