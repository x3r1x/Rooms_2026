import {getSocket} from "../model/di/webSocket/server.js";
import {switchWindowFromRegistrationToLobby} from "../view/app/windowSwitcher.js";

export function initRegistrationListeners(socket) {
    document.getElementById("nicknameButton").addEventListener("click", () => {
        const nicknameInput = document.getElementById("nicknameInput");
        socket = getSocket(nicknameInput.value);

        switchWindowFromRegistrationToLobby()
    })
}