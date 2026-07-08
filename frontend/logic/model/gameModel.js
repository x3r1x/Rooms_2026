import {GAME_CONSTANTS} from "../gameConstants.js";

export const lastState = {
    lastTime: 0,
    player: {
        x: GAME_CONSTANTS.PLAYER_START_X,
        y: GAME_CONSTANTS.PLAYER_START_Y,
    },
    bullets: []
};