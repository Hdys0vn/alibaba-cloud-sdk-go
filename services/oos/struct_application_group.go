package oos

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

// ApplicationGroup is a nested struct in oos response
type ApplicationGroup struct {
	Name             string `json:"Name" xml:"Name"`
	ImportTagKey     string `json:"ImportTagKey" xml:"ImportTagKey"`
	ApplicationName  string `json:"ApplicationName" xml:"ApplicationName"`
	DeployOutputs    string `json:"DeployOutputs" xml:"DeployOutputs"`
	DeployRegionId   string `json:"DeployRegionId" xml:"DeployRegionId"`
	Progress         string `json:"Progress" xml:"Progress"`
	StatusReason     string `json:"StatusReason" xml:"StatusReason"`
	CreateDate       string `json:"CreateDate" xml:"CreateDate"`
	ImportTagValue   string `json:"ImportTagValue" xml:"ImportTagValue"`
	UpdateDate       string `json:"UpdateDate" xml:"UpdateDate"`
	CreatedDate      string `json:"CreatedDate" xml:"CreatedDate"`
	CmsGroupId       string `json:"CmsGroupId" xml:"CmsGroupId"`
	DeployParameters string `json:"DeployParameters" xml:"DeployParameters"`
	Status           string `json:"Status" xml:"Status"`
	Description      string `json:"Description" xml:"Description"`
	UpdatedDate      string `json:"UpdatedDate" xml:"UpdatedDate"`
}