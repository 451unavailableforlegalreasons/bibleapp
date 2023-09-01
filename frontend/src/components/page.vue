<template>

    <div class="page  container my-auto p-4 h-5/6" v-bind:id="pageid">
      <!-- {{ verses.length === 0 ? $emit("getmoreverse") : "non"  }} -->
      <p v-for="(value, key) in localverses" class="verse" :id="value.vnum">
      <span class="chapter" v-if="value.vnum === 1">{{ chapter }} <span style="opacity: 0;font-size:0.7rem;" id="whitespace">a</span>  </span>
        <sup>{{ value.vnum }}</sup>{{ value.verse }}
      </p>
      <!-- <button @click="getmoreverse">more</button> -->
  </div>

</template>

<script>

export default {
  name: 'App',
  components: {
  },
  data: function () {
      return {
          localverses: [],
          pageid: "page"+this.pagenum,
          overflows: false,
          iterationCount: 0,
      }
  },
  props: ['verses', 'chapter', "pagenum", "lastinsertedindex"],
  methods: {

  },
    setup() {
    },
    mounted() {
        // fetch the first verse so update is called and the page gets filled with the next verses
        if (this.pagenum === 2 && this.lastinsertedindex != 0) {
            this.iterationCount +=  this.lastinsertedindex
        }
        let currentverse = this.verses[this.iterationCount]
        if (currentverse === undefined) {
            console.log("no verse error")
            return 
        }
        this.localverses.push(currentverse);
        this.iterationCount++;
    },
    updated() {
        this.overflows = isOverflown(document.getElementById("page"+this.pagenum))
        console.log("pagenum: ", this.pagenum)
        console.log("div overflow: ", this.overflows)
        let currentverse = this.verses[this.iterationCount]
        console.log("current verse:", currentverse)
        if (currentverse === undefined) {
            this.$emit("needmoreverse", this.iterationCount)
           return 
        }
        if (this.overflows == false) {
            this.localverses.push(currentverse);
            console.log("local verses: ", this.localverses)
            this.iterationCount++;
        } else {
            this.localverses.pop()
            this.iterationCount--;
            this.$emit("pagefilled", this.iterationCount, this.pagenum)
        }
    }
}



function isOverflown(element) {
  return element.scrollHeight > element.clientHeight || element.scrollWidth > element.clientWidth;
}




</script>

<style>
.page {
  border: 1px solid blue;
  height: 90%;
}

.verse {
    text-align: left;
    font-size: 1rem;
}

span.chapter {
  font-size: 2rem;
}
</style>
