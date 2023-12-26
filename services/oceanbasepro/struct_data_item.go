package oceanbasepro

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

// DataItem is a nested struct in oceanbasepro response
type DataItem struct {
	StepStatus                string                 `json:"StepStatus" xml:"StepStatus"`
	MaxCpu                    int64                  `json:"MaxCpu" xml:"MaxCpu"`
	LastExecutedTime          string                 `json:"LastExecutedTime" xml:"LastExecutedTime"`
	Inner                     bool                   `json:"Inner" xml:"Inner"`
	BlockCacheHit             string                 `json:"BlockCacheHit" xml:"BlockCacheHit"`
	AvgDbTime                 string                 `json:"AvgDbTime" xml:"AvgDbTime"`
	BlockIndexCacheHit        string                 `json:"BlockIndexCacheHit" xml:"BlockIndexCacheHit"`
	CpuTime                   string                 `json:"CpuTime" xml:"CpuTime"`
	ExecutorRpc               bool                   `json:"ExecutorRpc" xml:"ExecutorRpc"`
	ElapsedTime               string                 `json:"ElapsedTime" xml:"ElapsedTime"`
	ExecuteTime               string                 `json:"ExecuteTime" xml:"ExecuteTime"`
	FullSqlText               string                 `json:"FullSqlText" xml:"FullSqlText"`
	ProjectName               string                 `json:"ProjectName" xml:"ProjectName"`
	ProjectId                 string                 `json:"ProjectId" xml:"ProjectId"`
	ConcurrencyWaitTime       string                 `json:"ConcurrencyWaitTime" xml:"ConcurrencyWaitTime"`
	ConsistencyLevel          string                 `json:"ConsistencyLevel" xml:"ConsistencyLevel"`
	AvgGetPlanTime            string                 `json:"AvgGetPlanTime" xml:"AvgGetPlanTime"`
	RpcCount                  string                 `json:"RpcCount" xml:"RpcCount"`
	SsstoreReadRows           string                 `json:"SsstoreReadRows" xml:"SsstoreReadRows"`
	PartitionCount            string                 `json:"PartitionCount" xml:"PartitionCount"`
	StepName                  string                 `json:"StepName" xml:"StepName"`
	ExpectedWorkerCount       string                 `json:"ExpectedWorkerCount" xml:"ExpectedWorkerCount"`
	ProjectOwner              string                 `json:"ProjectOwner" xml:"ProjectOwner"`
	Executions                string                 `json:"Executions" xml:"Executions"`
	SqlId                     string                 `json:"SqlId" xml:"SqlId"`
	MemstoreReadRows          string                 `json:"MemstoreReadRows" xml:"MemstoreReadRows"`
	SqlTextShort              string                 `json:"SqlTextShort" xml:"SqlTextShort"`
	NetWaitTime               string                 `json:"NetWaitTime" xml:"NetWaitTime"`
	DbName                    string                 `json:"DbName" xml:"DbName"`
	RetCode                   string                 `json:"RetCode" xml:"RetCode"`
	FinishTime                string                 `json:"FinishTime" xml:"FinishTime"`
	AffectedRows              string                 `json:"AffectedRows" xml:"AffectedRows"`
	StepProgress              int                    `json:"StepProgress" xml:"StepProgress"`
	Interactive               bool                   `json:"Interactive" xml:"Interactive"`
	SqlType                   string                 `json:"SqlType" xml:"SqlType"`
	StartTime                 string                 `json:"StartTime" xml:"StartTime"`
	RetryCount                string                 `json:"RetryCount" xml:"RetryCount"`
	Suggestion                string                 `json:"Suggestion" xml:"Suggestion"`
	HitPlan                   bool                   `json:"HitPlan" xml:"HitPlan"`
	Statement                 string                 `json:"Statement" xml:"Statement"`
	ObServerId                string                 `json:"ObServerId" xml:"ObServerId"`
	SumDbTime                 string                 `json:"SumDbTime" xml:"SumDbTime"`
	EstimatedRemainingSeconds int64                  `json:"EstimatedRemainingSeconds" xml:"EstimatedRemainingSeconds"`
	BusinessName              string                 `json:"BusinessName" xml:"BusinessName"`
	ApplicationWaitTime       string                 `json:"ApplicationWaitTime" xml:"ApplicationWaitTime"`
	RequestTime               string                 `json:"RequestTime" xml:"RequestTime"`
	SumElapsedTime            string                 `json:"SumElapsedTime" xml:"SumElapsedTime"`
	ClientIp                  string                 `json:"ClientIp" xml:"ClientIp"`
	QueueTime                 string                 `json:"QueueTime" xml:"QueueTime"`
	TableScan                 bool                   `json:"TableScan" xml:"TableScan"`
	Diagnosis                 string                 `json:"Diagnosis" xml:"Diagnosis"`
	Metric                    string                 `json:"Metric" xml:"Metric"`
	Server                    string                 `json:"Server" xml:"Server"`
	ObUserId                  string                 `json:"ObUserId" xml:"ObUserId"`
	AvgCpuTime                string                 `json:"AvgCpuTime" xml:"AvgCpuTime"`
	DecodeTime                string                 `json:"DecodeTime" xml:"DecodeTime"`
	DiskReads                 string                 `json:"DiskReads" xml:"DiskReads"`
	MinCpu                    int64                  `json:"MinCpu" xml:"MinCpu"`
	ClientPort                string                 `json:"ClientPort" xml:"ClientPort"`
	ScheduleTime              string                 `json:"ScheduleTime" xml:"ScheduleTime"`
	DynamicSql                bool                   `json:"DynamicSql" xml:"DynamicSql"`
	WaitEvent                 string                 `json:"WaitEvent" xml:"WaitEvent"`
	GetPlanTime               string                 `json:"GetPlanTime" xml:"GetPlanTime"`
	WaitCount                 string                 `json:"WaitCount" xml:"WaitCount"`
	PlanId                    string                 `json:"PlanId" xml:"PlanId"`
	TransHash                 string                 `json:"TransHash" xml:"TransHash"`
	AvgElapsedTime            string                 `json:"AvgElapsedTime" xml:"AvgElapsedTime"`
	StepDescription           string                 `json:"StepDescription" xml:"StepDescription"`
	TraceId                   string                 `json:"TraceId" xml:"TraceId"`
	WaitTime                  string                 `json:"WaitTime" xml:"WaitTime"`
	Tags                      map[string]interface{} `json:"Tags" xml:"Tags"`
	ObDbId                    string                 `json:"ObDbId" xml:"ObDbId"`
	PlanType                  string                 `json:"PlanType" xml:"PlanType"`
	NetTime                   string                 `json:"NetTime" xml:"NetTime"`
	UserIoWaitTime            string                 `json:"UserIoWaitTime" xml:"UserIoWaitTime"`
	StepOrder                 int                    `json:"StepOrder" xml:"StepOrder"`
	UnitNum                   int64                  `json:"UnitNum" xml:"UnitNum"`
	UsedWorkerCount           string                 `json:"UsedWorkerCount" xml:"UsedWorkerCount"`
	ReturnRows                string                 `json:"ReturnRows" xml:"ReturnRows"`
	RequestId                 string                 `json:"RequestId" xml:"RequestId"`
	BloomFilterCacheHit       string                 `json:"BloomFilterCacheHit" xml:"BloomFilterCacheHit"`
	RowCacheHit               string                 `json:"RowCacheHit" xml:"RowCacheHit"`
	UserName                  string                 `json:"UserName" xml:"UserName"`
	RiskLevel                 string                 `json:"RiskLevel" xml:"RiskLevel"`
	DiagTypes                 []string               `json:"DiagTypes" xml:"DiagTypes"`
	DestConfig                DestConfig             `json:"DestConfig" xml:"DestConfig"`
	StepInfo                  StepInfo               `json:"StepInfo" xml:"StepInfo"`
	ExtraInfo                 ExtraInfo              `json:"ExtraInfo" xml:"ExtraInfo"`
	TransferStepConfig        TransferStepConfig     `json:"TransferStepConfig" xml:"TransferStepConfig"`
	SourceConfig              SourceConfig           `json:"SourceConfig" xml:"SourceConfig"`
	TransferMapping           TransferMapping        `json:"TransferMapping" xml:"TransferMapping"`
	DataPoints                []DataPoint            `json:"DataPoints" xml:"DataPoints"`
	Steps                     []Step                 `json:"Steps" xml:"Steps"`
	Labels                    []Label                `json:"Labels" xml:"Labels"`
	SqlList                   []SqlListItem          `json:"SqlList" xml:"SqlList"`
}
