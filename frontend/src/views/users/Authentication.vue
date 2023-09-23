<template>

    <div id="formwrapper">
        <div  v-if="logORreg === 0"><LoginForm @toggleForm="changeform" @submitform="submitform" /></div>
        <div   v-else><RegisterForm @toggleForm="changeform" @submitform="submitform" /></div>
    </div>

</template>




<script>
// import { Settings } from "../assets/js/settings.js"
import LoginForm from "./components/loginform.vue"
import RegisterForm from "./components/registerform.vue"


export default {
    name: 'AuthenticationForm',
    components: {
        LoginForm,
        RegisterForm,
    },
    data: function () {
        return {
            logORreg: 0, // 0 = login 1 = register

        }
    },
    methods: {
        changeform: function () {
            console.log("toggle form")
            if (this.logORreg === 0) {
                this.logORreg = 1
            } else { 
                this.logORreg = 0 
            }
        },
        submitform: function (obj) {
            // make request to server for login/register and associated logic (redirects, cookie storage...)
            console.log(obj)
        }
    }
}
</script>




<style>
</style>
