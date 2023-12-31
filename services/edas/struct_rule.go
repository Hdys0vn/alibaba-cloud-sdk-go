package edas

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

// Rule is a nested struct in edas response
type Rule struct {
	Step             int       `json:"Step" xml:"Step"`
	TemplateId       string    `json:"TemplateId" xml:"TemplateId"`
	MinReplicas      int       `json:"MinReplicas" xml:"MinReplicas"`
	LastDisableTime  int64     `json:"LastDisableTime" xml:"LastDisableTime"`
	TemplateVersion  int       `json:"TemplateVersion" xml:"TemplateVersion"`
	CreateTime       int64     `json:"CreateTime" xml:"CreateTime"`
	Mode             string    `json:"Mode" xml:"Mode"`
	Duration         int       `json:"Duration" xml:"Duration"`
	LoadNum          int       `json:"LoadNum" xml:"LoadNum"`
	GroupId          string    `json:"GroupId" xml:"GroupId"`
	VpcId            string    `json:"VpcId" xml:"VpcId"`
	Rt               int       `json:"Rt" xml:"Rt"`
	MetricType       string    `json:"MetricType" xml:"MetricType"`
	ScaleRuleEnabled bool      `json:"ScaleRuleEnabled" xml:"ScaleRuleEnabled"`
	UpdateTime       int64     `json:"UpdateTime" xml:"UpdateTime"`
	MaxReplicas      int       `json:"MaxReplicas" xml:"MaxReplicas"`
	SpecId           string    `json:"SpecId" xml:"SpecId"`
	VSwitchIds       string    `json:"VSwitchIds" xml:"VSwitchIds"`
	AppId            string    `json:"AppId" xml:"AppId"`
	ScaleRuleType    string    `json:"ScaleRuleType" xml:"ScaleRuleType"`
	ResourceFrom     string    `json:"ResourceFrom" xml:"ResourceFrom"`
	ScaleRuleName    string    `json:"ScaleRuleName" xml:"ScaleRuleName"`
	Enable           bool      `json:"Enable" xml:"Enable"`
	InstNum          int       `json:"InstNum" xml:"InstNum"`
	Cond             string    `json:"Cond" xml:"Cond"`
	Cpu              int       `json:"Cpu" xml:"Cpu"`
	MultiAzPolicy    string    `json:"MultiAzPolicy" xml:"MultiAzPolicy"`
	Metric           Metric    `json:"Metric" xml:"Metric"`
	Behaviour        Behaviour `json:"Behaviour" xml:"Behaviour"`
	Trigger          Trigger   `json:"Trigger" xml:"Trigger"`
}
