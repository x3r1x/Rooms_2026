import {GAME_CONSTANTS} from "../../model/game/gameConstants.js";

export function drawBackground(canvas, context) {
    context.fillStyle = '#1a1a2e';
    context.fillRect(GAME_CONSTANTS.CANVAS_START, GAME_CONSTANTS.CANVAS_START, canvas.width, canvas.height);
}