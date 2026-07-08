import {GAME_CONSTANTS} from "../gameConstants.js";

export const lastState = {
    lastTime: null,
    square: {
        x: GAME_CONSTANTS.SQUARE_START_X,
        y: GAME_CONSTANTS.SQUARE_START_Y,
        id: null
    },
    enemies: [],
    bullets: []
};

export function initLastState(dateNow, playerId) {
    lastState.lastTime = dateNow;
    lastState.square.id = playerId;
}