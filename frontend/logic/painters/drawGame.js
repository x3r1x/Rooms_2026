import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";
import {GAME_CONSTANTS} from "../gameConstants.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawSquare(canvas, context, state);

    context.fill();
}

function drawSquare(canvas, context, state) {
    context.fillStyle = "#FFFFFF";
    context.fillRect(state.square.x, state.square.y, GAME_CONSTANTS.SQUARE_VISUAL_SIZE, GAME_CONSTANTS.SQUARE_VISUAL_SIZE);
}