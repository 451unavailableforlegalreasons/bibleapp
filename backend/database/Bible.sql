CREATE TABLE IF NOT EXISTS BibleEdition (
    name VARCHAR(20) UNIQUE NOT NULL,
    language VARCHAR(20) NOT NULL,
    id TINYINT PRIMARY KEY UNIQUE NOT NULL 
);

CREATE TABLE IF NOT EXISTS BibleZip (
    edition TINYINT NOT NULL,
    zipcontent BLOB NOT NULL,


    FOREIGN KEY (edition) REFERENCES BibleEdition(id)
);




-- highlights and note table

CREATE TABLE IF NOT EXISTS Note (
    noteid INTEGER PRIMARY KEY NOT NULL ,--AUTOINCREMENT,
    author TINYINT NOT NULL, -- linked to users table

    bibleEdition TINYINT NOT NULL, -- linked to BibleEdition table
    bibleBook TINYINT NOT NULL, -- linked to book table
    bookChapterStart TINYINT NOT NULL,
    bookChapterEnd TINYINT NOT NULL,
    verseNumberStart TINYINT NOT NULL,
    verseNumberEnd TINYINT NOT NULL,

    CharNumStart INT NOT NULL,
    CharNumEnd INT NOT NULL,

    authorNote TEXT NOT NULL,
    highlightColor TINYINT NOT NULL, -- hard coded into the front end

    FOREIGN KEY(author) REFERENCES Users(id)
);

