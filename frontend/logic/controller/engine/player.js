import {GAME_CONSTANTS} from "../../model/game/gameConstants.js";

export function updatePlayer(direction, elapsedTime, state) {
    state.player.x += direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    state.player.y += direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
}

export function updatePlayerDirection(newDirection, state) {
    state.player.direction = newDirection;
}


export function updateEnemies(playersArray, state) {
    state.enemies = [];
    playersArray.forEach(function (enemy) {
        if (playersArray.player.id !== enemy.id) {
            //TODO: убрать, когда на беке начнёт посылаться direction
            if (!("direction" in enemy)) {
                enemy.direction = 0;
            }

            state.enemies.push(enemy);
        }
    })
}