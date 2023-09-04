package umanagement 

// this file is to add endpoints and fmtic for verifying the user email
// note: changing the user email in the changesensitiveinfo.go will change a database column to false and resend an email to the new email.
// But the link will redirect here (we only do the verification)

import (
	"bibleapp/services"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
    "net/mail"

	"github.com/gorilla/mux"
)



func AskEmailConf(w http.ResponseWriter, r *http.Request) {
    // button on account to send an email

    //based on uid check if user already verified his email
    //if not, retreive email and send him the verif
    var res struct {
        Email string `db:"email"`
        Verified bool `db:"emailverified"`
    }
    query := `SELECT email, emailverified FROM Users WHERE id=?;`
    err := services.DB.Get(&res, query, 1) // replace 1 with userid from r.context (created by authmiddleware)
    if err != nil {
        fmt.Println("error asking for email conf: ", err)
        http.Error(w, "error with your request for email confirmation", http.StatusInternalServerError)
        return
    }
    if res.Verified == true {
        w.WriteHeader(http.StatusNotModified); // bad usage but i'll do it anyway
        w.Write([]byte("your email is already verified"));
        return
    } else {
        err := SendVerificationEmail(res.Email) 
        if err != nil {
            fmt.Println(err)
            http.Error(w, "error sending verification email", http.StatusInternalServerError)
            return
        } else {
            w.WriteHeader(http.StatusNotModified); // bad usage but i'll do it anyway
            w.Write([]byte("email sent"));
            return
        }
    }

    
}


func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
    // endpoint that get the token from the link the user clicked on and verify that it is valid (then changes the database user email status)
    vars := mux.Vars(r);
    token := vars["token"]
    email := vars["email"]
    datelimit, err := strconv.ParseInt(vars["datelimit"], 10, 64)
    if err != nil {
        fmt.Println("invalid parameter", err);
        http.Error(w, "invalid parameter", http.StatusBadRequest)
        return
    }

    valid, err := verifylinktoken(token, email, datelimit)
    if err != nil {
        fmt.Println("error with verify email token: ", err);
        http.Error(w, "error with token", http.StatusBadRequest)
        return
    } 
    if valid == false {
        fmt.Println("invalid verify email token: ", err);
        http.Error(w, "invalid token", http.StatusBadRequest)
        return
    } else {
        // change database state
        dbquery := `UPDATE Users SET emailverified=TRUE WHERE email=?;`
        _, err := services.DB.Exec(dbquery, email)
        if err != nil {
            fmt.Println("error changing email status: ", err)
            http.Error(w, "error changing your email state, try again later", http.StatusInternalServerError)
            return
        } else {
            w.WriteHeader(http.StatusOK);
            w.Write([]byte("email verified"));
        }
    }
}




// =======================================
//           NON endpoint func
// =======================================




func SendVerificationEmail (email string) (error) {
    // this isn't a handler, just a function that send an email with the link to verify the email.
    _, err := mail.ParseAddress(email)
    if err != nil {
        return err
    }

    // compute token
    datelimit := time.Now().AddDate(0,1,0).Unix() // add 1 month to current time and convert to unix fmt
    token := generatetoken(email, datelimit)
    fmt.Println("generated token: ", token) // later send a real email

    return nil
}

func generatetoken(email string, limitdate int64) string {
    privatekey := "3qc1DaODDEFFWFdMSzZgrSIMNdaF43avz5isxXE9z7mWIH0lfHzxza7QllhwZo3F"
    h := sha256.New()
    s := strconv.FormatInt(limitdate, 10)
    h.Write([]byte(privatekey + email + s))
    outputtoken := hex.EncodeToString(h.Sum(nil))
    return outputtoken
}

func verifylinktoken(token string, email string, limitdate int64) (bool, error) {
    // based on the string passed as a parameter, verify that the token is valid (meaning the user clicked on the link in his inbox)
    // temporary thing for testing (will not be the production private key)
    currentdate := time.Now().Unix()
    if currentdate > limitdate {
        return false, errors.New("token is expired")
    }
    outputtoken := generatetoken(email, limitdate) 

    if outputtoken != token {
        // token don't match = modified token or value
        return false, errors.New("invalid token")
    }


    return true, nil
}
