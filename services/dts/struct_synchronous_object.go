package dts

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

// SynchronousObject is a nested struct in dts response
type SynchronousObject struct {
	DatabaseName  string    `json:"DatabaseName" xml:"DatabaseName"`
	WholeDatabase string    `json:"WholeDatabase" xml:"WholeDatabase"`
	TableList     TableList `json:"TableList" xml:"TableList"`
}
