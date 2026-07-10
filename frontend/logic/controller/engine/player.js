import {GAME_CONSTANTS} from "../../model/game/gameConstants.js";

export function updatePlayer(direction, elapsedTime, state) {
    state.player.x += direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    state.player.y += direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    updatePlayerDirection(state);
}

export function updatePlayerDirection(state) {
    state.player.direction = Math.atan2(state.mousePosition.y - state.player.y, state.mousePosition.x - state.player.x);
}


export function updateEnemies(playersArray, state) {
    state.enemies = [];
    playersArray.forEach(function (enemy) {
        if (state.player.id !== enemy.id) {
            state.enemies.push(enemy);
        }
    })
}