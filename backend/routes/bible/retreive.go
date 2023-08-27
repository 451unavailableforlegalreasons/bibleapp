package bible

import (
	"bibleapp/services"
	"encoding/json"
	"fmt"
	"net/http"
    "strconv"

	"github.com/gorilla/mux"
)


type verseResponseList struct {
    BibleEdition int64 `json:"bibleedition"`
    BibleBook int64 `json:"biblebook"`
    BookChapter int64 `json:"biblechapter"`
    Verses []dbVerseRetreive `json:"verses"`
}

type dbVerseRetreive struct {
    VerseNumber int `json:"vnum" db:"verseNumber"`
    Verse string `json:"verse" db:"verseContent"`
}



func GetVerses(w http.ResponseWriter, r *http.Request) {
    // public endpoint to retreive array of verses (limited to 15 per request)
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


    // can't request verse 15 to 2 but only 2 to 15
    // can only ask for 15 verses at a time
    diff := query["verseto"] - query["versefrom"]
    if  diff < 0 || diff > 15 || diff < -15 {
        http.Error(w, "verse number messed up", http.StatusBadRequest)
        return
    }
    var verses = make([]dbVerseRetreive, 0, diff)
    err = services.DB.Select(
        &verses, 
        "SELECT verseNumber, verseContent FROM Verse WHERE bibleEdition=? and bibleBook=? and bookChapter=? and verseNumber BETWEEN ? AND ? LIMIT 15;", 
        query["edition"],
        query["book"],
        query["chapter"],
        query["versefrom"],
        query["verseto"],
    )
    if err != nil {
        fmt.Printf("Error retreiving verses from req: %+v | err: %+v", query, err)
        http.Error(w, "error reading from database to retreive your verses", http.StatusInternalServerError)
        return
    }

    var responsecontent = verseResponseList{
        BibleEdition: query["edition"],
        BibleBook: query["book"],
        BookChapter: query["chapter"],
        Verses: verses,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(responsecontent)
}


