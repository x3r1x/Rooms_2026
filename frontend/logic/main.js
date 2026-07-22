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

if (!crypto.randomUUID) {
    crypto.randomUUID = function() {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
            const r = crypto.getRandomValues(new Uint8Array(1))[0] % 16 | 0;
            const v = c === 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    };
    console.log('✅ Polyfill для crypto.randomUUID() установлен');
}

async function loadData() {
    const response = await fetch('./assets/tile/allRoom.json');
    return await response.json();
}

async function loadTileInfo() {
    const response = await fetch('./assets/tile/tileInfo.json');
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