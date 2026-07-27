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

    const windowWidth = window.innerWidth;
    const windowHeight = window.innerHeight;

    const targetAspect = 900 / 756;
    let cssWidth = windowWidth;
    let cssHeight = windowWidth / targetAspect;

    if (cssHeight > windowHeight) {
        cssHeight = windowHeight;
        cssWidth = windowHeight * targetAspect;
    }
    canvas.style.width = `${cssWidth}px`;
    canvas.style.height = `${cssHeight}px`;
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

    const clientX = event.clientX - rect.left;
    const clientY = event.clientY - rect.top;

    const scaleX = canvas.width / rect.width;
    const scaleY = canvas.height / rect.height;

    return {
        x: clientX * scaleX,
        y: clientY * scaleY
    };
}