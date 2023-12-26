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

// Plan is a nested struct in swas_open response
type Plan struct {
	Bandwidth       int    `json:"Bandwidth" xml:"Bandwidth"`
	Core            int    `json:"Core" xml:"Core"`
	SupportPlatform string `json:"SupportPlatform" xml:"SupportPlatform"`
	Memory          int    `json:"Memory" xml:"Memory"`
	DiskType        string `json:"DiskType" xml:"DiskType"`
	DiskSize        int    `json:"DiskSize" xml:"DiskSize"`
	PlanId          string `json:"PlanId" xml:"PlanId"`
	Flow            int    `json:"Flow" xml:"Flow"`
	OriginPrice     string `json:"OriginPrice" xml:"OriginPrice"`
	Currency        string `json:"Currency" xml:"Currency"`
}
