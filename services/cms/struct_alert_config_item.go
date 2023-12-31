package cms

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

// AlertConfigItem is a nested struct in cms response
type AlertConfigItem struct {
	ComparisonOperator  string                                          `json:"ComparisonOperator" xml:"ComparisonOperator"`
	SilenceTime         string                                          `json:"SilenceTime" xml:"SilenceTime"`
	Webhook             string                                          `json:"Webhook" xml:"Webhook"`
	Times               string                                          `json:"Times" xml:"Times"`
	EscalationsLevel    string                                          `json:"EscalationsLevel" xml:"EscalationsLevel"`
	NoEffectiveInterval string                                          `json:"NoEffectiveInterval" xml:"NoEffectiveInterval"`
	EffectiveInterval   string                                          `json:"EffectiveInterval" xml:"EffectiveInterval"`
	Threshold           string                                          `json:"Threshold" xml:"Threshold"`
	Statistics          string                                          `json:"Statistics" xml:"Statistics"`
	TargetList          TargetListInDescribeGroupMonitoringAgentProcess `json:"TargetList" xml:"TargetList"`
}
