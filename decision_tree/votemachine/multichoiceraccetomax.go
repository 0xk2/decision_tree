package votemachine

import (
	"fmt"
	"log"

	"github.com/0xk2/decision_tree/decision_tree/utils"
)

type MultipleChoiceRaceToMaxData_Option struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MultipleChoiceRaceToMax struct {
	Options    []MultipleChoiceRaceToMaxData_Option `json:"options"`
	Max        int                                  `json:"max"`
	Top        int                                  `json:"top"`
	voted      map[int]int
	topChoices map[int]int
}

func (this *MultipleChoiceRaceToMax) init(data interface{}, isOuput bool, noOfChildren int) {
	tmp := data.(map[string]interface{})
	var max int
	var top int
	options := make([]MultipleChoiceRaceToMaxData_Option, 0)
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, MultipleChoiceRaceToMaxData_Option{
					Title:       opt.(map[string]interface{})["title"].(string),
					Description: opt.(map[string]interface{})["description"].(string),
				})
			}
		} else if key == "max" {
			max = utils.InterfaceToInt(value)
		} else if key == "top" {
			top = utils.InterfaceToInt(value)
		}
	}
	this.Options = options
	this.Max = max
	this.Top = top
}

func (this *MultipleChoiceRaceToMax) Start(from VoteMachineType, seed interface{}) {
	// aware of who is the seed sender
	if seed != nil {
		this.Options = seed.([]MultipleChoiceRaceToMaxData_Option)
	}
	this.topChoices = nil
	this.voted = make(map[int]int)
}

func (this *MultipleChoiceRaceToMax) Vote(who string, userSelectedOptions []interface{}) {
	choices := make([]int, 0)
	for _, opt := range userSelectedOptions {
		choices = append(choices, utils.InterfaceToInt(opt))
	}
	str := ""
	max := this.Max
	options := this.Options
	for _, choice := range choices {
		str += options[choice].Title + ","
		this.voted[choice] += 1
	}
	// tally
	topChoices := make(map[int]int)
	fmt.Printf("%s vote [%s]; top %d that reach %d choice win will\n", who, str, this.Top, this.Max)
	noOfTop := 0
	for opt, choice := range this.voted {
		if choice >= max {
			noOfTop += 1
			topChoices[opt] = choice
			if noOfTop == this.Top {
				this.topChoices = topChoices
				break
			}
		}
	}
}

func (this *MultipleChoiceRaceToMax) IsValidChoice(who string, userSelectedOptions []interface{}) bool {
	if len(userSelectedOptions) > this.Top {
		return false
	}
	return true
}

func (this *MultipleChoiceRaceToMax) GetTallyResult() (VoteMachineType, int, interface{}) {
	if this.topChoices != nil {
		log.Println(this.topChoices)
		// result is a portion of Options
		result := make([]MultipleChoiceRaceToMaxData_Option, 0)
		for opt := range this.topChoices {
			result = append(result, this.Options[opt])
		}
		return VM_MultipleChoiceRaceToMax, 0, result
	}
	return VM_MultipleChoiceRaceToMax, -1, nil
}

func (this *MultipleChoiceRaceToMax) GetCurrentVoteState() interface{} {
	return this.voted
}

func (this *MultipleChoiceRaceToMax) GetChoices() interface{} {
	return this.Options
}
