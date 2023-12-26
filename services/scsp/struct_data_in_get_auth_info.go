package scsp

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

// DataInGetAuthInfo is a nested struct in scsp response
type DataInGetAuthInfo struct {
	AppName   string `json:"AppName" xml:"AppName"`
	Time      int64  `json:"Time" xml:"Time"`
	AppKey    string `json:"AppKey" xml:"AppKey"`
	App       string `json:"App" xml:"App"`
	UserId    string `json:"UserId" xml:"UserId"`
	Code      string `json:"Code" xml:"Code"`
	SessionId string `json:"SessionId" xml:"SessionId"`
	UserName  string `json:"UserName" xml:"UserName"`
	TenantId  string `json:"TenantId" xml:"TenantId"`
}
