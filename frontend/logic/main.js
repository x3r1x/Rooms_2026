import {lastState} from "./model/gameModel.js";
import {updateGame} from "./engine/updateGame.js";
import {drawGame} from "./painters/drawGame.js";
import {initListeners} from "./engine/listeners.js";

export const canvas = document.getElementById("canvas");
const context = canvas.getContext('2d');

function loadGame() {
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;

    initListeners(canvas);

    lastState.lastTime = performance.now();
    requestAnimationFrame(startGameLoop);
}

function startGameLoop() {
    const currentTime = performance.now();
    const elapsedTime = currentTime - lastState.lastTime;

    updateGame(elapsedTime, lastState);
    drawGame(canvas, context, lastState);

    lastState.lastTime = currentTime;
    requestAnimationFrame(startGameLoop);
}

loadGame();