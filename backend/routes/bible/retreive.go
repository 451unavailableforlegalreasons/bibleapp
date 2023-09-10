package bible

import (
	"bibleapp/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type BibleZip struct {
    EditionId int `json:"editionid" db:"edition"`
    ZipContent []byte`json:"zipcontent" db:"zipcontent"`
}



func GetBible(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    edition := vars["edition"]
    if edition == "" {
        http.Error(w, "invalid editon", http.StatusBadRequest)
        return
    }



    var bible BibleZip
    err := services.DB.Get(
        &bible, 
        "SELECT * from BibleZip WHERE edition=?;", edition)
    if err != nil {
        fmt.Printf("Error retreiving verses from req: %+v | err: %+v", edition, err)
        http.Error(w, "error reading from database to retreive edition", http.StatusInternalServerError)
        return
    }


    w.Header().Set("Content-Type", "application/octet-stream")
    w.WriteHeader(http.StatusOK);
    w.Write(bible.ZipContent)
}

