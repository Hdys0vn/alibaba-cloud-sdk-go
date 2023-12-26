package linkvisual

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

// ResultListItem is a nested struct in linkvisual response
type ResultListItem struct {
	SlotEndTime   int            `json:"SlotEndTime" xml:"SlotEndTime"`
	ProductKey    string         `json:"ProductKey" xml:"ProductKey"`
	DeviceName    string         `json:"DeviceName" xml:"DeviceName"`
	SlotStartTime int            `json:"SlotStartTime" xml:"SlotStartTime"`
	Code          int            `json:"Code" xml:"Code"`
	SlotStatus    int            `json:"SlotStatus" xml:"SlotStatus"`
	IotId         string         `json:"IotId" xml:"IotId"`
	FileList      []FileListItem `json:"FileList" xml:"FileList"`
}