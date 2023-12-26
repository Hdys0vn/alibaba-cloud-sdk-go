package ehpc

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

// TypesInfo is a nested struct in ehpc response
type TypesInfo struct {
	Status              string  `json:"Status" xml:"Status"`
	InstanceTypeId      string  `json:"InstanceTypeId" xml:"InstanceTypeId"`
	InstanceBandwidthRx int     `json:"InstanceBandwidthRx" xml:"InstanceBandwidthRx"`
	GPUSpec             string  `json:"GPUSpec" xml:"GPUSpec"`
	InstanceBandwidthTx int     `json:"InstanceBandwidthTx" xml:"InstanceBandwidthTx"`
	InstancePpsRx       int     `json:"InstancePpsRx" xml:"InstancePpsRx"`
	InstancePpsTx       int     `json:"InstancePpsTx" xml:"InstancePpsTx"`
	GPUAmount           int     `json:"GPUAmount" xml:"GPUAmount"`
	CpuCoreCount        int     `json:"CpuCoreCount" xml:"CpuCoreCount"`
	MemorySize          int     `json:"MemorySize" xml:"MemorySize"`
	EniQuantity         int     `json:"EniQuantity" xml:"EniQuantity"`
	ZoneIds             ZoneIds `json:"ZoneIds" xml:"ZoneIds"`
}