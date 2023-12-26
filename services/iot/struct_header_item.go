package iot

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

// HeaderItem is a nested struct in iot response
type HeaderItem struct {
	FieldName    string `json:"FieldName" xml:"FieldName"`
	Alias        string `json:"Alias" xml:"Alias"`
	FieldType    string `json:"FieldType" xml:"FieldType"`
	TypeClass    string `json:"TypeClass" xml:"TypeClass"`
	TimeClass    string `json:"TimeClass" xml:"TimeClass"`
	GeoClass     string `json:"GeoClass" xml:"GeoClass"`
	DimDateClass string `json:"DimDateClass" xml:"DimDateClass"`
}
