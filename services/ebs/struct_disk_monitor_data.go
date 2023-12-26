package ebs

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

// DiskMonitorData is a nested struct in ebs response
type DiskMonitorData struct {
	WriteIOPS    int64  `json:"WriteIOPS" xml:"WriteIOPS"`
	ReadBPS      int64  `json:"ReadBPS" xml:"ReadBPS"`
	BPSPercent   int64  `json:"BPSPercent" xml:"BPSPercent"`
	Timestamp    string `json:"Timestamp" xml:"Timestamp"`
	IOPSPercent  int64  `json:"IOPSPercent" xml:"IOPSPercent"`
	BurstIOCount int64  `json:"BurstIOCount" xml:"BurstIOCount"`
	DiskId       string `json:"DiskId" xml:"DiskId"`
	WriteBPS     int64  `json:"WriteBPS" xml:"WriteBPS"`
	ReadIOPS     int64  `json:"ReadIOPS" xml:"ReadIOPS"`
}
