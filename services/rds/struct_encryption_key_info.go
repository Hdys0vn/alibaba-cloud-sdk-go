package rds

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

// EncryptionKeyInfo is a nested struct in rds response
type EncryptionKeyInfo struct {
	KeyType             string `json:"KeyType" xml:"KeyType"`
	EncryptionKey       string `json:"EncryptionKey" xml:"EncryptionKey"`
	Description         string `json:"Description" xml:"Description"`
	KeyUsage            string `json:"KeyUsage" xml:"KeyUsage"`
	DeleteDate          string `json:"DeleteDate" xml:"DeleteDate"`
	Creator             string `json:"Creator" xml:"Creator"`
	EncryptionKeyStatus string `json:"EncryptionKeyStatus" xml:"EncryptionKeyStatus"`
	Origin              string `json:"Origin" xml:"Origin"`
	MaterialExpireTime  string `json:"MaterialExpireTime" xml:"MaterialExpireTime"`
	AliasName           string `json:"AliasName" xml:"AliasName"`
	UsedBy              string `json:"UsedBy" xml:"UsedBy"`
}
