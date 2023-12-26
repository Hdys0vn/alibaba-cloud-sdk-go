package ehpc

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

// SoftwareInfo is a nested struct in ehpc response
type SoftwareInfo struct {
	SchedulerType    string                      `json:"SchedulerType" xml:"SchedulerType"`
	OsTag            string                      `json:"OsTag" xml:"OsTag"`
	SchedulerVersion string                      `json:"SchedulerVersion" xml:"SchedulerVersion"`
	AccountVersion   string                      `json:"AccountVersion" xml:"AccountVersion"`
	AccountType      string                      `json:"AccountType" xml:"AccountType"`
	EhpcVersion      string                      `json:"EhpcVersion" xml:"EhpcVersion"`
	Applications     ApplicationsInListSoftwares `json:"Applications" xml:"Applications"`
}
