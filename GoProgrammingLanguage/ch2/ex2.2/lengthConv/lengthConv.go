package lengthConv

import "fmt"

type KiloMeter float64
type Meter float64

func (km KiloMeter) String() string {
	return fmt.Sprintf("%g Km", km)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g M", m)
}

func KmtoM(km KiloMeter) Meter {
	return Meter(km * 1000)
}

func MtoKm(m Meter) KiloMeter {
	return KiloMeter(m / 1000)
}
