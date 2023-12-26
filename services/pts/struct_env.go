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

// Env is a nested struct in pts response
type Env struct {
	CreateTime    int64            `json:"CreateTime" xml:"CreateTime"`
	EnvVersion    string           `json:"EnvVersion" xml:"EnvVersion"`
	ModifiedTime  int64            `json:"ModifiedTime" xml:"ModifiedTime"`
	UsedCapacity  int64            `json:"UsedCapacity" xml:"UsedCapacity"`
	EnvName       string           `json:"EnvName" xml:"EnvName"`
	EnvId         string           `json:"EnvId" xml:"EnvId"`
	RunningScenes []string         `json:"RunningScenes" xml:"RunningScenes"`
	RelatedScenes []string         `json:"RelatedScenes" xml:"RelatedScenes"`
	Files         []File           `json:"Files" xml:"Files"`
	Properties    []PropertiesItem `json:"Properties" xml:"Properties"`
}