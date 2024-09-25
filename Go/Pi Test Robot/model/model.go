package model

import "github.com/googolgl/go-pca9685"

//------------------------------------------ Structs

type ServoArray struct {
	Servo1     *pca9685.Servo
	Servo1pos  int
	Servo2     *pca9685.Servo
	Servo2pos  int
	Servo3     *pca9685.Servo
	Servo3pos  int
	Servo4     *pca9685.Servo
	Servo4pos  int
	Servo5     *pca9685.Servo
	Servo5pos  int
	Servo6     *pca9685.Servo
	Servo6pos  int
	Servo7     *pca9685.Servo
	Servo7pos  int
	Servo8     *pca9685.Servo
	Servo8pos  int
	Servo9     *pca9685.Servo
	Servo9pos  int
	Servo10    *pca9685.Servo
	Servo10pos int
	Servo11    *pca9685.Servo
	Servo11pos int
	Servo12    *pca9685.Servo
	Servo12pos int
}
