import {GAME_CONSTANTS} from "../gameConstants.js";

export function updatePlayer(direction, elapsedTime, state) {
    state.square.x += direction.x * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
    state.square.y += direction.y * GAME_CONSTANTS.PLAYER_SPEED * elapsedTime;
}