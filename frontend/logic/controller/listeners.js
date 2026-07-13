import {createBullet} from "../model/factory/createBullet.js";
import {currentState} from "../model/storage/states.js";
import {GAME_CONSTANTS} from "../model/storage/gameConstants.js";

export const keys = {};

export function initListeners(canvas) {
    initWindowListeners()
    initCanvasListeners(canvas);
}

function initWindowListeners() {
    window.addEventListener('keydown', function (event) {
        keys[event.key.toLowerCase()] = true;
    });

    window.addEventListener('keyup', function (event) {
        keys[event.key.toLowerCase()] = false;
    });
}

function initCanvasListeners(canvas) {
    canvas.addEventListener('click', function (event) {
        const canvasRect = canvas.getBoundingClientRect();

        const x = event.clientX - canvasRect.left;
        const y = event.clientY - canvasRect.top;
        const direction = Math.atan2(y - currentState.player.y, x - currentState.player.x);

        const localX = GAME_CONSTANTS.PLAYER_VISUAL_SIZE / 2;
        const localY = GAME_CONSTANTS.PLAYER_VISUAL_SIZE / 2 - (GAME_CONSTANTS.BULLET_WIDTH + (GAME_CONSTANTS.PLAYER_VISUAL_SIZE * 0.1));

        const rotatedX = localX * Math.cos(direction) - localY * Math.sin(direction);
        const rotatedY = localX * Math.sin(direction) + localY * Math.cos(direction);

        const bulletStartX = currentState.player.x + rotatedX;
        const bulletStartY = currentState.player.y + rotatedY;

        createBullet(currentState, direction, bulletStartX, bulletStartY);
    });

    canvas.addEventListener('mousemove', function (event) {
        const canvasRect = canvas.getBoundingClientRect();

        const x = event.clientX - canvasRect.left;
        const y = event.clientY - canvasRect.top;

        currentState.mousePosition.x = x;
        currentState.mousePosition.y = y;
    })
}