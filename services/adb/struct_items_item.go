package adb

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

// ItemsItem is a nested struct in adb response
type ItemsItem struct {
	RollbackSQL     string `json:"RollbackSQL" xml:"RollbackSQL"`
	Region          string `json:"Region" xml:"Region"`
	User            string `json:"User" xml:"User"`
	DBType          string `json:"DBType" xml:"DBType"`
	ResultInfo      string `json:"ResultInfo" xml:"ResultInfo"`
	PageSize        int64  `json:"PageSize" xml:"PageSize"`
	AvgTaskCount    string `json:"AvgTaskCount" xml:"AvgTaskCount"`
	AccessIP        string `json:"AccessIP" xml:"AccessIP"`
	Pattern         string `json:"Pattern" xml:"Pattern"`
	Deadline        string `json:"Deadline" xml:"Deadline"`
	MaxStageCount   string `json:"MaxStageCount" xml:"MaxStageCount"`
	Reason          string `json:"Reason" xml:"Reason"`
	ModifiedTime    string `json:"ModifiedTime" xml:"ModifiedTime"`
	AdviceDate      string `json:"AdviceDate" xml:"AdviceDate"`
	CreatedTime     string `json:"CreatedTime" xml:"CreatedTime"`
	MaxPeakMemory   string `json:"MaxPeakMemory" xml:"MaxPeakMemory"`
	AdviceId        string `json:"AdviceId" xml:"AdviceId"`
	SubmitTime      string `json:"SubmitTime" xml:"SubmitTime"`
	AvgPeakMemory   string `json:"AvgPeakMemory" xml:"AvgPeakMemory"`
	SwitchTime      string `json:"SwitchTime" xml:"SwitchTime"`
	BuildSQL        string `json:"BuildSQL" xml:"BuildSQL"`
	JobStatus       string `json:"JobStatus" xml:"JobStatus"`
	MaxScanSize     string `json:"MaxScanSize" xml:"MaxScanSize"`
	PageNumber      int64  `json:"PageNumber" xml:"PageNumber"`
	Id              int    `json:"Id" xml:"Id"`
	MaxCpuTime      string `json:"MaxCpuTime" xml:"MaxCpuTime"`
	DBVersion       string `json:"DBVersion" xml:"DBVersion"`
	TotalCount      int64  `json:"TotalCount" xml:"TotalCount"`
	ReportDate      string `json:"ReportDate" xml:"ReportDate"`
	InstanceName    string `json:"InstanceName" xml:"InstanceName"`
	TableSchema     string `json:"TableSchema" xml:"TableSchema"`
	AvgScanSize     string `json:"AvgScanSize" xml:"AvgScanSize"`
	SQL             string `json:"SQL" xml:"SQL"`
	AvgCpuTime      string `json:"AvgCpuTime" xml:"AvgCpuTime"`
	TableName       string `json:"TableName" xml:"TableName"`
	DBClusterId     string `json:"DBClusterId" xml:"DBClusterId"`
	AdviceType      string `json:"AdviceType" xml:"AdviceType"`
	PrepareInterval string `json:"PrepareInterval" xml:"PrepareInterval"`
	AccessCount     string `json:"AccessCount" xml:"AccessCount"`
	Benefit         string `json:"Benefit" xml:"Benefit"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	AvgStageCount   string `json:"AvgStageCount" xml:"AvgStageCount"`
	MaxTaskCount    string `json:"MaxTaskCount" xml:"MaxTaskCount"`
	SubmitStatus    string `json:"SubmitStatus" xml:"SubmitStatus"`
	Status          string `json:"Status" xml:"Status"`
	TaskType        string `json:"TaskType" xml:"TaskType"`
	QueryCount      string `json:"QueryCount" xml:"QueryCount"`
}
