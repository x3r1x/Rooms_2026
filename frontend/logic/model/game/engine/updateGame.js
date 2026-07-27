import {keys} from '../../../controller/gameListeners.js';
import {extrapolateEnemies, lerpEnemies, updatePlayer} from "./players.js";
import {extrapolateNewBulletsMap, lerpNewBulletsMap} from "./bullet.js";
import {getNeighbouringSnapshots, getClientTime} from "../storage/interpolation.js";
import {GAME_CONSTANTS} from "../storage/gameConstants.js";

export function updateGame(elapsedTime, state) {
    updateMovementDirection(state.player.movementDirection);
    updatePlayer(state.player.movementDirection, elapsedTime, state.player);

    const renderTime = getClientTime() - GAME_CONSTANTS.INTERPOLATION_DELAY;
    const neighbouredSnapshots = getNeighbouringSnapshots(renderTime);

    if (neighbouredSnapshots.snapA && neighbouredSnapshots.snapB) {
        const lerpCoefficient = (renderTime - neighbouredSnapshots.snapA.t) / (neighbouredSnapshots.snapB.t - neighbouredSnapshots.snapA.t);
        lerpEnemies(lerpCoefficient, state.enemies, neighbouredSnapshots);
        state.bullets = lerpNewBulletsMap(lerpCoefficient, state.bullets, neighbouredSnapshots);
    } else if (neighbouredSnapshots.snapA && !neighbouredSnapshots.snapB
                && renderTime - neighbouredSnapshots.snapA.t <= GAME_CONSTANTS.MAX_EXTRAPOLATION_TIME)
    {
        const extrapolationTime = renderTime - neighbouredSnapshots.snapA.t;
        extrapolateEnemies(extrapolationTime, state.enemies, neighbouredSnapshots.snapA);
        state.bullets = extrapolateNewBulletsMap(extrapolationTime, state.bullets, neighbouredSnapshots.snapA);
    } else {
        console.log("Слишком огромная потеря пакетов, превышен лимит экстраполяции!");
    }
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