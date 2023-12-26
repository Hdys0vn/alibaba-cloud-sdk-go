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

// Content is a nested struct in mpaas response
type Content struct {
	ExecutionOrder          int64                                      `json:"ExecutionOrder" xml:"ExecutionOrder"`
	GreyUv                  int64                                      `json:"GreyUv" xml:"GreyUv"`
	PackageType             string                                     `json:"PackageType" xml:"PackageType"`
	UpgradeType             int64                                      `json:"UpgradeType" xml:"UpgradeType"`
	Expire                  string                                     `json:"Expire" xml:"Expire"`
	QrcodeUrl               string                                     `json:"QrcodeUrl" xml:"QrcodeUrl"`
	UpgradeContent          string                                     `json:"UpgradeContent" xml:"UpgradeContent"`
	IsOfficial              int64                                      `json:"IsOfficial" xml:"IsOfficial"`
	GreyConfigInfo          string                                     `json:"GreyConfigInfo" xml:"GreyConfigInfo"`
	DevicePercent           int64                                      `json:"DevicePercent" xml:"DevicePercent"`
	Accessid                string                                     `json:"Accessid" xml:"Accessid"`
	TaskStatus              int64                                      `json:"TaskStatus" xml:"TaskStatus"`
	Modifier                string                                     `json:"Modifier" xml:"Modifier"`
	Creator                 string                                     `json:"Creator" xml:"Creator"`
	NetType                 string                                     `json:"NetType" xml:"NetType"`
	Memo                    string                                     `json:"Memo" xml:"Memo"`
	GreyNotice              int64                                      `json:"GreyNotice" xml:"GreyNotice"`
	SyncResult              string                                     `json:"SyncResult" xml:"SyncResult"`
	Id                      int64                                      `json:"Id" xml:"Id"`
	CityContains            string                                     `json:"CityContains" xml:"CityContains"`
	ProductVersion          string                                     `json:"ProductVersion" xml:"ProductVersion"`
	Host                    string                                     `json:"Host" xml:"Host"`
	H5Id                    string                                     `json:"H5Id" xml:"H5Id"`
	AppId                   string                                     `json:"AppId" xml:"AppId"`
	UpgradeValidTime        int64                                      `json:"UpgradeValidTime" xml:"UpgradeValidTime"`
	ChannelExcludes         string                                     `json:"ChannelExcludes" xml:"ChannelExcludes"`
	SyncMode                string                                     `json:"SyncMode" xml:"SyncMode"`
	Dir                     string                                     `json:"Dir" xml:"Dir"`
	OsVersion               string                                     `json:"OsVersion" xml:"OsVersion"`
	ChannelContains         string                                     `json:"ChannelContains" xml:"ChannelContains"`
	GreyEndtimeData         string                                     `json:"GreyEndtimeData" xml:"GreyEndtimeData"`
	ProductId               string                                     `json:"ProductId" xml:"ProductId"`
	DownloadUrl             string                                     `json:"DownloadUrl" xml:"DownloadUrl"`
	WhitelistIds            string                                     `json:"WhitelistIds" xml:"WhitelistIds"`
	GreyNum                 int64                                      `json:"GreyNum" xml:"GreyNum"`
	PublishMode             int64                                      `json:"PublishMode" xml:"PublishMode"`
	MobileModelContains     string                                     `json:"MobileModelContains" xml:"MobileModelContains"`
	IsRelease               int64                                      `json:"IsRelease" xml:"IsRelease"`
	IsRc                    int64                                      `json:"IsRc" xml:"IsRc"`
	Policy                  string                                     `json:"Policy" xml:"Policy"`
	CityExcludes            string                                     `json:"CityExcludes" xml:"CityExcludes"`
	DeviceGreyNum           int64                                      `json:"DeviceGreyNum" xml:"DeviceGreyNum"`
	MobileModelExcludes     string                                     `json:"MobileModelExcludes" xml:"MobileModelExcludes"`
	H5Name                  string                                     `json:"H5Name" xml:"H5Name"`
	InnerVersion            string                                     `json:"InnerVersion" xml:"InnerVersion"`
	SilentType              int64                                      `json:"SilentType" xml:"SilentType"`
	IsPush                  int64                                      `json:"IsPush" xml:"IsPush"`
	GmtCreateStr            string                                     `json:"GmtCreateStr" xml:"GmtCreateStr"`
	PushContent             string                                     `json:"PushContent" xml:"PushContent"`
	ReleaseType             string                                     `json:"ReleaseType" xml:"ReleaseType"`
	Signature               string                                     `json:"Signature" xml:"Signature"`
	Platform                string                                     `json:"Platform" xml:"Platform"`
	IsEnterprise            int64                                      `json:"IsEnterprise" xml:"IsEnterprise"`
	EnableServerDomainCount string                                     `json:"EnableServerDomainCount" xml:"EnableServerDomainCount"`
	AppCode                 string                                     `json:"AppCode" xml:"AppCode"`
	Appstoreurl             string                                     `json:"Appstoreurl" xml:"Appstoreurl"`
	PackageInfoId           int64                                      `json:"PackageInfoId" xml:"PackageInfoId"`
	PublishType             int64                                      `json:"PublishType" xml:"PublishType"`
	PrivilegeSwitch         PrivilegeSwitch                            `json:"PrivilegeSwitch" xml:"PrivilegeSwitch"`
	ServerDomainConfigList  []ServerDomainConfigListItem               `json:"ServerDomainConfigList" xml:"ServerDomainConfigList"`
	ApiConfigList           []ApiConfigListItem                        `json:"ApiConfigList" xml:"ApiConfigList"`
	Whitelist               []WhitelistItemInQueryMdsUpgradeTaskDetail `json:"Whitelist" xml:"Whitelist"`
	RuleJsonList            []RuleJsonListItem                         `json:"RuleJsonList" xml:"RuleJsonList"`
	WebviewDomainConfigList []WebviewDomainConfigListItem              `json:"WebviewDomainConfigList" xml:"WebviewDomainConfigList"`
}