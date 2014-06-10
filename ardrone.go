package main

import (
	"math"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-ardrone"
	"github.com/hybridgroup/gobot-joystick"
)

type pair struct {
	x float64
	y float64
}

func main() {
	joystickAdaptor := new(gobotJoystick.JoystickAdaptor)
	joystickAdaptor.Name = "saitek"
	joystickAdaptor.Params = map[string]interface{}{
		"config": "./saitek-x52-reverse.json",
	}

	joystick := gobotJoystick.NewJoystick(joystickAdaptor)
	joystick.Name = "saitek"

	ardroneAdaptor := new(gobotArdrone.ArdroneAdaptor)
	ardroneAdaptor.Name = "Drone"

	drone := gobotArdrone.NewArdrone(ardroneAdaptor)
	drone.Name = "Drone"

	work := func() {

		offset := 32767.0
		right_stick := pair{x: 0, y: 0}
		left_stick := pair{x: 0, y: 0}

		gobot.On(joystick.Events["T1_press"], func(data interface{}) {
			drone.TakeOff()
		})
		gobot.On(joystick.Events["T3_press"], func(data interface{}) {
			drone.Hover()
		})
		gobot.On(joystick.Events["A_press"], func(data interface{}) {
			drone.Land()
		})
		gobot.On(joystick.Events["B_press"], func(data interface{}) {
			drone.Halt()
		})
		gobot.On(joystick.Events["right_x"], func(data interface{}) {
			val := float64(data.(int16))
			if left_stick.x-val < 500 {
				left_stick.x = val
			}
		})
		gobot.On(joystick.Events["right_y"], func(data interface{}) {
			val := float64(data.(int16))
			if left_stick.y-val < 500 {
				left_stick.y = val
			}
		})
		gobot.On(joystick.Events["right_rotate"], func(data interface{}) {
			val := float64(data.(int16))
			if right_stick.x-val < 500 {
				right_stick.x = val
			}
		})
		gobot.On(joystick.Events["left_throttle"], func(data interface{}) {
			val := float64(data.(int16))
			if right_stick.y-val < 100 {
				right_stick.y = val
			}
		})

		gobot.Every("0.01s", func() {
			pair := left_stick
			if pair.y < -10 {
				drone.Forward(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Backward(validatePitch(pair.y, offset))
			} else {
				drone.Forward(0)
			}

			if pair.x > 10 {
				drone.Right(validatePitch(pair.x, offset))
			} else if pair.x < -10 {
				drone.Left(validatePitch(pair.x, offset))
			} else {
				drone.Right(0)
			}
		})

		gobot.Every("0.01s", func() {
			pair := right_stick
			if pair.y < -10 {
				drone.Up(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Down(validatePitch(pair.y, offset))
			} else {
				drone.Up(0)
			}

			if pair.x > 10 {
				drone.Clockwise(validatePitch(pair.x, offset))
			} else if pair.x < -10 {
				drone.CounterClockwise(validatePitch(pair.x, offset))
			} else {
				drone.Clockwise(0)
			}
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{joystickAdaptor, ardroneAdaptor},
		Devices:     []gobot.Device{joystick, drone},
		Work:        work,
	}

	robot.Start()
}

func validatePitch(data float64, offset float64) float64 {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return float64(int(value*100)) / 100
		} else {
			return 1.0
		}
	} else {
		return 0.0
	}
}
