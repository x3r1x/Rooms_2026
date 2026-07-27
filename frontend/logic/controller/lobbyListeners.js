import {sendReadyState} from "../model/di/messages/clientMessages.js";
import {updateReadyText, selectedWeaponClass, updatePlayerClass} from "../view/app/lobbyView.js";

export function initLobbyListeners(socket, clientId) {
    let isReady = false;
    updateReadyText(isReady);
    updatePlayerClass();
    document.getElementById("readyButton").onclick = () => {
        isReady = !isReady;
        updateReadyText(isReady);
        sendReadyState(socket, selectedWeaponClass, isReady, clientId);
    }
}
