import {updateGame} from "./model/engine/updateGame.js";
import {drawGame} from "./view/painters/drawGame.js";
import {currentState, initLastState, room} from "./model/storage/states.js";
import {initListeners} from "./controller/listeners.js";
import {initSocket} from "./model/di/webSocket/server.js";
import {initSprites} from "./view/sprites.js";
import {initMap} from "./view/maps.js";
import {parserMapData, assemblyRoom} from "./model/game/layers.js";
export const canvas = document.getElementById("canvas");
export let socket = null;
const context = canvas.getContext('2d');

async function loadData() {
    const response = await fetch('../frontend/assets/tile/allRoom.json');
    return await response.json();
}

async function loadGame() {
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;
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