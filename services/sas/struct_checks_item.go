package sas

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

// ChecksItem is a nested struct in sas response
type ChecksItem struct {
	CheckId         int64          `json:"CheckId" xml:"CheckId"`
	CheckShowName   string         `json:"CheckShowName" xml:"CheckShowName"`
	Vendor          string         `json:"Vendor" xml:"Vendor"`
	VendorShowName  string         `json:"VendorShowName" xml:"VendorShowName"`
	InstanceType    string         `json:"InstanceType" xml:"InstanceType"`
	InstanceSubType string         `json:"InstanceSubType" xml:"InstanceSubType"`
	RiskLevel       string         `json:"RiskLevel" xml:"RiskLevel"`
	Status          string         `json:"Status" xml:"Status"`
	TaskId          string         `json:"TaskId" xml:"TaskId"`
	LastCheckTime   int64          `json:"LastCheckTime" xml:"LastCheckTime"`
	CheckPolicies   []CheckPolicie `json:"CheckPolicies" xml:"CheckPolicies"`
}
