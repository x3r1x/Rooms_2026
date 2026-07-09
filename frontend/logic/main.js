import {lastState} from "./model/gameModel.js";
import {updateGame} from "./engine/updateGame.js";
import {drawGame} from "./painters/drawGame.js";
import {initListeners} from "./engine/listeners.js";

export const canvas = document.getElementById("canvas");
export let socket = null;
const context = canvas.getContext('2d');

function loadGame() {
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;

    initListeners(canvas);
    socket = initSocket()

    GAME_SPRITES.PLAYER_GOES.src = GAME_CONSTANTS.PLAYER_SKIN_BLUE;
    GAME_SPRITES.BULLET_FLIES.src = GAME_CONSTANTS.BULLET_SKIN;
    initLastState(performance.now(), crypto.randomUUID());
    requestAnimationFrame(startGameLoop);
}

function startGameLoop() {
    const currentTime = performance.now();
    const elapsedTime = currentTime - lastState.lastTime;

    updateGame(elapsedTime, lastState, socket);
    drawGame(canvas, context, lastState);

    lastState.lastTime = currentTime;
    requestAnimationFrame(startGameLoop);
}

loadGame();