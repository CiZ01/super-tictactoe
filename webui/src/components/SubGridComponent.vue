<script>

// js classes
import SubGrid from '@/classes/subGrid.js';

export default {
  name: 'SubGridComponent',

  emits: ['game-over', 'next-move'],

  props:{
    player: {
      type: Number,
      required: true,
    },
    currentIndex: {
      type: Number,
      required: true,
    },
    myIndex:{
      type: Number,
      required: true,
    }
  },

  data() {
    return {
      subGrid: new SubGrid(),
      finished: false,
      myID: this.myIndex,
      marks: ['➖','❌', '⭕'],

      timeToShake: false,
    }
  },
  methods: {
    setBox(index, player) {
      if (this.currentIndex != -1 && this.myID !== this.currentIndex){
        this.timeToShake = true;
        setTimeout(() => {
          this.timeToShake = false
        }, 1500)
        return;
      }
      if (this.finished) {
        return;
      }
      const [win, err] = this.subGrid.setBox(index, player);
      if (err) {
        console.log(err.message);
        return;
      }
      if (win) {
        this.finished = true;
        this.$emit('game-over', this.subGrid.winner);
        return;
      }
      this.$emit('next-move', index);
    },

  },


}

</script>


<template>
  <div id="subgrid">
    <div class="subBoard">
      <div class="subCell" :class="{ shake: timeToShake }" v-for="(box, index) in subGrid.grid" :key="index" @click="setBox(index, (player%2)+1)">
        {{ marks[box] }}
      </div>
    </div>
  </div>
</template>


<style>
body {
  background-color: #2c3e50;
}


#subgrid {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #f1f1f1;
}


/* Stili per la griglia */
.subBoard {
  display: grid;
  grid-template-columns: repeat(3, auto);
  grid-gap: 0.1em;
  width: auto;
  height: auto;
  margin: 0 auto;
}

.subCell {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3em;
  color: #000;
  border: 2px solid #000;
  cursor: pointer;
  user-select: none;
  transition: background-color 0.2s;
}

.subCell:hover {
  background-color: #e6e6e6;
}


/* Animations */
.shake {
  animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
  transform: translate3d(0, 0, 0);
}

@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }

  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }

  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }

  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}
</style>
