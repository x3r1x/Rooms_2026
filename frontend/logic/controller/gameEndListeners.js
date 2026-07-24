import {switchToRegistrationAppState} from "../model/app/appState.js";

export function initGameEndListeners() {
    document.getElementById("restartGame").onclick = () => {
        switchToRegistrationAppState();
    }
}