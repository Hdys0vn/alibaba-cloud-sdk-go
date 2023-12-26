package eipanycast

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

// Anycast is a nested struct in eipanycast response
type Anycast struct {
	Status                 string               `json:"Status" xml:"Status"`
	CreateTime             string               `json:"CreateTime" xml:"CreateTime"`
	AnycastId              string               `json:"AnycastId" xml:"AnycastId"`
	AliUid                 int64                `json:"AliUid" xml:"AliUid"`
	Bid                    string               `json:"Bid" xml:"Bid"`
	ServiceLocation        string               `json:"ServiceLocation" xml:"ServiceLocation"`
	InstanceChargeType     string               `json:"InstanceChargeType" xml:"InstanceChargeType"`
	IpAddress              string               `json:"IpAddress" xml:"IpAddress"`
	Bandwidth              int                  `json:"Bandwidth" xml:"Bandwidth"`
	Description            string               `json:"Description" xml:"Description"`
	InternetChargeType     string               `json:"InternetChargeType" xml:"InternetChargeType"`
	BusinessStatus         string               `json:"BusinessStatus" xml:"BusinessStatus"`
	Name                   string               `json:"Name" xml:"Name"`
	ServiceManaged         int                  `json:"ServiceManaged" xml:"ServiceManaged"`
	ResourceGroupId        string               `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AnycastEipBindInfoList []AnycastEipBindInfo `json:"AnycastEipBindInfoList" xml:"AnycastEipBindInfoList"`
	Tags                   []Tag                `json:"Tags" xml:"Tags"`
}
