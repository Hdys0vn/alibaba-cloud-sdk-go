package swas_open

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

// InvokeInstance is a nested struct in swas_open response
type InvokeInstance struct {
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	StartTime        string `json:"StartTime" xml:"StartTime"`
	FinishTime       string `json:"FinishTime" xml:"FinishTime"`
	InvocationStatus string `json:"InvocationStatus" xml:"InvocationStatus"`
	Output           string `json:"Output" xml:"Output"`
	ExitCode         int64  `json:"ExitCode" xml:"ExitCode"`
	ErrorInfo        string `json:"ErrorInfo" xml:"ErrorInfo"`
	ErrorCode        string `json:"ErrorCode" xml:"ErrorCode"`
}
