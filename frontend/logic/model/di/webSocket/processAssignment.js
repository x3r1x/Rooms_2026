export function processAssignment(parsedMessage, state) {
    state.enemies = [];
    parsedMessage.p.forEach((player) => processPlayer(player, state));

    state.bullets = [];
    parsedMessage.b.forEach((bullet) => processBullet(bullet, state));
}

function processPlayer(player, state) {
    const newPlayerInModel = {
        x: player.x,
        y: player.y,
        direction: player.a,
        movementDirection: {
            x: player.mx,
            y: player.my
        },
        id: player.id,
        hp: player.h,
        rebornTime: player.rt
    }
    console.log(player.id);

    if (newPlayerInModel.id === state.player.id) {
        newPlayerInModel.mousePosition = state.player.mousePosition;
        state.player = newPlayerInModel;
    } else {
        state.enemies.push(newPlayerInModel);
    }
}

function processBullet(bullet, state) {
    state.bullets.push({
        x: bullet.x,
        y: bullet.y,
        direction: bullet.a,
        ownerId: bullet.oId
    })
}