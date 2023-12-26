package csb

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

// OrderInFindOrderedList is a nested struct in csb response
type OrderInFindOrderedList struct {
	Alias                 string              `json:"Alias" xml:"Alias"`
	ProjectName           string              `json:"ProjectName" xml:"ProjectName"`
	ServiceName           string              `json:"ServiceName" xml:"ServiceName"`
	ServiceVersion        string              `json:"ServiceVersion" xml:"ServiceVersion"`
	OrderStatus           int                 `json:"OrderStatus" xml:"OrderStatus"`
	AliveOrderCount       int                 `json:"AliveOrderCount" xml:"AliveOrderCount"`
	GmtCreate             int64               `json:"GmtCreate" xml:"GmtCreate"`
	MaxRT                 int                 `json:"MaxRT" xml:"MaxRT"`
	MinRT                 int                 `json:"MinRT" xml:"MinRT"`
	ServiceId             string              `json:"ServiceId" xml:"ServiceId"`
	ServiceStatus         int                 `json:"ServiceStatus" xml:"ServiceStatus"`
	Total                 Total               `json:"Total" xml:"Total"`
	ErrorTypeCatagoryList []ErrorTypeCatagory `json:"ErrorTypeCatagoryList" xml:"ErrorTypeCatagoryList"`
	Orders                []Order             `json:"Orders" xml:"Orders"`
}
