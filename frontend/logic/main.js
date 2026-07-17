import {updateGame} from "./model/engine/updateGame.js";
import {drawGame} from "./view/painters/drawGame.js";
import {currentState, initLastState} from "./model/storage/states.js";
import {initListeners, resizeCanvas} from "./controller/listeners.js";
import {getSocket} from "./model/di/webSocket/server.js";
import {initSprites} from "./view/sprites.js";
import {initMap} from "./view/maps.js";
import {assemblyRoom, parserMapData, parserTileInfo} from "./model/storage/layers.js";

export const canvas = document.getElementById("canvas");
export let socket = null;
const context = canvas.getContext('2d');

async function loadData() {
    const response = await fetch('../frontend/assets/tile/allRoom.json');
    return await response.json();
}

async function loadTileInfo() {
    const response = await fetch('../frontend/assets/tile/tileInfo.json');
    return await response.json();
}

function loadImages(){

}

async function loadGame() {
    resizeCanvas(canvas)
    initListeners(canvas);
    socket = getSocket()
    //TODO: id бы создавать на сервере...
    initLastState(performance.now(), crypto.randomUUID());
    // функция загрузки всех ресурсов на будущее
    const mapData = await loadData();
    const tileInfo = await loadTileInfo();
    parserMapData(mapData);
    parserTileInfo(tileInfo);
    initMap();
    initSprites();
    assemblyRoom();
    loadImages();

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