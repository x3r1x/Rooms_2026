import {switchWindowToRegistration} from "../view/app/windowSwitcher.js";

export function initGameEndListeners() {
    document.getElementById("restartGame").onclick = () => {
        switchWindowToRegistration();
    }
}