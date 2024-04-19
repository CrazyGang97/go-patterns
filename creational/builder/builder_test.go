package builder

import (
	"testing"
)

func TestNewBuilder(t *testing.T) {
	assembly := NewBuilder().Paint(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}
