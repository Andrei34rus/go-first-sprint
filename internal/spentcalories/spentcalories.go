package spentcalories

import (
	"errors"
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
	parts := strings.Split(data, " ")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("неверный формат: требуется 3 параметра")
	}
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, errors.New("неверный формат: " + err.Error())
	}
	activity := parts[1]
	if activity == "" {
		return 0, "", 0, errors.New("вид активности не может быть пустым")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, errors.New("неверный формат длительности: " + err.Error())
	}
	return count, activity, duration, nil
	// TODO: реализовать функцию
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanseMeters := float64(steps) * stepLength
	distanceKm := distanseMeters / mInKm
	return distanceKm
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	hours := duration.Hours()
	speed := dist / hours
	return speed
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	var (
		activityName string
		calories     float64
	)

	switch activityType {
	case "Бег", "бег":
		activityName = "Бег"
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба", "ходьба":
		activityName = "Ходьба"
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", errors.New("ошибка расчёта калорий")
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожжено калорий: %.2f",
		activityName,
		duration.Hours(),
		dist,
		speed,
		calories,
	)
	return result, nil
}

// TODO: реализовать функцию

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}
	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Hours()
	calories := (weight * speed * durationMinutes) / minInH
	return calories, nil
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}
	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	baseCalories := (weight * speed * durationMinutes) / minInH
	calories := baseCalories * walkingCaloriesCoefficient
	return calories, nil
	// TODO: реализовать функцию
}
