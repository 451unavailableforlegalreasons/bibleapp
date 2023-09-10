<template>

 
<main  class="flex align-center justify-center flex-col " style="border: 1px solid green">
    <div id="nav" class="d-flex flex flex-row justify-between align-middle">
        <Dmenu />
        <h1 id="bookname" class="text-xl font-bold leading-none tracking-tight text-gray-900 md:text-2xl lg:text-3xl text-center py-auto">{{ bibleedition.name + ': ' + biblebook.name + ' ' + chapter }}</h1>
        
        <div id="idk">
        --
        </div>
    </div>
    <div id="mainpage" class="mainpage mx-2 flex justify-evenly  columns-2xs h-full content-center">
        <Turnpage turndirection="left" 
                  @turnpage="turnpage"
        />
        <Page @pagefilled="pagefilled"
              @getmoreverses="getmoreverses" 
              :lastinsertedindex=lastinsertedindex
              :pagenum=1
              :chapter=chapter 
              :verses=allverses />
        <Page @pagefilled="pagefilled"
              @getmoreverses="getmoreverses" 
              :lastinsertedindex=lastinsertedindex
              :pagenum=2
              :chapter=chapter 
              :verses=allverses />
        <Turnpage turndirection="right" 
                  @turnpage="turnpage"
        />
    </div>
</main>

</template>

<script>

import Page from './page.vue'
import  Turnpage from "./turnpage.vue"
import Dmenu from "./dropdown.vue"

export default {
  name: 'main',
  components: {
      Page,
      Turnpage,
      Dmenu,
  },
  data: function() {
      return {
          previouspageinfo: { // when turning right, save the previous state here and when turn left if clicked, bring thoses back
              startvindex: -1, // if it doesn't exists in local storage, put the user on the 1st page of an edition
              endvindex: -1,
          }, // this should be saved in the local storage and loaded up when starting
          currentpageinfo: {
              gobalastverseinserted: 0, // index of the last inserted verse on currentpage. When clicking on turn right, start from this index
              lastinsertedindex: 0, // same as above but to sync page 2 with page 1 (page 2 continues where page 1 left)
          },
          bibleedition: {"name": "unknown", "id": 1},
          biblebook: {"name": "Genesis", "ordnum": 1},
          chapter: 1,
          allverses: [
              {"vnum":1,"verse":"In the beginning God created the heaven and the earth."},
              {"vnum":2,"verse":"And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters."},
              {"vnum":3,"verse":"And God said, Let there be light: and there was light."},
              {"vnum":4,"verse":"And God saw the light, that it was good: and God divided the light from the darkness."},
              {"vnum":5,"verse":"And God called the light Day, and the darkness he called Night. And the evening and the morning were the first day."},
              {"vnum":6,"verse":"And God said, Let there be a firmament in the midst of the waters, and let it divide the waters from the waters."},
              {"vnum":7,"verse":"And God made the firmament, and divided the waters which were under the firmament from the waters which were above the firmament: and it was so."},
              {"vnum":8,"verse":"And God called the firmament Heaven. And the evening and the morning were the second day."},
              {"vnum":9,"verse":"And God said, Let the waters under the heaven be gathered together unto one place, and let the dry land appear: and it was so."},
              {"vnum":10,"verse":"And God called the dry land Earth; and the gathering together of the waters called he Seas: and God saw that it was good."},
              {"vnum":11,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":12,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":13,"verse":"And the evening and the morning were the third day."},
              {"vnum":14,"verse":"And God said, Let there be lights in the firmament of the heaven to divide the day from the night; and let them be for signs, and for seasons, and for days, and years: 1:15 And let them be for lights in the firmament of the heaven to give light upon the earth: and it was so."},


              {"vnum":15,"verse":"In the beginning God created the heaven and the earth."},
              {"vnum":16,"verse":"And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters."},
              {"vnum":17,"verse":"And God said, Let there be light: and there was light."},
              {"vnum":18,"verse":"And God saw the light, that it was good: and God divided the light from the darkness."},
              {"vnum":19,"verse":"And God called the light Day, and the darkness he called Night. And the evening and the morning were the first day."},
              {"vnum":20,"verse":"And God said, Let there be a firmament in the midst of the waters, and let it divide the waters from the waters."},
              {"vnum":21,"verse":"And God made the firmament, and divided the waters which were under the firmament from the waters which were above the firmament: and it was so."},
              {"vnum":22,"verse":"And God called the firmament Heaven. And the evening and the morning were the second day."},
              {"vnum":23,"verse":"And God said, Let the waters under the heaven be gathered together unto one place, and let the dry land appear: and it was so."},
              {"vnum":24,"verse":"And God called the dry land Earth; and the gathering together of the waters called he Seas: and God saw that it was good."},
              {"vnum":25,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":26,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":27,"verse":"And the evening and the morning were the third day."},
              // filling with random thing just to overflow second page
              {"vnum":28,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":29,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":30,"verse":"And the evening and the morning were the third day."},
              {"vnum":31,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":32,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":33,"verse":"And the evening and the morning were the third day."},
              {"vnum":34,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":35,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":36,"verse":"And the evening and the morning were the third day."},
              {"vnum":37,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":38,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":39,"verse":"And the evening and the morning were the third day."},
              {"vnum":40,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":41,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":42,"verse":"And the evening and the morning were the third day."},
              {"vnum":43,"verse":"And God said, Let the earth bring forth grass, the herb yielding seed, and the fruit tree yielding fruit after his kind, whose seed is in itself, upon the earth: and it was so."},
              {"vnum":44,"verse":"And the earth brought forth grass, and herb yielding seed after his kind, and the tree yielding fruit, whose seed was in itself, after his kind: and God saw that it was good."},
              {"vnum":45,"verse":"And the evening and the morning were the third day."},
],
      }
  },
  methods: {
      fetchverses(bibleedition, biblebook, biblechapter, fromverse, toverse) {
      },
      turnpage(turndirection) {
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
              this.lastinsertedindex = iterationcount;
          } else if (pagenum === 2) {
              let lastinsertedindex = iterationcount - 1
              this.gobalastverseinserted = this.allverses[lastinsertedindex].vnum
              // console.log(this.allverses[lastinsertedindex])
              // console.log("global", this.gobalastverseinserted)
          } else {
              console.log("invalid pagenum")
          }
      }

  },
  updated () {
      console.log(this.allverses)
  }

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
