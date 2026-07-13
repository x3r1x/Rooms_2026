package infrastructure

//
//import (
//	"gamedevRooms/internal/game"
//	"gamedevRooms/internal/models"
//)
//
//func MakeServerMessage(players []*models.PlayerState) models.ServerMessage {
//	serverMessage := models.ServerMessage{
//		Players: make([]models.PlayerState, 0, len(game.Clients)),
//	}
//	for _, player := range players {
//		serverMessage.Players = append(serverMessage.Players, models.PlayerState{
//			X:         player.X,
//			Y:         player.Y,
//			Direction: player.Direction,
//			Id:        player.Id,
//			Bullets:   player.Bullets,
//		})
//	}
//	return serverMessage
//}
//
////25
