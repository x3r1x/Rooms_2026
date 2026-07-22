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
        hpSpriteIndex: 0,
        hp: null,
        rebornTime: null
    },
    enemies: [],
    bullets: [],

    map: [],
    mapWall: [],
    mapObject: [],
};
