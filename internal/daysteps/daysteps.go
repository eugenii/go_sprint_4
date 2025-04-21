package daysteps

import (
	"fmt"
	sc "go_sprint_4/internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parsedData := strings.Split(data, ",")
	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return 0, 0, err
	}
	if len(parsedData) != 3 {
		return 0, 0, fmt.Errorf("invalid data format")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return err.Error()
	}
	distance := float64(steps) * stepLength / mInKm
	calories, err := sc.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
}
