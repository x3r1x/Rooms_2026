import {GAME_CONSTANTS} from "../storage/gameConstants.js";
import {lerp} from "./interpolation.js";

export function updatePlayer(direction, elapsedTime, player) {
    updateVisualDirection(player);

    const playerSpeed = getPlayerSpeed(player.pc);

    const nextX = player.x + direction.x * playerSpeed * elapsedTime;
    const nextY = player.y + direction.y * playerSpeed * elapsedTime;

    // if (canMoveTo({x: nextX, y: nextY}, player)) {
    player.x = nextX;
    player.y = nextY;
    // }
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

export function getPlayerSpeed(playerType) {
    switch (playerType) {
        case GAME_CONSTANTS.PLAYER_TYPES.PISTOL:
            return GAME_CONSTANTS.PISTOL_PLAYER_SPEED;
        case GAME_CONSTANTS.PLAYER_TYPES.SHOTGUN:
            return GAME_CONSTANTS.SHOTGUN_PLAYER_SPEED;
        case GAME_CONSTANTS.PLAYER_TYPES.ROCKET_LAUNCHER:
            return GAME_CONSTANTS.ROCKET_PLAYER_SPEED;
    }

    return 0;
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
    modelEnemy.direction = lerp(enemyStart.a, enemyEnd.a, lerpCoefficient);
    modelEnemy.hp = lerp(enemyStart.h, enemyEnd.h, lerpCoefficient);
    modelEnemy.ps = lerp(enemyStart.ps, enemyEnd.ps, lerpCoefficient);
}

function extrapolateEnemy(modelEnemy, extrapolationTime, startEnemy) {
    const enemySpeed = getPlayerSpeed(modelEnemy.pc);

    modelEnemy.x = startEnemy.x + extrapolationTime * startEnemy.mx * enemySpeed;
    modelEnemy.y = startEnemy.y + extrapolationTime * startEnemy.my * enemySpeed;
}