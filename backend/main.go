package main

import (
	"bibleapp/routes/users"
	"bibleapp/routes/bible"
    "bibleapp/services"
	"fmt"
	"log"
	"net/http"
	"time"
    "context"

	"github.com/gorilla/mux"
)


func main() {
    r := mux.NewRouter()

    fmt.Println("running")
    services.Init_DB() // init the database connection here and check in handler if it is not nil & middleware
   
    // user routes logic
    userRouter := r.PathPrefix("/users").Subrouter();

    authRouter := userRouter.PathPrefix("/auth").Subrouter();
    authRouter.HandleFunc("/login", users.GetLogin).Methods("GET");
    authRouter.HandleFunc("/login", users.Login).Methods("POST");
    authRouter.HandleFunc("/register", users.GetRegister).Methods("GET");
    authRouter.HandleFunc("/register", users.Register).Methods("POST");
    authRouter.HandleFunc("/logout", users.Logout).Methods("POST");
    
    accountRouter := userRouter.PathPrefix("").Subrouter();
    accountRouter.Use(authenticationMiddleware); // for every requests made to /users/profile/*, the middel ware will do its job
    accountRouter.HandleFunc("/profile", users.GetProfile).Methods("GET");  // get profile info route
    accountRouter.HandleFunc("/profile", users.PostProfile).Methods("POST"); // update profile route
    accountRouter.HandleFunc("/settings", users.GetSettings).Methods("GET"); // get current settings route
    accountRouter.HandleFunc("/settings", users.PostSettings).Methods("POST");// change settings route
    // profileRouter.HandleFunc("/{uid:[0-9]+}", users.GetProfile).Methods("GET");
    // profileRouter.HandleFunc("/{uid:[0-9]+}", users.CreateProfile).Methods("POST");


    bibleRouter := r.PathPrefix("/bible").Subrouter()
    bibleRouter.HandleFunc("/getverses/{edition:[0-9]+}/{book:[0-9]+}/{chapter:[0-9]+}/{versefrom:[0-9]+}/{verseto:[0-9]+}", bible.GetVerses).Methods("GET");
    bibleRouter.HandleFunc("/geteditions", bible.GetBibleEditions)
    bibleRouter.HandleFunc("/getbooksofedition/{edition:[0-9]+}", bible.GetBibleBooksFromEdition)
    bibleRouter.HandleFunc("/getnumberofversefrombookfromedition/{edition:[0-9]+}/{book:[0-9]+}", bible.GetNumberOfVersesFromBookFromEdition)

    srv := &http.Server{
        Handler:      r,
        Addr:         "localhost:8080",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Fatal(srv.ListenAndServe());
}









/*
===================================

        MIDDLEWARE SECTION

===================================
*/

func requireDatabaseMiddleware(next http.Handler) http.Handler {
    // if the database connection doesn't exists, create it
    if services.DB == nil {
        fmt.Println("database is nil, creating it")
        services.Init_DB() // this will try 5 times to connect, spaced in time by 5 seconds
    }
    // check if connection was successful (if not return error else call next)
    if services.DB == nil {
        fmt.Println("Couldn't connect to database exiting")
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Server side error with the database connection sry"))
            // return
        })
    } else {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            next.ServeHTTP(w, r)
        })
    }
}

func authenticationMiddleware(next http.Handler) http.Handler {
    // if the database connection doesn't exists, create it
    if services.DB == nil {
        fmt.Println("database is nil, creating it")
        services.Init_DB() // this will try 5 times to connect, spaced in time by 5 seconds
    }
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if services.DB == nil {
            // we still coudln't connect to the database after 5 tries
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Server side error with the database connection sry"))
            return
        }




        token, err := r.Cookie("sid");
        if err != nil {
            fmt.Println("Cookie not found");
            w.WriteHeader(http.StatusForbidden);
            w.Write([]byte("Cookie not found"));
            return
        }


        fmt.Println(token)
        var storage  struct {
            Sid string `db:"sid"`
            Value string `db:"value"` // base64 encoded json data
        }
        err = services.DB.Get(&storage, "SELECT * FROM Sessions WHERE sid=$1", token.Value) // ignore the lsp error
        if err != nil {
            w.WriteHeader(http.StatusForbidden);
            w.Write([]byte("invalid cookie"));
            return
        }
        fmt.Println(storage)
        // pass storage to next handler
        // it is the next handler responsability to authorize the user
        r = r.Clone(context.WithValue(r.Context(), "Sid", storage.Sid))
        r = r.Clone(context.WithValue(r.Context(), "Value", storage.Value))
        fmt.Println(r.Context().Value("Sid"))
        fmt.Println(r.Context().Value("Value"))
        next.ServeHTTP(w, r)
    })
}




// in the future a middleware that parse the context and lays it out in
// a way that is standardized (so that every handler knows where/how to get data)
