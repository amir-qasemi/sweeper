package sweeper

import (
	"fmt"
	"log"
	"math"
	"time"
)

type Agent interface {
	Move(Target)
	Id() int
	IsFree() bool
	SetFree(bool)
	Loc() Location
	Log(string)
}

type ThreeMoveAgent struct {
	CurrentTarget Target
	CurrentLoc    Location
	AgentId       int
	Free          bool // TODO: race condition
}

func (a *ThreeMoveAgent) Loc() Location {
	return a.CurrentLoc
}
func (a *ThreeMoveAgent) IsFree() bool {
	return a.Free
}

func (a *ThreeMoveAgent) SetFree(f bool) {
	a.Free = f
}
func (a *ThreeMoveAgent) Move(target Target) {
	for {
		if a.CurrentLoc == target.Loc {
			a.Log(fmt.Sprintf("In the loc{%d, %d}", target.Loc.X, target.Loc.Y))
			break
		}

		horiz, vert := DetermineStep(a.CurrentLoc, target.Loc)
		a.CurrentLoc.X += horiz
		a.CurrentLoc.Y += vert

		a.Log(fmt.Sprintf("Moving with x: %d, y: %d, currentLoc:%d, %d", horiz, vert, a.CurrentLoc.X, a.CurrentLoc.Y))

		time.Sleep(1 * time.Second)
	}
}

// TODO : lift to struct
func (a *ThreeMoveAgent) Id() int {
	return a.AgentId
}

func (a *ThreeMoveAgent) Log(msg string) {
	log.Printf("Agent#%d, %s", a.Id(), msg)
}

func DetermineStep(srcLoc Location, dstLoc Location) (int, int) {
	verticalMove, horizMove := 0, 0

	if srcLoc.X < dstLoc.X {
		horizMove = 1
	} else if srcLoc.X > dstLoc.X {
		horizMove = -1
	}

	if srcLoc.Y < dstLoc.Y {
		verticalMove = 1
	} else if srcLoc.Y > dstLoc.Y {
		verticalMove = -1
	}

	return horizMove, verticalMove
}

func EstimateDistance(srcLoc Location, dstLoc Location) int {
	return int(math.Abs(float64(dstLoc.X-srcLoc.X)) + math.Abs(float64(dstLoc.Y-srcLoc.Y)))
}
