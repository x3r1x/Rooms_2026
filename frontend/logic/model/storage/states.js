import {GAME_CONSTANTS} from "./gameConstants.js";

export const currentState = {
    lastTime: null,

    player: {
        x: GAME_CONSTANTS.PLAYER_START_X,
        y: GAME_CONSTANTS.PLAYER_START_Y,
        mousePosition: {
            x: 0,
            y: 0
        },
        direction: null,
        movementDirection: {
            x: 0,
            y: 0
        },
        didShoot: false,
        id: null,
        spriteIndex: 0,
    },
    enemies: [],
    bullets: [],

    map: [],
    mapWall: [],
    mapObject: [],
};

export function initLastState(dateNow, playerId) {
    currentState.lastTime = dateNow;

    currentState.player.mousePosition.x = 0;
    currentState.player.mousePosition.y = 0;

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