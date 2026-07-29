import {lerp} from "./interpolation.js";
import {GAME_CONSTANTS} from "../storage/gameConstants.js";

export function handleBullets(bullets, snaps, extrapolationTime, lerpCoefficient, didRoomChange) {
    clearBullets(bullets, snaps);

    snaps.snapA.b.forEach((bulletStart) => {
        if (bulletStart.id in bullets  && !didRoomChange) {
            const bulletEnd = snaps.snapB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id);

            if (bulletEnd) {
                lerpBullet(bullets[bulletStart.id], bulletStart, bulletEnd, lerpCoefficient);
            }

            if (!bulletEnd && extrapolationTime < GAME_CONSTANTS.MAX_EXTRAPOLATION_TIME) {
                extrapolateBullet(bullets[bulletStart.id], extrapolationTime, bulletStart);
            }
        }
    })
}

function clearBullets(bullets, snaps) {
    for (const id in bullets) {
        const inSnapA = snaps.snapA.b.some((bullet) => bullet.id === id);
        const inSnapB = snaps.snapB.b.some((bullet) => bullet.id === id);

        if (!inSnapA && !inSnapB) {
            delete bullets[id];
        }
    }
}

function lerpBullet(modelBullet, bulletStart, bulletEnd, lerpCoefficient) {
    modelBullet.x = lerp(bulletStart.x, bulletEnd.x, lerpCoefficient);
    modelBullet.y = lerp(bulletStart.y, bulletEnd.y, lerpCoefficient);
}

function extrapolateBullet(modelBullet, extrapolationTime, startBullet) {
    modelBullet.x = startBullet.x + extrapolationTime * Math.cos(startBullet.direction) * GAME_CONSTANTS.BULLET_SPEED;
    modelBullet.y = startBullet.y + extrapolationTime * Math.sin(startBullet.direction) * GAME_CONSTANTS.BULLET_SPEED;
}