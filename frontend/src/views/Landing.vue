<template>
    <p>Display thoses links only if the user is not authenticated (for him to register, login, select and edition (it will then redirect automatically))</p>
    <router-link to="/">Home</router-link> |
    <router-link to="/chooseedition">Choose  edition to read</router-link>
    <router-view/>
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
