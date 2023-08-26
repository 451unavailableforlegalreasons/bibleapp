CREATE TABLE IF NOT EXISTS BibleEdition (
    name VARCHAR(20) UNIQUE NOT NULL,
    id TINYINT PRIMARY KEY UNIQUE NOT NULL 
);


CREATE TABLE IF NOT EXISTS BibleBook (
    name VARCHAR(20) unique not null,
    id TINYINT PRIMARY KEY UNIQUE NOT NULL 
);


CREATE TABLE IF NOT EXISTS Verse (
    bibleEdition TINYINT NOT NULL, -- linked to BibleEdition table
    bibleBook TINYINT NOT NULL, -- linked to book table
    bookChapter TINYINT NOT NULL,
    verseNumber TINYINT NOT NULL,


    verseContent TEXT NOT NULL,



    FOREIGN KEY(bibleEdition) REFERENCES BibleEdition(id),
    FOREIGN KEY(bibleBook) REFERENCES BibleBook(id)
);




-- highlights and note table


CREATE TABLE IF NOT EXISTS Note (
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

