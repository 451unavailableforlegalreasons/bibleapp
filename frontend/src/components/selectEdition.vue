<template>

    <ul id="editionSelection">
        <li v-for="edition in this.editionlist" :key=edition.editionid
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
        editionlist: [],
    }
  },
  methods: {
        selectEd: function (editionid) {
            console.log("edition selected: ", editionid)
            this.$emit("selectEd", editionid)
            if (editionid == 2) {
                // fetch te corresponding bible edition and store it inside the index db 
                // the other views/components will look inside the database to retreive the verses
                // and they will look inside the localStorage for the bookmark/lastvisitedpage/prefered edition/...
                fetch("es_rvr.json")
                    .then(response => response.json())
                    .then(content => console.log(content))
            } else {
                console.log("sry can't fetch this edition'")
            }
        }
  },
  beforeMount () {
      this.editionlist = [
          {"editionid": 1, "name": "bruh", "lang": "fr"},
          {"editionid": 2, "name": "es_rvr", "lang": "es"},
          {"editionid": 3, "name": "brou", "lang": "en"},
      ]

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
