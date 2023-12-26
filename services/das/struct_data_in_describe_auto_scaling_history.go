package das

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

// DataInDescribeAutoScalingHistory is a nested struct in das response
type DataInDescribeAutoScalingHistory struct {
	InstanceId  string                   `json:"InstanceId" xml:"InstanceId"`
	Storage     []map[string]interface{} `json:"Storage" xml:"Storage"`
	Resource    []map[string]interface{} `json:"Resource" xml:"Resource"`
	Shard       []map[string]interface{} `json:"Shard" xml:"Shard"`
	Bandwidth   []map[string]interface{} `json:"Bandwidth" xml:"Bandwidth"`
	SpecHistory []SpecHistoryItem        `json:"SpecHistory" xml:"SpecHistory"`
}