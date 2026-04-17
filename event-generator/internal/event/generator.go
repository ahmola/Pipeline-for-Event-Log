package event

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func generateMetadata(eType string) map[string]interface{} {
	meta := make(map[string]interface{})
	switch eType {
	case "page_view":
		pages := []string{"/home", "/about", "/contact", "/products", "/cart"}
		meta["url"] = pages[rand.Intn(len(pages))]
	case "purchase":
		meta["amount"] = rand.Intn(100000)
		meta["item_id"] = rand.Intn(500)
	case "error":
		errorTypes := []struct {
			code int
			msg  string
		}{
			{400, "bad_request"},
			{401, "unauthorized"},
			{403, "forbidden"},
			{404, "not_found"},
			{429, "too_many_requests"},
			{500, "internal_server_error"},
			{502, "bad_gateway"},
			{503, "service_unavailable"},
		}

		selected := errorTypes[rand.Intn(len(errorTypes))]
		meta["code"] = selected.code
		meta["msg"] = selected.msg
	}
	return meta
}

func GenerateRandomEvent() Event {
	eType := eventTypes[rand.Intn(len(eventTypes))]
	return Event{
		EventID:   uuid.New().String(),
		UserID:    rand.Intn(1000) + 1,
		EventType: eType,
		CreatedAt: time.Now(),
		Metadata:  generateMetadata(eType),
	}
}
