package main

import (
	"github.com/saintbyte/TimeOfDay"
	"log/slog"
	"time"
)

func main() {
	currentTime := time.Now()
	partOfDay := TimeOfDay.Parse(currentTime)
	slog.Info("Current time:", currentTime)
	if partOfDay.PartOfDay1 == TimeOfDay.Day {
		slog.Info("Добрый день")
	} else {
		slog.Info("Доброй ночи")
	}
}
