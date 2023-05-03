package flights

type Flights struct {
	Route       [][]string `json:"route"`
	FlightsList [][]string `json:"flights"`
}

type Calculation [][]string

type InvalidFlight struct{}

func (i InvalidFlight) Error() string {
	return "invalid flight"
}

func (f Flights) CalculateRoute() (Calculation, error) {
	if len(f.Route) == 0 || len(f.FlightsList) == 0 {
		return Calculation{}, InvalidFlight{}
	}

	if !f.validateFlights() {
		return Calculation{}, InvalidFlight{}
	}

	start := f.Route[0][0]
	var result Calculation

	return calculate(f.FlightsList, f.Route, start, result), nil
}

func (f Flights) validateFlights() bool {
	var existStart bool
	var existEnd bool

	for _, flight := range f.FlightsList {
		if flight[0] == f.Route[0][0] {
			existStart = true
		}
		if flight[1] == f.Route[0][1] {
			existEnd = true
		}
	}

	return existStart && existEnd
}

func calculate(flightsList [][]string, route [][]string, currentStart string, result Calculation) Calculation {
	end := route[0][1]

	for {
		for _, flight := range flightsList {
			if flight[0] == currentStart {
				result = append(result, flight)

				resultStart := result[0][0]
				resultEnd := result[len(result)-1][1]

				if resultStart == route[0][0] && resultEnd == end {
					return result
				}

				currentStart = flight[1]

				return calculate(flightsList, route, currentStart, result)
			}
		}
	}
}
