package main

import (
	events "event-generator/internal/event"
	pkg "event-generator/pkg"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 랜덤 이벤트 생성 스크립트
func main() {
	// db 초기화
	repo := pkg.DBInit()

	// 헬스체크 엔드포인트 생성
	slog.Info("Health Check Endpoint Init")
	healthCheck := gin.Default()

	healthCheck.GET("/health", func(c *gin.Context) {
		// db상태 확인
		db, _ := pkg.DB.DB()
		if err := db.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "unhealthy", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	})

	// 주기의 티커
	env_sec, err := strconv.ParseInt(os.Getenv("TICKER_SECONDS"), 10, 32)
	if err != nil {
		slog.Error("Failed to parse TICKER_SECONDS", "error", err)
		env_sec = 5
	}
	ticker := time.NewTicker(time.Duration(env_sec) * time.Second)
	defer ticker.Stop()
	slog.Info("Event Generator started: random events will be generated!", "Seconds", env_sec)

	var count int = 1
	// 주기마다 랜덤하게 이벤트 생성
	for range ticker.C {
		// 고루틴을 활용한 병렬적으로 이벤트 생성
		var wg sync.WaitGroup

		randomNum := rand.Intn(10) + 1
		for i := 0; i < randomNum; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				event := events.GenerateRandomEvent()
				if err := repo.Create(&event); err != nil {
					slog.Error("Failed to create event", "error", err)
					return
				}
				slog.Info("Event Created", "ID", event.EventID, "Type", event.EventType)
			}()
		}

		// 고루틴이 완료될 때까지 대기
		wg.Wait()
		slog.Info("Random Event Go Routine Done! ", "Count", count, "Generated", randomNum)
		count++
	}
}
