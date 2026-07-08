import {lastState} from "../model/gameModel.js";
import {createBullet} from "../factory/createBullet.js";
import {updatePlayerDirection} from "./player.js";

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

        //FIXME: those numbers are that magic that even Harry Potter himself is shaking
        createBullet(lastState, direction, lastState.player.x - 5, lastState.player.y - 20)
    });

    canvas.addEventListener('mousemove', function (event) {
        const canvasRect = canvas.getBoundingClientRect();

        const x = event.clientX - canvasRect.left;
        const y = event.clientY - canvasRect.top;
        const direction = Math.atan2(y - lastState.player.y, x - lastState.player.x);

        updatePlayerDirection(direction, lastState);
    })
}