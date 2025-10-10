package utils

import (
	"context"
	"log"
	"time"
)

type CooldownManager struct{}

func NewCooldownManager() *CooldownManager {
	return &CooldownManager{}
}

func (cm *CooldownManager) Wait(ctx context.Context, seconds int, who string) error {
	if seconds <= 0 {
		return nil
	}
	log.Printf("[%s] cooldown: waiting %ds", who, seconds)

	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		log.Printf("[%s] cooldown interrupted (shutdown)", who)
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (cm *CooldownManager) WaitTo(ctx context.Context, until time.Time, who string) error {
	duration := time.Until(until.Add(time.Second))
	if duration <= 0 {
		return nil
	}

	log.Printf("[%s] cooldown: waiting until %s (%v left)", who, until.Format(time.RFC3339), duration)

	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		log.Printf("[%s] cooldown interrupted (shutdown)", who)
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
