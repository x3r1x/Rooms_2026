import {lastState} from "../model/gameModel.js";
import { keys } from './listeners.js';
import {isBulletOutOfBounds} from "./cleaner.js";


export function updateGame(elapsedTime, state) {
    //TODO: написать игру =)
    let dx = 0;
    let dy = 0;
     state.bullets = state.bullets.filter(bullet => {
        bullet.x += Math.cos(bullet.direction) * elapsedTime * 1.5;
        bullet.y += Math.sin(bullet.direction) * elapsedTime * 1.5;

        return !isBulletOutOfBounds(bullet);
    })

    console.log(state.bullets);

    let speed = 0.1;
    if (keys['w'] || keys['ц'] || keys['arrowup']) {
        dy -= 1
    }
    if (keys['s'] || keys['ы'] || keys['arrowdown']){
        dy += 1
    }
    if (keys['a'] || keys['ф'] || keys['arrowleft']){
        dx -= 1
    }
    if (keys['d'] || keys['в'] || keys['arrowright']){
        dx += 1
    }
    let lengthVector = (Math.sqrt(dx * dx + dy * dy));
    if (lengthVector > 0) {
        dx /= lengthVector;
        dy /= lengthVector;
    }
    //Отладочный код для перемещения квадратика
    state.square.x += dx * speed * elapsedTime;
    state.square.y += dy * speed * elapsedTime;
}