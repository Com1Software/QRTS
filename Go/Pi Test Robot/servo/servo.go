package servo

import (
	"robot/model"
	"time"
)

func ServoTesta(sa model.ServoArray) {
	for i := 0; i < 130; i++ {
		sa = SetAllServoPositions(sa, i)
		SetServoPositions(sa)
		time.Sleep(10 * time.Millisecond)
	}
}
func SetAllServoPositions(sa model.ServoArray, setpos int) model.ServoArray {
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
	SetServoPositions(sa)
	return sa

}

func SetLevelServoPositions(sa model.ServoArray, setpos int, level int) model.ServoArray {
	switch {
	case level == 1:
		sa.Servo1pos = ReversePosition(setpos)
		sa.Servo4pos = setpos
		sa.Servo7pos = ReversePosition(setpos)
		sa.Servo10pos = setpos

	case level == 2:
		sa.Servo2pos = setpos
		sa.Servo5pos = ReversePosition(setpos)
		sa.Servo8pos = setpos
		sa.Servo11pos = ReversePosition(setpos)

	case level == 3:
		sa.Servo3pos = setpos
		sa.Servo6pos = ReversePosition(setpos)
		sa.Servo9pos = setpos
		sa.Servo12pos = ReversePosition(setpos)

	}
	SetServoPositions(sa)
	return sa

}

func ReversePosition(pos int) int {
	rpos := 130 - pos
	return rpos
}

func SetServoPositions(sa model.ServoArray) {
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

func MoveLevelServoPositionsDown(sa model.ServoArray, level int, startpos int, setpos int, mstimer time.Duration) model.ServoArray {

	for i := startpos; i >= setpos; i-- {

		switch {
		case level == 1:
			if i <= ReversePosition(sa.Servo1pos) {
				sa.Servo1.Angle(ReversePosition(i))
				sa.Servo1pos = ReversePosition(i)
			}

			if i <= sa.Servo4pos {
				sa.Servo4.Angle(i)
				sa.Servo4pos = i
			}

			if i <= ReversePosition(sa.Servo7pos) {
				sa.Servo7.Angle(ReversePosition(i))
				sa.Servo7pos = ReversePosition(i)
			}
			if i <= sa.Servo10pos {
				sa.Servo10.Angle(i)
				sa.Servo10pos = i
			}

		case level == 2:

			if i <= sa.Servo2pos {
				sa.Servo2.Angle(i)
				sa.Servo2pos = i

			}
			if i <= ReversePosition(sa.Servo5pos) {
				sa.Servo5.Angle(ReversePosition(i))
				sa.Servo5pos = ReversePosition(i)
			}

			if i <= sa.Servo8pos {
				sa.Servo8.Angle(i)
				sa.Servo8pos = i
			}

			if i <= ReversePosition(sa.Servo11pos) {
				sa.Servo11.Angle(ReversePosition(i))
				sa.Servo11pos = ReversePosition(i)
			}
		case level == 3:
			if i <= sa.Servo3pos {
				sa.Servo3.Angle(i)
				sa.Servo3pos = i
			}

			if i <= sa.Servo6pos {
				sa.Servo6.Angle(i)
				sa.Servo6pos = i
			}

			if i <= sa.Servo9pos {
				sa.Servo9.Angle(i)
				sa.Servo9pos = i
			}

			if i <= sa.Servo12pos {
				sa.Servo12.Angle(i)
				sa.Servo12pos = i
			}

		}
		time.Sleep(mstimer * time.Millisecond)

	}
	return sa

}

func MoveLevelServoPositionsUp(sa model.ServoArray, level int, startpos int, setpos int, mstimer time.Duration) model.ServoArray {

	for i := startpos; i <= setpos; i++ {

		switch {
		case level == 1:
			if i >= ReversePosition(sa.Servo1pos) {
				sa.Servo1.Angle(ReversePosition(i))
				sa.Servo1pos = ReversePosition(i)
			}
			if i >= sa.Servo4pos {
				sa.Servo4.Angle(i)
				sa.Servo4pos = i
			}
			if i >= ReversePosition(sa.Servo7pos) {
				sa.Servo7.Angle(ReversePosition(i))
				sa.Servo7pos = ReversePosition(i)
			}

			if i >= sa.Servo10pos {
				sa.Servo10.Angle(i)
				sa.Servo10pos = i
			}

		case level == 2:
			if i >= sa.Servo2pos {
				sa.Servo2.Angle(i)
				sa.Servo2pos = i

			}
			if i >= ReversePosition(sa.Servo5pos) {
				sa.Servo5.Angle(ReversePosition(i))
				sa.Servo5pos = ReversePosition(i)
			}

			if i >= sa.Servo8pos {
				sa.Servo8.Angle(i)
				sa.Servo8pos = i
			}

			if i >= ReversePosition(sa.Servo11pos) {
				sa.Servo11.Angle(ReversePosition(i))
				sa.Servo11pos = ReversePosition(i)
			}

		case level == 3:
			if i >= sa.Servo3pos {
				sa.Servo3.Angle(i)
				sa.Servo3pos = i
			}

			if i >= sa.Servo6pos {
				sa.Servo6.Angle(i)
				sa.Servo6pos = i
			}

			if i >= sa.Servo9pos {
				sa.Servo9.Angle(i)
				sa.Servo9pos = i
			}

			if i >= sa.Servo12pos {
				sa.Servo12.Angle(i)
				sa.Servo12pos = i
			}

		}
		time.Sleep(mstimer * time.Millisecond)

	}
	return sa

}

func MoveServo1and2BiDirectionalPositionDown(sa model.ServoArray, startpos int, setpos int, bistartpos int, bisetpos int, mstimer time.Duration) model.ServoArray {
	ii := bistartpos
	for i := startpos; i >= setpos; i-- {

		if i <= ReversePosition(sa.Servo1pos) {
			sa.Servo1.Angle(ReversePosition(ii))
			sa.Servo1pos = ReversePosition(ii)
		}

		if i <= sa.Servo4pos {
			sa.Servo4.Angle(ii)
			sa.Servo4pos = ii
		}

		if i <= ReversePosition(sa.Servo7pos) {
			sa.Servo7.Angle(ReversePosition(ii))
			sa.Servo7pos = ReversePosition(ii)
		}
		if i <= sa.Servo10pos {
			sa.Servo10.Angle(ii)
			sa.Servo10pos = ii
		}

		if i <= sa.Servo2pos {
			sa.Servo2.Angle(i)
			sa.Servo2pos = i

		}
		if i <= ReversePosition(sa.Servo5pos) {
			sa.Servo5.Angle(ReversePosition(i))
			sa.Servo5pos = ReversePosition(i)
		}

		if i <= sa.Servo8pos {
			sa.Servo8.Angle(i)
			sa.Servo8pos = i
		}

		if i <= ReversePosition(sa.Servo11pos) {
			sa.Servo11.Angle(ReversePosition(i))
			sa.Servo11pos = ReversePosition(i)
		}
		ii++
		time.Sleep(mstimer * time.Millisecond)

	}
	return sa

}

func MoveServo1and2BiDirectionalPositionUp(sa model.ServoArray, startpos int, setpos int, bistartpos int, bisetpos int, mstimer time.Duration) model.ServoArray {

	for i := startpos; i <= setpos; i++ {
		if i >= ReversePosition(sa.Servo1pos) {
			sa.Servo1.Angle(ReversePosition(i))
			sa.Servo1pos = ReversePosition(i)
		}
		if i >= sa.Servo4pos {
			sa.Servo4.Angle(i)
			sa.Servo4pos = i
		}
		if i >= ReversePosition(sa.Servo7pos) {
			sa.Servo7.Angle(ReversePosition(i))
			sa.Servo7pos = ReversePosition(i)
		}

		if i >= sa.Servo10pos {
			sa.Servo10.Angle(i)
			sa.Servo10pos = i
		}

		if i >= sa.Servo2pos {
			sa.Servo2.Angle(i)
			sa.Servo2pos = i

		}
		if i >= ReversePosition(sa.Servo5pos) {
			sa.Servo5.Angle(ReversePosition(i))
			sa.Servo5pos = ReversePosition(i)
		}

		if i >= sa.Servo8pos {
			sa.Servo8.Angle(i)
			sa.Servo8pos = i
		}

		if i >= ReversePosition(sa.Servo11pos) {
			sa.Servo11.Angle(ReversePosition(i))
			sa.Servo11pos = ReversePosition(i)
		}

		time.Sleep(mstimer * time.Millisecond)

	}
	return sa

}
