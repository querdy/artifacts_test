package agents

import (
	"artifacts/client"
	"artifacts/state"
	"artifacts/utils"
	"context"
	"log"
)

type BaseAgent struct {
	client      *client.ArtifactsMMOClient
	ctx         context.Context
	cancel      context.CancelFunc
	cooldowns   *utils.CooldownManager
	running     bool
	processFunc func(*state.Character)
}

func NewBaseAgent(c *client.ArtifactsMMOClient) *BaseAgent {
	ctx, cancel := context.WithCancel(context.Background())
	return &BaseAgent{
		client:    c,
		ctx:       ctx,
		cancel:    cancel,
		cooldowns: utils.NewCooldownManager(),
	}
}

func (ag *BaseAgent) Stop() {
	if !ag.running {
		return
	}
	ag.cancel()
	ag.running = false
}

func (ag *BaseAgent) Run(character *state.Character) {
	if ag.running {
		log.Println("Already running")
		return
	}
	ag.running = true
	for {
		select {
		case <-ag.ctx.Done():
			log.Printf("[%s] stopped", character.Name)
			return
		default:
			ag.processFunc(character)
		}
	}
}

func (ag *BaseAgent) process(character *state.Character) {}
