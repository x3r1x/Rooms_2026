import {drawBackground} from "./background.js";
import {drawBullets} from "./drawBullets.js";

export function drawGame(canvas, context, state) {
    //TODO: написать игру =)
    drawBackground(canvas, context);
    drawBullets(canvas, context, state);
    drawDot(canvas, context, state);

    context.fill();

    //Отладочный код для перемещения квадратика
    context.fillStyle = "#FFFFFF";
    context.fillRect(state.square.x, state.square.y, 20, 20);
}