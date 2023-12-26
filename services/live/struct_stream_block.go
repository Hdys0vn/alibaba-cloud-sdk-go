package live

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

// StreamBlock is a nested struct in live response
type StreamBlock struct {
	Status       int    `json:"Status" xml:"Status"`
	BlockType    string `json:"BlockType" xml:"BlockType"`
	AppName      string `json:"AppName" xml:"AppName"`
	UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
	StreamName   string `json:"StreamName" xml:"StreamName"`
	ReleaseTime  string `json:"ReleaseTime" xml:"ReleaseTime"`
	LocationList string `json:"LocationList" xml:"LocationList"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
}