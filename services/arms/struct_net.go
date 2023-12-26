package arms

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

// Net is a nested struct in arms response
type Net struct {
	NetIcmpSwitch        int64  `json:"NetIcmpSwitch" xml:"NetIcmpSwitch"`
	NetIcmpActive        int64  `json:"NetIcmpActive" xml:"NetIcmpActive"`
	NetIcmpTimeout       int64  `json:"NetIcmpTimeout" xml:"NetIcmpTimeout"`
	NetIcmpInterval      int64  `json:"NetIcmpInterval" xml:"NetIcmpInterval"`
	NetIcmpNum           int64  `json:"NetIcmpNum" xml:"NetIcmpNum"`
	NetIcmpSize          int64  `json:"NetIcmpSize" xml:"NetIcmpSize"`
	NetIcmpDataCut       int64  `json:"NetIcmpDataCut" xml:"NetIcmpDataCut"`
	NetDnsQueryMethod    string `json:"NetDnsQueryMethod" xml:"NetDnsQueryMethod"`
	NetDnsSwitch         int64  `json:"NetDnsSwitch" xml:"NetDnsSwitch"`
	NetTraceRouteSwitch  int64  `json:"NetTraceRouteSwitch" xml:"NetTraceRouteSwitch"`
	NetTraceRouteTimeout int64  `json:"NetTraceRouteTimeout" xml:"NetTraceRouteTimeout"`
	NetTraceRouteNum     int64  `json:"NetTraceRouteNum" xml:"NetTraceRouteNum"`
	WhiteList            string `json:"WhiteList" xml:"WhiteList"`
	NetDnsNs             string `json:"NetDnsNs" xml:"NetDnsNs"`
	NetDigSwitch         int64  `json:"NetDigSwitch" xml:"NetDigSwitch"`
	NetDnsServer         int64  `json:"NetDnsServer" xml:"NetDnsServer"`
	NetDnsTimeout        string `json:"NetDnsTimeout" xml:"NetDnsTimeout"`
}