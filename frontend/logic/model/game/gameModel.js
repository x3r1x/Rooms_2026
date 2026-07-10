import {GAME_CONSTANTS} from "./gameConstants.js";

export const lastState = {
    lastTime: null,
    player: {
        x: GAME_CONSTANTS.PLAYER_START_X,
        y: GAME_CONSTANTS.PLAYER_START_Y,
        direction: null,
        id: null
    },
    map: [],
    mapCollisian: [],
    enemies: [],
    bullets: []
};

export function initLastState(dateNow, playerId) {
    lastState.lastTime = dateNow;
    lastState.player.id = playerId;
    lastState.player.direction = 0;
}