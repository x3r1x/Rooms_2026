import {canvas} from "../main.js";
import {lastState} from "../model/gameModel.js";
import {createBullet} from "../factory/createBullet.js";

canvas.addEventListener('click', function (event) {
    const canvasRect = canvas.getBoundingClientRect();

    const x = event.clientX - canvasRect.left;
    const y = event.clientY - canvasRect.top;
    const direction = Math.atan2(y - lastState.dot.y, x - lastState.dot.x);

    //FIXME: those numbers are that magic that even Harry Potter himself is shaking
    createBullet(lastState, direction, lastState.dot.x - 5, lastState.dot.y - 20)
});

export const keys = {};
window.addEventListener('keydown', function(event) {
    keys[event.key.toLowerCase()] = true;
    console.log("key " + event.key);
});
window.addEventListener('keyup', function(event) {
    keys[event.key.toLowerCase()] = false;
});