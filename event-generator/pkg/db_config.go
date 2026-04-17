package pkg

import (
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	events "event-generator/internal/event"
	_ "time/tzdata"
)

var DB *gorm.DB

func DBInit() *events.EventRepository {
	// docker compose에서 가져올 때 환경 변수 선언
	destination := os.Getenv("DB_DSN")
	if destination == "" {
		destination = "host=localhost user=root password=1234 dbname=event port=5432 sslmode=disable TimeZone=Asia/Seoul"
	}
	slog.Info("DB Destination : ", "DSN", destination)

	// 데이터베이스 연결
	var err error

	slog.Info("Start Connection")
	for i := 0; i < 10; i++ {
		slog.Info("Try Connection... #", "attempt", i+1)
		DB, err = gorm.Open(postgres.Open(destination), &gorm.Config{})

		if err == nil {
			break
		}

		slog.Warn("DB Connection failed, Restart... ", "error", err)
		time.Sleep(time.Second)
	}

	if err != nil || DB == nil {
		slog.Error("Eventually DB Connection failed.")
		os.Exit(1)
	}
	slog.Info("Event DB Connected! Start migrate")

	// DB 테이블 생성
	if err := DB.AutoMigrate(&events.Event{}); err != nil {
		slog.Error("Event DB Migration Failed!", "error", err)
		os.Exit(1)
	}
	slog.Info("Event DB Migration Success!")

	// Event Repository 생성
	repo := &events.EventRepository{GormRepository: events.GormRepository[events.Event]{DB: DB}}
	slog.Info("Event Repository is ready: ", "DB Name", repo.DB.Name())

	return repo
}
