package users

import (
	// "bibleapp/services"
	"fmt"
	"net/http"
	"time"
)


func Logout(w http.ResponseWriter, r *http.Request) {
    context := r.Context() // dict contaiing data
    fmt.Println(context.Value("tokenContent"))
    cookie := http.Cookie{
        Name: "test",
        Value: "chocolate-cookie",

        Path: "/",
        // Domain: ".",
        Expires: time.Now().Add(time.Hour * 2).Add(time.Second * 60),

        Secure: true,
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("Logout route"));
}


func Login(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("Login route"));
}
