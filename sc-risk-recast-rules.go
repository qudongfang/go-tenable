package go_tenable

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func (sc *TenableSC) ListRiskRecastRules() RecastRiskRuleResponse {
	var params = "fields=id,repository,organization,user,plugin,newSeverity,hostType,hostValue,port,protocol,order,status," +
		"comments,createdTime,modifiedTime"

	resp, err := sc.Get("recastRiskRule", params)
	tmp, _ := ioutil.ReadAll(resp.Body)
	var Rules = RecastRiskRuleResponse{}
	err = json.Unmarshal(tmp, &Rules)

	if err != nil {
		log.Printf("Unable to unmarshal Risk Recast Rules: %v\n", err)
	}

	return Rules
}

type RecastRiskRuleResponse struct {
	Type        string           `json:"type"`
	RecastRules []RecastRiskRule `json:"response"`
	ErrorCode   int              `json:"error_code"`
	ErrorMsg    string           `json:"error_msg"`
	Warnings    []interface{}    `json:"warnings"`
	Timestamp   int              `json:"timestamp"`
}

type RecastRiskRule struct {
	HostValue   interface{} `json:"hostValue,omitempty"`
	HostType    string      `json:"hostType"`
	Port        string      `json:"Port,omitempty"`
	Protocol    string      `json:"protocol,omitempty"`
	NewSeverity string      `json:"newSeverity"`
	CreatedTime string      `json:"createdTime"`
	Order       string      `json:"order"`
	Status      string      `json:"status"`
	ID          string      `json:"id"`
	Comments    string      `json:"comments"`
	Plugin      struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	} `json:"plugin"`
	Repository struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	} `json:"repository"`
	User struct {
		ID        string `json:"id"`
		Username  string `json:"username"`
		Firstname string `json:"firstname,omitempty"`
		Lastname  string `json:"lastname,omitempty"`
	} `json:"User"`
}
