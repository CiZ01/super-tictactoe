<script>

// js classes
import masterGrid from '@/classes/masterGrid.js';

import SubGridComponent from '@/components/SubGridComponent.vue';

export default {
  components: {
    SubGridComponent,
  },

  name: 'MasterGrid',
  data() {
    return {
      masterGrid: new masterGrid(),
      lastPlayer: 0,
      message: '',
      finished: false,
      nextBox: -1,
      marks: ['➖','❌', '⭕']
    }
  },
  methods: {
    setBox(index, player) {
      if (this.finished) {
        return;
      }
      let [win, err] = this.masterGrid.setBox(index, player);
      if (err) {
        console.log(err.message);
        return;
      }
      if (win) {
        this.finished = true;
        this.message = `Ha vinto il giocatore ${player}`;
      }

      this.nextMove(index);
    },

    nextMove(newIndex){
      if (this.masterGrid.grid[newIndex] !== 0) {
        this.nextBox = -1;
      } else {
        this.nextBox = newIndex;
      }
      this.lastPlayer = (this.lastPlayer%2)+1;
      return;
    },

    resetBoard() {
      this.masterGrid.resetGrid();
      this.lastPlayer = 0;
    },
  },
}

</script>

<template>
  <div id="app">
    <h1>Tic Tac Toe</h1>
    <div class="board">
      <div class="cell" :class="{ 'completed-cell': masterGrid.grid[index] != 0 }" v-for="(box, index) in masterGrid.grid" :key="index">
        <div class="shadow-box" v-if="nextBox != -1 && nextBox != index">
        </div>
        <SubGridComponent :player="lastPlayer" :currentIndex="nextBox" :myIndex="index" 
                          @game-over="(player) => setBox(index, player)" 
                          @next-move="(index) => nextMove(index)" v-if="masterGrid.grid[index] == 0"
        />
        <div v-if="masterGrid.grid[index] != 0" class="mark">
          {{ marks[masterGrid.grid[index]] }}
        </div>
      </div>
    </div>
    <p>{{ message }}</p>
    <button @click="resetBoard">Resetta Griglia</button>
  </div>
</template>

<style>
body {
  background-color: #2c3e50;
}


#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #f1f1f1;
  margin-top: 60px;
}


/* Stili per la griglia */
.board {
  display: grid;
  grid-template-columns: repeat(3, auto);
  grid-gap: 0.1em;
  width: 50em;
  height: 50em;
  margin: 0 auto;
}

.cell {
  display: flex;
  align-items: center;
  justify-content: center;

  color: #000;
  border: 2px solid #000;
  cursor: pointer;
  user-select: none;
  transition: background-color 0.2s;

  position: relative;

  padding: 2em;

  z-index: 1;
}

.completed-cell{
  background-color: #e6e6e6;
  cursor: default;
}


.shadow-box {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #010101;
  opacity: 0.5;
  z-index: 2;

  pointer-events: none;
}

.mark{
  position: relative;

  font-size: 8em;
  z-index: 3;
  pointer-events: none;
}
</style>
