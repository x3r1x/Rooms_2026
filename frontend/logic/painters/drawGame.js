import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawSquare(canvas, context, state);

    context.fill();
}

function drawSquare(canvas, context, state) {
    context.fillStyle = "#FFFFFF";
    context.fillRect(state.square.x, state.square.y, 20, 20);
}