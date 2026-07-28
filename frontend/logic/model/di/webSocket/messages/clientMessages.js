export function registerClient(socket, nickname) {
    const message = {
        "n": nickname
    }

    socket.send(JSON.stringify(message))
}

export function sendReadyState(socket, playerClass, ready, clientId) {
    const message = {
        "id": clientId,
        "r": ready,
        "pc": playerClass,
    }

    socket.send(JSON.stringify(message))
}

export function sendGameInfo(socket, state) {
    const message = {
        "id": state.player.id,
        "a": state.player.direction,
        "mx": state.player.movementDirection.x,
        "my": state.player.movementDirection.y,
        "s": state.player.didShoot
    }

    socket.send(JSON.stringify(message));
}