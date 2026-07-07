import {lastState} from "./model/gameModel.js";
import {updateGame} from "./engine/updateGame.js";
import {drawGame} from "./painters/drawGame.js";

function loadGame() {
    lastState.lastTime = Date.now();

    startGameLoop();
}

function startGameLoop() {
    const canvas = document.getElementById("canvas");
    const context = canvas.getContext('2d');
    const currentTime = Date.now();
    const elapsedTime = currentTime - lastState.lastTime;

    updateGame(elapsedTime, lastState);
    drawGame(canvas, context, lastState);

    lastState.lastTime = currentTime;
    requestAnimationFrame(startGameLoop);
}

loadGame();