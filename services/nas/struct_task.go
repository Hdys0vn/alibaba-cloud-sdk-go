package nas

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

// Task is a nested struct in nas response
type Task struct {
	FilesystemId   string `json:"FilesystemId" xml:"FilesystemId"`
	DataFlowId     string `json:"DataFlowId" xml:"DataFlowId"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	SourceStorage  string `json:"SourceStorage" xml:"SourceStorage"`
	FileSystemPath string `json:"FileSystemPath" xml:"FileSystemPath"`
	Originator     string `json:"Originator" xml:"Originator"`
	TaskAction     string `json:"TaskAction" xml:"TaskAction"`
	DataType       string `json:"DataType" xml:"DataType"`
	Progress       int64  `json:"Progress" xml:"Progress"`
	Status         string `json:"Status" xml:"Status"`
	ReportPath     string `json:"ReportPath" xml:"ReportPath"`
	CreateTime     string `json:"CreateTime" xml:"CreateTime"`
	StartTime      string `json:"StartTime" xml:"StartTime"`
	EndTime        string `json:"EndTime" xml:"EndTime"`
	FsPath         string `json:"FsPath" xml:"FsPath"`
}
