<template>

    <div class="page  container my-auto h-5/6" v-bind:id="pageid">
      <p v-if="(this.lastinsertedindex !== 0 && this.pagenum === 2) || (this.pagenum === 1)" v-for="(value, key) in localverses" class="verse" :id="value.vnum">
      <span class="chapter" v-if="value.vnum === 1">{{ chapter }} <span style="opacity: 0;font-size:0.7rem;" id="whitespace">a</span>  </span>
        <sup>{{ value.vnum }}</sup>{{ value.verse }}
      </p>
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
          iterationCount: 0,
          pagefilled: false,
          firstiterationdone: false,
      }
  },
  props: ['verses', 'chapter', "pagenum", "lastinsertedindex"],
  methods: {

  },
    setup() {
    },
    mounted() {
        // put something in page so update gets called and the page is filled until overflow
        // this.iterationCount = this.globalverseoffset
        if (this.pagenum === 1) {
            let currentverse = this.verses[this.iterationCount]
            if (currentverse === undefined) {
                console.log("no verse error")
                return 
            } else {
                this.localverses.push(currentverse);
                this.iterationCount++;
            }
        } else {
            //wait for this.lastinsertindex to change (event caused by page 1)
            // it will trigger update in page1 (but nochange) & update in page2 -> pick verse up from where page1 left
            return
        }


    },
    watch: {
        verses(newv, oldv) {
            //clean from old verses and push new ones
            this.localverses = []
            this.pagefilled = false
            this.firstiterationdone = false
            // put something in page so update gets called and the page is filled until overflow
            this.iterationCount = this.lastinsertedindex
            console.log(this.iterationCount)
            if (this.pagenum === 1) {
                let currentverse = this.verses[this.iterationCount]
                if (currentverse === undefined) {
                    console.log("no verse error")
                    return 
                } else {
                    this.localverses.push(currentverse);
                    this.iterationCount++;
                }
            } else {
                //wait for this.lastinsertindex to change (event caused by page 1)
                // it will trigger update in page1 (but nochange) & update in page2 -> pick verse up from where page1 left
                return
            }
        }

    },
    updated() {
        if (this.pagenum === 2 && this.lastinsertedindex !== 0 && this.firstiterationdone === false) {
            // page 1 was filled (lastinsertedindex is not longer 0 && it caused an update on this doc)
            // but its first iteration so increment iterationCount
            this.iterationCount += this.lastinsertedindex
            this.firstiterationdone = true
        }




        let overflows = isOverflown(document.getElementById("page"+this.pagenum))
        let currentverse = this.verses[this.iterationCount]
        if (currentverse === undefined) {
            this.$emit("needmoreverse", this.iterationCount)
           return 
        }
        if (this.pagefilled == false && overflows == false) {
            this.localverses.push(currentverse);
            this.iterationCount++;
        } else if (this.pagefilled === true && overflows == false){
            return
        } else if (this.pagefilled ==  false && overflows == true) {
            this.pagefilled = true
            this.localverses.pop()
            this.iterationCount--;
            console.log("iteration count filled p2: ", this.iterationCount)
            this.$emit("pagefilled", this.iterationCount, this.pagenum)
        }
        console.log(overflows, currentverse, this.pagefilled, this.iterationCount, this.localverses)

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
