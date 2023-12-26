package csas

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

// Device is a nested struct in csas response
type Device struct {
	DeviceTag     string             `json:"DeviceTag" xml:"DeviceTag"`
	DeviceType    string             `json:"DeviceType" xml:"DeviceType"`
	DeviceModel   string             `json:"DeviceModel" xml:"DeviceModel"`
	DeviceVersion string             `json:"DeviceVersion" xml:"DeviceVersion"`
	Hostname      string             `json:"Hostname" xml:"Hostname"`
	Username      string             `json:"Username" xml:"Username"`
	SaseUserId    string             `json:"SaseUserId" xml:"SaseUserId"`
	Department    string             `json:"Department" xml:"Department"`
	InnerIP       string             `json:"InnerIP" xml:"InnerIP"`
	SrcIP         string             `json:"SrcIP" xml:"SrcIP"`
	Memory        string             `json:"Memory" xml:"Memory"`
	CPU           string             `json:"CPU" xml:"CPU"`
	Disk          string             `json:"Disk" xml:"Disk"`
	Mac           string             `json:"Mac" xml:"Mac"`
	AppVersion    string             `json:"AppVersion" xml:"AppVersion"`
	DeviceBelong  string             `json:"DeviceBelong" xml:"DeviceBelong"`
	SharingStatus bool               `json:"SharingStatus" xml:"SharingStatus"`
	DeviceStatus  string             `json:"DeviceStatus" xml:"DeviceStatus"`
	AppStatus     string             `json:"AppStatus" xml:"AppStatus"`
	PaStatus      string             `json:"PaStatus" xml:"PaStatus"`
	IaStatus      string             `json:"IaStatus" xml:"IaStatus"`
	DlpStatus     string             `json:"DlpStatus" xml:"DlpStatus"`
	NacStatus     string             `json:"NacStatus" xml:"NacStatus"`
	CreateTime    string             `json:"CreateTime" xml:"CreateTime"`
	UpdateTime    string             `json:"UpdateTime" xml:"UpdateTime"`
	HistoryUsers  []HistoryUsersItem `json:"HistoryUsers" xml:"HistoryUsers"`
}