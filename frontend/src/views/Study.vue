<template>

 
<main  class="flex align-center justify-center flex-col ">
    <Navbar :maintitle=maintitle>
        <template class="quicktools">
            <div id="userconnected" class="flex -space-x-4">
                <img class="userimg connected rounded-full ring-2 ring-blue-300" src="data:image/jpeg;base64,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" />
                <img class="userimg afk rounded-full ring-2 ring-gray-300" src="data:image/jpeg;base64,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" />
            </div>
            <div id="toolbar">
                <font-awesome-icon @click="search()" class="quicktoolsicons" icon="fa-solid fa-magnifying-glass" size="xl" style="color: #6c6f65;"/>
                <font-awesome-icon @click="showgoto=true" class="quicktoolsicons" icon="fa-solid fa-arrows-to-circle" size="xl" style="color: #6c6f65;"/>
                <font-awesome-icon @click="highlight()" class="quicktoolsicons" icon="fa-solid fa-highlighter" size="xl" style="color: #6c6f65;"/>
                <router-link to="/"><font-awesome-icon  class="quicktoolsicons" icon="fa-solid fa-house" size="xl" style="color: #6c6f65;"/></router-link>
            </div>
        </template>
    </Navbar>
    <div id="mainpage" class="mainpage mx-2 flex justify-evenly  columns-2xs h-full content-center">

        <PopUp :show=show @close="show=false"/>
        <PopUp :show=showgoto @close="showgoto=false; this.$refs.focusMe.focus()">
            <template v-slot:header >
                <p>Go to your the book you want</p>
            </template>
            <template v-slot:body>
                <input type="text" id="goto" name="goto" placeholder="Ex: Mattew" class="border border-slate-400 focus:border-slate-400 hover:border-slate-400 rounded-lg m-0 w-full px-1 py-5" autofocus v-model=gotosearch @keyup.enter="gotobook()"/>
            </template>
            <template v-slot:footer>
                <ul>
                    <li v-for="(book, index) in gotobooks" :key=index>
                        <p v-show="index <= 10"  @click="this.bookindex = indexofBookByName(this.bible, book.name); showgoto=false">
                        {{ book.name }}
                        </p>
                        <p v-show="index === 11">...</p>
                    </li>
                </ul>
            </template>
        </PopUp>

        <Turnpage turndirection="left" 
                  :class="layoutstyle"
                  @turnpage="turnpage"
        />
        <template v-if="layout === 'page'">
            <Page @pagefilled="pagefilled"
            @getmoreverses="getmoreverses" 
            :lastinsertedindex=currentpageinfo.lastinsertedindex
            :pagenum=1
            :chapter=chapter 
            :verses=allverses 
            :class="{ 'displaybible': true }"
            />
            <Page @pagefilled="pagefilled"
            @getmoreverses="getmoreverses" 
            :lastinsertedindex=currentpageinfo.lastinsertedindex
            :pagenum=2
            :chapter=chapter 
            :verses=allverses
            :class="{ 'displaybible': true }"
            />
        </template>
        <template v-else-if="layout === 'scroll'">
            <ScrollPage :bible=bible :bookindex=bookindex
            :class="{ 'displaybible': true }"
            />
        </template>
        <Turnpage turndirection="right" 
                  :class="layoutstyle"
                  @turnpage="turnpage"
        />
    </div>
</main>

</template>

<script>
import { Settings } from "../assets/js/settings.js"

import Page from '../components/page.vue'
import ScrollPage from '../components/scrollpage.vue'
import Turnpage from "../components/turnpage.vue"
import Dmenu from "../components/dropdown.vue"
import Navbar from "../components/navbar.vue"

import PopUp from "../components/popup.vue"

export default {
  name: 'main',
  components: {
      Page,
      Turnpage,
      Navbar,
      ScrollPage,
      PopUp,
  },
  data: function() {
      return {
          // Popup data
          show: false,
          showgoto: false,
          gotosearch: "",
          // page/scroll data
          bible: undefined,
          layout: 'scroll', // default
          bookindex: 0,
          previouspageinfo: { // when turning right, save the previous state here and when turn left if clicked, bring thoses back
              startvindex: -1, // if it doesn't exists in local storage, put the user on the 1st page of an edition
              endvindex: -1,
          }, // this should be saved in the local storage and loaded up when starting
          currentpageinfo: {
              gobalastverseinserted: 0, // index of the last inserted verse on currentpage. When clicking on turn right, start from this index
              lastinsertedindex: 0, // same as above but to sync page 2 with page 1 (page 2 continues where page 1 left)
          },
          bibleedition: {"name": "unknown", "id": 1},
          chapter: 1,
          allverses: [
              {"vnum":1,"verse":"In the beginning God created the heaven and the earth."},
          ],
      }
  },
  methods: {
      gotobook() {
          // when enter is pressed look for a book that resemble to what user searched and go to it
          if (this.gotobooks.length >= 1) {
              this.bookindex = this.indexofBookByName(this.bible, this.gotobooks[0].name)
          }
          this.showgoto=false; 
          this.gotosearch='' 
      },
      turnpage(turndirection) {
          // logic for turning pages depending on the layout used

          // turnpage for scrollpage format (change the bookindex)
          if (turndirection == "right") {
              if (this.bookindex + 1 >= this.bible.length - 1) {
                  this.bookindex = 0
              } else {
                  this.bookindex +=1
              }
          } else if (turndirection == "left") {
              if (this.bookindex === 0) {
                  this.bookindex = this.bible.length - 1
              } else {
                  this.bookindex -=1
              }
          } else {
              console.log("error turning page")
          }


        // turnpage for 2 page format
          // if (turndirection === "left") {
          //     console.log("sorry can't turn left for now")
          // } else if (turndirection == "right") {
          //     // update lastpagefirstvnum var before turning 
          //     this.previouspageinfo.startvindex = 0
          //     this.previouspageinfo.endvindex = this.currentpageinfo.lastinsertedindex
          //
          //
          //     this.currentpageinfo.lastinsertedindex = 0 // index is relative to the array here (nothing global to the bible like vnum)
          //     this.allverses = this.allverses.slice(this.currentpageinfo.gobalastverseinserted)
          // } else {
          //     console.log("which way are you turning ?")
          // }
      },
      pagefilled (iterationcount, pagenum) {
          if (pagenum === 1) {
              this.currentpageinfo.lastinsertedindex = iterationcount;
          } else if (pagenum === 2) {
              let lastinsertedindex = iterationcount - 1
              this.currentpageinfo.gobalastverseinserted = this.allverses[lastinsertedindex].vnum
              // console.log(this.allverses[lastinsertedindex])
              // console.log("global", this.gobalastverseinserted)
          } else {
              console.log("invalid pagenum")
          }
      },

      indexofBookByName (bible, bookname) {
          for (let i = 0; i < bible.length; i++) {
              if (bible[i].name.toLowerCase() === bookname.toLowerCase()) {
                  return i
              }
          }
      },
      keydown(evt){
          // detect keyboard shortcuts and update data accordingly
          if (!evt) evt = event;
          if (evt.ctrlKey && evt.key =="k"){ 
              evt.preventDefault()
                this.showgoto = true
          } else if (evt.key === "Escape") {
              this.showgoto = false
              // do this for every popup
              this.showgoto = false
              this.showgoto = false
          }
          // else if (evt.shiftKey && evt.keyCode == 9){ //Shif+TAB
              //   alert("Shift+TAB");
              // }
      }


  },
    created () {
        // look for past usage of the app (edition that has been visited before: restore it (push route to it))
        // or display welcome page | select edition

        // based on usersettings lastinsertedindex... update the allverses and other fields for pages components to render where the user left last time
        this.bibleedition = this.$route.params.edition









        if (this.bibleedition == 2) {
            // fetch te corresponding bible edition and store it inside the index db 
            // the other views/components will use the dible object to lookup verses...
            // the indexed db is just for persistent localstorage (more space)
            // and they will look inside the localStorage for the bookmark/lastvisitedpage/prefered edition/...


                // don't use fecth, its async and we can't do it here
            const request = new XMLHttpRequest();
            request.open("GET", "http://localhost:8080/es_rvr.json", false); 
            request.send(null);

            if (request.status === 200) {
                // console.log(request.response);
                this.bible = JSON.parse(request.response)
            } else {
                console.log(request.status)
                console.log(request.response)
            }

        } else {
            console.log("sry can't fetch this edition'")
        }




        /*
            Books, chapter, verse number information should be in the url or localstorage    
        */

            // lastvisitdata: {
            //     edition: null,
            //     book: 0,
            //     chapter: 0,
            //     firstdisplayedverse: 0,
            // },
        let s = new Settings
        s.loadfromlocalStorage()
        s.settings.lastvisitdata.edition = this.bibleedition
        s.writetolocalStorage()
        let book = this.bible[s.settings.lastvisitdata.book]
        this.biblebook = book.name
        this.chapter = s.settings.lastvisitdata.chapter+1

        let tmparr = []
        for (const [index, verse] of book["chapters"][this.chapter-1].entries()) {
            let obj = {"vnum": index+1, "verse": verse}
            tmparr.push(obj)
        }
        this.allverses = tmparr
    },
    watch: {
        pageloadingdata: {
            handler (old, newone) {
                let usettings = new Settings
                usettings.settings = newone
                usettings.writetolocalStorage()
            },
            deep:true,
        }
    },
    computed: {
        // pageloadingdata() {
        //     let s = new Settings
        //     s.loadfromlocalStorage()
        //     return s.settings
        // },

        biblebook: function () {
            console.log(this.bookindex)
            return this.bible[this.bookindex].name
        },
        chapter: function () {
            // look for the chapter number on screen (visible or last seen | think about screen having two number diaplyed --> breakpoint to switch from one to the other)
            return ""
        },
        maintitle: function () {
            return this.bibleedition + ': ' + this.biblebook+ ' ' + this.chapter 
        },
        layoutstyle: function () {
            if (this.usersettings === undefined) {
                return {
                    'basis-2/12': true,
                    'sticky': true,
                }
            }
            else if (this.usersettings.preferences.layout | "scroll" === 'scroll') {
                return {
                    'basis-2/12': true,
                    'sticky': true,
                }

            }
        },
        gotobooks: function() {
            let list = []
            for (const book of this.bible) {
                if (book.name === undefined) {
                    return list
                }
                if (book.name.includes(this.gotosearch) && this.gotosearch !== "") {
                    list.push(book)
                }
            }
            return list
        }
    },


































/*
=========================================================


Highlight functionality


=========================================================
*/



  mounted () {
      let maindiv = document.getElementById("mainpage")
      maindiv.addEventListener(
          "mouseup",
          handlemouseup,
          // passiveSupported ? { passive: true } : false,
      );
      document.onkeydown = this.keydown;
  },
}






function handlemouseup() {
    let maindiv = document.getElementById("mainpage")
    let children = maindiv.children
    for(var i=0; i<children.length; i++){
        var child = children[i];
        // if the element contains the displaybible class, it contains the bible verses
        if (child.className.includes("displaybible")) {
            let selection = window.getSelection()
            // console.log(selection)
            //basenode is the start HTMLElement ---> .baseOffest === char0 based index
            //extentnode is the end HTMLElement ---> .extentoffset === char 0 based index 

        }
    }
}

</script>

<style scoped>
main{
    height: 100vh;
}
div.mainpage {
 height: 90vh;
}
div.nav {
    height: 10vh;
}


.quicktools{
    display: flex;
    flex-direction: row;
    justify-content: end;
    align-content: center;
}
.quicktoolsicons {
    margin: 0.3em;
}
.quicktoolsicons:hover {
    cursor: pointer;
}


.userimg {
   width: 2em; 
   height: 2em; 
}
.afk {

}
</style>
