import {GAME_CONSTANTS} from "../storage/gameConstants.js";
import {canMoveTo} from "./collision.js";

export function updatePlayer(direction, elapsedTime, player) {
    updateVisualDirection(player);

    const nextX = player.x + direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    const nextY = player.y + direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;

    if (canMoveTo({x: nextX, y: nextY}, player)){
        player.x = nextX;
        player.y = nextY;
    }
}

export function updateVisualDirection(player) {
    player.direction = Math.atan2(player.mousePosition.y - player.y, player.mousePosition.x - player.x);
}

export function updateEnemies(elapsedTime, enemies) {
    enemies.forEach(function (enemy) {
        enemy.x += enemy.movementDirection.x * elapsedTime * GAME_CONSTANTS.PLAYER_SPEED;
        enemy.y += enemy.movementDirection.y * elapsedTime * GAME_CONSTANTS.PLAYER_SPEED;
    })
}