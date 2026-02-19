// Package wallpapercalc предоставляет функции для расчёта обоев.
package wallpapercalc

import (
	"fmt"
	"math"
)

// Perimeter вычисляет периметр комнаты.
func Perimeter(length, width float64) (float64, error) {
	if length <= 0 || width <= 0 {
		return 0, fmt.Errorf("length and width must be positive")
	}
	return 2 * (length + width), nil
}

// RollsNeeded рассчитывает количество рулонов.
func RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error) {
	if perimeter <= 0 || height <= 0 || rollWidth <= 0 || rollLength <= 0 {
		return 0, fmt.Errorf("all values must be positive")
	}

	wallArea := perimeter * height
	rollArea := rollWidth * rollLength

	rolls := math.Ceil(wallArea / rollArea)
	return int(rolls), nil
}

// AddExtraRolls добавляет рулоны через указатель.
func AddExtraRolls(rolls *int, extra int) error {
	if rolls == nil {
		return fmt.Errorf("nil pointer")
	}
	if extra < 0 {
		return fmt.Errorf("extra cannot be negative")
	}

	*rolls += extra
	return nil
}

// FormatWallpaperReport формирует отчёт.
func FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error) {
	if room == "" {
		return "", fmt.Errorf("room cannot be empty")
	}
	if rolls <= 0 || rollPrice <= 0 {
		return "", fmt.Errorf("invalid rolls or price")
	}

	total := float64(rolls) * rollPrice

	report := fmt.Sprintf(
		"Room: %s\nRolls: %d\nPrice per roll: %.2f\nTotal: %.2f\n",
		room, rolls, rollPrice, total,
	)

	return report, nil
}
