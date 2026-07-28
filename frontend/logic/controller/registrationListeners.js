import {getSocket} from "../model/di/webSocket/server.js";

export function initRegistrationListeners(socket) {
    const button = document.getElementById("nicknameButton");
    const nicknameInput = document.getElementById("nicknameInput");
    button.onclick = () => {
        socket = getSocket(nicknameInput.value);
        button.onclick = null;
    }
    nicknameInput.onkeydown = (event) => {
        if (event.key === "Enter") {
            socket = getSocket(nicknameInput.value);
            button.onclick = null;
            nicknameInput.onkeydown = null;
        }
    };
}