import {lastState} from "../model/gameModel.js";

export function createBullet(state, bulletDirection, shotX, shotY) {
    const newBullet = {
        x: shotX,
        y: shotY,
        direction: bulletDirection
    };

    lastState.bullets.push(newBullet);
}

/*

interface bullet
{
    x: number,
    y: number,
    direction: number - в радианах, от -PI до PI
}

 */