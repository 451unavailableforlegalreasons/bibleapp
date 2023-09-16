<template>

 
<main  class="flex align-center justify-center flex-col " style="border: 1px solid green">
    <div id="nav" class="d-flex flex flex-row justify-between align-middle">
        <Dmenu />
        <h1 id="bookname" class="text-xl font-bold leading-none tracking-tight text-gray-900 md:text-2xl lg:text-3xl text-center py-auto">{{ bibleedition+ ': ' + biblebook+ ' ' + chapter }}</h1>
        
        <div id="quicktools">
            icon
            <font-awesome-icon :icon="'fas hammer'" />
        </div>
    </div>
    <div id="mainpage" class="mainpage mx-2 flex justify-evenly  columns-2xs h-full content-center">
        <Turnpage turndirection="left" 
                  :class="layoutstyle"
                  @turnpage="turnpage"
        />
        <template v-if="pageloadingdata.preferences.layout=== 'page'">
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
        <template v-else-if="pageloadingdata.preferences.layout === 'scroll'">
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

import Page from './page.vue'
import ScrollPage from './scrollpage.vue'
import  Turnpage from "./turnpage.vue"
import Dmenu from "./dropdown.vue"

export default {
  name: 'main',
  components: {
      Page,
      Turnpage,
      Dmenu,
      ScrollPage,
  },
  props: ["bible", "pageloadingdata"],
  data: function() {
      return {
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
              {"vnum":3,"verse":"And God said, Let there be light: and there was light."},
              {"vnum":4,"verse":"And God saw the light, that it was good: and God divided the light from the darkness."},
          ],
      }
  },
  methods: {
      fetchverses(bibleedition, biblebook, biblechapter, fromverse, toverse) {
      },
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
  beforeMount() {
     // based on pageloadingdata lastinsertedindex... update the allverses and other fields for pages components to render where the user left last time
     this.bibleedition = this.pageloadingdata.lastvisitdata.edition.name
     let book = this.bible[this.pageloadingdata.lastvisitdata.book]
     this.biblebook = book.name
     this.chapter = this.pageloadingdata.lastvisitdata.chapter+1



     let tmparr = []
     for (const [index, verse] of book["chapters"][this.chapter-1].entries()) {
         let obj = {"vnum": index+1, "verse": verse}
         tmparr.push(obj)
     }
     console.log(tmparr)
     this.allverses = tmparr
  },

  computed: {
    biblebook: function () {
        return this.bible[this.bookindex].name
    },
    chapter: function () {
        // look for the chapter number on screen (visible or last seen | think about screen having two number diaplyed --> breakpoint to switch from one to the other)
        return ""
    },
    layoutstyle: function () {
        if (this.pageloadingdata.preferences.layout === 'scroll') {
            return {
                'basis-2/12': true,
                'sticky': true,
            }

        }
    }
  }



































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
