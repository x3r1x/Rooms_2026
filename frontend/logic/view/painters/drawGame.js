import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";
import {drawPlayers} from "./drawPlayers.js";

export function drawGame(canvas, context, state) {
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawPlayers(canvas, context, state);

    context.fill();
}