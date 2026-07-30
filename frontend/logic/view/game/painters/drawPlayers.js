import {GAME_CONSTANTS, GAME_SPRITES} from "../../../model/game/storage/gameConstants.js";
import {gameNicknames} from "../../../model/game/storage/gameState.js";

export function drawPlayers(context, mainPlayer, enemies) {
    drawMainPlayer(context, mainPlayer);

    for (const enemy of Object.values(enemies)) {
        if (mainPlayer.roomId === enemy.roomId) {
            drawEnemy(context, enemy)
        }
    }
}

function drawMainPlayer(context, player) {
    const deathScreen = document.getElementById("death-screen");
    const respawnTimer = document.getElementById("respawn-timer");
    const spriteSheet = GAME_SPRITES.PLAYER[player.pc];
    const frameWidth = spriteSheet.width / 6;
    const frameHeight = spriteSheet.height;
    if (player.hp > 0) {
        deathScreen.style.display = "none"
        if (player.movementDirection.x !== 0 || player.movementDirection.y !== 0) {
            player.spriteIndex = Math.floor(Date.now() / 200) % 6;
        } else {
            player.spriteIndex = 0;
        }
        if (!(player.ps && Math.floor(Date.now() / 50) % 2 === 0)) {
            context.save();
            context.translate(player.x, player.y);
            context.rotate(player.direction);
            context.drawImage(
                spriteSheet,
                player.spriteIndex * frameWidth, 0, frameWidth, frameHeight,
                -GAME_CONSTANTS.PLAYER_VISUAL_WIDTH / 2, -GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT / 2, GAME_CONSTANTS.PLAYER_VISUAL_WIDTH, GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT
            );
            context.restore();
        }
        drawHealthBar(context, player);
        drawName(context, player);
    } else {
        if (player.hp != null) {
            deathScreen.style.display = "flex";
            respawnTimer.textContent = Math.floor(player.rebornTime / 50);
        }
    }
}

function drawEnemy(context, enemy) {
    const spriteSheet = GAME_SPRITES.ENEMY[enemy.pc];
    const frameWidth = spriteSheet.width / 6;
    const frameHeight = spriteSheet.height;
    if (enemy.hp > 0) {
        if (enemy.movementDirection.x !== 0 || enemy.movementDirection.y !== 0) {
            enemy.spriteIndex = Math.floor(Date.now() / 200) % 6;
        } else {
            enemy.spriteIndex = 0;
        }

        if (!(enemy.ps && Math.floor(Date.now() / 50) % 2 === 0)) {
            context.save();
            context.translate(enemy.x, enemy.y);
            context.rotate(enemy.direction);
            context.drawImage(
                spriteSheet,
                enemy.spriteIndex * frameWidth, 0, frameWidth, frameHeight,
                -GAME_CONSTANTS.PLAYER_VISUAL_WIDTH / 2, -GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT / 2, GAME_CONSTANTS.PLAYER_VISUAL_WIDTH, GAME_CONSTANTS.PLAYER_VISUAL_HEIGHT
            );
            context.restore();
        }
        drawHealthBar(context, enemy);
        drawName(context, enemy)
    }

}

function drawHealthBar(context, player) {
    const sprite = GAME_SPRITES.HEALTH_BAR;
    const frameWidth = sprite.width / 11;
    const scale = 0.8;
    const frameHeight = sprite.height;
    player.hpSpriteIndex = 10 - Math.floor(player.hp / 10);
    context.drawImage(
        sprite,
        player.hpSpriteIndex * frameWidth, 0, frameWidth, frameHeight,
        player.x - (scale * frameWidth / 2), player.y - (GAME_CONSTANTS.PLAYER_VISUAL_WIDTH), scale * frameWidth, scale * frameHeight
    );

}

function drawName(context, player) {
    const nickname = gameNicknames[player.id];

    context.fillStyle = "white";
    context.font = "16px monospace";
    context.textAlign = "center";

    const textX = player.x;
    const textY = player.y - GAME_CONSTANTS.PLAYER_VISUAL_WIDTH - 3;

    context.fillText(nickname, textX, textY);
}