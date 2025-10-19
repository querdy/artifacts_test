package agents

import (
	"artifacts/client"
	"artifacts/state"
	"log"
)

type WoodcutterAgent struct {
	*BaseAgent
}

func NewWoodcutterAgent(c *client.ArtifactsMMOClient) *WoodcutterAgent {
	ba := NewBaseAgent(c)
	ag := &WoodcutterAgent{BaseAgent: NewBaseAgent(c)}
	ba.processFunc = ag.process
	return &WoodcutterAgent{BaseAgent: ba}
}

func (ag *WoodcutterAgent) process(character *state.Character) {
	log.Printf("[%s] started woodcutting loop", character.Name)
	_ = ag.cooldowns.WaitTo(ag.ctx, character.CooldownExpiration, character.Name)
	for {
		if character.GetInventoryFillLevel() > 0.5 {
			bank := state.GameStateData.GetNearestMapByContentCode(character.MapId, "bank")
			if bank.MapID != character.MapId {
				movementData, err := ag.client.ActionMove(character, bank)
				if err != nil {
					return
				}
				*character = movementData.Data.Character
				_ = ag.cooldowns.Wait(ag.ctx, character.Cooldown, character.Name)
			}
			depositItemData, err := ag.client.ActionDepositItem(character, character.GetInventoryItems())
			if err != nil {
				return
			}
			*character = depositItemData.Data.Character
			_ = ag.cooldowns.Wait(ag.ctx, character.Cooldown, character.Name)
		}
		maxLvlResource := state.Resource{}
		for i := range state.GameStateData.Resources {
			if state.GameStateData.Resources[i].Skill == "woodcutting" && state.GameStateData.Resources[i].Level <= character.WoodcuttingLevel && state.GameStateData.Resources[i].Code != "magic_tree" {
				if maxLvlResource.Level < state.GameStateData.Resources[i].Level {
					maxLvlResource = state.GameStateData.Resources[i]
				}
			}
		}
		nearestMap := state.GameStateData.GetNearestMapByContentCode(character.MapId, maxLvlResource.Code)
		if nearestMap != nil && nearestMap.MapID != character.MapId {
			movementData, err := ag.client.ActionMove(character, nearestMap)
			if err != nil {
				return
			}
			*character = movementData.Data.Character
			_ = ag.cooldowns.Wait(ag.ctx, character.Cooldown, character.Name)
		}
		gatheringData, err := ag.client.ActionGathering(character)
		if err != nil {
			return
		}
		*character = gatheringData.Data.Character
		log.Printf("[%s] inventory filled to %.2f", character.Name, character.GetInventoryFillLevel())
		_ = ag.cooldowns.Wait(ag.ctx, character.Cooldown, character.Name)
	}
}
