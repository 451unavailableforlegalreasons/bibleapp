<template>

 
    <div @click="turnpage" class="turnpage m-0 px-6 flex flex-col align-center justify-center">
  <p class="align-self-center justify-self-center">
  {{ turndirection==="left" ? '<' : '>'}}
  </p>
</div>

</template>

<script>


export default {
  name: 'App',
  components: {
  },
  props: ['turndirection'],
  methods: {
      turnpage: function (event) {
          console.log(this.turndirection)
          if (this.turndirection=="left") {
              console.log("turning left")
          } else if (this.turndirection=="right"){
              console.log("turning right")
          } else {
              console.log("unknown direction")
          }
      }
  },
}
</script>

<style scoped>
.turnpage {
height: 100%;
transition: 0.4s background-color;
}
.turnpage:hover {
  background-color: #f0e7da;
}

</style>
