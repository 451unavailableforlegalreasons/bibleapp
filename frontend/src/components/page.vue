<template>

  <div class="page  container my-auto p-4 h-5/6">
      <p v-for="(value, key) in verses" class="verse" :id="value.vnum">
      <span class="chapter" v-if="value.vnum === 1">{{ chapter }} <span style="opacity: 0;font-size:0.7rem;" id="whitespace">a</span>  </span>
        <sup>{{ value.vnum }}</sup>{{ value.verse }}
        {{ incrementlastinsertedindex() }}
      </p>
  </div>

</template>

<script>

export default {
  name: 'App',
  components: {
  },
  props: ['verses', 'page1lastinsertedvnum', 'chapter'],
  data: function() {

  },
  methods: {
      incrementlastinsertedindex: function() {
          this.$emit("incrementlastinsertedindex")
      },
      getnextverse: function() {
          // iterator over verses 
          // we need to check if the last verse overflowed the content or not
          // if it overflowed, delete it and emit event for page 2 to start filling where page1 stopped -- i have no idea how to do that with vue
          // if not continue inserting a new element (verse)
          // if we inserted all verses, emit event to fetch new verses
          // if there isn't new elements appended to the verses array, stop
          // if there is, continue adding
      }
  }
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
