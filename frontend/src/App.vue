<template>

    <div id="App">
        <MainPage/>
    </div>
</template>

<script>


import MainPage from './components/main.vue'

export default {
  name: 'App',
  components: {
      MainPage,
  },
  data: function () {
      return {

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
  background-color: #FCE9C5;
  /* margin-top: 60px; */
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
