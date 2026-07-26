import {gameState} from "../model/game/storage/gameState.js";

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
    canvas.addEventListener('click', () => gameState.player.didShoot = true);

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