package mpaas

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

// ServerDomainConfigListItem is a nested struct in mpaas response
type ServerDomainConfigListItem struct {
	Description  string `json:"Description" xml:"Description"`
	ConfigType   string `json:"ConfigType" xml:"ConfigType"`
	AppCode      string `json:"AppCode" xml:"AppCode"`
	ConfigValue  string `json:"ConfigValue" xml:"ConfigValue"`
	H5Name       string `json:"H5Name" xml:"H5Name"`
	H5Id         string `json:"H5Id" xml:"H5Id"`
	GmtCreate    string `json:"GmtCreate" xml:"GmtCreate"`
	ConfigStatus int64  `json:"ConfigStatus" xml:"ConfigStatus"`
	GmtModified  string `json:"GmtModified" xml:"GmtModified"`
	Id           int64  `json:"Id" xml:"Id"`
}