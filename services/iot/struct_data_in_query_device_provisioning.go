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

// DataInQueryDeviceProvisioning is a nested struct in iot response
type DataInQueryDeviceProvisioning struct {
	GmtCreate           int64  `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified         int64  `json:"GmtModified" xml:"GmtModified"`
	AliyunUid           string `json:"AliyunUid" xml:"AliyunUid"`
	ProductKey          string `json:"ProductKey" xml:"ProductKey"`
	DeviceName          string `json:"DeviceName" xml:"DeviceName"`
	SourceIotInstanceId string `json:"SourceIotInstanceId" xml:"SourceIotInstanceId"`
	TargetIotInstanceId string `json:"TargetIotInstanceId" xml:"TargetIotInstanceId"`
	SourceRegion        string `json:"SourceRegion" xml:"SourceRegion"`
	TargetRegion        string `json:"TargetRegion" xml:"TargetRegion"`
}