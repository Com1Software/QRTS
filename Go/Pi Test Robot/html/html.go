package html

import (
	"net"
	"robot/model"
	"strconv"

	"fmt"
	"os"
)

// ----------------------------------------------------------------
func ServoTest(xip string, sa model.ServoArray, name string) string {

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
	xdata = xdata + "    background-color: green;"
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
	xdata = xdata + "<fieldset>"
	xdata = xdata + "<legend>"
	xdata = xdata + name

	xdata = xdata + "</legend>"

	xdata = xdata + "<p>Pi Test Robot</p>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<BR>  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  <BR><BR>"
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/QRTS'> [ QRTS GitHub Repository ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<table>"
	xdata = xdata + "<td width='50'>"

	xdata = xdata + "Leg 1<BR><BR><BR><BR>"
	xdata = xdata + "Leg 2<BR><BR><BR><BR>"
	xdata = xdata + "Leg 3<BR><BR><BR><BR>"
	xdata = xdata + "Leg 4<BR><BR>"

	xdata = xdata + "</td>"

	xdata = xdata + "<td width='200'>"

	xdata = xdata + "Servo 1 Position " + strconv.Itoa(sa.Servo1pos) + "<BR>"
	xdata = xdata + "Servo 2 Position " + strconv.Itoa(sa.Servo2pos) + "<BR>"
	xdata = xdata + "Servo 3 Position " + strconv.Itoa(sa.Servo3pos) + "<BR><BR>"
	xdata = xdata + "Servo 4 Position " + strconv.Itoa(sa.Servo4pos) + "<BR>"
	xdata = xdata + "Servo 5 Position " + strconv.Itoa(sa.Servo5pos) + "<BR>"
	xdata = xdata + "Servo 6 Position " + strconv.Itoa(sa.Servo6pos) + "<BR><BR>"
	xdata = xdata + "Servo 7 Position " + strconv.Itoa(sa.Servo7pos) + "<BR>"
	xdata = xdata + "Servo 8 Position " + strconv.Itoa(sa.Servo8pos) + "<BR>"
	xdata = xdata + "Servo 9 Position " + strconv.Itoa(sa.Servo9pos) + "<BR><BR>"
	xdata = xdata + "Servo 10 Position " + strconv.Itoa(sa.Servo10pos) + "<BR>"
	xdata = xdata + "Servo 11 Position " + strconv.Itoa(sa.Servo11pos) + "<BR>"
	xdata = xdata + "Servo 12 Position " + strconv.Itoa(sa.Servo12pos) + "<BR><BR>"

	xdata = xdata + "</td>"

	xdata = xdata + "</table>"

	xdata = xdata + "</fieldset>"

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
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest1?name=Servo Test 1'> [ Servo Test 1 ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest2?name=Servo Test 2''> [ Servo Test 2 ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest3?name=Servo Test 3''> [ Servo Test 3 ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest4?name=Servo Test 4''> [ Servo Test 4 ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest5?name=Servo Test 5''> [ Servo Test 5 ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/servotest6?name=Servo Test 6''> [ Servo Test 6 ] </A>  "
	xdata = xdata + "<BR><HR><BR>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/setstand?name=Set Stand'> [ Set Stand ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/movestand?name=Move Stand'> [ Move Stand ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/movelay?name=Move Lay'> [ Move Lay ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/movesit?name=Move Sit'> [ Move Sit ] </A>  "

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
