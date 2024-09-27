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
		sa.Servo1pos = setpos
		sa.Servo4pos = setpos
		sa.Servo7pos = setpos
		sa.Servo10pos = setpos

	case level == 2:
		sa.Servo2pos = setpos
		sa.Servo5pos = setpos
		sa.Servo8pos = setpos
		sa.Servo11pos = setpos

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
	rpos := 0
	switch {
	case pos > 65:
		rpos = pos - 65
	case pos < 65:
		rpos = 130 - pos
	default:
		rpos = pos
	}
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
			if i <= sa.Servo1pos {
				sa.Servo1.Angle(i)
				sa.Servo1pos = i
			}
			if i <= sa.Servo4pos {
				sa.Servo4.Angle(ReversePosition(i))
				sa.Servo4pos = ReversePosition(i)
			}
			if i <= sa.Servo7pos {
				sa.Servo7.Angle(i)
				sa.Servo7pos = i
			}

			if i <= sa.Servo10pos {
				sa.Servo10.Angle(ReversePosition(i))
				sa.Servo10pos = ReversePosition(i)
			}

		case level == 2:
			if i <= sa.Servo2pos {
				sa.Servo2.Angle(ReversePosition(i))
				sa.Servo2pos = ReversePosition(i)
			}

			if i <= sa.Servo5pos {
				sa.Servo5.Angle(i)
				sa.Servo5pos = i

			}

			if i <= sa.Servo8pos {
				sa.Servo8.Angle(ReversePosition(i))
				sa.Servo8pos = ReversePosition(i)
			}

			if i <= sa.Servo11pos {
				sa.Servo11.Angle(i)
				sa.Servo11pos = i
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
			if i >= sa.Servo1pos {
				sa.Servo1.Angle(i)
				sa.Servo1pos = i
			}
			if i >= sa.Servo4pos {
				sa.Servo4.Angle(ReversePosition(i))
				sa.Servo4pos = ReversePosition(i)
			}
			if i >= sa.Servo7pos {
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
			if i >= sa.Servo5pos {
				sa.Servo5.Angle(ReversePosition(i))
				sa.Servo5pos = i
			}

			if i >= sa.Servo8pos {
				sa.Servo8.Angle(ReversePosition(i))
				sa.Servo8pos = i
			}

			if i >= sa.Servo11pos {
				sa.Servo11.Angle(i)
				sa.Servo11pos = i
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
