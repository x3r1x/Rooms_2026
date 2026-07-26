import {sendReadyState} from "../model/di/messages/clientMessages.js";
import {updateReadyText, selectedWeaponClass, updatePlayerClass} from "../view/app/lobbyView.js";

let isReady = false;

export function initLobbyListeners(socket, clientId) {
    updatePlayerClass();
    document.getElementById("readyButton").onclick = () => {
        isReady = !isReady;
        updateReadyText(isReady);
        sendReadyState(socket, selectedWeaponClass, isReady, clientId);
    }
}
