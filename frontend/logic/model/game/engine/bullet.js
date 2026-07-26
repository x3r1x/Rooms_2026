import {lerp} from "../storage/interpolation.js";

export function updateBullets(dt, bullets, snaps) {
    snaps.stateA.b.forEach((bulletStart) => {
        if (bulletStart.id in bullets) {
            if (!(snaps.stateB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id))) {
                delete bullets[bulletStart.id];
            } else {
                const bulletEnd = snaps.stateB.b.find((bulletEnd) => bulletEnd.id === bulletStart.id);

                bullets[bulletStart.id].x = lerp(bulletStart.x, bulletEnd.x, dt);
                bullets[bulletStart.id].y = lerp(bulletStart.y, bulletEnd.y, dt);
            }
        }
    })
}