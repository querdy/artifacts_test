package game

import (
	"artifacts/client"
	"artifacts/game/agents"
	"artifacts/state"
)

func CharactersController(ac *client.ArtifactsMMOClient) {
	damager := state.GameStateData.GetCharacterByName("Marnie")
	damagerAgent := agents.NewMinerAgent(ac)
	miner := state.GameStateData.GetCharacterByName("Milva")
	minerAgent := agents.NewMinerAgent(ac)
	woodcutter := state.GameStateData.GetCharacterByName("Beevee")
	woodcutterAgent := agents.NewWoodcutterAgent(ac)
	alchemist := state.GameStateData.GetCharacterByName("Veebee")
	alchemistAgent := agents.NewWoodcutterAgent(ac)
	chief := state.GameStateData.GetCharacterByName("Caleb")
	chiefAgent := agents.NewMinerAgent(ac)
	go damagerAgent.Run(damager)
	go minerAgent.Run(miner)
	go woodcutterAgent.Run(woodcutter)
	go alchemistAgent.Run(alchemist)
	go chiefAgent.Run(chief)
}
