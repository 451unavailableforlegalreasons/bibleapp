<template>

    <div id="App">
        <template v-if="this.hasselectededition === true">
            <MainPage/>
        </template>
        <template v-else>
            <SelectBibleEdition @selectEd=selectEd />
        </template>
    </div>
</template>

<script>


import MainPage from './components/main.vue'
import SelectBibleEdition from './components/selectEdition.vue'

export default {
  name: 'App',
  components: {
      MainPage,
      SelectBibleEdition,
  },
  data: function () {
      return {
          // should never be access directly by child components but through computed propreties (for reading) | writing to this is ok
          pageloadingdata: {
              hasselectededition: false,
              preferededition: {id: null, name:null, lang:null},
              lastvisitdata: {
                  edition: null,
                  book: null,
                  chapter:null,
                  firstdisplayedverse: null,
              },
          },
      }
  },
  methods: {
      selectEd: function (editionid) {
          console.log('received edition:', editionid)
          this.pageloadingdata.hasselectededition = true
          this.pageloadingdata.preferededition = editionid 
          localStorage.setItem("pageloadingdata", this.pageloadingdata)

      }
  },
  beforeMount() {
      // see if user has already used the app and put his prefered edition at last visited page or ask him to select edition
      let pageloadingdata = localStorage.getItem("pageloadingdata")
      console.log(pageloadingdata)
      if (pageloadingdata === null) {
          // create the object model and store it with no content
          pageloadingdata = {
              hasselectededition: false,
              preferededition: {id: null, name:null, lang:null},
              lastvisitdata: {
                  edition: null,
                  book: null,
                  chapter:null,
                  firstdisplayedverse: null,
              },
          }
          let pageloadstr = JSON.stringify(pageloadingdata)
          localStorage.setItem("pageloadingdata", pageloadstr)
      } else {
          // if the user has already an pageloadingdata object stored, retreive change the component object
          let parsedpageloadingdata = JSON.parse(pageloadingdata)
          this.pageloadingdata = parsedpageloadingdata 
      }
  },
    watch: {
        pageloadingdata: {
            handler (old, newone) {
                let pageloadstr = JSON.stringify(newone)
                console.log("into str: ", pageloadstr)
                localStorage.setItem("pageloadingdata", pageloadstr)
            },
            deep:true,
        }
    },
    computed: {
        pageloadingdata() {
            let pageloadingdata = localStorage.getItem("pageloadingdata")
            let parsedpageloadingdata = JSON.parse(pageloadingdata)
            return parsedpageloadingdata
        }
    }

   
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /* color: #fff5e8 */
  /* margin-top: 60px; */
}
html {
  background-color: #FCE9C5;
}

html * {
    border: 0px solid !important;
    font-family: "dayrom";
}


@font-face {
  font-family: "dayrom";
  src: local("dayrom"),
   url(./assets/fonts/dayroman/DAYROM__.ttf) format("truetype");
}
</style>
