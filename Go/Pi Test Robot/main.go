package main

import (
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"fmt"
	"os"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)

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

// ----------------------------------------------------------------
func ServoTest1(xip string) string {

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Servo Test 1</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: red;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>Pi Test Robot</p>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<BR>  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/QRTS'> [ QRTS GitHub Repository ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "Pi Test Robot"
	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

func InitPage(xip string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xxip := ""
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>PI Test Robot</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body>"
	xdata = xdata + "<H1>Pi Test Robot</H1>"
	//---------
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xxip = fmt.Sprintf("%s", ipv4)
		}
	}
	xdata = xdata + "<p> Host Port IP : " + xip
	xdata = xdata + "<BR> Machine IP : " + xxip + "</p>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/about'> [ About ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/test'> [ Testing ] </A><BR>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest1'> [ Servo Test 1 ] </A>  "

	xdata = xdata + "<BR><BR>Pi Test Robot"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// ----------------------------------------------------------------
func AboutPage(xip string) string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>About Page</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: lightblue;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>Pi Test Robot</p>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + " <BR> <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/QRTS'> [ QRTS GitHub Repository ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "Pi Test Robot"
	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func TestPage(xip string) string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Test Page</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = LoopDisplay(xdata)

	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: grey;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: white;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startLoop()'>"
	xdata = xdata + "<H1>Test Page</H1>"
	xdata = xdata + "<div id='txtloop'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p1>Pi Test Robot</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p2>Pi Test Robot</p2>"
	xdata = xdata + "</center>"

	//------------------------------------------------------------------------
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// -------------------------------------------------- Functions
// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start msedge"}
	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func DateTimeDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm' : 'am';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtdt').innerHTML = ' - '+h + ':' + m + ':' + s+' '+ampm+' - '+day+', '+month+' '+dm+', '+yr;"
	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}

func LoopDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startLoop() {"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtloop').innerHTML = Math.random();"
	xdata = xdata + "  var t = setTimeout(startLoop, 500);"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func ServoTesta(sa ServoArray) {
	for i := 0; i < 130; i++ {
		sa = SetAllServoPositions(sa, i)
		SetServoPositions(sa)
		time.Sleep(10 * time.Millisecond)
	}
}
func SetAllServoPositions(sa ServoArray, setpos int) ServoArray {
	sa.Servo1pos = setpos
	sa.Servo2pos = setpos
	sa.Servo3pos = setpos
	sa.Servo4pos = setpos
	sa.Servo5pos = setpos
	sa.Servo6pos = setpos
	sa.Servo7pos = setpos
	sa.Servo8pos = setpos
	sa.Servo9pos = setpos
	sa.Servo10pos = setpos
	sa.Servo11pos = setpos
	sa.Servo12pos = setpos
	return sa

}

func SetServoPositions(sa ServoArray) {
	sa.Servo1.Angle(sa.Servo1pos)
	sa.Servo2.Angle(sa.Servo2pos)
	sa.Servo3.Angle(sa.Servo3pos)
	sa.Servo4.Angle(sa.Servo4pos)
	sa.Servo5.Angle(sa.Servo5pos)
	sa.Servo6.Angle(sa.Servo6pos)
	sa.Servo7.Angle(sa.Servo7pos)
	sa.Servo8.Angle(sa.Servo8pos)
	sa.Servo9.Angle(sa.Servo9pos)
	sa.Servo10.Angle(sa.Servo10pos)
	sa.Servo11.Angle(sa.Servo11pos)
	sa.Servo12.Angle(sa.Servo12pos)
}

// ----------------------------------------------------------------
func main() {
	fmt.Println("Pi Test Robot")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	xip := fmt.Sprintf("%s", GetOutboundIP())
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
	sa := []ServoArray{}
	saa := ServoArray{}
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
			xdata := InitPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page Handler
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := AboutPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Test Page Handler
		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			xdata := TestPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ Servo Test 1
		http.HandleFunc("/servotest1", func(w http.ResponseWriter, r *http.Request) {
			xdata := ServoTest1(xip)
			sa[0] = SetAllServoPositions(sa[0], 65)
			SetServoPositions(sa[0])
			time.Sleep(1000 * time.Millisecond)
			ServoTesta(sa[0])
			//sa[0] = SetAllServoPositions(sa[0], 75)
			//SetServoPositions(sa[0])
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------- Start Server
		Openbrowser(xip + ":" + port)
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
