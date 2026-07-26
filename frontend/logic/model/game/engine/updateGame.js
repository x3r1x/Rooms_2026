import {keys} from '../../../controller/gameListeners.js';
import {updateEnemies, updatePlayer} from "./players.js";
import {updateBullets} from "./bullet.js";
import {getNeighbouringSnapshots, getClientTime} from "../storage/interpolation.js";
import {GAME_CONSTANTS} from "../storage/gameConstants.js";

export function updateGame(elapsedTime, state) {
    updateMovementDirection(state.player.movementDirection);
    updatePlayer(state.player.movementDirection, elapsedTime, state.player);

    const renderTime = getClientTime() - GAME_CONSTANTS.INTERPOLATION_DELAY;
    const neighbouredSnapshots = getNeighbouringSnapshots(renderTime);

    if (!neighbouredSnapshots.stateA || !neighbouredSnapshots.stateB) return

    const dt = (renderTime - neighbouredSnapshots.stateA.t) / (neighbouredSnapshots.stateB.t - neighbouredSnapshots.stateA.t);
    updateEnemies(dt, state.enemies, neighbouredSnapshots);
    updateBullets(dt, state.bullets, neighbouredSnapshots);
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