package types

import "github.com/0xk2/decision_tree/decision_tree/votemachine"

type Vote struct {
	Who       string        `json:"who"`
	Options   []interface{} `json:"options"`
	MissionId string        `json:"missionId"`
}

type VoteResponse struct {
	Status      bool                        `json:"status"`
	Options     []interface{}               `json:"options"`
	Message     string                      `json:"message"`
	MachineType votemachine.VoteMachineType `json:"machine_type"`
}
