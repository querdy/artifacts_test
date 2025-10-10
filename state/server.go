package state

import "time"

type Account struct {
	Username           string        `json:"username"`
	Email              string        `json:"email"`
	Member             bool          `json:"member"`
	MemberExpiration   *string       `json:"member_expiration"`
	Status             string        `json:"status"`
	Badges             []interface{} `json:"badges"`
	Skins              []interface{} `json:"skins"`
	Gems               int           `json:"gems"`
	EventToken         int           `json:"event_token"`
	AchievementsPoints int           `json:"achievements_points"`
	Banned             bool          `json:"banned"`
	BanReason          string        `json:"ban_reason"`
}

type ServerStatus struct {
	Version         string    `json:"version"`
	ServerTime      time.Time `json:"server_time"`
	MaxLvl          int       `json:"max_lvl"`
	MaxSkillLvl     int       `json:"max_skill_lvl"`
	CharacterOnline int       `json:"character_online"`
}
