import {initGameListeners, resizeCanvas} from "../../controller/gameListeners.js";
import {assemblyRoom, parseMapData} from "./storage/layers.js";
import {initMap} from "../../view/game/maps.js";
import {initSprites} from "../../view/game/sprites.js";
import {gameState} from "./storage/gameState.js";
import {updateGame} from "./engine/updateGame.js";
import {drawGame} from "../../view/game/painters/drawGame.js";
import {socket} from "../app/appState.js";

export const canvas = document.getElementById("canvas");
const context = canvas.getContext('2d');

export async function loadGame() {
    resizeCanvas(canvas)
    initGameListeners(canvas);
    const mapData = await loadData();
    const tileInfo = await loadTileInfo();
    parseMapData(mapData);
    initMap();
    initSprites();
    assemblyRoom();
}

async function loadData() {
    const response = await fetch('./assets/tile/allRoom.json');
    return await response.json();
}

async function loadTileInfo() {
    const response = await fetch('./assets/tile/tileInfo.json');
    return await response.json();
}

export function startGameLoop() {
    const currentTime = performance.now();
    const elapsedTime = currentTime - gameState.lastTime;

    updateGame(elapsedTime, gameState, socket);
    drawGame(canvas, context, gameState);

    gameState.lastTime = currentTime;
    requestAnimationFrame(startGameLoop);
}