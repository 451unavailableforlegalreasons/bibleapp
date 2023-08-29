<template>

    <div class="page  container my-auto p-4 h-5/6" v-bind:id="pageid">
      <!-- {{ verses.length === 0 ? $emit("getmoreverse") : "non"  }} -->
      <p v-for="(value, key) in localverses" class="verse" :id="value.vnum">
      <span class="chapter" v-if="value.vnum === 1">{{ chapter }} <span style="opacity: 0;font-size:0.7rem;" id="whitespace">a</span>  </span>
        <sup>{{ value.vnum }}</sup>{{ value.verse }}
        {{ $emit("getmoreverse") }}
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
          pageid: "page"+this.pagenum
      }
  },
  props: ['verses', 'chapter', "pagenum"],
  methods: {
      yeildlastinsertedversenum() {
          // console.log(this.localverses.length)
          this.$emit("lastinsertedvnum", this.pagenum, this.localverses.length)//[this.localverses.length-1].vnum)
      },
      getnextverse () {
          // iterator over verses 
          // we need to check if the last verse overflowed the content or not
          // if it overflowed, delete it and emit event for page 2 to start filling where page1 stopped -- i have no idea how to do that with vue
          // if not continue inserting a new element (verse)
          // if we inserted all verses, emit event to fetch new verses
          // if there isn't new elements appended to the verses array, stop
          // if there is, continue adding
      },
      getmoreverse () {
          console.log("page"+this.pagenum)
        let overflow = isOverflown(document.getElementById("page"+this.pagenum))
        if (overflow == false) {
            this.$emit("getmoreverses", 1)
            this.localverses = this.verses
        } else {
            console.log("div is overflowing so no more verse for this page")
            this.yeildlastinsertedversenum()
        }
      },
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
