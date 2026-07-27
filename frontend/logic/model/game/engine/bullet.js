import {lerp} from "../storage/interpolation.js";

export function getNewBulletsMap(dt, bullets, snaps) {
    const newBulletsMap = {}

    snaps.stateA.b.forEach((bulletStart) => {
        if (bulletStart.id in bullets && snaps.stateB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id)) {
            const bulletEnd = snaps.stateB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id);

            newBulletsMap[bulletStart.id] = bullets[bulletStart.id];
            newBulletsMap[bulletStart.id].x = lerp(bulletStart.x, bulletEnd.x, dt);
            newBulletsMap[bulletStart.id].y = lerp(bulletStart.y, bulletEnd.y, dt);
        }
    })

    return newBulletsMap;
}