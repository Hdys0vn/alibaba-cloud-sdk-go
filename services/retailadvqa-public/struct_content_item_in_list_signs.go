package retailadvqa_public

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

// ContentItemInListSigns is a nested struct in retailadvqa_public response
type ContentItemInListSigns struct {
	Id             string `json:"Id" xml:"Id"`
	PlatformName   string `json:"PlatformName" xml:"PlatformName"`
	GmtCreate      int64  `json:"GmtCreate" xml:"GmtCreate"`
	SignStatus     int64  `json:"SignStatus" xml:"SignStatus"`
	QaAccountId    string `json:"QaAccountId" xml:"QaAccountId"`
	CreateUserName string `json:"CreateUserName" xml:"CreateUserName"`
	SignName       string `json:"SignName" xml:"SignName"`
	ErrDescription string `json:"ErrDescription" xml:"ErrDescription"`
}
