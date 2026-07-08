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
    context.save();
    context.translate(state.player.x, state.player.y)
    context.rotate(state.player.direction + Math.PI / 2);
    context.fillStyle = "#FFFFFF";
    context.fillRect(-GAME_CONSTANTS.PLAYER_VISUAL_SIZE / 2, -GAME_CONSTANTS.PLAYER_VISUAL_SIZE / 2, GAME_CONSTANTS.PLAYER_VISUAL_SIZE, GAME_CONSTANTS.PLAYER_VISUAL_SIZE);
    context.restore();
}