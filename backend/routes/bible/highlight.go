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



type Highlight struct {
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


func CreateHighlight(w http.ResponseWriter, r *http.Request) {
    // insert a new highligh in database for user (based on sid cookie to get userid)
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


func DeleteHighlight(w http.ResponseWriter, r *http.Request) {
    // delete highligth by passing its id (check if the user owns it)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    // json.NewEncoder(w).Encode()
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
    var query string = `SELECT * FROM Note WHERE author=?`
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
    // user can only change text and color
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK);
    // json.NewEncoder(w).Encode()
}
