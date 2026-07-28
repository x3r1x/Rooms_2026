import {GAME_SPRITES} from "../../../model/game/storage/gameConstants.js";

const effects = [];

function spawnEffect(x, y, type, direction = 0) {
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

export function checkAndSpawnNewBulletEffects(stateA, stateB) {
    if (!stateA || !stateB || !stateA.b || !stateB.b) return;

    for (let i = 0; i < stateB.b.length; i++) {
        const bulletB = stateB.b[i];
        let existedBefore = false;

        for (let j = 0; j < stateA.b.length; j++) {
            if (stateA.b[j].id === bulletB.id) {
                existedBefore = true;
                break;
            }
        }
        if (!existedBefore) {
            spawnEffect(
                bulletB.x,
                bulletB.y,
                "MUZZLE_FLASH",
                bulletB.a
            );
        }
    }
}

export function checkAndSpawnExplosions(stateA, stateB) {
    if (!stateA || !stateB || !stateA.b || !stateB.b) return;

    for (let i = 0; i < stateA.b.length; i++) {
        const oldBullet = stateA.b[i];
        let isAliveInNewState = false;

        for (let j = 0; j < stateB.b.length; j++) {
            if (stateB.b[j].id === oldBullet.id) {
                isAliveInNewState = true;
                break;
            }
        }

        if (!isAliveInNewState) {
            spawnEffect(
                oldBullet.x,
                oldBullet.y,
                "EXPLOSION",
                oldBullet.a
            );
        }
    }
}