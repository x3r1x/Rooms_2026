import {canvas} from "../main.js";
import {bulletDrawProperties} from "../model/gameModel.js";

export function isBulletOutOfBounds(bullet) {
    const didPassLeft = bullet.x + bulletDrawProperties.x < 0;
    const didPassRight = bullet.x > canvas.width;
    const didPassTop = bullet.y + bulletDrawProperties.y < 0;
    const didPassBottom = bullet.y > canvas.height;

    return didPassLeft || didPassRight || didPassTop || didPassBottom;
}