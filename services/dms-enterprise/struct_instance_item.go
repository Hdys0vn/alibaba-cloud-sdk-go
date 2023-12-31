package dms_enterprise

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

// InstanceItem is a nested struct in dms_enterprise response
type InstanceItem struct {
	BusinessTime       string `json:"BusinessTime" xml:"BusinessTime"`
	CheckStatus        int64  `json:"CheckStatus" xml:"CheckStatus"`
	DagId              int64  `json:"DagId" xml:"DagId"`
	Delete             string `json:"Delete" xml:"Delete"`
	EndTime            string `json:"EndTime" xml:"EndTime"`
	GmtCreate          string `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified        string `json:"GmtModified" xml:"GmtModified"`
	HistoryDagId       int64  `json:"HistoryDagId" xml:"HistoryDagId"`
	Id                 int64  `json:"Id" xml:"Id"`
	LastRunningContext string `json:"LastRunningContext" xml:"LastRunningContext"`
	Msg                string `json:"Msg" xml:"Msg"`
	Status             int64  `json:"Status" xml:"Status"`
	TaskType           int64  `json:"TaskType" xml:"TaskType"`
	TenantId           string `json:"TenantId" xml:"TenantId"`
	TriggerType        int64  `json:"TriggerType" xml:"TriggerType"`
	Version            string `json:"Version" xml:"Version"`
}
