import {GAME_CONSTANTS} from "../gameConstants.js";

export function updatePlayer(direction, elapsedTime, state) {
    state.player.x += direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    state.player.y += direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
}

export function updatePlayerDirection(newDirection, state) {
    state.player.direction = newDirection;
}

