import {GAME_CONSTANTS} from "./gameConstants.js";

export const lobbyState = {
    clientId: null,
    players: {},
    countdown: null
}

export let gameState = null;

export function initGameState(playerId) {
    gameState = {
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
            id: playerId,
            spriteIndex: 0,
            hp: null,
            rebornTime: null
        },
        enemies: [],
        bullets: [],

        map: [],
        mapWall: [],
        mapObject: [],
    }
}

export function startGameState(dateNow, playerId) {
    gameState.lastTime = dateNow;

    gameState.player.mousePosition.x = 0;
    gameState.player.mousePosition.y = 0;

    gameState.player.id = playerId;
    gameState.player.direction = 0;
    gameState.previousVisualDirection = 0;
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