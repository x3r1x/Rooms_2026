export const GAME_CONSTANTS = {
    PISTOL_PLAYER_SPEED: 0.6,
    SHOTGUN_PLAYER_SPEED: 0.4,
    ROCKET_PLAYER_SPEED: 0.2,
    BULLET_SPEED: 1.5,
    CANVAS_START: 0,

    PLAYER_VISUAL_HEIGHT: 50,
    PLAYER_VISUAL_WIDTH: 44,

    PLAYER_HITBOX_SIZE: 50,

    PLAYER_HITBOX_HEIGHT: 35,
    PLAYER_HITBOX_WIDTH: 30,
    HALF: 17,

    MUZZLE_FLASH_PATH: 'assets/images/muzzleFlash.png',
    MUZZLE_FLASH_S_PATH: 'assets/images/muzzleFlash_s.png',
    EXPLOSION_S_PATH: 'assets/images/explosion_s.png',
    EXPLOSION_PATH: 'assets/images/explosion.png',
    PLAYER_DEATH_PATH: 'assets/images/death.png',
    PLAYER_HIT_PATH: 'assets/images/hit.png',
    SHIELD_PATH: 'assets/images/shield.png',
    HEALTH_BAR_PATH: 'assets/images/health-bar_sheet.png',

    PLAYER_TYPES: {
        PISTOL: "g",
        SHOTGUN: "r",
        ROCKET_LAUNCHER: "s"
    },

    IS_PLAYER: {
        NO_PLAYER: "n",
        KILLER: "k",
        VICTIM: "v",
        BOTH: "b"
    },

    BULLET: {
        PLAYER: {},
        ENEMY: {},
    },
    ENEMY_PATH: {
        g: 'assets/images/enemy_g.png',
        r: 'assets/images/enemy_r.png',
        s: 'assets/images/enemy_s.png',
        bg: 'assets/images/bullet_e_g.png',
        br: 'assets/images/bullet_e_r.png',
        bs: 'assets/images/bullet_e_s.png',
    },
    PLAYER_PATH: {
        g: 'assets/images/player_g.png',
        r: 'assets/images/player_r.png',
        s: 'assets/images/player_s.png',
        bg: 'assets/images/bullet_p_g.png',
        br: 'assets/images/bullet_p_r.png',
        bs: 'assets/images/bullet_p_s.png',
    },

    SNAPSHOTS_AMOUNT: 100,
    INTERPOLATION_DELAY: 100,
    MAX_EXTRAPOLATION_TIME: 300,

    KILL_FEED_DELETE_TIME: 5000,

    TILE_IMG_PATH: 'assets/tile/tileset x1.png'
}

export const PLAYER_LOCAL_POINTS = [
    {x: -GAME_CONSTANTS.HALF, y: -GAME_CONSTANTS.HALF},
    {x: GAME_CONSTANTS.HALF, y: -GAME_CONSTANTS.HALF},
    {x: GAME_CONSTANTS.HALF, y: GAME_CONSTANTS.HALF},
    {x: -GAME_CONSTANTS.HALF, y: GAME_CONSTANTS.HALF}
];

export const GAME_SPRITES = {
    PLAYER_DIE: new Image(),

    HEALTH_BAR: new Image(),

    BULLET_HIT: new Image(),
    BULLET_FLIES: new Image(),

    ENEMY: {
        g: new Image(),
        r: new Image(),
        s: new Image(),
        bg: {
            img: new Image(),
            w: 11,
            h: 3
        },
        br: {
            img: new Image(),
            w: 5,
            h: 5
        },
        bs: {
            img: new Image(),
            w: 16,
            h: 7
        },
    },
    PLAYER: {
        g: new Image(),
        r: new Image(),
        s: new Image(),
        bg: {
            img: new Image(),
            w: 11,
            h: 3
        },
        br: {
            img: new Image(),
            w: 5,
            h: 5
        },
        bs: {
            img: new Image(),
            w: 16,
            h: 7
        },
    },
    EFFECTS: {
        MUZZLE_FLASH: {
            img: new Image(),
            scale: 1,
            maxFrames: 3,
        },
        MUZZLE_FLASH_G: {
            img: new Image(),
            scale: 1,
            maxFrames: 3,
        },
        MUZZLE_FLASH_R: {
            img: new Image(),
            scale: 1,
            maxFrames: 3,
        },
        MUZZLE_FLASH_S: {
            img: new Image(),
            scale: 1,
            maxFrames: 3,
        },
        EXPLOSION_G: {
            img: new Image(),
            scale: 10,
            maxFrames: 9,
        },
        EXPLOSION_R: {
            img: new Image(),
            scale: 10,
            maxFrames: 9,
        },
        EXPLOSION_S: {
            img: new Image(),
            scale: 10,
            maxFrames: 9,
        },
        EXPLOSION: {
            img: new Image(),
            scale: 0.3,
            maxFrames: 3,
        },

        PLAYER_DEATH: {
            img: new Image(),
            scale: 2.5,
            maxFrames: 12,
        },

        SOM_EXPLOSION: {
            img: new Image(),
            scale: 2.5,
            maxFrames: 8,
        },
        SHIELD: {
            img: new Image(),
            scale: 1.5,
            maxFrames: 24,
        },

        PLAYER_HIT: {
            img: new Image(),
            scale: 1.5,
            maxFrames: 9,
        },
    },
}

export const WEAPON_SPRITES = {
    PISTOL: "assets/images/pistol.png",
    SHOTGUN: "assets/images/shotgun.png",
    ROCKET_LAUNCHER: "assets/images/rocket_launcher.png"
}

export const TILE_IMG = {
    TILE: new Image(),
}