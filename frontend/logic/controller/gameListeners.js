import {createBullet} from "../model/game/factory/createBullet.js";
import {gameState} from "../model/game/storage/gameState.js";
import {GAME_CONSTANTS} from "../model/game/storage/gameConstants.js";

export const keys = {};

export function initGameListeners(canvas) {
    initWindowListeners()
    initCanvasListeners(canvas);

    resizeCanvas(canvas);
    
    window.addEventListener('resize', () => resizeCanvas(canvas));
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
    canvas.width = 900;
    canvas.height = 756;
    
    canvas.style.width = '100%';
    canvas.style.height = '100%';
}

function initCanvasListeners(canvas) {
    canvas.addEventListener('click', function (event) {
        const {x, y} = getMousePos(canvas, event);
        const direction = Math.atan2(y - gameState.player.y, x - gameState.player.x);

        const localX = GAME_CONSTANTS.PLAYER_VISUAL_WIDTH / 2;
        const localY = GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT / 2;

        const rotatedX = localX * Math.cos(direction) - localY * Math.sin(direction);
        const rotatedY = localX * Math.sin(direction) + localY * Math.cos(direction);

        const bulletStartX = gameState.player.x + rotatedX;
        const bulletStartY = gameState.player.y + rotatedY;

        createBullet(gameState, direction, bulletStartX, bulletStartY);
    });

    canvas.addEventListener('mousemove', function (event) {
        const {x, y} = getMousePos(canvas, event);

        gameState.player.mousePosition.x = x;
        gameState.player.mousePosition.y = y;
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