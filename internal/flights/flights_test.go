package flights

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRoute(t *testing.T) {
	tests := []struct {
		name        string
		flights     Flights
		expected    Calculation
		expectedErr error
	}{
		{
			name: "valid route",
			flights: Flights{
				Route: [][]string{{"MEX", "YYZ"}},
				FlightsList: [][]string{
					{"MEX", "LAX"},
					{"LAX", "SFO"},
					{"SFO", "SEA"},
					{"SEA", "YYZ"},
				},
			},
			expected: Calculation{
				{"MEX", "LAX"},
				{"LAX", "SFO"},
				{"SFO", "SEA"},
				{"SEA", "YYZ"},
			},
			expectedErr: nil,
		},
		{
			name: "valid route with a small route",
			flights: Flights{
				Route: [][]string{{"MEX", "SFO"}},
				FlightsList: [][]string{
					{"MEX", "LAX"},
					{"SFO", "SEA"},
					{"SEA", "YYZ"},
					{"LAX", "SFO"},
				},
			},
			expected: Calculation{
				{"MEX", "LAX"},
				{"LAX", "SFO"},
			},
			expectedErr: nil,
		},
		{
			name: "invalid route",
			flights: Flights{
				Route:       [][]string{},
				FlightsList: [][]string{},
			},
			expected:    Calculation{},
			expectedErr: InvalidFlight{},
		},
		{
			name: "invalid flights",
			flights: Flights{
				Route: [][]string{{"MEX", "YYZ"}},
				FlightsList: [][]string{
					{"MEX", "LAX"},
					{"LAX", "SFO"},
				},
			},
			expected:    Calculation{},
			expectedErr: InvalidFlight{},
		},
		{
			name: "no route found",
			flights: Flights{
				Route: [][]string{{"MEX", "SEA"}},
				FlightsList: [][]string{
					{"MEX", "LAX"},
					{"LAX", "SFO"},
					{"SEA", "YYZ"},
				},
			},
			expected:    Calculation{},
			expectedErr: InvalidFlight{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.flights.CalculateRoute()
			assert.Equal(t, tt.expected, result)
			assert.IsType(t, tt.expectedErr, err)
			if err != nil {
				assert.Equal(t, "invalid flight", err.Error())
			}
		})
	}
}

func TestValidateFlights(t *testing.T) {
	tests := []struct {
		name     string
		flights  Flights
		expected bool
	}{
		{
			name: "valid flights",
			flights: Flights{
				Route: [][]string{{"MEX", "YYZ"}},
				FlightsList: [][]string{
					{"MEX", "LAX"},
					{"LAX", "SFO"},
					{"SFO", "SEA"},
					{"SEA", "YYZ"},
				},
			},
			expected: true,
		},
		{
			name: "invalid flights",
			flights: Flights{
				Route: [][]string{{"MEX", "YYZ"}},
				FlightsList: [][]string{
					{"MEX", "LAX"},
					{"LAX", "SEA"},
					{"SFO", "YYZ"},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.flights.validateFlights()
			assert.Equal(t, tt.expected, result)
		})
	}
}
