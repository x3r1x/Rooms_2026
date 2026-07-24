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

    ENEMY_SPRITE_LIST_PATH: 'assets/images/enemy.png',
    ENEMY_BULLET_SKIN_PATH: 'assets/images/enemy_bullet_classic.png',

    TILE_IMG_PATH: 'assets/tile/tileset x1.png'
}

export const PLAYER_LOCAL_POINTS = [
    {x: -GAME_CONSTANTS.HALF, y: -GAME_CONSTANTS.HALF},
    {x: GAME_CONSTANTS.HALF, y: -GAME_CONSTANTS.HALF},
    {x: GAME_CONSTANTS.HALF, y: GAME_CONSTANTS.HALF},
    {x: -GAME_CONSTANTS.HALF, y: GAME_CONSTANTS.HALF}
];

export const GAME_SPRITES = {
    PLAYER_GOES: new Image(),
    PLAYER_DIE: new Image(),

    HEALTH_BAR: new Image(),

    BULLET_HIT: new Image(),
    BULLET_FLIES: new Image(),

    ENEMY_GOES: new Image(),
    ENEMY_BULLET_FLIES: new Image()
}

export const TILE_IMG = {
    TILE: new Image(),
}