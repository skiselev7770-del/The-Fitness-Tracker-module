package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	// lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("ожидалось элементов: 3, получено: %d", len(parts))
	}
	numSteps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("не удалось преобразовать %s в число шагов: %v", parts[0], err)
	}
	if numSteps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным, получено: %d", numSteps)
	}

	durationWalk, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("не удалось преобразовать %s в длительность: %v", parts[1], err)
	}
	if durationWalk <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность не может быть меньше нуля, получено: %v", durationWalk)
	}

	return numSteps, parts[1], durationWalk, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}

	calculatedDistance := distance(steps, height)
	hours := duration.Hours()

	return calculatedDistance / hours

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	numSteps, typeActivity, durationTraining, err := parseTraining(data)
	if err != nil {
		log.Println("ошибка обработки данных тренировки:", err)
		return "", err
	}
	if numSteps <= 0 {
		log.Println("количество шагов должно быть положительным, получено:", numSteps)
		return "", err
	}
	if durationTraining <= 0 {
		log.Println("длительность тренировки должна быть больше нуля, получено:", durationTraining)
		return "", err
	}
	if weight <= 0 || height <= 0 {
		return "", fmt.Errorf("вес и рост не могут быть меньше нуля, получено: %.2f и %.2f", weight, height)
	}

	var calculatedDistance float64
	var calculatedMeanSpeed float64
	var calculatedCalories float64

	switch typeActivity {
	case "Ходьба":
		calculatedDistance = distance(numSteps, height)
		calculatedMeanSpeed = meanSpeed(numSteps, height, durationTraining)
		calculatedCalories, err = WalkingSpentCalories(numSteps, weight, height, durationTraining)
		if err != nil {
			return "", fmt.Errorf("Ошибка: %v", err)
		}

		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeActivity,
			durationTraining.Hours(),
			calculatedDistance,
			calculatedMeanSpeed,
			calculatedCalories), nil

	case "Бег":
		calculatedDistance = distance(numSteps, height)
		calculatedMeanSpeed = meanSpeed(numSteps, height, durationTraining)
		calculatedCalories, err = RunningSpentCalories(numSteps, weight, height, durationTraining)
		if err != nil {
			return "", fmt.Errorf("Ошибка: %v", err)
		}

		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeActivity,
			durationTraining.Hours(),
			calculatedDistance,
			calculatedMeanSpeed,
			calculatedCalories), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0, fmt.Errorf("недопустимые входные данные: steps = %d, duration = %v", steps, duration)
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("вес и рост не могут быть меньше нуля, получено: %.2f и %.2f", weight, height)
	}

	calculatedMeanSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	countedCalories := (weight * calculatedMeanSpeed * minutes) / minInH

	return countedCalories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0, fmt.Errorf("недопустимые входные данные: steps = %d, duration = %v", steps, duration)
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("вес и рост не могут быть меньше нуля, получено: %.2f и %.2f", weight, height)
	}

	calculatedMeanSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	countedCalories := ((weight * calculatedMeanSpeed * minutes) / minInH) * walkingCaloriesCoefficient

	return countedCalories, nil

}
