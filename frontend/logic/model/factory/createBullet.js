export function createBullet(state, bulletDirection, shotX, shotY) {
    const newBullet = {
        x: shotX,
        y: shotY,
        direction: bulletDirection
    };

    state.player.bullets[crypto.randomUUID()] = newBullet;
    state.newBulletsDirection.push(bulletDirection);
}