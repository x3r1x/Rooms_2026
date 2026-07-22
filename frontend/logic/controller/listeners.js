import {createBullet} from "../model/factory/createBullet.js";
import {currentState} from "../model/storage/states.js";
import {GAME_CONSTANTS} from "../model/storage/gameConstants.js";

export const keys = {};

export function initListeners(canvas) {
    initWindowListeners()
    initCanvasListeners(canvas);

    canvas.addEventListener('resize', resizeCanvas)
}

function initWindowListeners() {
    window.addEventListener('keydown', function (event) {
        keys[event.key.toLowerCase()] = true;
    });

    window.addEventListener('keyup', function (event) {
        keys[event.key.toLowerCase()] = false;
    });
}

export function resizeCanvas(canvas) {
    const BASE_WIDTH = 900;
    const BASE_HEIGHT = 756;
    const container = canvas.parentElement;
    const containerWidth = container.clientWidth;
    const containerHeight = container.clientHeight;

    const aspect = BASE_WIDTH / BASE_HEIGHT;
    let newWidth = containerWidth;
    let newHeight = containerWidth / aspect;

    if (newHeight > containerHeight) {
        newHeight = containerHeight;
        newWidth = containerHeight * aspect;
    }
    canvas.width = BASE_WIDTH;
    canvas.height = BASE_HEIGHT;
    canvas.style.width = `${newWidth}px`;
    canvas.style.height = `${newHeight}px`;
}

function initCanvasListeners(canvas) {
    canvas.addEventListener('click', function (event) {
        const {x, y} = getMousePos(canvas, event);
        const direction = Math.atan2(y - currentState.player.y, x - currentState.player.x);

        const localX = GAME_CONSTANTS.PLAYER_VISUAL_WIDTH / 2;
        const localY = GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT / 2;

        const rotatedX = localX * Math.cos(direction) - localY * Math.sin(direction);
        const rotatedY = localX * Math.sin(direction) + localY * Math.cos(direction);

        const bulletStartX = currentState.player.x + rotatedX;
        const bulletStartY = currentState.player.y + rotatedY;

        createBullet(currentState, direction, bulletStartX, bulletStartY);
    });

    canvas.addEventListener('mousemove', function (event) {
        const {x, y} = getMousePos(canvas, event);

        currentState.player.mousePosition.x = x;
        currentState.player.mousePosition.y = y;
    })
}
function getMousePos(canvas, event) {
    const rect = canvas.getBoundingClientRect();
    const scaleX = canvas.width / rect.width;
    const scaleY = canvas.height / rect.height;
    return {
        x: (event.clientX - rect.left) * scaleX,
        y: (event.clientY - rect.top) * scaleY
    };
}