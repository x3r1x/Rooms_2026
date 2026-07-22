import {initRegistrationListeners} from "./controller/registrationListeners.js";
import {initAppState} from "./model/app/appState.js";

export let socket = null;

if (!crypto.randomUUID) {
    crypto.randomUUID = function () {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            const r = crypto.getRandomValues(new Uint8Array(1))[0] % 16 | 0;
            const v = c === 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    };
    console.log('✅ Polyfill для crypto.randomUUID() установлен');
}

function startApp() {
    initAppState();
    initRegistrationListeners(socket);
}

startApp();