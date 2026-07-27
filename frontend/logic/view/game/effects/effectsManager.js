import {GAME_SPRITES} from "../../../model/game/storage/gameConstants.js";

const effects = [];

export function spawnEffect(x, y, type, direction = 0) {
    console.log("Spawn effect direction:", direction);
    effects.push({
        x: x,
        y: y,
        type: type,
        direction: direction,
        currentFrame: 0,
        totalFrames: 3,
        startTime: Date.now(),
        frameDuration: 1000 / 30
    });
    console.log(effects);
}

function drawEffect(context, fx) {
    const spriteSheet = GAME_SPRITES.EFFECTS[fx.type];
    const frameWidth = spriteSheet.width / fx.totalFrames;
    const frameHeight = spriteSheet.height;

    const scale = 0.5;
    const drawWidth = frameWidth * scale;
    const drawHeight = frameHeight * scale;

    context.save();
    context.translate(fx.x, fx.y);
    context.rotate(fx.direction + Math.PI);

    context.drawImage(
        spriteSheet,
        fx.currentFrame * frameWidth, 0, frameWidth, frameHeight,
        -drawWidth / 2, -drawHeight / 2, drawWidth, drawHeight
    );

    context.restore();
}
export function updateAndDrawEffects(context) {
    const now = Date.now();

    for (let i = effects.length - 1; i >= 0; i--) {
        const fx = effects[i];

        const elapsed = now - fx.startTime;
        fx.currentFrame = Math.floor(elapsed / fx.frameDuration);

        if (fx.currentFrame >= fx.totalFrames) {
            effects.splice(i, 1);
            continue;
        }
        drawEffect(context, fx);
    }
}