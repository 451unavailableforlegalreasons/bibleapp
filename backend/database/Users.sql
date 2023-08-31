CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY UNIQUE ,--NOT NULL, -- index table based on id
    fullname VARCHAR(30) NOT NULL,
    email VARCHAR(35) NOT NULL UNIQUE, -- index table based on this too
    password BINARY(32) NOT NULL,
    age TINYINT NOT NULL,
    gender TINYINT NOT NULL, -- 1 man 2 woman 3 other (even if doesn't exists)
    country VARCHAR(2) NOT NULL, -- country code list
    language VARCHAR(2) NOT NULL -- same as for country code but this way we only distribute the bible in the desired language
);



-- cerate user
INSERT INTO Users (fullname, email, password, age, gender, country, language) VALUES (?, ?, ?, ?, ?, ?, ?);

-- read / verify user exists (login) and get non PII data - (log out is an other table) 
SELECT id, age, gender, country, language FROM Users WHERE email=? AND password=?;

-- read/ verify user exists (login) and get User with PII data - (log out is an other table) 
SELECT id, email, fullname, age, gender, country, language FROM Users WHERE email=? AND password= ?;

-- update user (no email, id, password here)
UPDATE Users SET fullname=?, age=?, gender=?, country=?, language=? WHERE id=?;

-- delete user 
DELETE FROM Users WHERE id=?;

------------------------------
--  non basic CRUD queries  --
------------------------------

-- update user email - change email endpoint
UPDATE Users SET email=? WHERE id=?;

-- update user password  - reset password endpoint
UPDATE Users SET password=? WHERE id=?;


-- ================================ --
--          Session table           --
-- ================================ --


CREATE TABLE IF NOT EXISTS Session(
    sid VARCHAR(20) PRIMARY KEY UNIQUE NOT NULL, 
    b64content TEXT not null
);

--------------------------------------------
--  turning above queries into procedures --
--------------------------------------------
-- i don't plan to do it for now
