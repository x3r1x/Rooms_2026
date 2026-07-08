import {lastState} from "../model/gameModel.js";
import {createBullet} from "../factory/createBullet.js";
import {updatePlayerDirection} from "./player.js";
import {GAME_CONSTANTS} from "../gameConstants.js";

export const keys = {};

export function initListeners(canvas) {
    initWindowListeners()
    initCanvasListeners(canvas);
}

function initWindowListeners() {
    window.addEventListener('keydown', function(event) {
        keys[event.key.toLowerCase()] = true;
    });

    window.addEventListener('keyup', function(event) {
        keys[event.key.toLowerCase()] = false;
    });
}

function initCanvasListeners(canvas) {
    canvas.addEventListener('click', function (event) {
        const canvasRect = canvas.getBoundingClientRect();

        const x = event.clientX - canvasRect.left;
        const y = event.clientY - canvasRect.top;
        const direction = Math.atan2(y - lastState.player.y, x - lastState.player.x);

        const localX = GAME_CONSTANTS.PLAYER_VISUAL_SIZE / 2;
        const localY = GAME_CONSTANTS.PLAYER_VISUAL_SIZE / 2 - (GAME_CONSTANTS.BULLET_WIDTH+(GAME_CONSTANTS.PLAYER_VISUAL_SIZE*0.1));

        const rotatedX = localX * Math.cos(direction) - localY * Math.sin(direction);
        const rotatedY = localX * Math.sin(direction) + localY * Math.cos(direction);

        const bulletStartX = lastState.player.x + rotatedX;
        const bulletStartY = lastState.player.y + rotatedY;
        createBullet(lastState, direction, bulletStartX, bulletStartY);
    });

    canvas.addEventListener('mousemove', function (event) {
        const canvasRect = canvas.getBoundingClientRect();

        const x = event.clientX - canvasRect.left;
        const y = event.clientY - canvasRect.top;
        const direction = Math.atan2(y - lastState.player.y, x - lastState.player.x);

        updatePlayerDirection(direction, lastState);
    })
}