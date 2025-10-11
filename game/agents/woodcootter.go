package agents

import (
	"artifacts/client"
	"artifacts/state"
	"artifacts/utils"
	"context"
	"log"
)

type WoodcutterAgent struct {
	client    *client.ArtifactsMMOClient
	ctx       context.Context
	cancel    context.CancelFunc
	cooldowns *utils.CooldownManager
	running   bool
}

func NewWoodcutterAgent(c *client.ArtifactsMMOClient) *WoodcutterAgent {
	ctx, cancel := context.WithCancel(context.Background())
	return &WoodcutterAgent{
		client:    c,
		ctx:       ctx,
		cancel:    cancel,
		cooldowns: utils.NewCooldownManager(),
	}
}

func (ag *WoodcutterAgent) Stop() {
	if !ag.running {
		return
	}
	ag.cancel()
	ag.running = false
	log.Printf("[MinerAgent] stopped")
}

func (ag *WoodcutterAgent) Run(character *state.Character) {
	if ag.running {
		return
	}
	ag.running = true
	log.Printf("[%s] started woodcutting loop", character.Name)

	for {
		select {
		case <-ag.ctx.Done():
			log.Printf("[%s] stopped", character.Name)
			return
		default:
			ag.process(character)
		}
	}
}

func (ag *WoodcutterAgent) process(character *state.Character) {
	_ = ag.cooldowns.WaitTo(ag.ctx, character.CooldownExpiration, character.Name)
	for {
		if character.GetInventoryFillLevel() > 0.3 {
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
			if state.GameStateData.Resources[i].Skill == "woodcutting" && state.GameStateData.Resources[i].Level <= character.WoodcuttingLevel {
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
