import {updateGame} from "./model/engine/updateGame.js";
import {drawGame} from "./view/painters/drawGame.js";
import {currentState, initLastState} from "./model/storage/states.js";
import {initListeners} from "./controller/listeners.js";
import {initSocket} from "./model/di/webSocket/server.js";
import {initSprites} from "./view/sprites.js";
import {initMap} from "./view/maps.js";
import {assemblyRoom, parserMapData} from "./model/storage/layers.js";

export const canvas = document.getElementById("canvas");
export let socket = null;
const context = canvas.getContext('2d');

async function loadData() {
    const response = await fetch('../frontend/assets/tile/allRoom.json');
    return await response.json();
}

async function loadGame() {
    canvas.width = 36*25;
    canvas.height = 36*21;
    initListeners(canvas);
    socket = initSocket()
    //TODO: id бы создавать на сервере...
    initLastState(performance.now(), crypto.randomUUID());
    // функция загрузки всех ресурсов на будущее
    const mapData = await loadData();
    parserMapData(mapData);
    initMap();
    initSprites();
    assemblyRoom();

    requestAnimationFrame(startGameLoop);
}

function startGameLoop() {
    const currentTime = performance.now();
    const elapsedTime = currentTime - currentState.lastTime;

    updateGame(elapsedTime, currentState, socket);
    drawGame(canvas, context, currentState);

    currentState.lastTime = currentTime;
    requestAnimationFrame(startGameLoop);
}

loadGame();