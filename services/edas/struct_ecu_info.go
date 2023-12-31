package edas

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

// EcuInfo is a nested struct in edas response
type EcuInfo struct {
	VpcId         string `json:"VpcId" xml:"VpcId"`
	UpdateTime    int64  `json:"UpdateTime" xml:"UpdateTime"`
	IpAddr        string `json:"IpAddr" xml:"IpAddr"`
	AvailableCpu  int    `json:"AvailableCpu" xml:"AvailableCpu"`
	CreateTime    int64  `json:"CreateTime" xml:"CreateTime"`
	UserId        string `json:"UserId" xml:"UserId"`
	InstanceId    string `json:"InstanceId" xml:"InstanceId"`
	RegionId      string `json:"RegionId" xml:"RegionId"`
	EcuId         string `json:"EcuId" xml:"EcuId"`
	Online        bool   `json:"Online" xml:"Online"`
	DockerEnv     bool   `json:"DockerEnv" xml:"DockerEnv"`
	AvailableMem  int    `json:"AvailableMem" xml:"AvailableMem"`
	ZoneId        string `json:"ZoneId" xml:"ZoneId"`
	Name          string `json:"Name" xml:"Name"`
	HeartbeatTime int64  `json:"HeartbeatTime" xml:"HeartbeatTime"`
}
