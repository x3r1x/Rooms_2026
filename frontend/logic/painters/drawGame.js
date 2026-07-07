import {drawBackground} from "./background.js";

export function drawGame(canvas, context, state) {
    //TODO: написать игру =)
    drawBackground(canvas, context);

    //Отладочный код для перемещения квадратика
    context.fillStyle = "#FFFFFF";
    context.fillRect(state.square.x, state.square.y, 20, 20);
}