package bible

import (
	"bibleapp/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)



type DBHighlight struct {
    // in database its called a note but its also a highlight so... sry for the confusion
    HighligthId int `json:"hid" db:"noteid"`
    Author int `json:"author" db:"author"`

    BibleEdition int `json:"edition" db:"bibleEdition"`
    BibleBook int `json:"book" db:"bibleBook"`
    BookChapterStart int `json:"chapterstart" db:"bookChapterStart"`
    BookChapterEnd int `json:"chapterend" db:"bookChapterEnd"`
    
    VerseNumberStart int `json:"vnumstart" db:"verseNumberStart"`
    VerseNumberEnd int `json:"vnumend" db:"verseNumberEnd"`

    CharNumStart int `json:"charnumstart" db:"CharNumStart"`
    CharNumEnd int `json:"charnumend" db:"CharNumEnd"`


    AuthorNote string `json:"note" db:"authorNote"`
    HighlighColor int `json:"color" db:"highlightColor"`
}
// i know... i should use composition but i'm too lazy to do it
type Highlight struct {
    BibleEdition int `json:"edition" db:"bibleEdition"`
    BibleBook int `json:"book" db:"bibleBook"`
    BookChapterStart int `json:"chapterstart" db:"bookChapterStart"`
    BookChapterEnd int `json:"chapterend" db:"bookChapterEnd"`
    
    VerseNumberStart int `json:"vnumstart" db:"verseNumberStart"`
    VerseNumberEnd int `json:"vnumend" db:"verseNumberEnd"`

    CharNumStart int `json:"charnumstart" db:"CharNumStart"`
    CharNumEnd int `json:"charnumend" db:"CharNumEnd"`


    AuthorNote string `json:"note" db:"authorNote"`
    HighlighColor int `json:"color" db:"highlightColor"`
}


func CreateHighlight(w http.ResponseWriter, r *http.Request) {
    // insert a new highligh in database for user (based on sid cookie to get userid)
    var h Highlight
    err := json.NewDecoder(r.Body).Decode(&h)
    if err != nil {
        http.Error(w, "error parsing request body", http.StatusBadRequest)
        return
    }


    // please check that there isn't already a near identical note on the same verse...
    // to avoid dos by retreiving 1000000 highlights in a single verse
    _, err = services.DB.Exec(`INSERT INTO Note(
    author,
    bibleEdition,
    bibleBook,
    bookChapterStart,
    bookChapterEnd,
    verseNumberStart,
    verseNumberEnd,
    CharNumStart,
    CharNumEnd,
    authorNote,
    highlightColor) Values(?,?,?,?,?,?,?,?,?,?,?);`, 
    1, // 1 is for the author id change it later to sid cookie value
    h.BibleEdition, 
    h.BibleBook,
    h.BookChapterStart,
    h.BookChapterEnd,
    h.VerseNumberStart,
    h.VerseNumberEnd,
    h.CharNumStart,
    h.CharNumEnd,
    h.AuthorNote,
    h.HighlighColor,
)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "error inserting into databsae", http.StatusBadRequest)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(``)
}


func DeleteHighlight(w http.ResponseWriter, r *http.Request) {
    // delete highligth by passing its id (check if the user owns it)
    query := `DELETE FROM Note WHERE author=? AND noteid=?` // pass authorid (sid cookie value) and noteid from body
    var vars = mux.Vars(r)
    noteid, err := strconv.ParseInt(vars["noteid"], 10, 64)
    if err != nil {
        http.Error(w, "error parsing noteid parameter", http.StatusBadRequest)
        return
    }
    // replace the 1 with the sid cookie next
    _, err = services.DB.Exec(query, 1, noteid)
    if err != nil {
        http.Error(w, "error deleting highlight from database", http.StatusBadRequest)
        return
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK);
        json.NewEncoder(w).Encode(`{"err": ""}`)
    }

}



func ReadHighlight(w http.ResponseWriter, r *http.Request) {
    // retreive all highligth of user (sid cookie), based on filter provided in url (see bellow)
    type queryparms struct {
        urlparamname string
        parsingerror error
        rawurldata string
        parseddata int64 
    }
    params := make([]queryparms, 6)


    urlparams := r.URL.Query()
    var keys = []string{"edition", "book", "chapterfrom", "chapterto", "vfrom", "vto"}
    for index, key :=range keys {
        params[index].urlparamname = key
        params[index].rawurldata = urlparams.Get(params[index].urlparamname)
        params[index].parseddata, params[index].parsingerror = strconv.ParseInt(params[index].rawurldata, 10, 64)
    }
    
    // filter invalid parameters (param exiss but invalid value) and return error
    for index:=0; index < len(params); index++{
        tmperror := params[index].parsingerror
        if  tmperror != nil && params[index].rawurldata != "" {
            // if the parameter is used but the value couldn't be parsed then its an error - compared to params not provided so error when parsing
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("invalid parameters"))
            return
        }
    }


    // clean the array from non passed parameters
    // and build the query based on supplied parameters 
    var filteredparams []interface{}
    filteredparams = append(filteredparams, 1) // 1 is for the author id then it will be replaced with the id from the sid cookie
    var query string = `SELECT noteid, author, 
    bibleEdition,
    bibleBook,
    bookChapterStart,
    bookChapterEnd,
    verseNumberStart,
    verseNumberEnd,
    CharNumStart,
    CharNumEnd,
    authorNote,
    highlightColor,
    FROM Note WHERE author=?`




    querypieces := []string{ 
        `AND bibleEdition=?`,
        `AND bibleBook=?`,
        `AND bookChapterStart>=?`,
        `AND bookChapterEnd<=?`,
        `AND verseNumberStart>=?`,
        `AND verseNumberEnd<=?`,
    }
    for index, value := range params{
        tmperror := value.parsingerror
        if tmperror != nil && value.rawurldata == "" {
            continue
        } else {
            filteredparams = append(filteredparams, int(params[index].parseddata))
            query = strings.Join([]string{query, querypieces[index]}, " ")
        }
    }
    // future me will need to make sure user don't retreive all their highlight from any edition, any book ...
    // but he will do it once someone is exploiting my lazyness and crashing the servers
    // make sure edition, book and chapter start and end are provided. the rest is not that important in order to prevent DOS
    var minqueryparams = []struct {
        pname string
        validated bool
    }{{"edition", false}, {"book", false}, {"chapterfrom", false}, {"chapterto", false}}
    for _, value := range params{
        for i, v := range minqueryparams {
            if v.pname == value.urlparamname && value.parsingerror == nil {
                minqueryparams[i].validated = true
            }
        }
    }
    for _,v := range minqueryparams {
        if v.validated == false {
            fmt.Println("error not all min query params provided");
            http.Error(w, "error not all min query params provided", http.StatusInternalServerError)
            return
        }
    }
    

    query = strings.Join([]string{query, ";"}, "")

    // fmt.Println(query)
    var highlights []Highlight
    if services.DB == nil {
        services.Init_DB()
        if services.DB == nil {
            fmt.Println("database connection error")
            http.Error(w, "error connecting to database", http.StatusInternalServerError)
            return
        }
    }
    fmt.Println(filteredparams...) 
    rows, err := services.DB.Queryx(query, filteredparams...)
    defer rows.Close()
    for rows.Next() {
        var highlight Highlight
        err = rows.StructScan(&highlight)
        if err != nil {
            fmt.Println("error rowscan : ", err);
            http.Error(w, "error reading from database to retreive your highlights", http.StatusInternalServerError)
            return
        } else {
            highlights = append(highlights, highlight)
        }
    }


    if rows.Err() != nil {
        fmt.Println("error rows: ", err);
        http.Error(w, "error reading from database to retreive your highlights", http.StatusInternalServerError)
        return
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(highlights)
}


func EditHighlight(w http.ResponseWriter, r *http.Request) {
    // based on highligth id edit text or any field associated except the beginning and end (for that user needs to delete the highlith and create a new one)
    // user can only change text and color - fail silently
    type DeleteHreq struct {
        Noteid int `json:"noteid" db:"noteid"`
        Newcolor int `json:"newcolor" db:"highlightColor"`
        Newtext string `json:"newtext" db:"authorNote"`
    }
    var delreq DeleteHreq;
    json.NewDecoder(r.Body).Decode(&delreq)
    query := `UPDATE Note SET highlightColor=?, authorNote=? WHERE noteid=? AND author=?;`

    _, err := services.DB.Exec(query, delreq.Newcolor, delreq.Newtext, delreq.Noteid, 1) // replace 1 with userid from sid cookie
    if err != nil {
        fmt.Println(err)
        http.Error(w, "error updating highlight in database", http.StatusInternalServerError)
        return
    }


    // w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    // json.NewEncoder(w).Encode()
}
