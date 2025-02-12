package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, errors.New("ошибка")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, errors.New("ошибка")
	}

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, errors.New("ошибка")
	}

	return steps, slice[1], duration, nil
}

func distance(steps int) float64 {
	return (float64(steps) * lenStep) / mInKm
}

func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := distance(steps)

	meanSp := distance / duration.Hours()

	return meanSp
}

func TrainingInfo(data string, weight, height float64) string {
	steps, typeTr, duration, err := parseTraining(data)
	if err != nil {
		return ""
	}

	switch {
	case typeTr == "Ходьба":
		final := fmt.Sprintf("Тип тренировки: %s/nДлительность: %.2f ч./nДистанция: %.2f км./nСкорость: %.2f км/ч/nСожгли калорий: %.2f ", typeTr, duration.Hours(), distance(steps), meanSpeed(steps, duration), WalkingSpentCalories(steps, weight, height, duration))
		return final
	case typeTr == "Бег":
		final := fmt.Sprintf("Тип тренировки: %s/nДлительность: %.2f ч./nДистанция: %.2f км./nСкорость: %.2f км/ч/nСожгли калорий: %.2f ", typeTr, duration.Hours(), distance(steps), meanSpeed(steps, duration), RunningSpentCalories(steps, weight, duration))
		return final
	}
	return "неизвестный тип тренировки"
}

const (
	runningCaloriesMeanSpeedMultiplier = 18.0
	runningCaloriesMeanSpeedShift      = 20.0
)

func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	meanSp := meanSpeed(steps, duration)

	return ((runningCaloriesMeanSpeedMultiplier * meanSp) - runningCaloriesMeanSpeedShift) * weight
}

const (
	walkingCaloriesWeightMultiplier = 0.035
	walkingSpeedHeightMultiplier    = 0.029
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	meanSp := meanSpeed(steps, duration)

	return ((walkingCaloriesWeightMultiplier * weight) + (meanSp*meanSp/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}
