<template>

    <section class="bg-gray-50">
        <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
            <a href="#" class="flex items-center mb-6 text-2xl font-semibold text-gray-900">
                <div class="flex flex-row">
                    <Applogo /> 
                    <p>Simple Bible </p>
                </div>
            </a>
            <div class="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0">
                <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
                        Create an account
                    </h1>
                    <form class="space-y-4 md:space-y-6" @submit.prevent="send">
                        <div>
                            <label for="email" class="block mb-2 text-sm font-medium text-gray-900">Your email</label>
                            <input v-model="email" type="email" name="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5" required="">
                        </div>
                        <div>
                            <label for="password" class="block mb-2 text-sm font-medium text-gray-900 ">Password</label>
                            <input v-model="password1" type="password" name="password" id="password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5" required="">
                        </div>
                        <div>
                            <label for="confirm-password" class="block mb-2 text-sm font-medium text-gray-900 ">Confirm password</label>
                            <input v-model="password2" type="password" name="confirm-password" id="confirm-password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5" required="">
                        </div>
                        <!-- <div class="flex items-start"> -->
                            <!-- <div class="flex items-center h-5"> -->
                            <!--     <input id="terms" aria-describedby="terms" type="checkbox" class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300" required=""> -->
                            <!-- </div> -->
                            <!-- <div class="ml-3 text-sm"> -->
                            <!--     <label for="terms" class="font-light text-gray-500" href="#"><a id="terms" href="terms">Terms and Conditions</a></label> -->
                            <!-- </div> -->
                        <!-- </div> -->
                        <button  type="submit" class="w-full border border-black-500 bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center">Create an account</button>
                        <p class="text-sm font-light text-gray-500">
                        Already have an account? <button  @click="toggle()"  class="font-medium text-primary-600 hover:underline dark:text-primary-500">Login here</button>
                        </p>
                    </form>
                </div>
            </div>
        </div>
    </section>



</template>




<script>
import Applogo from "../../../components/appsvg.vue"


export default {
    components: {
        Applogo,
    },
    data: function () {
        return {
            email: "",
            password1: "",
            password2: ""
        }
    },
    methods: {
        toggle: function() {
            console.log("zeofhauiezpfhjo")
            this.$emit("toggleForm")
        },
        send: function() {
            if (this.password1 !== this.password2) {
                alert("password do not match")
            } else {
                const obj = {"type": "register", "data": {"email": this.email, "password": this.password1}}
                this.$emit("submitform", obj)
            }
        },
    }

}
</script>




<style>
</style>
