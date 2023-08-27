<template>

 
<main class="flex align-center justify-center" style="border: 1px solid green">
    <div class="mainpage mx-2 flex justify-evenly py-auto columns-2xs h-screen content-center">
        <Turnpage turndirection="left" />
        <Page :chapter=chapter :verses=verses :page1lastinsertedvnum=page1lastinsertedvnum />
        <Page :chapter=chapter :verses=verses :page1lastinsertedvnum=page1lastinsertedvnum />
        <Turnpage turndirection="right" />
    </div>
</main>

</template>

<script>

import Page from './page.vue'
import  Turnpage from "./turnpage.vue"

export default {
  name: 'App',
  components: {
      Page,
      Turnpage,
  },
  data: function() {
      return {
          bibleedition: {"name": "unknown", "id": 1},
          biblebook: {"name": "Genesis", "ordnum": 1},
          chapter: 1,
          verses: [
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
              {"vnum":14,"verse":"And God said, Let there be lights in the firmament of the heaven to divide the day from the night; and let them be for signs, and for seasons, and for days, and years: 1:15 And let them be for lights in the firmament of the heaven to give light upon the earth: and it was so."}],
        page1lastinsertedvnum: 0, // the page component will edit it

      }
  },
  methods: {
      incrementlastinsertedindex: function() {
          this.page1lastinsertedvnum+=1
      }
  }
}
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
</style>
