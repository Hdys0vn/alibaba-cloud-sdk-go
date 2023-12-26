package vs

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

// RegionProportionData is a nested struct in vs response
type RegionProportionData struct {
	TotalQuery      string `json:"TotalQuery" xml:"TotalQuery"`
	TotalBytes      string `json:"TotalBytes" xml:"TotalBytes"`
	AvgResponseRate string `json:"AvgResponseRate" xml:"AvgResponseRate"`
	AvgResponseTime string `json:"AvgResponseTime" xml:"AvgResponseTime"`
	ReqErrRate      string `json:"ReqErrRate" xml:"ReqErrRate"`
	AvgObjectSize   string `json:"AvgObjectSize" xml:"AvgObjectSize"`
	Bps             string `json:"Bps" xml:"Bps"`
	Qps             string `json:"Qps" xml:"Qps"`
	RegionEname     string `json:"RegionEname" xml:"RegionEname"`
	Region          string `json:"Region" xml:"Region"`
	Proportion      string `json:"Proportion" xml:"Proportion"`
	BytesProportion string `json:"BytesProportion" xml:"BytesProportion"`
}
