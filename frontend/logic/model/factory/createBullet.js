export function createBullet(state, bulletDirection, shotX, shotY) {
    state.player.bullets[crypto.randomUUID()] = {
        x: shotX,
        y: shotY,
        direction: bulletDirection
    };

    state.newBulletsDirection.push(bulletDirection);
}