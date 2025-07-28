package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65
)

func parsePackage(data string) (int, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, fmt.Errorf("the number of parameters must be 2")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("сan not convert string into integer")
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("can not use negative value of steps")
	}

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, errors.New("сan not convert string into time.Duration")
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("can not use negative value of duration")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceKm := (StepLength * float64(steps)) / 1000
	kalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	final := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKm, kalories)

	return final
}
