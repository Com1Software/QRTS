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

	// Sets a single PWM channel 0
	//	pca0.SetChannel(1, 0, 130)
	pca0.SetChannel(0, 0, 130)
	//	pca0.SetChannel(1, 2, 130)

	// Servo on channel 0
	servo0 := pca0.ServoNew(0, nil)
	servo1 := pca0.ServoNew(1, nil)
	servo2 := pca0.ServoNew(2, nil)
	// Angle in degrees. Must be in the range `0` to `Range`
	for i := 0; i < 130; i++ {
		servo0.Angle(i)
		servo1.Angle(i)
		servo2.Angle(i)

		time.Sleep(10 * time.Millisecond)
	}

	// Fraction as pulse width expressed between 0.0 `MinPulse` and 1.0 `MaxPulse`
	servo0.Fraction(0.5)
}
