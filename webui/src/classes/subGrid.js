import * as err from './myError.js';

export default class SubGrid {
    /*
    *   @param {int} id_grid: 0-8
    *   grid: array of 9 elements
    *   gridSize: number of elements in grid
    *   winner: null, 1, 2
    * 
    */
    constructor(id_grid) {
        this.id = id_grid;
        this.grid = [0, 0, 0, 0, 0, 0, 0, 0, 0];
        this.gridSize = 0;
        this.winner = null;
    }

    /*
    *   set the specified box to the specified player
    *   @param {int} box: 0-8
    *   @param {int} player: 1 or 2
    *   @return {int} 0 if the game is not over, 1 if the game is over
    *   @return {string} error message if there is an error
    */
    setBox(box, player) {
        if (box < 0 || box > 8) return [null, err.id_box_error];
        if (player !== 1 && player !== 2) return [null, err.id_player_error];
        if (this.grid[box] !== 0) return [null, err.box_full_error];

        this.grid[box] = player;
        this.gridSize++;

        // if the match is a draw, the winner is the player who played last
        if (this.checkWin() || this.gridSize === 9) {
            this.winner = player;
            return [1, null];
        }


        return [0, null];
    }

    checkWin() {
        if (this.gridSize < 3) return false;

        // check row
        if (this.grid[0] === this.grid[1] && this.grid[1] === this.grid[2] && this.grid[0] !== 0) return true;
        if (this.grid[3] === this.grid[4] && this.grid[4] === this.grid[5] && this.grid[3] !== 0) return true;
        if (this.grid[6] === this.grid[7] && this.grid[7] === this.grid[8] && this.grid[6] !== 0) return true;

        // check column
        if (this.grid[0] === this.grid[3] && this.grid[3] === this.grid[6] && this.grid[0] !== 0) return true;
        if (this.grid[1] === this.grid[4] && this.grid[4] === this.grid[7] && this.grid[1] !== 0) return true;
        if (this.grid[2] === this.grid[5] && this.grid[5] === this.grid[8] && this.grid[2] !== 0) return true;

        // check diagonal
        if (this.grid[0] === this.grid[4] && this.grid[4] === this.grid[8] && this.grid[0] !== 0) return true;
        if (this.grid[2] === this.grid[4] && this.grid[4] === this.grid[6] && this.grid[2] !== 0) return true;

        return false;
    }

    resetGrid() {
        this.grid = [0, 0, 0, 0, 0, 0, 0, 0, 0];
        this.gridSize = 0;
    }
}