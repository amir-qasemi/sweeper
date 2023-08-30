package main

import (
	"it/sweeper"
)

func main() {
	numAg := 2
	agents := make([]sweeper.Agent, 0)

	// Init Agents
	for i := 0; i < numAg; i++ {
		agent := &sweeper.ThreeMoveAgent{
			CurrentLoc: sweeper.Location{X: 0, Y: 0},
			AgentId:    i,
			Free:       true,
		}
		agents = append(agents, agent)
	}

	// Init targets TODO: can be inited with parameter
	targets := [3]sweeper.Target{}
	targets[0] = sweeper.Target{
		Loc: sweeper.Location{X: 10, Y: 7},
	}
	targets[1] = sweeper.Target{
		Loc: sweeper.Location{X: 11, Y: 11},
	}
	targets[2] = sweeper.Target{
		Loc: sweeper.Location{X: -8, Y: 20},
	}

	c := sweeper.Coordinator{Agents: agents, Targets: targets[:]}
	c.Coordinate()
}
