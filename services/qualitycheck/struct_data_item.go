package qualitycheck

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

// DataItem is a nested struct in qualitycheck response
type DataItem struct {
	Id                  int64  `json:"Id" xml:"Id"`
	AliUid              int64  `json:"AliUid" xml:"AliUid"`
	SubAliUid           int64  `json:"SubAliUid" xml:"SubAliUid"`
	RoleName            string `json:"RoleName" xml:"RoleName"`
	LoginUserType       int    `json:"LoginUserType" xml:"LoginUserType"`
	BusinessSpaceCode   string `json:"BusinessSpaceCode" xml:"BusinessSpaceCode"`
	BusinessSpaceName   string `json:"BusinessSpaceName" xml:"BusinessSpaceName"`
	OrderInstanceId     string `json:"OrderInstanceId" xml:"OrderInstanceId"`
	XspaceProductCode   string `json:"XspaceProductCode" xml:"XspaceProductCode"`
	XspaceCommodityCode string `json:"XspaceCommodityCode" xml:"XspaceCommodityCode"`
	ProductType         string `json:"ProductType" xml:"ProductType"`
	Language            string `json:"Language" xml:"Language"`
	StartTime           string `json:"StartTime" xml:"StartTime"`
	EndTime             string `json:"EndTime" xml:"EndTime"`
	CurrentStatus       int    `json:"CurrentStatus" xml:"CurrentStatus"`
}