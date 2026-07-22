import {APP_STATES} from "./appConstants.js";

export let appState = null

export function initAppState() {
    appState = APP_STATES.REGISTRATION
}
