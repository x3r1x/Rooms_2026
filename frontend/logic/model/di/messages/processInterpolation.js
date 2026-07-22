export function processInterpolation(playerInterpolations, state) {
    for (const [id, playerInterpolation] of Object.entries(playerInterpolations)) {
        const playerInModel = getPlayerFromModelById(state, id);

        if (playerInModel === null) {
            console.log(`processIntepolation: Unknown id: ${id}`);
        } else {
            interpolatePlayerInModel(playerInModel, playerInterpolation);
        }
    }
}

function interpolatePlayerInModel(player, interpolation) {
    player.x += interpolation.dx;
    player.y += interpolation.dy;
    player.direction += interpolation.deltaDirection;

    const newBulletsList = {};
    for (const [id, bullet] of Object.entries(player.bullets)) {
        if (id in interpolation.deltaBullets) {
            newBulletsList[id] = {
                x: bullet.x + interpolation.deltaBullets[id].dx,
                y: bullet.y + interpolation.deltaBullets[id].dy,
                movementDirection: bullet.movementDirection
            }

        }
    }

    for (const [id, newBullet] of Object.entries(interpolation.newBullets)) {
        newBulletsList[id] = newBullet;
    }

    player.bullets = newBulletsList;
}