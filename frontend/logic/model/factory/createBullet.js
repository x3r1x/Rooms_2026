export function createBullet(state, bulletDirection, shotX, shotY) {
    state.bullets.push({
        x: shotX,
        y: shotY,
        direction: bulletDirection,
        ownerId: state.player.id
    })

    state.player.didShot = true;
}