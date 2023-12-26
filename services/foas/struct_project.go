package foas

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

// Project is a nested struct in foas response
type Project struct {
	Name            string `json:"Name" xml:"Name"`
	Region          string `json:"Region" xml:"Region"`
	ModifyTime      int64  `json:"ModifyTime" xml:"ModifyTime"`
	ClusterId       string `json:"ClusterId" xml:"ClusterId"`
	ManagerIds      string `json:"ManagerIds" xml:"ManagerIds"`
	GlobalJobConfig string `json:"GlobalJobConfig" xml:"GlobalJobConfig"`
	CreateTime      int64  `json:"CreateTime" xml:"CreateTime"`
	DeployType      string `json:"DeployType" xml:"DeployType"`
	State           string `json:"State" xml:"State"`
	Modifier        string `json:"Modifier" xml:"Modifier"`
	Creator         string `json:"Creator" xml:"Creator"`
	Id              string `json:"Id" xml:"Id"`
	Description     string `json:"Description" xml:"Description"`
}
