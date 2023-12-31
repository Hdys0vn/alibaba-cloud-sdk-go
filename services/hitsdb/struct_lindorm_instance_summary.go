package hitsdb

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

// LindormInstanceSummary is a nested struct in hitsdb response
type LindormInstanceSummary struct {
	VpcId               string `json:"VpcId" xml:"VpcId"`
	EngineType          string `json:"EngineType" xml:"EngineType"`
	ExpireTime          string `json:"ExpireTime" xml:"ExpireTime"`
	CreateTime          string `json:"CreateTime" xml:"CreateTime"`
	PayType             string `json:"PayType" xml:"PayType"`
	AliUid              int64  `json:"AliUid" xml:"AliUid"`
	InstanceStorage     string `json:"InstanceStorage" xml:"InstanceStorage"`
	InstanceId          string `json:"InstanceId" xml:"InstanceId"`
	NetworkType         string `json:"NetworkType" xml:"NetworkType"`
	ServiceType         string `json:"ServiceType" xml:"ServiceType"`
	RegionId            string `json:"RegionId" xml:"RegionId"`
	CreateMilliseconds  int64  `json:"CreateMilliseconds" xml:"CreateMilliseconds"`
	InstanceAlias       string `json:"InstanceAlias" xml:"InstanceAlias"`
	ZoneId              string `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus      string `json:"InstanceStatus" xml:"InstanceStatus"`
	ExpiredMilliseconds int64  `json:"ExpiredMilliseconds" xml:"ExpiredMilliseconds"`
	EnableStream        bool   `json:"EnableStream" xml:"EnableStream"`
	EnableCompute       bool   `json:"EnableCompute" xml:"EnableCompute"`
	ResourceGroupId     string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags                []Tag  `json:"Tags" xml:"Tags"`
}
