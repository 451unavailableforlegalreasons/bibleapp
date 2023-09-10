# Notes about the Bible app project

# Changes
- the server won't distribute verses with an api anymore but give the entire bible version to the browser (to be stored locally (an edition never changes so its ok)).
- new way of doing things:
    - user request a default formated bibile edition (on first load, then user can change the edition)
    - server gives it
    - edition stored locally on user storage
    - highlights... are still pushed to the server (since the edition is formated, the char count, verse number will be the same)
    - navigation doesn't iumplies requests anymore but saving browser state...

## Functionalities to implement:
- users acounts (storing the right data in permanent database - sqlite for now)
    - DONE:define user requirements: (email, password, fullname, age, gender, country, language)
    - DONE:login with email and password
    - verify email function
    - change email (send and email to new address and if user clicks on link: change to new email)
    - reset password function
    - crud for all other fields (fullname, age, sex, country, language)
- user authentication and authorization


- highlight make sure user does store the same highlight twice (because they curently can)
- add emailverified column in database schema for user


## Work in progress
- DONE: defining how to store bible  (sql table)
- DONE: parsing a bible and put it in the database
- DONE: Endpoint to retreive verses list from a book chapter
    (req: Mattew 12 1-25 --> Book: Mattew, Chapter:12, verse:1-15)
    (resp: json with all verses)

