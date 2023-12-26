package aliyuncvc

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

// Data is a nested struct in aliyuncvc response
type Data struct {
	DeviceModel     string                `json:"DeviceModel" xml:"DeviceModel"`
	ConferenceCode  string                `json:"ConferenceCode" xml:"ConferenceCode"`
	CreateTime      string                `json:"CreateTime" xml:"CreateTime"`
	ConferenceName  string                `json:"ConferenceName" xml:"ConferenceName"`
	DeviceErrorCode int                   `json:"DeviceErrorCode" xml:"DeviceErrorCode"`
	PageSize        int                   `json:"PageSize" xml:"PageSize"`
	SerialNumber    string                `json:"SerialNumber" xml:"SerialNumber"`
	PictureUrl      string                `json:"PictureUrl" xml:"PictureUrl"`
	Total           int                   `json:"Total" xml:"Total"`
	TotalCount      int                   `json:"TotalCount" xml:"TotalCount"`
	Manufacturer    string                `json:"Manufacturer" xml:"Manufacturer"`
	ScreenCode      string                `json:"ScreenCode" xml:"ScreenCode"`
	PageNumber      int                   `json:"PageNumber" xml:"PageNumber"`
	DeviceMessage   string                `json:"DeviceMessage" xml:"DeviceMessage"`
	ActivationCode  string                `json:"ActivationCode" xml:"ActivationCode"`
	Status          string                `json:"Status" xml:"Status"`
	CastScreenCode  string                `json:"CastScreenCode" xml:"CastScreenCode"`
	SN              string                `json:"SN" xml:"SN"`
	ActiveCode      string                `json:"ActiveCode" xml:"ActiveCode"`
	UserInfos       []UserInfo            `json:"UserInfos" xml:"UserInfos"`
	Devices         []DeviceInListDevices `json:"Devices" xml:"Devices"`
}