import {GAME_CONSTANTS} from "../storage/gameConstants.js";

export function updatePlayer(direction, elapsedTime, player) {
    player.x += direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    player.y += direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    updateVisualDirection(player);
}

export function updateVisualDirection(player) {
    console.log(player)
    player.direction = Math.atan2(player.mousePosition.y - player.y, player.mousePosition.x - player.x);
}

export function updateEnemies(elapsedTime, enemies) {
    enemies.forEach(function (enemy) {
        enemy.x += enemy.movementDirection.x * elapsedTime * GAME_CONSTANTS.PLAYER_SPEED;
        enemy.y += enemy.movementDirection.y * elapsedTime * GAME_CONSTANTS.PLAYER_SPEED;
    })
}