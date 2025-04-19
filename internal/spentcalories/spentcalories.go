package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parsedData := strings.Split(data, ",")
	if len(parsedData) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("steps must be positive")
	}

	activity := parsedData[1]

	duration, err := time.ParseDuration(parsedData[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format")
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("duration must be positive")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration.Hours() <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input data")
	}
	if weight <= 0 || height <= 0 || weight > 200 || height > 300 {
		return 0, fmt.Errorf("invalid input data")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * speed * weight) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || duration.Minutes() <= 0 {
		return 0, fmt.Errorf("invalid input data")
	}
	if weight <= 0 || height <= 0 || weight > 200 || height > 300 {
		return 0, fmt.Errorf("invalid input data")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * speed * weight) / minInH * walkingCaloriesCoefficient
	return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	if weight <= 0 || height <= 0 || weight > 200 || height > 300 {
		return "", fmt.Errorf("invalid input data")
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	hours := duration.Hours()

	var calories float64
	var calcErr error

	switch activity {
	case "Ходьба":
		calories, calcErr = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, calcErr = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "неизвестный тип тренировки", fmt.Errorf("неизвестный тип тренировки")
	}

	if calcErr != nil {
		return "", calcErr
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.3f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f",
		activity,
		hours,
		dist,
		speed,
		calories,
	)

	return info, nil
}
