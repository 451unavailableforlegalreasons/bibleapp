package bible 

import (
	"bibleapp/services"
    "encoding/json"
	"fmt"
	"net/http"
)

type verseRequest struct {
    BibleEdition int `json:"bibleedition"`
    BookName int `json:"bookname"`
    BookChapter int `json:"bookchapter"`
    VerseNumFrom int `json:"versenumfrom"`
    VerseNumTo int `json:"versenumto"`
}

type verseResponseList struct {
    BibleEdition int `json:"bibleedition"`
    BibleBook int `json:"biblebook"`
    BookChapter int `json:"biblechapter"`
    Verses []dbVerseRetreive `json:"verses"`
}

type dbVerseRetreive struct {
    VerseNumber int `json:"vnum" db:"verseNumber"`
    Verse string `json:"verse" db:"verseContent"`
}



func GetVerses(w http.ResponseWriter, r *http.Request) {
    // public endpoint to retreive array of verses (limited to 20 per request)


    var requestedVerses verseRequest


    err := json.NewDecoder(r.Body).Decode(&requestedVerses)
    if err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }
    // if there is a major error requestedVerses will be a nil pointer
    if &requestedVerses == nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    // if feilds not filled or inexistant they are set to 0
    // check every field to see if they were filled proprely
    if requestedVerses.BookName == 0 || requestedVerses.BookChapter == 0 || requestedVerses.VerseNumFrom == 0 || requestedVerses.VerseNumTo == 0 || requestedVerses.BibleEdition == 0 {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }
    
    // can't request verse 15 to 2 but only 2 to 15
    // can only ask for 15 verses at a time
    var diff = requestedVerses.VerseNumTo - requestedVerses.VerseNumFrom;
    if  diff < 0 || diff > 15 || diff < -15 {
        http.Error(w, "verse number messed up", http.StatusBadRequest)
        return
    }
    var verses = make([]dbVerseRetreive, 0, diff)
    err = services.DB.Select(
        &verses, 
        "SELECT verseNumber, verseContent FROM Verse WHERE bibleEdition=? and bibleBook=? and bookChapter=? and verseNumber BETWEEN ? AND ? LIMIT 15;", 
        requestedVerses.BibleEdition,
        requestedVerses.BookName,
        requestedVerses.BookChapter,
        requestedVerses.VerseNumFrom,
        requestedVerses.VerseNumTo,
    )
    if err != nil {
        fmt.Printf("Error retreiving verses from req: %+v | err: %+v", requestedVerses, err)
        http.Error(w, "error reading from database to retreive your verses", http.StatusInternalServerError)
        return
    }
    
    var responsecontent = verseResponseList{
        BibleEdition: requestedVerses.BibleEdition,
        BibleBook: requestedVerses.BookName,
        BookChapter: requestedVerses.BookChapter,
        Verses: verses,
    }

    // fmt.Printf("%v", requestedVerses)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(responsecontent)
}


