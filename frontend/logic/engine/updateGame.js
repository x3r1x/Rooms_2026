import {lastState} from "../model/gameModel";
import { keys } from './listeners.js';
export function updateGame(elapsedTime, state) {
    //TODO: написать игру =)
    let dx = 0;
    let dy = 0;
    let speed = 0.5;
    if (keys['w']){
        dy -= 1
    }
    if (keys['s']){
        dy += 1
    }
    if (keys['a']){
        dx -= 1
    }
    if (keys['d']){
        dx += 1
    }
    //Отладочный код для перемещения квадратика
    state.square.x += dx * speed * elapsedTime;
    state.square.y += dy * speed * elapsedTime;
}