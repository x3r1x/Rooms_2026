import {initGameState} from "../../game/storage/gameState.js";
import {updateStatisticView} from "../../../view/app/statisticsView.js";
import {updateLobbyView} from "../../../view/app/lobbyView.js";
import {changeCountdownTimer} from "../../../view/app/countdownView.js";

export function processWaitingMessage(parsedMessage, lobbyState) {
    if (!parsedMessage.oId || !parsedMessage.p) {
        console.log(`Не получены некоторые поля! ${parsedMessage}`);
        return;
    }

    lobbyState.clientId = parsedMessage.oId;
    lobbyState.players = parsedMessage.p;
    lobbyState.players.sort((player1, player2) => player1.n.localeCompare(player2.n))

    updateLobbyView(lobbyState.clientId, lobbyState.players);
}

export function processReadyMessage(parsedMessage, lobbyState, gameState) {
    if (!parsedMessage.c || !parsedMessage.m) {
        console.log(`Не получены некоторые поля! ${parsedMessage}`);
        return;
    }

    if (gameState == null) {
        initGameState(lobbyState.clientId);
        parseMap(parsedMessage);
        //TODO: parse map
    }

    lobbyState.countdown = parsedMessage.c;

}

export function processCountdownMessage(parsedMessage, lobbyState) {
    if (!parsedMessage.c) {
        console.log(`Не получены некоторые поля! ${parsedMessage}`);
        return;
    }

    lobbyState.countdown = parsedMessage.c;
    changeCountdownTimer(lobbyState.countdown);
}

export function processGameAssignment(parsedMessage, gameState) {
    gameState.enemies = [];
    parsedMessage.p.forEach((player) => processPlayer(player, gameState));

    gameState.bullets = [];
    parsedMessage.b.forEach((bullet) => processBullet(bullet, gameState));

    updateStatisticView(gameState);
    gameState.didShoot = false;
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