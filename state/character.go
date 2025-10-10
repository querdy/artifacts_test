package state

import (
	"time"
)

type Effect struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type Character struct {
	Name                 string          `json:"name"`
	Account              string          `json:"account"`
	Skin                 string          `json:"skin"`
	Level                int             `json:"level"`
	Xp                   int             `json:"xp"`
	MaxXp                int             `json:"max_xp"`
	Gold                 int             `json:"gold"`
	Speed                int             `json:"speed"`
	MiningLevel          int             `json:"mining_level"`
	MiningXp             int             `json:"mining_xp"`
	MiningMaxXp          int             `json:"mining_max_xp"`
	WoodcuttingLevel     int             `json:"woodcutting_level"`
	WoodcuttingXp        int             `json:"woodcutting_xp"`
	WoodcuttingMaxXp     int             `json:"woodcutting_max_xp"`
	FishingLevel         int             `json:"fishing_level"`
	FishingXp            int             `json:"fishing_xp"`
	FishingMaxXp         int             `json:"fishing_max_xp"`
	WeaponcraftingLevel  int             `json:"weaponcrafting_level"`
	WeaponcraftingXp     int             `json:"weaponcrafting_xp"`
	WeaponcraftingMaxXp  int             `json:"weaponcrafting_max_xp"`
	GearcraftingLevel    int             `json:"gearcrafting_level"`
	GearcraftingXp       int             `json:"gearcrafting_xp"`
	GearcraftingMaxXp    int             `json:"gearcrafting_max_xp"`
	JewelrycraftingLevel int             `json:"jewelrycrafting_level"`
	JewelrycraftingXp    int             `json:"jewelrycrafting_xp"`
	JewelrycraftingMaxXp int             `json:"jewelrycrafting_max_xp"`
	CookingLevel         int             `json:"cooking_level"`
	CookingXp            int             `json:"cooking_xp"`
	CookingMaxXp         int             `json:"cooking_max_xp"`
	AlchemyLevel         int             `json:"alchemy_level"`
	AlchemyXp            int             `json:"alchemy_xp"`
	AlchemyMaxXp         int             `json:"alchemy_max_xp"`
	Hp                   int             `json:"hp"`
	MaxHp                int             `json:"max_hp"`
	Haste                int             `json:"haste"`
	CriticalStrike       int             `json:"critical_strike"`
	Wisdom               int             `json:"wisdom"`
	Prospecting          int             `json:"prospecting"`
	Initiative           int             `json:"initiative"`
	Threat               int             `json:"threat"`
	AttackFire           int             `json:"attack_fire"`
	AttackEarth          int             `json:"attack_earth"`
	AttackWater          int             `json:"attack_water"`
	AttackAir            int             `json:"attack_air"`
	Dmg                  int             `json:"dmg"`
	DmgFire              int             `json:"dmg_fire"`
	DmgEarth             int             `json:"dmg_earth"`
	DmgWater             int             `json:"dmg_water"`
	DmgAir               int             `json:"dmg_air"`
	ResFire              int             `json:"res_fire"`
	ResEarth             int             `json:"res_earth"`
	ResWater             int             `json:"res_water"`
	ResAir               int             `json:"res_air"`
	Effects              []Effect        `json:"effects"`
	X                    int             `json:"x"`
	Y                    int             `json:"y"`
	Layer                string          `json:"layer"`
	MapId                int             `json:"map_id"`
	Cooldown             int             `json:"cooldown"`
	CooldownExpiration   time.Time       `json:"cooldown_expiration"`
	WeaponSlot           string          `json:"weapon_slot"`
	RuneSlot             string          `json:"rune_slot"`
	ShieldSlot           string          `json:"shield_slot"`
	HelmetSlot           string          `json:"helmet_slot"`
	BodyArmorSlot        string          `json:"body_armor_slot"`
	LegArmorSlot         string          `json:"leg_armor_slot"`
	BootsSlot            string          `json:"boots_slot"`
	Ring1Slot            string          `json:"ring1_slot"`
	Ring2Slot            string          `json:"ring2_slot"`
	AmuletSlot           string          `json:"amulet_slot"`
	Artifact1Slot        string          `json:"artifact1_slot"`
	Artifact2Slot        string          `json:"artifact2_slot"`
	Artifact3Slot        string          `json:"artifact3_slot"`
	Utility1Slot         string          `json:"utility1_slot"`
	Utility1SlotQuantity int             `json:"utility1_slot_quantity"`
	Utility2Slot         string          `json:"utility2_slot"`
	Utility2SlotQuantity int             `json:"utility2_slot_quantity"`
	BagSlot              string          `json:"bag_slot"`
	Task                 string          `json:"task"`
	TaskType             string          `json:"task_type"`
	TaskProgress         int             `json:"task_progress"`
	TaskTotal            int             `json:"task_total"`
	InventoryMaxItems    int             `json:"inventory_max_items"`
	Inventory            []InventoryItem `json:"inventory"`
}

func (c *Character) GetInventoryCount() int {
	counter := 0
	for i := range c.Inventory {
		counter += c.Inventory[i].Quantity
	}
	return counter
}

func (c *Character) GetInventoryFillLevel() float64 {
	if c.InventoryMaxItems == 0 {
		return float64(0)
	}
	return float64(c.GetInventoryCount()) / float64(c.InventoryMaxItems)
}

func (c *Character) GetInventoryItems() []Item {
	var items []Item
	for _, invItem := range c.Inventory {
		if invItem.Code != "" && invItem.Quantity != 0 {
			items = append(items, Item{
				Code:     invItem.Code,
				Quantity: invItem.Quantity,
			})
		}
	}
	return items
}
