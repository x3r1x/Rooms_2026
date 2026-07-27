import {lerp} from "../storage/interpolation.js";
import {GAME_CONSTANTS} from "../storage/gameConstants.js";

export function lerpNewBulletsMap(dt, bullets, snaps) {
    const newBulletsMap = {};

    snaps.snapA.b.forEach((bulletStart) => {
        if (bulletStart.id in bullets && snaps.snapB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id)) {
            const bulletEnd = snaps.snapB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id);

            newBulletsMap[bulletStart.id] = bullets[bulletStart.id];
            newBulletsMap[bulletStart.id].x = lerp(bulletStart.x, bulletEnd.x, dt);
            newBulletsMap[bulletStart.id].y = lerp(bulletStart.y, bulletEnd.y, dt);
        }
    })

    return newBulletsMap;
}

export function extrapolateNewBulletsMap(extrapolationTime, bullets, closestSnap) {
    const newBulletsMap = {};

    closestSnap.b.forEach((bullet) => {
        if (bullet.id in bullets) {
            newBulletsMap[bullet.id] = bullets[bullet.id];
            newBulletsMap[bullet.id].x = bullet.x + extrapolationTime * Math.cos(bullet.direction) * GAME_CONSTANTS.BULLET_SPEED;
            newBulletsMap[bullet.id].y = bullet.y + extrapolationTime * Math.sin(bullet.direction) * GAME_CONSTANTS.BULLET_SPEED;
        }
    })

    return newBulletsMap;
}