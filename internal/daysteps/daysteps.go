package daysteps

import (
	"errors"
	"fmt"
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
		return 0, 0, errors.New("ошибка")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, errors.New("ошибка")
	}

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, errors.New("ошибка")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceKm := (StepLength * float64(steps)) / 1000
	kalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	final := fmt.Sprintf("Количество шагов: %d./nДистанция составила %.2f км./nВы сожгли %.2f ккал.", steps, distanceKm, kalories)

	return final
}
