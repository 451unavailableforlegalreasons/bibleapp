<template>

    <ul id="editionSelection">
        <li v-for="(edition, index) in this.editionlist" :key=index
            @click="selectEd(edition.editionid)"
            > 
            {{ edition.name }} ---- {{edition.lang}}
        </li>
    </ul>




</template>

<script>


export default {
  name: 'selectEdition',
  data: function () {
    return {
        editionlist: [
          {"editionid": 1, "name": "bruh", "lang": "fr"},
          {"editionid": 2, "name": "es_rvr", "lang": "es"},
          {"editionid": 3, "name": "brou", "lang": "en"},
        ],
    }
  }, 
  methods: {
      selectEd: function (editionindex) {
          this.$emit("selectEd", this.editionlist[editionindex-1])
          console.log(this.editionlist[editionindex-1])
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
