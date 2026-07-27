import {GAME_CONSTANTS} from "../storage/gameConstants.js";
import {canMoveTo} from "./collision.js";
import {lerp, lerpDirection} from "../storage/interpolation.js";

export function updatePlayer(direction, elapsedTime, player) {
    updateVisualDirection(player);

    const nextX = player.x + direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    const nextY = player.y + direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;

    if (canMoveTo({x: nextX, y: nextY}, player)) {
        player.x = nextX;
        player.y = nextY;
    }
}

export function updateVisualDirection(player) {
    player.direction = Math.atan2(player.mousePosition.y - player.y, player.mousePosition.x - player.x);
}

export function handleEnemies(enemies, snaps, extrapolationTime, lerpCoefficient) {
    clearEnemies(enemies, snaps);

    snaps.snapA.p.forEach((playerStart) => {
        if (playerStart.id in enemies) {
            const playerEnd = snaps.snapB.p.find((playerEnd) => playerEnd.id === playerStart.id);

            if (playerEnd) {
                lerpEnemy(enemies[playerStart.id], playerStart, playerEnd, lerpCoefficient);
            }

            if (!playerEnd && extrapolationTime < GAME_CONSTANTS.MAX_EXTRAPOLATION_TIME) {
                extrapolateEnemy(enemies[playerStart.id], extrapolationTime, playerStart);
            }
        }
    })
}

function clearEnemies(enemies, snaps) {
    for (const id in enemies) {
        const inSnapA = snaps.snapA.p.some((enemy) => enemy.id === id);
        const inSnapB = snaps.snapB.p.some((enemy) => enemy.id === id);

        if (!inSnapA && !inSnapB) {
            delete enemies[id];
        }
    }
}

function lerpEnemy(modelEnemy, enemyStart, enemyEnd, lerpCoefficient) {
    modelEnemy.x = lerp(enemyStart.x, enemyEnd.x, lerpCoefficient);
    modelEnemy.y = lerp(enemyStart.y, enemyEnd.y, lerpCoefficient);
    modelEnemy.direction = lerpDirection(enemyStart.a, enemyEnd.a, lerpCoefficient);
    modelEnemy.hp = lerp(enemyStart.h, enemyEnd.h, lerpCoefficient);
}

function extrapolateEnemy(modelEnemy, extrapolationTime, startEnemy) {
    modelEnemy.x = startEnemy.x + extrapolationTime * startEnemy.mx * GAME_CONSTANTS.PLAYER_SPEED;
    modelEnemy.y = startEnemy.y + extrapolationTime * startEnemy.my * GAME_CONSTANTS.PLAYER_SPEED;
}