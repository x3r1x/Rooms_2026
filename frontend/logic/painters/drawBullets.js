import {bulletDrawProperties} from "../model/gameModel.js";

export function drawBullets(canvas, context, state) {
    state.bullets.forEach(function (bullet) {
        context.save();
        context.translate(bullet.x + bulletDrawProperties.x / 2, bullet.y + bulletDrawProperties.y / 2);
        context.rotate(bullet.direction + Math.PI / 2);
        context.fillStyle = bulletDrawProperties.color;
        context.fillRect(-bulletDrawProperties.x / 2, -bulletDrawProperties.y / 2, bulletDrawProperties.x, bulletDrawProperties.y);
        context.restore();
    })
}

/*

interface bullet
{
    x: number,
    y: number,
    direction: number - в радианах, от -PI до PI
}

 */