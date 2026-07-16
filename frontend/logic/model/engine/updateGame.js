import {keys} from '../../controller/listeners.js';
import {updateEnemies, updatePlayer} from "./players.js";
import {getUpdatedBullets} from "./bullet.js";

export function updateGame(elapsedTime, state) {
    updateMovementDirection(state.player.movementDirection);
    updatePlayer(state.player.movementDirection, elapsedTime, state.player);
    updateEnemies(elapsedTime, state.enemies);
    state.bullets = getUpdatedBullets(elapsedTime, state.bullets);
}

function updateMovementDirection(direction) {
    direction.x = 0;
    direction.y = 0;

    if (keys['w'] || keys['ц'] || keys['arrowup']) {
        direction.y = -1;
    }
    if (keys['s'] || keys['ы'] || keys['arrowdown']) {
        direction.y = 1;
    }
    if (keys['a'] || keys['ф'] || keys['arrowleft']) {
        direction.x = -1;
    }
    if (keys['d'] || keys['в'] || keys['arrowright']) {
        direction.x = 1;
    }

    const directionVectorLength = Math.sqrt(direction.x * direction.x + direction.y * direction.y);
    if (directionVectorLength !== 0) {
        direction.x /= directionVectorLength;
        direction.y /= directionVectorLength;
    }
}