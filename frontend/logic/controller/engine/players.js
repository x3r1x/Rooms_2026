import {GAME_CONSTANTS} from "../../model/game/gameConstants.js";
import {updateBullets} from "./bullet.js";

export function updatePlayer(direction, elapsedTime, state) {
    state.player.x += direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    state.player.y += direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    updatePlayerDirection(state);
    state.player.bullets = updateBullets(elapsedTime, state.player.bullets)
}

export function updatePlayerDirection(state) {
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