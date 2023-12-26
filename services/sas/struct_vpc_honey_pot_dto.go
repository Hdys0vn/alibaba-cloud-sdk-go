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

// VpcHoneyPotDTO is a nested struct in sas response
type VpcHoneyPotDTO struct {
	VpcName                   string          `json:"VpcName" xml:"VpcName"`
	VpcId                     string          `json:"VpcId" xml:"VpcId"`
	HoneyPotEniInstanceId     string          `json:"HoneyPotEniInstanceId" xml:"HoneyPotEniInstanceId"`
	CidrBlock                 string          `json:"CidrBlock" xml:"CidrBlock"`
	VpcStatus                 string          `json:"VpcStatus" xml:"VpcStatus"`
	CreateTime                int64           `json:"CreateTime" xml:"CreateTime"`
	HoneyPotVpcSwitchId       string          `json:"HoneyPotVpcSwitchId" xml:"HoneyPotVpcSwitchId"`
	HoneyPotExistence         bool            `json:"HoneyPotExistence" xml:"HoneyPotExistence"`
	VpcRegionId               string          `json:"VpcRegionId" xml:"VpcRegionId"`
	HoneyPotEcsInstanceStatus string          `json:"HoneyPotEcsInstanceStatus" xml:"HoneyPotEcsInstanceStatus"`
	HoneyPotInstanceStatus    string          `json:"HoneyPotInstanceStatus" xml:"HoneyPotInstanceStatus"`
	VpcSwitchIdList           []VpcSwitchInfo `json:"VpcSwitchIdList" xml:"VpcSwitchIdList"`
}