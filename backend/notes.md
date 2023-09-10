# Notes about the Bible app project

# clientstorage branch fucntionalities
The clientstorage branch is a new and lighter approach for the server
Instead of distributing verses in groups of 15 os a specific edition, it sends once the whole bible edition 
and let the client do its own stuff.

Highlights don't change. (We suppose the client hasn't altered the bible data so there will be no verification 
if overflows in highlights...).

The bible edition will not be stored verse by verse in the database but edition will be formated, to contain
enough information for the client to know what verse is what... accros every edition, and then zipped to be sent to the client.

On the client it will stay in zipped format and be unzipped at every start (i know not that efficient but easier to implement)

## Functionalities to implement:
- users acounts (storing the right data in permanent database - sqlite for now)
    - DONE:define user requirements: (email, password, fullname, age, gender, country, language)
    - DONE:login with email and password
    - DONE:verify email function
    - DONE:reset password function
    - change email (send and email to new address and if user clicks on link: change to new email)
    - crud for all other fields (fullname, age, sex, country, language)
- user authentication and authorization


- highlight make sure user does store the same highlight twice (because they curently can)
- add emailverified column in database schema for user



