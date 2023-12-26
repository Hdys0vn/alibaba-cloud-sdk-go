package ens

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

// ActivityInfos is a nested struct in ens response
type ActivityInfos struct {
	ActivityId   string `json:"ActivityId" xml:"ActivityId"`
	ActivityName string `json:"ActivityName" xml:"ActivityName"`
	State        string `json:"State" xml:"State"`
	GmtStart     string `json:"GmtStart" xml:"GmtStart"`
	GmtEnd       string `json:"GmtEnd" xml:"GmtEnd"`
	Duration     string `json:"Duration" xml:"Duration"`
	GmtCreate    string `json:"GmtCreate" xml:"GmtCreate"`
	Method       string `json:"Method" xml:"Method"`
	Input        string `json:"Input" xml:"Input"`
	Output       string `json:"Output" xml:"Output"`
	Error        string `json:"Error" xml:"Error"`
	WorkerNode   string `json:"WorkerNode" xml:"WorkerNode"`
}