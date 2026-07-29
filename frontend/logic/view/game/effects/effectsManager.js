import {GAME_SPRITES} from "../../../model/game/storage/gameConstants.js";
import {handleEnemies} from "../../../model/game/engine/players";
import {handleBullets} from "../../../model/game/engine/bullet";

const effects = [];

function spawnEffect(x, y, type, direction = 0) {
    const spriteInfo = GAME_SPRITES.EFFECTS[type];
    effects.push({
        x: x,
        y: y,
        type: type,
        direction: direction,
        currentFrame: 0,
        scale: spriteInfo.scale,
        maxFrames: spriteInfo.maxFrames,
        startTime: Date.now(),
        frameDuration: 1000 / 30,
        img: GAME_SPRITES.EFFECTS[type].img,
    });
}

function drawEffect(context, fx) {

    const spriteSheet = fx.img;
    const frameWidth = spriteSheet.width / fx.maxFrames;
    const frameHeight = spriteSheet.height;

    const drawWidth = frameWidth * fx.scale;
    const drawHeight = frameHeight * fx.scale;

    context.save();
    context.translate(fx.x, fx.y);
    context.rotate(fx.direction);

    context.drawImage(
        spriteSheet,
        fx.currentFrame * frameWidth, 0, frameWidth, frameHeight,
        -drawWidth / 2, -drawHeight / 2, drawWidth, drawHeight
    );

    context.restore();
}

export function checkAndSpawnEffects(neighbouredSnapshots) {
    checkAndSpawnNewBulletEffects(neighbouredSnapshots.snapA, neighbouredSnapshots.snapB);
    checkAndSpawnExplosions(neighbouredSnapshots.snapA, neighbouredSnapshots.snapB);
    checkAndSpawnDie(neighbouredSnapshots.snapA, neighbouredSnapshots.snapB);
    checkAndSpawnHit(neighbouredSnapshots.snapA, neighbouredSnapshots.snapB);
    checkAndSpawnShield(neighbouredSnapshots.snapA, neighbouredSnapshots.snapB);
}

export function updateAndDrawEffects(context) {
    const now = Date.now();

    for (let i = effects.length - 1; i >= 0; i--) {
        const fx = effects[i];

        const elapsed = now - fx.startTime;
        fx.currentFrame = Math.floor(elapsed / fx.frameDuration);

        if (fx.currentFrame >= fx.maxFrames) {
            effects.splice(i, 1);
            continue;
        }
        drawEffect(context, fx);
    }
}

function checkAndSpawnNewBulletEffects(stateA, stateB) {
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
            switch (bulletB.t) {
                case 'g':
                    spawnEffect(
                        bulletB.x,
                        bulletB.y,
                        "MUZZLE_FLASH",
                        bulletB.a
                    );
                    break;
                case 'r':
                    spawnEffect(
                        bulletB.x,
                        bulletB.y,
                        "MUZZLE_FLASH",
                        bulletB.a
                    );
                    break;
                case 's':
                    spawnEffect(
                        bulletB.x,
                        bulletB.y,
                        "MUZZLE_FLASH_S",
                        bulletB.a
                    );
                    break;
            }

        }
    }
}

function checkAndSpawnExplosions(stateA, stateB) {
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
            switch (oldBullet.t) {
                case 'g':
                    spawnEffect(
                        oldBullet.x,
                        oldBullet.y,
                        "EXPLOSION",
                        oldBullet.a
                    );
                    break;
                case 'r':
                    spawnEffect(
                        oldBullet.x,
                        oldBullet.y,
                        "EXPLOSION",
                        oldBullet.a
                    );
                    break;
                case 's':
                    spawnEffect(
                        oldBullet.x,
                        oldBullet.y,
                        "EXPLOSION_S",
                        oldBullet.a
                    );
                    break;
            }
        }
    }
}

function checkAndSpawnDie(stateA, stateB){
    if (!stateA || !stateB) return;

    const oldPlayers = stateA.p;
    const newPlayers = stateB.p;

    for (let i = 0; i < newPlayers.length; i++) {
        const newEntity = newPlayers[i];
        const oldEntity = oldPlayers.find(e => e.id === newEntity.id);
        if (oldEntity && oldEntity.h > 0 && newEntity.h <= 0) {
            spawnEffect(
                newEntity.x,
                newEntity.y,
                "PLAYER_DEATH",
                newEntity.a || 0
            );
        }
    }
}

function checkAndSpawnHit(stateA, stateB){
    if (!stateA || !stateB) return;

    const oldPlayers = stateA.p;
    const newPlayers = stateB.p;

    for (let i = 0; i < newPlayers.length; i++) {
        const newPlayer = newPlayers[i];
        const oldPlayer = oldPlayers.find(e => e.id === newPlayer.id);
        if (oldPlayer && oldPlayer.h > newPlayer.h) {
            spawnEffect(
                newPlayer.x,
                newPlayer.y,
                "PLAYER_HIT",
                newPlayer.a || 0
            );
        }
    }
}

function checkAndSpawnShield(stateA, stateB){
    if (!stateA || !stateB) return;

    const oldPlayers = stateA.p;
    const newPlayers = stateB.p;

    for (let i = 0; i < newPlayers.length; i++) {
        const newPlayer = newPlayers[i];
        const oldPlayer = oldPlayers.find(e => e.id === newPlayer.id);
        if (oldPlayer && oldPlayer.h <= 0 && newPlayer.h > 0) {
            spawnEffect(
                newPlayer.x,
                newPlayer.y,
                "SHIELD",
                newPlayer.a || 0
            );
        }
    }
}