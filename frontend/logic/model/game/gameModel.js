import {GAME_CONSTANTS} from "./gameConstants.js";

export const currentState = {
    lastTime: null,
    mousePosition: {
        x: null,
        y: null
    },
    player: {
        x: GAME_CONSTANTS.PLAYER_START_X,
        y: GAME_CONSTANTS.PLAYER_START_Y,
        direction: null,
        bullets: {},
        id: null
    },
    map: [],
    mapCollisian: [],
    enemies: [],
};

export const lastState = {
    player: {
        x: GAME_CONSTANTS.PLAYER_START_X,
        y: GAME_CONSTANTS.PLAYER_START_Y,
        direction: null,
        bullets: {},
        id: null
    },
    enemies: [],
}

export function initLastState(dateNow, playerId) {
    currentState.lastTime = dateNow;

    currentState.mousePosition.x = 0;
    currentState.mousePosition.y = 0;

    currentState.player.id = playerId;
    currentState.player.direction = 0;
}