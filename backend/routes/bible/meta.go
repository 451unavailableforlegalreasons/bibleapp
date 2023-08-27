package bible

// theses routes gives meta data for bible editions, books ... in the database


import (
	"bibleapp/services"
	"net/http"
    "encoding/json"
    "fmt"
    "strconv"

	"github.com/gorilla/mux"
)


type editionStruct struct {
    Name string `json:"name" db:"name"`
    Id int `json:"id" db:"id"`
}

type bookByEdition struct {
    Name string `json:"name" db:"name"`
    Ordnum int `json:"ordnum" db:"ordnum"`
}







func GetBibleEditions(w http.ResponseWriter, r *http.Request) {
    var editions []editionStruct


    rows, err := services.DB.Query("SELECT * FROM BibleEdition;")
    if err != nil {
        fmt.Println("error reading bible editions from database: ", err);
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    }
    defer rows.Close()


    var edition editionStruct
    for rows.Next() {
        err := rows.Scan(&edition.Name, &edition.Id)
        if err != nil {
            fmt.Println("error reading bible editions from database: ", err);
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode("{'err': 'internal database error'}")
            return
        }

        editions = append(editions, edition)
    }
    if err := rows.Err(); err != nil {
        fmt.Println("error reading bible editions from database: ", err);
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    }



    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(editions)
}










func GetBibleBooksFromEdition(w http.ResponseWriter, r *http.Request) {
    // get list of all bible book for a certain edition
    vars := mux.Vars(r)
    var query = make(map[string]int64)
    var err error

    for key, val := range vars{  
        query[key], err = strconv.ParseInt(val, 10, 64)
        if err != nil {
            http.Error(w, "invalid parameter", http.StatusBadRequest)
            return
        }
    }


    var books []bookByEdition
    rows, err := services.DB.Query("SELECT name, ordnum FROM BibleBook WHERE edition=?;", query["edition"])
    if err != nil {
        fmt.Println("error reading bible editions from database: ", err);
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    }
    defer rows.Close()


    var book bookByEdition 
    for rows.Next() {
        err := rows.Scan(&book.Name, &book.Ordnum)
        if err != nil {
            fmt.Println("error reading bible editions from database: ", err);
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode("{'err': 'internal database error'}")
            return
        }

        books = append(books, book)
    }
    if err := rows.Err(); err != nil {
        fmt.Println("error reading bible editions from database: ", err);
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(books)
}



func GetNumberOfVersesFromBookFromEdition(w http.ResponseWriter, r *http.Request) {
    // get number of verse in a book from a certain edition
    vars := mux.Vars(r)
    var query = make(map[string]int64)
    var err error

    for key, val := range vars{  
        query[key], err = strconv.ParseInt(val, 10, 64)
        if err != nil {
            http.Error(w, "invalid parameter", http.StatusBadRequest)
            return
        }
    }


    var count int
    err = services.DB.Get(
        &count, 
        "SELECT COUNT(*) FROM Verse WHERE BibleEdition=? AND BibleBook=?;", 
        query["edition"],
        query["book"],
    )
    if err != nil {
        fmt.Println("error reading bible books from database");
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("{'err': 'internal database error'}")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(fmt.Sprintf("{'count': %d}", count))
}

