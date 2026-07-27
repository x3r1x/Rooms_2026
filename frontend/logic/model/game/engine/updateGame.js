import {keys} from '../../../controller/gameListeners.js';
import {handleEnemies, updatePlayer} from "./players.js";
import {handleBullets} from "./bullet.js";
import {getNeighbouringSnapshots, getClientTime} from "../storage/interpolation.js";
import {GAME_CONSTANTS} from "../storage/gameConstants.js";
import {checkAndSpawnNewBulletEffects, checkAndSpawnExplosions} from "../../../view/game/effects/effectsManager.js";

export function updateGame(elapsedTime, state) {
    updateMovementDirection(state.player.movementDirection);
    updatePlayer(state.player.movementDirection, elapsedTime, state.player);

    const renderTime = getClientTime() - GAME_CONSTANTS.INTERPOLATION_DELAY;
    const neighbouredSnapshots = getNeighbouringSnapshots(renderTime);

    if (!neighbouredSnapshots.snapA) return;
    const lerpCoefficient = (renderTime - neighbouredSnapshots.snapA.t) / (neighbouredSnapshots.snapB.t - neighbouredSnapshots.snapA.t);
    const extrapolationTime = renderTime - neighbouredSnapshots.snapA.t;
    checkAndSpawnNewBulletEffects(neighbouredSnapshots.stateA, neighbouredSnapshots.stateB, state);
    checkAndSpawnExplosions(neighbouredSnapshots.stateA, neighbouredSnapshots.stateB);

    handleEnemies(state.enemies, neighbouredSnapshots, extrapolationTime, lerpCoefficient);
    handleBullets(state.bullets, neighbouredSnapshots, extrapolationTime, lerpCoefficient);
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