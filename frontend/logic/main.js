import {updateGame} from "./controller/engine/updateGame.js";
import {drawGame} from "./view/painters/drawGame.js";
import {initLastState, lastState} from "./model/game/gameModel.js";
import {initListeners} from "./controller/engine/listeners.js";
import {initSocket} from "./model/di/server.js";
import {initSprites} from "./view/sprites.js";

export const canvas = document.getElementById("canvas");
export let socket = null;
const context = canvas.getContext('2d');

function loadGame() {
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;

    initListeners(canvas);
    socket = initSocket()
    initLastState(performance.now(), crypto.randomUUID());
    initSprites();

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