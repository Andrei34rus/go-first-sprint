package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err //удален лополнительный контекст ошибки
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err //удален лополнительный контекст ошибки
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("длительность не может быть отрицательной")
	}
	return steps, duration, nil
	// TODO: реализовать функцию
}

func DayActionInfo(data string, weight, height float64) string {
	if data == "" {
		log.Println("Ошибка: пустые данные")
		return ""
	}
	if weight <= 0 || height <= 0 {
		log.Println("Ошибка: некорректные параметры веса или роста")
		return ""
	}

	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Ошибка парсинга: %v", err)
		return ""
	}

	distanseMeters := float64(steps) * stepLength
	distanceKm := distanseMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Ошибка расчёта калорий: %v", err) //Исправлено:добавлено логирование конкретной ошибки
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
	return result
	// TODO: реализовать функцию
}
