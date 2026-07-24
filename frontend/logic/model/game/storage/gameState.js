import {GAME_CONSTANTS} from "./gameConstants.js";

export let gameState = null;
export let gameNicknames = {};
export let finalStatistics = null;

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
            hpSpriteIndex: 0,
            hp: null,
            rebornTime: null,
        },
        enemies: [],
        bullets: [],

        map: [],
        mapWall: [],
        mapObject: [],
    }
}

export function startGameState(dateNow) {
    gameState.lastTime = dateNow;

    gameState.player.mousePosition.x = 0;
    gameState.player.mousePosition.y = 0;

    gameState.player.direction = 0;
    gameState.previousVisualDirection = 0;
}

export const gameMap= {}

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

export function setFinalStatistics(statistic) {
    finalStatistics = statistic;
    finalStatistics.sort((player1, player2) => player2.k - player1.k);
}

export function resetGameStateStorage() {
    gameState = null;
    gameNicknames = {};
    finalStatistics = null;
}