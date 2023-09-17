<template>

    <div class="scrollpage  basis-8/12 container my-auto h-full">
      <p v-for="(chapter, key) in book.chapters" class="chapter" >
      <span class="chapter">{{ key + 1 }}<span style="opacity: 0;font-size:0.7rem;" id="whitespace">_</span></span>
      <span class="verse" v-for="(verse, index) in chapter" >
          <sup>{{ index +1 }}</sup>{{ " " + verse }}
      </span>
      </p>
  </div>

</template>

<script>

export default {
  name: 'scrollpage',
  components: {
  },
  data: function () {
      return {
          book: undefined,
      }
  },
  props: ['bible', "bookindex"],
  beforeMount() {
     this.book = this.bible[this.bookindex]
  },
  updated() {
     this.book = this.bible[this.bookindex]
  }
}

</script>

<style>
.scrollpage {
    overflow-y: scroll;
    overflow-x: hidden;
}


/* Hide scrollbar for Chrome, Safari and Opera */
.scrollpage::-webkit-scrollbar {
  display: none;
}

/* Hide scrollbar for IE, Edge and Firefox */
.scrollpage{
  -ms-overflow-style: none;  /* IE and Edge */
  scrollbar-width: none;  /* Firefox */
}



.verse {
    text-align: left;
    font-size: 1rem;
}

span.chapter {
  font-size: 2rem;
}
</style>
