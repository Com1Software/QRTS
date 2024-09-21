package main

import (
	"log"
	"time"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)
type ServoArray struct {
	Servo1   int
	Servo2   int
	Servo3   int
	Servo4   int
	Servo5   int
	Servo6   int
	Servo7   int
	Servo8   int
	Servo9   int
	Servo10   int
	Servo11   int
	Servo12   int
}

// ----------------------------------------- Main
func main() {
	i2c, err := i2c.New(pca9685.Address, "/dev/i2c-1")
	if err != nil {
		log.Fatal(err)
	}
	pca0, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Fatal(err)
	}
	pca0.SetChannel(0, 0, 130)
	leg := 1
	servo1 := pca0.ServoNew(1, nil)
	servo2 := pca0.ServoNew(2, nil)
	servo3 := pca0.ServoNew(3, nil)
	servo4 := pca0.ServoNew(4, nil)
	servo5 := pca0.ServoNew(5, nil)
	servo6 := pca0.ServoNew(6, nil)
	servo7 := pca0.ServoNew(7, nil)
	servo8 := pca0.ServoNew(8, nil)
	servo9 := pca0.ServoNew(9, nil)
	servo10 := pca0.ServoNew(10, nil)
	servo11 := pca0.ServoNew(11, nil)
	servo12 := pca0.ServoNew(12, nil)

	// Angle in degrees. Must be in the range `0` to `Range`
	for i := 0; i < 130; i++ {
		if leg == 1 || leg == 0 {
			servo1.Angle(i)
			servo2.Angle(i)
			servo3.Angle(i)
		}
		if leg == 2 || leg == 0 {
			servo4.Angle(i)
			servo5.Angle(i)
			servo6.Angle(i)
		}
		if leg == 3 || leg == 0 {
			servo7.Angle(i)
			servo8.Angle(i)
			servo9.Angle(i)
		}
		if leg == 4 || leg == 0 {
			servo10.Angle(i)
			servo11.Angle(i)
			servo12.Angle(i)
		}

		time.Sleep(10 * time.Millisecond)
	}

	// Fraction as pulse width expressed between 0.0 `MinPulse` and 1.0 `MaxPulse`

	servo1.Fraction(0.5)
	servo2.Fraction(0.5)
	servo3.Fraction(0.5)
	servo4.Fraction(0.5)
	servo5.Fraction(0.5)
	servo6.Fraction(0.5)
	servo7.Fraction(0.5)
	servo8.Fraction(0.5)
	servo9.Fraction(0.5)
	servo10.Fraction(0.5)
	servo11.Fraction(0.5)
	servo12.Fraction(0.5)
}
