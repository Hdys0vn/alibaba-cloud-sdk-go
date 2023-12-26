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

// DataArchiveOrderDetail is a nested struct in dms_enterprise response
type DataArchiveOrderDetail struct {
	Comment             string          `json:"Comment" xml:"Comment"`
	Committer           string          `json:"Committer" xml:"Committer"`
	CommitterId         int64           `json:"CommitterId" xml:"CommitterId"`
	GmtCreate           string          `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified         string          `json:"GmtModified" xml:"GmtModified"`
	Id                  int64           `json:"Id" xml:"Id"`
	PluginType          string          `json:"PluginType" xml:"PluginType"`
	StatusCode          string          `json:"StatusCode" xml:"StatusCode"`
	StatusDesc          string          `json:"StatusDesc" xml:"StatusDesc"`
	WorkflowInstanceId  int64           `json:"WorkflowInstanceId" xml:"WorkflowInstanceId"`
	WorkflowStatusDesc  string          `json:"WorkflowStatusDesc" xml:"WorkflowStatusDesc"`
	RelatedUserList     []int64         `json:"RelatedUserList" xml:"RelatedUserList"`
	RelatedUserNickList []string        `json:"RelatedUserNickList" xml:"RelatedUserNickList"`
	PluginExtraData     PluginExtraData `json:"PluginExtraData" xml:"PluginExtraData"`
	PluginParam         PluginParam     `json:"PluginParam" xml:"PluginParam"`
}