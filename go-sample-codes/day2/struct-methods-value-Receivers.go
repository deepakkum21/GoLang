// The methods are not defined within the struct, however. Instead, they're associated with the struct.
// Here, you can either let these methods just access data via working as a copy of the struct, or you can actually point through and modify the object.
// Methods that just access values are called value receivers and methods that can modify information are pointer receivers.

package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16  //min: 0,      max: 65535    16bit
	brake_pedal    uint16  //min: 0,      max: 65535
	steering_wheel int16   //min: -32768  max: 32768
	top_speed_kmh  float64 //what's our top speed?
}

func main() {
	//a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	a_car := car{22314, 0, 12562, 225.0}
	fmt.Println("gas_pedal:", a_car.gas_pedal, "brake_pedal:", a_car.brake_pedal, "steering_wheel:", a_car.steering_wheel)
	fmt.Println("Speed in kmh: ", a_car.kmh())
	fmt.Println("Speed in kmh: ", a_car.mph())
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / kmh_multiple / usixteenbitmax)
}
