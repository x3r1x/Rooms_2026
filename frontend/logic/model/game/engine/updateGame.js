import {keys} from '../../../controller/gameListeners.js';
import {updateEnemies, updatePlayer} from "./players.js";
import {updateBullets} from "./bullet.js";
import {getNeighbouringSnapshots, getClientTime} from "../storage/interpolation.js";
import {GAME_CONSTANTS} from "../storage/gameConstants.js";

import { spawnEffect } from "../../../view/game/effects/effectsManager.js";

export function updateGame(elapsedTime, state) {
    updateMovementDirection(state.player.movementDirection);
    updatePlayer(state.player.movementDirection, elapsedTime, state.player);

    const renderTime = getClientTime() - GAME_CONSTANTS.INTERPOLATION_DELAY;
    const neighbouredSnapshots = getNeighbouringSnapshots(renderTime);

    if (!neighbouredSnapshots.stateA || !neighbouredSnapshots.stateB) return;

    checkAndSpawnNewBulletEffects(neighbouredSnapshots.stateA, neighbouredSnapshots.stateB, state);

    const dt = (renderTime - neighbouredSnapshots.stateA.t) / (neighbouredSnapshots.stateB.t - neighbouredSnapshots.stateA.t);
    updateEnemies(dt, state.enemies, neighbouredSnapshots);
    updateBullets(dt, state.bullets, neighbouredSnapshots);
}

function checkAndSpawnNewBulletEffects(stateA, stateB) {
    if (!stateA || !stateB || !stateA.b || !stateB.b) return;

    for (let i = 0; i < stateB.b.length; i++) {
        const bulletB = stateB.b[i];
        let existedBefore = false;

        for (let j = 0; j < stateA.b.length; j++) {
            if (stateA.b[j].id === bulletB.id) {
                existedBefore = true;
                break;
            }
        }

        if (!existedBefore) {
            spawnEffect(
                bulletB.x,
                bulletB.y,
                "MUZZLE_FLASH",
                bulletB.a
            );
        }
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