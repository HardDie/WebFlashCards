package web_server

import (
	"fmt"
)

type (
	State int
)

const (
	StateMainPage State = iota
	StateSettings
	StateSelfCheckStraight
	StateSelfCheckBackward
	StateSelfCheckBoth
	StateTypingStraight
	StateTypingBackward
	StateTypingBoth
	StateNotAuthorised
)

func (s State) String() string {
	switch s {
	case StateMainPage:
		return "MainPage"
	case StateSettings:
		return "Settings"
	case StateSelfCheckStraight:
		return "SelfCheck(straight)"
	case StateSelfCheckBackward:
		return "SelfCheck(backward)"
	case StateSelfCheckBoth:
		return "SelfCheck(both)"
	case StateTypingStraight:
		return "Typing(straight)"
	case StateTypingBackward:
		return "Typing(backward)"
	case StateTypingBoth:
		return "Typing(both)"
	case StateNotAuthorised:
		return "Not authorised"
	}
	return fmt.Sprintf("unknown(%d)", s)
}
