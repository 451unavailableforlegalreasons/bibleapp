import sqlite3
import re
import json
import pprint
import base64



def connectToDb():
    con = sqlite3.connect("../../database.db")
    cur = con.cursor()
    # res = cur.execute("SELECT * FROM BibleEdition")
    # print(res.fetchone())
    return con, cur

con, cur = connectToDb()
files = ["fr_apee.json.gz", "en_bbe.json.gz", "en_kjv.json.gz"] #"en_kjv.json"
index = 1
for file in files:
	f = open(file, "rb")
	data= f.read()
	#b64 = base64.b64encode(data)
	f.close()
	cur.execute("INSERT INTO BibleZip (edition, zipcontent) Values (?, ?);",(index, data))
	index+=1
con.commit()
#con, cur = connectToDb()


#bookcounter = 1
#for key in map:
#    print("Loading book of :", key)
#    cur.execute("INSERT INTO BibleBook Values (\"{}\", {}, {}, {});".format(key, bookcounter, bookcounter, 1))

#    for i in range(0,len(map[key])):
        # print(verse)
#        cur.execute("INSERT INTO Verse Values ({}, {}, {}, {}, \"{}\");".format(1,bookcounter,map[key][i]['chap'], map[key][i]['num'], map[key][i]['verse']))
#    bookcounter+=1
    
#print("Commiting changes")
#con.commit()
#print("job done")


