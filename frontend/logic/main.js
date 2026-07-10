import {updateGame} from "./controller/engine/updateGame.js";
import {drawGame} from "./view/painters/drawGame.js";
import {initLastState, lastState} from "./model/game/gameModel.js";
import {initListeners} from "./controller/engine/listeners.js";
import {initSocket} from "./model/di/server.js";
import {initSprites} from "./view/sprites.js";
import {initMap} from "./view/maps.js";

export const canvas = document.getElementById("canvas");
export let socket = null;
const context = canvas.getContext('2d');

async function loadData() {
    const response = await fetch('./logic/room1.json');
    return await response.json();
}

async function loadGame() {
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;

    const mapData = await loadData();
    initListeners(canvas);
    socket = initSocket()
    initLastState(performance.now(), crypto.randomUUID());
    lastState.map = mapData.layers[0].data;
    lastState.mapCollisian = mapData.layers[1].data;
    initMap();
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