<template>

    <template v-if="this.pageloadingdata.hasselectededition === true">
        <MainPage :bible=bible :pageloadingdata=pageloadingdata />
    </template>
    <template v-else>
        <h1> Sorry </h1>
    </template>
</template>

<script>


import MainPage from '../components/main.vue'
import SelectBibleEdition from '../components/selectEdition.vue'
import { Settings } from "../assets/js/settings.js"

export default {
    name: 'App',
    components: {
        MainPage,
        SelectBibleEdition,
    },
    data: function () {
        return {
            pageloadingdata: undefined,
            bible: undefined, // this contains the json bible
        }
    },
    methods: {
    },
    created() {
        if (this.$route.edition === 2) {
            // fetch te corresponding bible edition and store it inside the index db 
            // the other views/components will use the dible object to lookup verses...
            // the indexed db is just for persistent localstorage (more space)
            // and they will look inside the localStorage for the bookmark/lastvisitedpage/prefered edition/...


            // don't use fecth, its async and we can't do it here
            const request = new XMLHttpRequest();
            request.open("GET", "es_rvr.json", false); 
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

    },
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
    url(../assets/fonts/dayroman/DAYROM__.ttf) format("truetype");
}
</style>
