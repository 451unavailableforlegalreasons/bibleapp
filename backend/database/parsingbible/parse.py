import sqlite3
import re
import pprint

# this program does not function well
# it cut off some verses




def connectToDb():
    con = sqlite3.connect("../database.db")
    cur = con.cursor()
    # res = cur.execute("SELECT * FROM BibleEdition")
    # print(res.fetchone())
    return con, cur


# connectToDb()
map = {}
currentbook = ""
linecounter= 0
con, cur = connectToDb()

f = open("tempbible.txt")
line = f.readline()
while line != "":
    if re.search("^@",line):
        currentbook = line[1:-1]
        # print(currentbook)
        map[currentbook] = [] # array of verses
    elif re.search("^[0-9]+:[0-9]+", line):
        try:
            chapter = line.split(':')[0]
            versenum = re.split(" ", line.split(':')[1])[0]
        except:
            # tiny exeption which mess everything (the verse is split on two lines so we don't look of a new chapter or versenum)
            # print(chapter, versenum)
            # print(linecounter, line.split(':'))
            pass
        try:
            versecontent = re.split("^[0-9]+:[0-9]+ ", line)[1]
            # print(versecontent)
        except:
            # print(re.split("^[0-9]+:[0-9]+ ", line))
            versecontent += re.split("^[0-9]+:[0-9]+ ", line)[0]
        
        try:
            map[currentbook].append({
                "chap": int(chapter),
                "num": int(versenum),
                "verse": re.sub("\n", "", versecontent)
                })
        except:
            print(chapter, versenum, versecontent)
    line = f.readline()
    linecounter+=1

f.close()

# pprint.pprint(map)


print("Done reading and parsing file")
# for key in map:
# print(map["Genesis"])


bookcounter = 1
for key in map:
    print("Loading book of :", key)
    cur.execute("INSERT INTO BibleBook Values (\"{}\", {}, {}, {});".format(key, bookcounter, bookcounter, 1))

    for i in range(0,len(map[key])):
        # print(verse)
        cur.execute("INSERT INTO Verse Values ({}, {}, {}, {}, \"{}\");".format(1,bookcounter,map[key][i]['chap'], map[key][i]['num'], map[key][i]['verse']))
    bookcounter+=1
    
print("Commiting changes")
con.commit()
print("job done")
