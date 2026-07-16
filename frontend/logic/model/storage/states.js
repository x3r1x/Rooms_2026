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
    enemies: [],

    movementDirection: {
        x: 0,
        y: 0
    },
    newBulletsDirection: [],

    map: [],
    mapWall: [],
    mapObject: [],
};

export const previousState = {
    player: {
        x: null,
        y: null,
        direction: null,
        bullets: {},
        id: null
    },
    enemies: []
}

export function initLastState(dateNow, playerId) {
    currentState.lastTime = dateNow;

    currentState.mousePosition.x = 0;
    currentState.mousePosition.y = 0;

    currentState.player.id = playerId;
    currentState.player.direction = 0;
    currentState.previousVisualDirection = 0;
}

export const room = {
    exits: [],
    floors: [],
    walls: [],
    object: [],
    exit: {top: false, left: false, down: false, right: false},
    flap: {top: false, left: false, down: false, right: false},
    collision: []
}

export const layersForRoom = {
    width: 0,
    height:0,
    tileSize: 0,
    floor: [],
    walls: [],
    exit: [],
    flap: [],
    objects: [],
    tilesInfo: [],
}