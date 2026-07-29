import {getSocket} from "../model/di/webSocket/server.js";

const button = document.getElementById("nicknameButton");
const nicknameInput = document.getElementById("nicknameInput");

export function initRegistrationListeners(socket) {
    button.onclick = () => {
        socket = getSocket(nicknameInput.value);
        setRegistrationStates(true);

        socket.onerror = () => {
            setRegistrationStates(false);
        }
    }
    nicknameInput.onkeydown = (event) => {
        if (event.key === "Enter") {
            socket = getSocket(nicknameInput.value);
            setRegistrationStates(true);

            socket.onerror = () => {
                setRegistrationStates(false);
            }
        }
    };
}

function setRegistrationStates(isDisabled) {
    button.disabled = isDisabled;
    nicknameInput.disabled = isDisabled;
}