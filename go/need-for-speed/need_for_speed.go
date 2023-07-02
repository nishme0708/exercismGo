package speed

// TODO: define the 'Car' type struct

// NewCar creates a new remote controlled car with full battery and given specifications.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery || car.battery == 0 {
		return car
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	if car.batteryDrain > car.battery || car.battery == 0 {
		return false
	}
	ticks := (track.distance - car.distance) / car.speed
	car.batteryDrain = (ticks * car.batteryDrain)
	if car.batteryDrain > car.battery {
		return false
	}
	return true
}
