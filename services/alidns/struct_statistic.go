package alidns

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

// Statistic is a nested struct in alidns response
type Statistic struct {
	HttpCount    int64  `json:"HttpCount" xml:"HttpCount"`
	TotalCount   int64  `json:"TotalCount" xml:"TotalCount"`
	Timestamp    int64  `json:"Timestamp" xml:"Timestamp"`
	DomainType   string `json:"DomainType" xml:"DomainType"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
	V6HttpCount  int64  `json:"V6HttpCount" xml:"V6HttpCount"`
	SubDomain    string `json:"SubDomain" xml:"SubDomain"`
	V4HttpCount  int64  `json:"V4HttpCount" xml:"V4HttpCount"`
	V6HttpsCount int64  `json:"V6HttpsCount" xml:"V6HttpsCount"`
	IpCount      int64  `json:"IpCount" xml:"IpCount"`
	HttpsCount   int64  `json:"HttpsCount" xml:"HttpsCount"`
	Count        string `json:"Count" xml:"Count"`
	V4HttpsCount int64  `json:"V4HttpsCount" xml:"V4HttpsCount"`
}
