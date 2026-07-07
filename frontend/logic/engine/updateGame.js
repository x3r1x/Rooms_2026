import {isBulletOutOfBounds} from "./cleaner.js";


export function updateGame(elapsedTime, state) {
     state.bullets = state.bullets.filter(bullet => {
        bullet.x += Math.cos(bullet.direction) * elapsedTime * 1.5;
        bullet.y += Math.sin(bullet.direction) * elapsedTime * 1.5;

        return !isBulletOutOfBounds(bullet);
    })

    console.log(state.bullets);
}