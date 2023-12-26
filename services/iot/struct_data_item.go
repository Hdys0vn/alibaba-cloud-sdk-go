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

// DataItem is a nested struct in iot response
type DataItem struct {
	ProductKey      string `json:"ProductKey" xml:"ProductKey"`
	ProductName     string `json:"ProductName" xml:"ProductName"`
	DeviceName      string `json:"DeviceName" xml:"DeviceName"`
	Nickname        string `json:"Nickname" xml:"Nickname"`
	DeviceSecret    string `json:"DeviceSecret" xml:"DeviceSecret"`
	IotId           string `json:"IotId" xml:"IotId"`
	UtcCreate       string `json:"UtcCreate" xml:"UtcCreate"`
	GmtCreate       string `json:"GmtCreate" xml:"GmtCreate"`
	UtcActive       string `json:"UtcActive" xml:"UtcActive"`
	GmtActive       string `json:"GmtActive" xml:"GmtActive"`
	Status          string `json:"Status" xml:"Status"`
	FirmwareVersion string `json:"FirmwareVersion" xml:"FirmwareVersion"`
	NodeType        int    `json:"NodeType" xml:"NodeType"`
	Region          string `json:"Region" xml:"Region"`
}
