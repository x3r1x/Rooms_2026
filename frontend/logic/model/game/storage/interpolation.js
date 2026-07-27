import {GAME_CONSTANTS} from "./gameConstants.js";

let snapshots = new Array(GAME_CONSTANTS.SNAPSHOTS_AMOUNT);
let head = 0;
let snapshotsAmount = 0;
let clientTime = 0;

export function lerp(start, end, dt) {
    return start + (end - start) * dt;
}

export function lerpDirection(start, end, dt) {
    let diff = start - end;

    if (diff < -Math.PI) {
        diff += 2 * Math.PI;
    }

    if (diff > Math.PI) {
        diff -= 2 * Math.PI;
    }

    return start + diff * dt;
}

export function getSnapshotsAmount() {
    return snapshotsAmount;
}

export function saveSnapshot(snapshot) {
    snapshots[head++] = snapshot;
    head = head % GAME_CONSTANTS.SNAPSHOTS_AMOUNT;
    snapshotsAmount++;
}

export function addToClientTime(time) {
    clientTime += time;
}

export function resetInterpolationModule() {
    clientTime = 0;
    snapshotsAmount = 0;
    snapshots = new Array(GAME_CONSTANTS.SNAPSHOTS_AMOUNT);
    head = 0;
}

export function getClientTime() {
    return clientTime;
}

export function getNeighbouringSnapshots(targetTime) {
    let stateA = null;
    let stateB = null;

    let stateAIndex = (head - 1 + GAME_CONSTANTS.SNAPSHOTS_AMOUNT) % GAME_CONSTANTS.SNAPSHOTS_AMOUNT;

    for (let i = 0; i< GAME_CONSTANTS.SNAPSHOTS_AMOUNT; i++) {
        const currentSnapshot = snapshots[stateAIndex];
        if (!currentSnapshot) {
            break;
        }

        if (currentSnapshot.t <= targetTime) {
            stateA = currentSnapshot;
            stateB = snapshots[(stateAIndex + 1) % GAME_CONSTANTS.SNAPSHOTS_AMOUNT];

            if (stateB && stateB.t > targetTime) {
                return {
                    stateA: stateA,
                    stateB: stateB
                }
            }
        }

        stateAIndex = (stateAIndex - 1 + GAME_CONSTANTS.SNAPSHOTS_AMOUNT) % GAME_CONSTANTS.SNAPSHOTS_AMOUNT;
    }

    return {
        stateA: null,
        stateB: null
    }
}