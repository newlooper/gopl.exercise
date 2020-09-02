package convlib

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToM converts a Feet to Meter
func FToM(ft Feet) Meter { return Meter(ft * 0.3048) }

// MToF converts a Meter to Feet
func MToF(m Meter) Feet { return Feet(m * 3.2808) }
