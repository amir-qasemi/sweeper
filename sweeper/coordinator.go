package sweeper

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Coordinator struct {
	Agents  []Agent
	Targets []Target
}

func (c *Coordinator) Coordinate() {
	var wg sync.WaitGroup
	wg.Add(len(c.Targets))

	for i, t := range c.Targets {
		for {
			agent := c.FindBestSuiter(t)
			if agent != nil {
				agent.Log(fmt.Sprintf("Got target %d", i))
				agent.SetFree(false)

				go func(i int, t Target, agent Agent) {
					defer wg.Done()
					defer agent.SetFree(true)
					agent.Move(t)

				}(i, t, agent)
				break
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}

	wg.Wait()
}

func (c *Coordinator) FindBestSuiter(target Target) Agent {
	freeAgents := make([]Agent, 0)
	for _, agent := range c.Agents {
		if agent.IsFree() {
			freeAgents = append(freeAgents, agent)
		}
	}

	if len(freeAgents) == 0 {
		return nil
	}

	if len(freeAgents) == 1 {
		return freeAgents[0]
	}

	agentDistances := make(map[int][]Agent)
	minDist := math.MaxInt
	for _, agent := range freeAgents {
		dist := EstimateDistance(agent.Loc(), target.Loc)
		if minDist >= dist {
			minDist = dist
		}

		_, ok := agentDistances[dist]
		if !ok {
			agentDistances[dist] = make([]Agent, 0)
		}
		agentDistances[dist] = append(agentDistances[dist], agent)
	}

	closestAgents := agentDistances[minDist]
	if len(closestAgents) == 1 {
		return closestAgents[0]
	}

	return highestPriorityAgent(closestAgents)
}

func highestPriorityAgent(agents []Agent) Agent {
	highestAgentId, highestIndex := agents[0].Id(), 0

	for i, a := range agents {
		if a.Id() >= highestAgentId {
			highestAgentId = a.Id()
			highestIndex = i
		}
	}

	return agents[highestIndex]
}
