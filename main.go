package main

import (
	"fmt"
)

func findRoute(flights [][]string, route [][]string) [][]string {
	routes := make(map[string]string)
	for _, flight := range flights {
		origin, destination := flight[0], flight[1]
		routes[origin] = destination
	}
	for origin, destination := range routes {
		for {
			if _, ok := routes[destination]; !ok {
				break
			}
			if routes[destination] == "" {
				routes[destination] = origin
				break
			}
			destination = routes[destination]
		}
	}
	var start string
	if len(route) > 0 {
		start = route[0][0]
	} else {
		for k := range routes {
			if _, ok := routes[k]; !ok {
				start = k
				break
			}
		}
	}
	var finalDest string
	if len(route) > 0 {
		finalDest = route[len(route)-1][1]
	} else {
		for _, dest := range routes {
			if _, ok := routes[dest]; !ok {
				finalDest = dest
				break
			}
		}
	}
	var finalDestReached bool
	var newRoute [][]string
	for {
		if dest, ok := routes[start]; ok {
			if !finalDestReached && dest == finalDest {
				newRoute = append(newRoute, route...)
				finalDestReached = true
			}
			newRoute = append(newRoute, []string{start, dest})
			start = dest
		} else {
			break
		}
	}
	return newRoute
}

func main() {
	flightList := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	route := [][]string{{"SFO", "EWR"}}
	fmt.Println(findRoute(flightList, route)) // Output: [[SFO ATL] [ATL GSO] [GSO IND] [IND EWR]]
}
