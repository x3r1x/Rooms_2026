import {lastState} from "../model/gameModel.js";
import {createBullet} from "../factory/createBullet.js";

export const keys = {};

export function initListeners(canvas) {
    window.addEventListener('keydown', function(event) {
        keys[event.key.toLowerCase()] = true;
        console.log("key " + event.key);
    });

    window.addEventListener('keyup', function(event) {
        keys[event.key.toLowerCase()] = false;
    });

    canvas.addEventListener('click', function (event) {
        const canvasRect = canvas.getBoundingClientRect();

        const x = event.clientX - canvasRect.left;
        const y = event.clientY - canvasRect.top;
        const direction = Math.atan2(y - lastState.player.y, x - lastState.player.x);

        //FIXME: those numbers are that magic that even Harry Potter himself is shaking
        createBullet(lastState, direction, lastState.player.x - 5, lastState.player.y - 20)
    });
}

