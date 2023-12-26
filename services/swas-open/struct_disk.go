package swas_open

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

// Disk is a nested struct in swas_open response
type Disk struct {
	CreationTime   string `json:"CreationTime" xml:"CreationTime"`
	Status         string `json:"Status" xml:"Status"`
	Device         string `json:"Device" xml:"Device"`
	Size           int    `json:"Size" xml:"Size"`
	DiskName       string `json:"DiskName" xml:"DiskName"`
	DiskChargeType string `json:"DiskChargeType" xml:"DiskChargeType"`
	DiskType       string `json:"DiskType" xml:"DiskType"`
	Category       string `json:"Category" xml:"Category"`
	DiskId         string `json:"DiskId" xml:"DiskId"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	RegionId       string `json:"RegionId" xml:"RegionId"`
	Remark         string `json:"Remark" xml:"Remark"`
	InstanceName   string `json:"InstanceName" xml:"InstanceName"`
}