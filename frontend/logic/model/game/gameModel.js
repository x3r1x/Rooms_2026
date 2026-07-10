import {GAME_CONSTANTS} from "./gameConstants.js";

export const lastState = {
    lastTime: null,
    mousePosition: {
        x: null,
        y: null
    },
    player: {
        x: GAME_CONSTANTS.PLAYER_START_X,
        y: GAME_CONSTANTS.PLAYER_START_Y,
        direction: null,
        bullets: [],
        id: null
    },
    map: [],
    mapCollisian: [],
    enemies: [],
};

export function initLastState(dateNow, playerId) {
    lastState.lastTime = dateNow;

    lastState.mousePosition.x = 0;
    lastState.mousePosition.y = 0;

    lastState.player.id = playerId;
    lastState.player.direction = 0;
}