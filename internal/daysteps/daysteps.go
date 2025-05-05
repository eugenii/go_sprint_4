package daysteps

import (
	"fmt"
	sc "go_sprint_4/internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	stepLength = 0.65
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format")
	}
	if strings.Contains(parts[0], " ") {
		return 0, 0, fmt.Errorf("invalid steps format")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format")
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("invalid steps value")
	}

	durStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format")
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("invalid duration value")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Print(err) // Логируем ошибку как ожидают тесты
		return ""
	}

	if weight <= 0 || height <= 0 || weight > 200 || height > 300 {
		log.Print("invalid weight or height value")
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	calories, err := sc.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Print(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, distance, calories)
}
