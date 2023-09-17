<template>

 
<main  class="flex align-center justify-center flex-col " style="border: 1px solid green">
    <div id="nav" class="d-flex flex flex-row justify-between align-middle">
        <Dmenu />
        <h1 id="bookname" class="text-xl font-bold leading-none tracking-tight text-gray-900 md:text-2xl lg:text-3xl text-center py-auto">{{ bibleedition+ ': ' + biblebook+ ' ' + chapter }}</h1>
        
        <div id="quicktools">
        search, goto, highlight, bookmark ?, 
        </div>
    </div>
    <div id="mainpage" class="mainpage mx-2 flex justify-evenly  columns-2xs h-full content-center">
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
            :verses=allverses />
            <Page @pagefilled="pagefilled"
            @getmoreverses="getmoreverses" 
            :lastinsertedindex=currentpageinfo.lastinsertedindex
            :pagenum=2
            :chapter=chapter 
            :verses=allverses />
        </template>
        <template v-else-if="layout === 'scroll'">
            <ScrollPage :bible=bible :bookindex=bookindex />
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

export default {
  name: 'main',
  components: {
      Page,
      Turnpage,
      Dmenu,
      ScrollPage,
  },
  data: function() {
      return {
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
      turnpage(turndirection) {

        // turnpage for scrollpage format (change the bookindex)
        if (turndirection == "right") {
            this.bookindex +=1
        } else if (turndirection == "left") {
            this.bookindex -=1
        } else {
            console.log("error turning page")
        }


      // turnpage for 2 page format
          if (turndirection === "left") {
              console.log("sorry can't turn left for now")

          } else if (turndirection == "right") {
              // update lastpagefirstvnum var before turning 
              this.previouspageinfo.startvindex = 0
              this.previouspageinfo.endvindex = this.currentpageinfo.lastinsertedindex
              

              this.currentpageinfo.lastinsertedindex = 0 // index is relative to the array here (nothing global to the bible like vnum)
              this.allverses = this.allverses.slice(this.currentpageinfo.gobalastverseinserted)
          } else {
              console.log("which way are you turning ?")
          }
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
            return this.bible[this.bookindex].name
        },
        chapter: function () {
            // look for the chapter number on screen (visible or last seen | think about screen having two number diaplyed --> breakpoint to switch from one to the other)
            return ""
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
        }
    },


































/*
=========================================================


Highlight functionality


=========================================================
*/



  // mounted () {
  //     let maindiv = document.getElementById("mainpage")
  //     maindiv.addEventListener(
  //         "click",
  //         handleMousedown,
  //         // passiveSupported ? { passive: true } : false,
  //     );
  // },
}


// function handleMousedown() {
//     let maindiv = document.getElementById("mainpage")
//     let children = maindiv.children
//     for(var i=0; i<children.length; i++){
//         var child = children[i];
//         // child.style.color = "red";
//         if (child.id) {
//             console.log(child.id)
//             let selection = window.getSelection()
//             console.log(selection)
//
//             let verses = child.children
//             for (var j=0; j < verses.length; j++) {
//                 if (verses[j].parentElement.id)
//             }
//         }
//     }
// }
//
</script>

<style scoped>
html {
  background-color: #FADEAA;
}
/* toggle border here*/
html * {
  border: 0px !important;
}
.mainpage {
  border: 1px solid red;
  /*height: 80vh;*/
}
main {
    height: 100vh;
}
div.mainpage {
 height: 90vh;
}
div.nav {
    height: 10vh;
}
</style>
