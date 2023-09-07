import * as err from './myError.js';
import SubGrid from './subGrid.js';

export default class MasterGrid extends SubGrid{
    constructor(){
        super();
        this.grid = [0, 0, 0, 0, 0, 0, 0, 0, 0];
        this.gridSize = 0;
    }

    /*
    *   set the specified box to the specified player
    *   @param {int} box: 0-8
    *   @param {int} player: 1 or 2
    *   @return {int} 0 if the game is not over, 1 if the game is over
    *   @return {string} error message if there is an error
    */
    setBox(box, player){
        if (box < 0 || box > 8) return [null, err.id_box_error];
        if (player !== 1 && player !== 2) return [null, err.id_player_error];
        if (this.grid[box] !== 0) return [null, err.box_full_error ];

        this.grid[box] = player;
        this.gridSize++;
        
        if (this.checkWin()) return [1, null];

        if (this.gridSize === 9) return [2, null];

        return [0, null];
    }


    resetGrid(){
        this.grid = [0, 0, 0, 0, 0, 0, 0, 0, 0];
        this.gridSize = 0;
    }
} 