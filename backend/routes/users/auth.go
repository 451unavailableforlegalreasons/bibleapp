package users

import (
	"bibleapp/services"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	// "time"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)


func Logout(w http.ResponseWriter, r *http.Request) {
    context := r.Context() // dict contaiing data
    fmt.Println(context.Value("tokenContent"))
    cookie := http.Cookie{
        Name: "test",
        Value: "chocolate-cookie",

        Path: "/",
        // Domain: ".",
        // Expires: time.Now().Add(time.Hour * 2).Add(time.Second * 60), // it is a session cookie so no expire
        // MaxAge:, max age is expire but in second and has the priority

        Secure: true,
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("Logout route"));
}



type LoginForm struct {
    Email string `json:"email"`
    Password string `json:"password"` 
}


type RegisterForm struct {
    Email string `json:"email"`
    Password string `json:"password"` 
    FullName string `json:"fullname"`
    Age int `json:"age"`
    Gender string `json:"gender"`
    Language string `json:"favlang"` 
    Country string `json:"country"` 
}


func parseRegisterForm(r *http.Request) (RegisterForm, error){
    var regform RegisterForm 
    err := json.NewDecoder(r.Body).Decode(&regform)
    if err != nil {
        fmt.Println("Error reading form from body")
        return regform, errors.New("invalid form")
    }

    // sorry for that horrible code 
    if regform.Email == "" ||
    regform.Password == "" || 
    regform.FullName == "" || 
    regform.Age == 0 ||
    regform.Gender == "" ||
    regform.Language == "" ||
    regform.Country == "" {
        return regform, errors.New("invalid register form")
    } else {
        return regform, nil
    }
}


func parseLoginForm(r *http.Request) (LoginForm, error){
    var logform LoginForm
    err := json.NewDecoder(r.Body).Decode(&logform)
    if err != nil {
        fmt.Println("Error reading form from body")
        return logform, errors.New("invalid form")
    }

    
    if logform.Email == "" || logform.Password == "" {
        return logform, errors.New("invalid login form")
    } else {
        return logform, nil
    }
}










func GetLogin(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("Get form"));
    return
}

func Login(w http.ResponseWriter, r *http.Request) {

    logform, err := parseLoginForm(r)
    if err != nil {
        fmt.Println("error parsing login form")
        w.WriteHeader(http.StatusBadRequest);
        w.Write([]byte("error with your form"));
        return
    }



    //check email and passwod match in database
    // if they do , get id from db and create sid cookie
    var user struct {
        Id int `db:"id"`
        Age int `db:"age"`
        Gender string `db:"gender"`
        Country string `db:"country"`
        Language string `db:"language"`
        Password string `db:"password"`
    }
    err = services.DB.Get(
        &user, 
        "SELECT rowid as id, password, age, gender, country, language FROM Users WHERE email=?", 
        logform.Email,
    )
    if err == sql.ErrNoRows {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode("{'err': 'try again'}")
        return
    } else if err != nil {
        fmt.Println("error login database");
        fmt.Println(err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    } 





    // bcryp return nil if password match
    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logform.Password)) == nil {
        fmt.Println("login successful");
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode("{'err': 'none'}")
        return
    } else {
        // create cookie in Session database and send it to the client
        fmt.Println("login failed");
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode("{'err': 'invalid credentials'}")
        return
    }
}






func GetRegister(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("Get form register"));
    return
}


func Register(w http.ResponseWriter, r *http.Request) {
    regform , err := parseRegisterForm(r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest);
        w.Write([]byte(err.Error()));
        return
    }
    pass, err := HashPassword(regform.Password)
    if err != nil {
        fmt.Println("error hashing register")
        w.WriteHeader(http.StatusBadRequest);
        w.Write([]byte("error with your password"));
        return
    }







    var user struct {
        Id int `db:"id"`
        Email string `db:"email"`
        Password string `db:"password"`
    } 
    err = services.DB.Get(
        &user, 
        "SELECT id, email, password FROM Users WHERE email=?;", 
        regform.Email,
    )
    fmt.Printf("%+v\n",user)
    if err == sql.ErrNoRows {
        // user doesn't exists so everything is ok
        fmt.Println("inserting new user")
        res, err := services.DB.Exec(
            "INSERT INTO Users (fullname, email, password, age, gender, country, language) VALUES (?, ?, ?, ?, ?, ?, ?);",
            regform.FullName,
            regform.Email,
            pass,
            regform.Age,
            regform.Gender,
            regform.Country,
            regform.Language,
        )
        if err != nil {
            fmt.Println(res, err)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode("{'err': 'internal database error'}")
            return
        }
    } else if err != nil {
        fmt.Println("error login database");
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    } else {
        // user exists so don't put it inside db
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode("{'err': 'error with your form'}")
        return
    }


    //check users email isn't already in the database
    // if not insert new user and add a sid cookie

    fmt.Printf("%+v\n", regform)
    fmt.Printf("%+v\n", user)

    w.WriteHeader(http.StatusOK);
    w.Write([]byte("success"));
    return
}






// ============================== //
//          Utils functions       //
// ============================== //


func HashPassword(password string) ([]byte, error) {
    // note: bcrypt accept password that are less than 72 bytes long
    pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return pass, err

}


