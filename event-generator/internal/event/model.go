package event

import (
	"time"
)

type Event struct {
	EventID   string                 `gorm:"primaryKey" json:"event_id"`
	UserID    int                    `json:"user_id"`
	EventType string                 `json:"event_type"`
	CreatedAt time.Time              `json:"created_at"`
	Metadata  map[string]interface{} `gorm:"type:jsonb" json:"metadata"`
}

// 조회, 단순 클릭, 구매, 에러
var eventTypes = []string{"page_view", "click", "purchase", "error"}
