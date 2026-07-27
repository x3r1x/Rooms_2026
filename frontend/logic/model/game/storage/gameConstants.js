export const GAME_CONSTANTS = {
    PLAYER_SPEED: 0.2,
    BULLET_SPEED: 1.5,
    CANVAS_START: 0,

    PLAYER_VISUAL_SIZE: 40,
    PLAYER_VISUAL_HEIGHT: 50,
    PLAYER_VISUAL_WIDTH: 44,

    PLAYER_HITBOX_SIZE: 50,

    PLAYER_HITBOX_HEIGHT: 35,
    PLAYER_HITBOX_WIDTH: 30,

    HALF: 17,
    PLAYER_START_X: 500,
    PLAYER_START_Y: 400,
    PLAYER_SKIN_BLUE_PATH: 'assets/images/person.png',

    BULLET_WIDTH: 5,
    BULLET_HEIGHT: 25,
    BULLET_COLOR: "#cdcbcb",

    HEALTH_BAR_PATH: 'assets/images/health-bar_sheet.png',

    BULLET_SKIN_PATH: 'assets/images/bullet_classic.png',

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
            w: 3,
            h: 3
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
            w: 3,
            h: 3
        },
        bs: {
            img: new Image(),
            w: 16,
            h: 7
        },
    },



    ENEMY_BULLET_FLIES: new Image()
}

export const TILE_IMG = {
    TILE: new Image(),
}