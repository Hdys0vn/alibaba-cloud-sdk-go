package pts

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

// LoadConfig is a nested struct in pts response
type LoadConfig struct {
	MaxRunningTime         int                  `json:"MaxRunningTime" xml:"MaxRunningTime"`
	TestMode               string               `json:"TestMode" xml:"TestMode"`
	AgentCount             int                  `json:"AgentCount" xml:"AgentCount"`
	KeepTime               int                  `json:"KeepTime" xml:"KeepTime"`
	Increment              int                  `json:"Increment" xml:"Increment"`
	AutoStep               bool                 `json:"AutoStep" xml:"AutoStep"`
	Configuration          Configuration        `json:"Configuration" xml:"Configuration"`
	VpcLoadConfig          VpcLoadConfig        `json:"VpcLoadConfig" xml:"VpcLoadConfig"`
	RelationLoadConfigList []RelationLoadConfig `json:"RelationLoadConfigList" xml:"RelationLoadConfigList"`
	ApiLoadConfigList      []ApiLoadConfig      `json:"ApiLoadConfigList" xml:"ApiLoadConfigList"`
}
