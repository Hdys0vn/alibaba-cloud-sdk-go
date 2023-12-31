package cloud_siem

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

// PlaybookListItem is a nested struct in cloud_siem response
type PlaybookListItem struct {
	OpCode      string `json:"OpCode" xml:"OpCode"`
	OpLevel     string `json:"OpLevel" xml:"OpLevel"`
	Description string `json:"Description" xml:"Description"`
	DisplayName string `json:"DisplayName" xml:"DisplayName"`
	TaskConfig  string `json:"TaskConfig" xml:"TaskConfig"`
	Name        string `json:"Name" xml:"Name"`
	WafPlaybook bool   `json:"WafPlaybook" xml:"WafPlaybook"`
}
