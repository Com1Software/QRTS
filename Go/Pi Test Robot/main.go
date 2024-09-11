package main

import (
	"log"
	"time"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)

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
	servo0 := pca0.ServoNew(0, nil)
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

	// Angle in degrees. Must be in the range `0` to `Range`
	for i := 0; i < 130; i++ {
		if leg == 1 || leg == 0 {
			servo0.Angle(i)
			servo1.Angle(i)
			servo2.Angle(i)
		}
		if leg == 2 || leg == 0 {
			servo3.Angle(i)
			servo4.Angle(i)
			servo5.Angle(i)
		}
		if leg == 3 || leg == 0 {
			servo6.Angle(i)
			servo7.Angle(i)
			servo8.Angle(i)
		}
		if leg == 4 || leg == 0 {
			servo9.Angle(i)
			servo10.Angle(i)
			servo11.Angle(i)
		}

		time.Sleep(10 * time.Millisecond)
	}

	// Fraction as pulse width expressed between 0.0 `MinPulse` and 1.0 `MaxPulse`
	servo0.Fraction(0.5)
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

}
