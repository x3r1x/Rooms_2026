import {GAME_CONSTANTS} from "../storage/gameConstants.js";
import {updateBullets} from "./bullet.js";

export function updatePlayer(direction, elapsedTime, state) {
    const directionVectorLength = Math.sqrt(direction.x * direction.x + direction.y * direction.y);
    const normalizedDirection = {
        x: direction.x,
        y: direction.y
    }
    if (directionVectorLength !== 0) {
        normalizedDirection.x = normalizedDirection.x / directionVectorLength;
        normalizedDirection.y = normalizedDirection.y / directionVectorLength;
    }

    state.player.x += normalizedDirection.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    state.player.y += normalizedDirection.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    updateVisualDirection(state);
    state.player.bullets = updateBullets(elapsedTime, state.player.bullets)
}

export function updateVisualDirection(state) {
    state.player.direction = Math.atan2(state.mousePosition.y - state.player.y, state.mousePosition.x - state.player.x);
}

export function updateEnemies(elapsedTime, enemies) {
    for (const enemy of Object.values(enemies)) {
        enemy.x += enemy.direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
        enemy.y += enemy.direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
        enemy.bullets = updateBullets(elapsedTime, enemy.bullets);
    }
}

export function getPlayerFromModelById(state, id) {
    if (state.player.id === id) {
        return state.player;
    }

    if (id in state.enemies) {
        return state.enemies[id]
    }

    return null;
}