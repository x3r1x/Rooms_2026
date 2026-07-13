export function createBullet(state, bulletDirection, shotX, shotY) {
    const newBullet = {
        x: shotX,
        y: shotY,
        direction: bulletDirection
    };

    //FIXME: потом убрать генерацию IDшника, когда будет связь с фронтом
    state.player.bullets[crypto.randomUUID()] = newBullet;
}