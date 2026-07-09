export function createBullet(state, bulletDirection, shotX, shotY) {
    const newBullet = {
        x: shotX,
        y: shotY,
        direction: bulletDirection
    };
    state.bullets.push(newBullet);
}