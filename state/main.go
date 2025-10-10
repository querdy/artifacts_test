package state

import (
	"math"
)

type GameState struct {
	Server     ServerStatus
	Account    Account
	Characters []Character
	Maps       []Map
	Resources  []Resource
}

func (gs GameState) GetCharacterByName(Name string) *Character {
	for i := range gs.Characters {
		if gs.Characters[i].Name == Name {
			return &gs.Characters[i]
		}
	}
	return nil
}

func (gs GameState) CalculateDistance(startMap *Map, endMap *Map) float64 {
	//fmt.Println("a", utils.Stringify(startMap))
	//fmt.Println("b", utils.Stringify(endMap))
	deltaX := float64(endMap.X - startMap.X)
	deltaY := float64(endMap.Y - startMap.Y)

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func (gs GameState) GetMapById(id int) *Map {
	for i := range gs.Maps {
		if gs.Maps[i].MapID == id {
			return &gs.Maps[i]
		}
	}
	return nil
}

func (gs GameState) GetNearestMapByContentCode(currentMapId int, code string) *Map {
	currentMap := gs.GetMapById(currentMapId)
	var nearest *Map
	nearestDistance := math.MaxFloat64
	for i := range gs.Maps {
		if gs.Maps[i].Interactions.Content != nil && gs.Maps[i].Interactions.Content.Code == code {
			currentDistance := gs.CalculateDistance(currentMap, &gs.Maps[i])
			if currentDistance < nearestDistance {
				nearest = &gs.Maps[i]
				nearestDistance = currentDistance
			}
		}
	}
	return nearest
}

var GameStateData = GameState{}
