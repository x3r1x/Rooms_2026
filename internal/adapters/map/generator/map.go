package generator

import (
	"gamedevRooms/internal/domain"
	"math/rand/v2"
)

type Map struct {
	gameMap map[string]*Room
}

type roomDirection struct {
	dx, dy   int
	name     string
	opposite string
}

type mapPoint struct {
	x, y int
}

var directions = []roomDirection{
	{dx: 0, dy: 1, name: domain.BottomMarker, opposite: domain.TopMarker},
	{dx: -1, dy: 0, name: domain.LeftMarker, opposite: domain.RightMarker},
	{dx: 0, dy: -1, name: domain.TopMarker, opposite: domain.BottomMarker},
	{dx: 1, dy: 0, name: domain.RightMarker, opposite: domain.LeftMarker},
}

func NewMap(roomsCount int) *Map {
	return &Map{
		gameMap: generateShuffledRooms(roomsCount),
	}
}

func (m *Map) GetGameMap() map[string]*Room {
	return m.gameMap
}

func generateShuffledRooms(roomsCount int) map[string]*Room {
	newMap := make(map[string]*Room)
	roomsGrid := make(map[mapPoint]*Room)
	rg := getRoomGenerator()

	startRoom := rg.CreateRoomWithRandomBarrier()
	startPoint := mapPoint{0, 0}
	roomsGrid[startPoint] = startRoom
	newMap[startRoom.GetId()] = startRoom

	createdPoints := []mapPoint{startPoint}

	generateRandomRoomsGrid(newMap, roomsGrid, rg, createdPoints, roomsCount)
	connectNeighbouredRooms(roomsGrid)
	applyExitChangesToMatrix(rg, newMap)

	return newMap
}

func generateRandomRoomsGrid(newMap map[string]*Room, grid map[mapPoint]*Room, rg *RoomGenerator, createdPoints []mapPoint, roomsCount int) {
	for len(newMap) < roomsCount {
		parentMapPoint := createdPoints[rand.IntN(len(createdPoints))]
		parentRoom := grid[parentMapPoint]
		randomDirections := getRandomDirections()

		for _, direction := range randomDirections {
			newMapPoint := mapPoint{
				x: parentMapPoint.x + direction.dx,
				y: parentMapPoint.y + direction.dy,
			}

			if _, exists := grid[newMapPoint]; !exists {
				newRoom := rg.CreateRoomWithRandomBarrier()

				parentRoom.SetExit(direction.name, newRoom.GetId())
				newRoom.SetExit(direction.opposite, parentRoom.GetId())

				grid[newMapPoint] = newRoom
				newMap[newRoom.id] = newRoom
				createdPoints = append(createdPoints, newMapPoint)

				break
			}
		}
	}
}

func connectNeighbouredRooms(grid map[mapPoint]*Room) {
	for point, room := range grid {
		for _, direction := range directions {
			neighborMapPoint := mapPoint{
				x: point.x + direction.dx,
				y: point.y + direction.dy,
			}

			if neighbour, exists := grid[neighborMapPoint]; exists {
				if room.GetExit(direction.name) == "" && rand.Float64() < domain.ConnectNeighbouredRoomChance {
					room.SetExit(direction.name, neighbour.id)
					neighbour.SetExit(direction.opposite, room.id)
				}
			}
		}
	}
}

func applyExitChangesToMatrix(rg *RoomGenerator, newMap map[string]*Room) {
	for _, room := range newMap {
		rg.ProcessRoomExits(room)
	}
}

func getRandomDirections() []roomDirection {
	randomDirections := make([]roomDirection, len(directions))
	copy(randomDirections, directions)
	rand.Shuffle(len(directions), func(i, j int) {
		randomDirections[i], randomDirections[j] = randomDirections[j], randomDirections[i]
	})

	return randomDirections
}

func getRoomGenerator() *RoomGenerator {
	parser := NewJsonParser()

	return NewRoomGenerator(parser)
}
