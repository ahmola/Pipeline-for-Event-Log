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

// 조회, 구매, 환불, 에러
var eventTypes = []string{"page_view", "purchase", "refund", "error"}
