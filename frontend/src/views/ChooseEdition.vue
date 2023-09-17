<template>

    <Navbar :maintitle=navbartitle />
    <h1>Download and read the bible in the version you like</h1>
    <ul id="editionSelection">
        <li v-for="(edition, index) in this.editionlist" :key=index
            @click="selectEd(edition.editionid)"
            > 
            {{ edition.name }} ---- {{edition.lang}}
        </li>
    </ul>




</template>



<script>
import { Settings } from "../assets/js/settings.js"

import { Navbar } from "../components/navbar.vue"

export default {
  name: 'ChooseBiblePage',
  components: {
      Navbar,
  },
  data: function () {
    return {
        navbartitle: "Edition picker",
        editionlist: [
          {"editionid": 1, "name": "bruh", "lang": "fr"},
          {"editionid": 2, "name": "es_rvr", "lang": "es"},
          {"editionid": 3, "name": "brou", "lang": "en"},
        ],
    }
  }, 
  methods: {
      selectEd: function (editionindex) {
          // change route to /study/{bibleeditionID}
          this.$router.push({ name: 'study', params: { edition: this.editionlist[editionindex-1].editionid} })
          let settings = new Settings
          settings.loadfromlocalStorage()
          console.log(settings.settings)
          settings.settings.preferededition.id = this.editionlist[editionindex-1].editionid
          settings.settings.preferededition.name = this.editionlist[editionindex-1].name
          settings.settings.preferededition.lang = this.editionlist[editionindex-1].lang
          settings.settings.hasselectededition = true
          console.log(settings.settings)
          settings.writetolocalStorage()

      }
  },
  created() {

      // fetch all bible editon available from sever
      // and update editionlist


  }
}
</script>

<style scoped>

</style>
