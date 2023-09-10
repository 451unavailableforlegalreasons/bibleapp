package bible

// Meta data about bibles in database


import (
	"bibleapp/services"
	"net/http"
    "encoding/json"
    "fmt"
)


type editionStruct struct {
    Name string `json:"name" db:"name"`
    Language string `json:"lang" db:"language"`
    Id int `json:"id" db:"id"`
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
        err := rows.Scan(&edition.Name, &edition.Language, &edition.Id)
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
