package main

import (
	"log"
	"net/http"
	"robot/html"
	"robot/model"
	"robot/servo"
	"robot/utils"
	"runtime"

	"fmt"
	"os"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)

// ----------------------------------------------------------------
func main() {
	fmt.Println("Pi Test Robot")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	xip := fmt.Sprintf("%s", utils.GetOutboundIP())
	port := "8080"
	i2c, err := i2c.New(pca9685.Address, "/dev/i2c-1")
	if err != nil {
		log.Fatal(err)
	}
	pca0, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Fatal(err)
	}
	pca0.SetChannel(0, 0, 130)
	sa := []model.ServoArray{}
	saa := model.ServoArray{}
	sa = append(sa, saa)
	sa[0].Servo1 = pca0.ServoNew(1, nil)
	sa[0].Servo2 = pca0.ServoNew(2, nil)
	sa[0].Servo3 = pca0.ServoNew(3, nil)
	sa[0].Servo4 = pca0.ServoNew(4, nil)
	sa[0].Servo5 = pca0.ServoNew(5, nil)
	sa[0].Servo6 = pca0.ServoNew(6, nil)
	sa[0].Servo7 = pca0.ServoNew(7, nil)
	sa[0].Servo8 = pca0.ServoNew(8, nil)
	sa[0].Servo9 = pca0.ServoNew(9, nil)
	sa[0].Servo10 = pca0.ServoNew(10, nil)
	sa[0].Servo11 = pca0.ServoNew(11, nil)
	sa[0].Servo12 = pca0.ServoNew(12, nil)

	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:

		fmt.Println("Not")

		//-------------------------------------------------------------
	default:

		fmt.Println("Server running....")
		fmt.Println("Listening on " + xip + ":" + port)

		fmt.Println("")
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			xdata := html.InitPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page Handler
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := html.AboutPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Test Page Handler
		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			xdata := html.TestPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Servo Test
		http.HandleFunc("/servotest1", func(w http.ResponseWriter, r *http.Request) {
			sa[0] = servo.SetAllServoPositions(sa[0], 65)
			xdata := html.ServoTest(xip, sa[0])
			fmt.Fprint(w, xdata)
		})
		// ------------------------------------------------ Servo Test
		http.HandleFunc("/servotest2", func(w http.ResponseWriter, r *http.Request) {
			sa[0] = servo.SetLevelServoPositions(sa[0], 115, 3)
			xdata := html.ServoTest(xip, sa[0])
			fmt.Fprint(w, xdata)
		})
		// ------------------------------------------------ Servo Test
		http.HandleFunc("/servotest3", func(w http.ResponseWriter, r *http.Request) {
			sa[0] = servo.MoveLevelServoPositionsDown(sa[0], 2, 65, 0, 100)
			xdata := html.ServoTest(xip, sa[0])
			fmt.Fprint(w, xdata)
		})
		// ------------------------------------------------ Servo Test
		http.HandleFunc("/servotest4", func(w http.ResponseWriter, r *http.Request) {
			sa[0] = servo.MoveLevelServoPositionsDown(sa[0], 1, 65, 40, 100)
			xdata := html.ServoTest(xip, sa[0])
			fmt.Fprint(w, xdata)
		})
		// ------------------------------------------------ Servo Test
		http.HandleFunc("/servotest5", func(w http.ResponseWriter, r *http.Request) {
			sa[0] = servo.MoveLevelServoPositionsUp(sa[0], 2, 65, 130, 100)
			xdata := html.ServoTest(xip, sa[0])
			fmt.Fprint(w, xdata)
		})
		// ------------------------------------------------ Servo Test
		http.HandleFunc("/servotest6", func(w http.ResponseWriter, r *http.Request) {
			sa[0] = servo.MoveLevelServoPositionsUp(sa[0], 1, 65, 130, 100)
			xdata := html.ServoTest(xip, sa[0])
			fmt.Fprint(w, xdata)
		})

		//------------------------------------------------- Start Server
		utils.Openbrowser(xip + ":" + port)
		if err := http.ListenAndServe(xip+":"+port, nil); err != nil {
			panic(err)
		}
	}
	sa[0].Servo1.Fraction(0.5)
	sa[0].Servo2.Fraction(0.5)
	sa[0].Servo3.Fraction(0.5)
	sa[0].Servo4.Fraction(0.5)
	sa[0].Servo5.Fraction(0.5)
	sa[0].Servo6.Fraction(0.5)
	sa[0].Servo7.Fraction(0.5)
	sa[0].Servo8.Fraction(0.5)
	sa[0].Servo9.Fraction(0.5)
	sa[0].Servo10.Fraction(0.5)
	sa[0].Servo11.Fraction(0.5)
	sa[0].Servo12.Fraction(0.5)

}
