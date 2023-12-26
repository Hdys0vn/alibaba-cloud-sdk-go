package opensearch

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

// ResultInGetFunctionTask is a nested struct in opensearch response
type ResultInGetFunctionTask struct {
	EndTime      int64  `json:"EndTime" xml:"EndTime"`
	ExtendInfo   string `json:"ExtendInfo" xml:"ExtendInfo"`
	FunctionName string `json:"FunctionName" xml:"FunctionName"`
	Generation   string `json:"Generation" xml:"Generation"`
	Progress     int64  `json:"Progress" xml:"Progress"`
	RunId        string `json:"RunId" xml:"RunId"`
	StartTime    int64  `json:"StartTime" xml:"StartTime"`
	Status       string `json:"Status" xml:"Status"`
}
