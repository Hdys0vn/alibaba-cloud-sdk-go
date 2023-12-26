package mpaas

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

// MiniPackageListItem is a nested struct in mpaas response
type MiniPackageListItem struct {
	Status           int64  `json:"Status" xml:"Status"`
	DownloadUrl      string `json:"DownloadUrl" xml:"DownloadUrl"`
	AppCode          string `json:"AppCode" xml:"AppCode"`
	Memo             string `json:"Memo" xml:"Memo"`
	H5Id             string `json:"H5Id" xml:"H5Id"`
	H5Version        string `json:"H5Version" xml:"H5Version"`
	PublishPeriod    int64  `json:"PublishPeriod" xml:"PublishPeriod"`
	GmtModified      string `json:"GmtModified" xml:"GmtModified"`
	PackageType      int64  `json:"PackageType" xml:"PackageType"`
	ExtendInfo       string `json:"ExtendInfo" xml:"ExtendInfo"`
	FallbackBaseUrl  string `json:"FallbackBaseUrl" xml:"FallbackBaseUrl"`
	ResourceType     int64  `json:"ResourceType" xml:"ResourceType"`
	AutoInstall      int64  `json:"AutoInstall" xml:"AutoInstall"`
	InstallType      int64  `json:"InstallType" xml:"InstallType"`
	ClientVersionMax string `json:"ClientVersionMax" xml:"ClientVersionMax"`
	Platform         string `json:"Platform" xml:"Platform"`
	H5Name           string `json:"H5Name" xml:"H5Name"`
	GmtCreate        string `json:"GmtCreate" xml:"GmtCreate"`
	ClientVersionMin string `json:"ClientVersionMin" xml:"ClientVersionMin"`
	MainUrl          string `json:"MainUrl" xml:"MainUrl"`
	Id               int64  `json:"Id" xml:"Id"`
	ExtraData        string `json:"ExtraData" xml:"ExtraData"`
}