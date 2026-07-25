package daysteps

import (
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
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		log.Printf("ожидалось элементов: 2, получено %d. Входные данные: %q", len(parts), data)
		return 0, 0, fmt.Errorf("ожидалось элементов: 2, получено: %d", len(parts))
	}

	numSteps, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("не удалось преобразовать %s в число шагов: %v", parts[0], err)
		return 0, 0, fmt.Errorf("не удалось преобразовать %s в число шагов: %v", parts[0], err)
	}
	if numSteps <= 0 {
		log.Printf("количество шагов должно быть положительным, получено: %d", numSteps)
		return 0, 0, fmt.Errorf("количество шагов должно быть положительным, получено: %d", numSteps)
	}

	durationWalk, err := time.ParseDuration(parts[1])
	if err != nil {
		log.Printf("не удалось преобразовать %s в длительность: %v", parts[1], err)
		return 0, 0, fmt.Errorf("не удалось преобразовать %s в длительность: %v", parts[1], err)

	}
	if durationWalk <= 0 {
		log.Printf("продолжительность должна быть больше нуля, получено: %d", durationWalk)
		return 0, 0, fmt.Errorf("продолжительность должна быть больше нуля, получено: %d", durationWalk)
	}

	return numSteps, durationWalk, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	sumSteps, durationWalk, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}
	if sumSteps <= 0 {
		return ""
	}

	distance := (float64(sumSteps) * stepLength) / float64(mInKm)

	caloriesSpent, err := spentcalories.WalkingSpentCalories(sumSteps, weight, height, durationWalk)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		sumSteps, distance, caloriesSpent)

}
