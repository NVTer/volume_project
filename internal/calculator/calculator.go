// Package calculator has code, which describe
// volume calculator.
package calculator

import (
	"github.com/NVTer/volume_project/internal/builder"
	"github.com/shopspring/decimal"
)

// Calculate is the implementation of counting the volume of all shapes.
func Calculate(shape string, radius, length, width, hieght decimal.Decimal) (decimal.Decimal, error) {
	b := builder.NewBuilder(shape, radius, length, width, hieght)
	sh, err := b.BuildShape()
	if err != nil {
		return decimal.Decimal{}, err
	}
	return sh.CalculateVolume(), nil
}
