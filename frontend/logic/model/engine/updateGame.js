import {keys} from '../../controller/listeners.js';
import {updateEnemies, updatePlayer} from "./players.js";

export function updateGame(elapsedTime, state) {
    updateMovementDirection(state.movementDirection);
    updatePlayer(state.movementDirection, elapsedTime, state);
    updateEnemies(elapsedTime, state.enemies);

    state.movementDirection = {
        x: 0,
        y: 0
    }

}

function updateMovementDirection(direction) {
    if (keys['w'] || keys['ц'] || keys['arrowup']) {
        direction.y -= 1
    }
    if (keys['s'] || keys['ы'] || keys['arrowdown']) {
        direction.y += 1
    }
    if (keys['a'] || keys['ф'] || keys['arrowleft']) {
        direction.x -= 1
    }
    if (keys['d'] || keys['в'] || keys['arrowright']) {
        direction.x += 1
    }
}