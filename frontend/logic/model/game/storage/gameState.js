export let gameState = null;
export let gameNicknames = {};
export let finalStatistics = null;

export function initGameState(playerId) {
    gameState = {
        lastTime: null,
        player: {
            x: 0,
            y: 0,
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
            roomId: null,
            spriteIndex: 0,
            hpSpriteIndex: 0,
            hp: null,
            rebornTime: null,
            pc: null,
        },
        enemies: {},
        bullets: {},
    }
}

export function startGameState(dateNow) {
    gameState.lastTime = dateNow;

    gameState.player.mousePosition.x = 0;
    gameState.player.mousePosition.y = 0;

    gameState.player.direction = 0;
    gameState.player.roomId = 0;

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