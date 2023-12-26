package config

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

// CompliancePacksItem is a nested struct in config response
type CompliancePacksItem struct {
	CompliancePackId         string `json:"CompliancePackId" xml:"CompliancePackId"`
	RiskLevel                int    `json:"RiskLevel" xml:"RiskLevel"`
	CreateTimestamp          int64  `json:"CreateTimestamp" xml:"CreateTimestamp"`
	AccountId                int64  `json:"AccountId" xml:"AccountId"`
	AggregatorId             string `json:"AggregatorId" xml:"AggregatorId"`
	Status                   string `json:"Status" xml:"Status"`
	CompliancePackTemplateId string `json:"CompliancePackTemplateId" xml:"CompliancePackTemplateId"`
	Description              string `json:"Description" xml:"Description"`
	CompliancePackName       string `json:"CompliancePackName" xml:"CompliancePackName"`
}
