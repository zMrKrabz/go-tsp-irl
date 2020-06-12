package solver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "encoding/json"
)

type Resp struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

// Solver is the solution to the traveling salesman problem,
// starting from the first coordinate and ending at the last coordinate
type Solver struct {
	coordinates []Coordinate
}

func (s Solver) String() string {
	var str string
	for _, c := range s.coordinates {
		str = str + c.String() + "\n"
	}

	return str
}

// Coordinate is one location
type Coordinate struct {
	X float64
	Y float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%v,%v", c.X, c.Y)
}

// AddCoordinate allows you to adds a destination to the solver
func (s *Solver) AddCoordinate(c Coordinate) {
	s.coordinates = append(s.coordinates, c)
}

// Solve requires at least 2 coordinates to be in the Solver Object
func (s *Solver) Solve() []Coordinate {

	// Start from the first coordinate, and see which next coordiante is the closest
	// based on distance
	for i, c := range s.coordinates {
		if i == len(s.coordinates)-1 {
			break
		}

		closest := s.coordinates[i+1] // the closest coordinate
		d := getDistance(c, closest)  // the distance of the closest coordinate
		index := i + 1                // the index of the closest coordinate

		// make a slice and look at every aforementioned coordinate
		slice := s.coordinates[i : len(s.coordinates)-1]
		for j, co := range slice {
			dis := getDistance(c, co)
			if dis < d {
				closest = s.coordinates[j]
				d = dis
				index = j
			}
		}

		// perform swap
		original := s.coordinates[i]
		s.coordinates[i] = closest
		s.coordinates[index] = original
	}

	return s.coordinates
}

// getDistance returns the distance between two coordinates
func getDistance(c Coordinate, d Coordinate) float64 {
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%v;%v?overview=false", c.String(), d.String())

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var r Resp
	json.NewDecoder(resp.Body).Decode(&r)

	return r.Routes[0].Distance
}
