import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";

export function drawGame(canvas, context, state) {
    //TODO: написать игру =)
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawDot(canvas, context, state);

    context.fill();
}

function drawDot(canvas, context, state) {
    context.fillStyle = "red";
    context.fillRect(state.dot.x, state.dot.y, 5, 5);
}