import {initGameState} from "../../game/storage/gameState.js";
import {updateStatisticView} from "../../../view/app/statisticsView.js";
import {updateLobbyView} from "../../../view/app/lobbyView.js";
import {changeCountdownTimer} from "../../../view/app/countdownView.js";
import {parseMap} from "../preloadingResources/mapHandler.js";
import {getSnapshotsAmount, saveSnapshot} from "../../game/storage/interpolation.js";

export function processWaitingMessage(parsedMessage, lobbyState) {
    if (!parsedMessage.oId || !parsedMessage.p) {
        console.log(`Не получены некоторые поля! ${parsedMessage}`);
        return;
    }

    lobbyState.clientId = parsedMessage.oId;
    lobbyState.players = parsedMessage.p;

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

export function processGameAssignment(parsedMessage, gameState, gameNicknames) {
    parsedMessage.p.forEach((player) => processPlayer(player, gameState));
    parsedMessage.b.forEach((bullet) => processBullet(bullet, gameState));

    saveSnapshot(parsedMessage);
    updateStatisticView(parsedMessage.stat, gameNicknames, gameState.player.id);
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
        rebornTime: player.rt,
        roomId: player.room_id,
        pc: player.pc,
    };

    if (newPlayerInModel.id === state.player.id) {
        newPlayerInModel.mousePosition = state.player.mousePosition;
        state.player = newPlayerInModel;
    } else if (getSnapshotsAmount() < 2 || !(newPlayerInModel.id in state.enemies)) {
        state.enemies[newPlayerInModel.id] = newPlayerInModel;
    }
}

function processBullet(bullet, state) {
    if (!(bullet.id in state.bullets)) {
        state.bullets[bullet.id] = {
            x: bullet.x,
            y: bullet.y,
            direction: bullet.a,
            ownerId: bullet.oId,
            id: bullet.id,
            type: bullet.t
        }
    }
}