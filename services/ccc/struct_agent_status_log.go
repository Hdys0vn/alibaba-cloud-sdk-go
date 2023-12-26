package ccc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AgentStatusLog is a nested struct in ccc response
type AgentStatusLog struct {
	PhoneNo             string `json:"PhoneNo" xml:"PhoneNo"`
	TargetRequest       string `json:"TargetRequest" xml:"TargetRequest"`
	TransferNumber      string `json:"TransferNumber" xml:"TransferNumber"`
	CallId              string `json:"CallId" xml:"CallId"`
	AliHangupCause      string `json:"AliHangupCause" xml:"AliHangupCause"`
	RamId               string `json:"RamId" xml:"RamId"`
	StatisticDate       string `json:"StatisticDate" xml:"StatisticDate"`
	CallType            string `json:"CallType" xml:"CallType"`
	Extend4             string `json:"Extend4" xml:"Extend4"`
	SkillGroupId        string `json:"SkillGroupId" xml:"SkillGroupId"`
	InstanceId          string `json:"InstanceId" xml:"InstanceId"`
	CalleeId            string `json:"CalleeId" xml:"CalleeId"`
	PressKey            string `json:"PressKey" xml:"PressKey"`
	KeyMarkRelation     string `json:"KeyMarkRelation" xml:"KeyMarkRelation"`
	StatisticTime       int    `json:"StatisticTime" xml:"StatisticTime"`
	MonitedAgentPhoneNo string `json:"MonitedAgentPhoneNo" xml:"MonitedAgentPhoneNo"`
	GroupNo             string `json:"GroupNo" xml:"GroupNo"`
	Extend2             string `json:"Extend2" xml:"Extend2"`
	ConnId              string `json:"ConnId" xml:"ConnId"`
	TenantId            string `json:"TenantId" xml:"TenantId"`
	Type                string `json:"Type" xml:"Type"`
	Extend1             string `json:"Extend1" xml:"Extend1"`
	TransferNo          string `json:"TransferNo" xml:"TransferNo"`
	MonitedAgentNo      string `json:"MonitedAgentNo" xml:"MonitedAgentNo"`
	Id                  int64  `json:"Id" xml:"Id"`
	ParentNote          string `json:"ParentNote" xml:"ParentNote"`
	CallDir             string `json:"CallDir" xml:"CallDir"`
	Note                string `json:"Note" xml:"Note"`
	CallerId            string `json:"CallerId" xml:"CallerId"`
	ContactType         string `json:"ContactType" xml:"ContactType"`
	AgentNo             string `json:"AgentNo" xml:"AgentNo"`
	AgentDropCall       string `json:"AgentDropCall" xml:"AgentDropCall"`
	OutboundScenario    bool   `json:"OutboundScenario" xml:"OutboundScenario"`
	Extend3             string `json:"Extend3" xml:"Extend3"`
	TargetSelect        string `json:"TargetSelect" xml:"TargetSelect"`
	Status              string `json:"Status" xml:"Status"`
	Acid                string `json:"Acid" xml:"Acid"`
}