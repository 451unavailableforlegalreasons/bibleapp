package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func GetProfile(w http.ResponseWriter, r *http.Request) {
    context := r.Context()
    w.WriteHeader(http.StatusOK);
    vars := mux.Vars(r);


// please figure this shit out | id don't know how to retreive the data in the context
    newcontent := context.Value("Value")
    if newcontent == nil {
        return
    }
    fmt.Println(vars["uid"])
    fmt.Println(vars)
    fmt.Println(newcontent)
    if newcontent == vars["uid"] {
        w.Write([]byte("this was your uid"))
    } else {
        w.Write([]byte("try again"))
    }

}



func CreateProfile(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK);
    w.Write([]byte("profile route Post"));
}
