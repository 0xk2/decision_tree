package network

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0xk2/decision_tree/decision_tree/dtree"
	. "github.com/0xk2/decision_tree/decision_tree/types"
	"github.com/0xk2/decision_tree/decision_tree/utils"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var missionData MissionData
	err := decoder.Decode(&missionData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	startId := missionData.Start
	var checkPoints = make(map[string]*dtree.CheckPoint)
	for _, checkPointData := range missionData.CheckPoints {
		checkPoints[checkPointData.Id] = dtree.CreateEmptyCheckPoint(checkPointData.Name, checkPointData.IsOuput, checkPointData.VoteMachineType, checkPointData.Data)
		checkPoints[checkPointData.Id].Id = checkPointData.Id
	}
	for _, checkPointData := range missionData.CheckPoints {
		if checkPointData.Children != nil {
			// log.Println("id: ", checkPointData.Id, "; children: ", checkPointData.Children)
			for _, childId := range checkPointData.Children {
				checkPoints[checkPointData.Id].Attach(checkPoints[childId])
			}
		}
	}
	for _, checkPoint := range checkPoints {
		checkPoint.FinishAttachChildren()
	}
	missionId := utils.RandString(16)
	Missions[missionId] = dtree.CreateMission(checkPoints[startId], missionData.Name, missionData.Description)
	Missions[missionId].PrintFromCurrent()
	// for demo, start now
	Missions[missionId].Current.Start("", nil)

	CreateResponse := CreateResponse{
		Id: missionId,
	}
	jsonData, err := json.Marshal(CreateResponse)
	if err != nil {
		http.Error(w, "Error crafting response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsonData)
}
