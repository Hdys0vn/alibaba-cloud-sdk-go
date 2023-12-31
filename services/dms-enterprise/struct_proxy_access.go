package dms_enterprise

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

// ProxyAccess is a nested struct in dms_enterprise response
type ProxyAccess struct {
	ProxyAccessId int64  `json:"ProxyAccessId" xml:"ProxyAccessId"`
	GmtCreate     string `json:"GmtCreate" xml:"GmtCreate"`
	AccessId      string `json:"AccessId" xml:"AccessId"`
	OriginInfo    string `json:"OriginInfo" xml:"OriginInfo"`
	IndepAccount  string `json:"IndepAccount" xml:"IndepAccount"`
	UserId        int64  `json:"UserId" xml:"UserId"`
	UserUid       string `json:"UserUid" xml:"UserUid"`
	UserName      string `json:"UserName" xml:"UserName"`
	ProxyId       int64  `json:"ProxyId" xml:"ProxyId"`
	InstanceId    int64  `json:"InstanceId" xml:"InstanceId"`
}
