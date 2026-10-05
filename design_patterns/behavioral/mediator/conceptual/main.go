package main

func main() {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

// PassengerTrain: Arrived
// FreightTrain: Arrival blocked, waiting
// PassengerTrain: Leaving
// FreightTrain: Arrival permitted
// FreightTrain: Arrived
