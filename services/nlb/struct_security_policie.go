package nlb

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

// SecurityPolicie is a nested struct in nlb response
type SecurityPolicie struct {
	RegionId             string            `json:"RegionId" xml:"RegionId"`
	ResourceGroupId      string            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Ciphers              string            `json:"Ciphers" xml:"Ciphers"`
	SecurityPolicyId     string            `json:"SecurityPolicyId" xml:"SecurityPolicyId"`
	SecurityPolicyName   string            `json:"SecurityPolicyName" xml:"SecurityPolicyName"`
	SecurityPolicyStatus string            `json:"SecurityPolicyStatus" xml:"SecurityPolicyStatus"`
	TlsVersion           string            `json:"TlsVersion" xml:"TlsVersion"`
	RelatedListeners     []RelatedListener `json:"RelatedListeners" xml:"RelatedListeners"`
	Tags                 []Tag             `json:"Tags" xml:"Tags"`
}
