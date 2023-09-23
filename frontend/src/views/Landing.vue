<template>
    <h1>Bible Study</h1>


    <h2>A minimalist application to read the bible and tool to never forget the <br/> verses that resonate with you</h2>
    <p>You can use the application without creating an account by first selecting the edition you want to read by clicking 
    <router-link to="/chooseedition">here</router-link>
    </p>

    <p>
    If you don't know what this application is about, i highly encourage you to start by 
    <router-link to="/discover">going throught the discover guide</router-link>
    It will show you all the functionality of the app so you can start on the right foot in your bible study
    </p>

    <p>
    If you want to do something else, here are the links you can click
    <router-link to="/users/auth">Login/Register</router-link>
    <router-link to="/discover">link</router-link>
    <router-link to="/discover">link</router-link>
    <router-link to="/discover">link</router-link>
    <router-link to="/discover">link</router-link>
    <router-link to="/discover">link</router-link>
    </p>


    <!-- <router-view/> -->
</template>




<script>
import { Settings } from "../assets/js/settings.js"
export default {
    name: 'App',
    components: {
    },
    data: function () {
        return {}
    },
    created () {
        // look for past usage of the app (edition that has been visited before: restore it (push route to it))
        // or display welcome page | select edition

        let usettings = new Settings
        usettings.loadfromlocalStorage
        this.pageloadingdata = usettings.settings
        console.log("homepage: ", usettings.settings)


        if (usettings.settings.hasselectededition === true) {
            this.$route.push("/study/"+usettings.settings.preferededition.id)
        } else {
            //
        }
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
        pageloadingdata() {
            let s = new Settings
            s.loadfromlocalStorage()
            return s.settings
        }
    },
}
</script>




<style>
</style>
