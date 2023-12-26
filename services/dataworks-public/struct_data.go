package dataworks_public

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

// Data is a nested struct in dataworks_public response
type Data struct {
	LastModifyTime                int64                          `json:"LastModifyTime" xml:"LastModifyTime"`
	CreatorName                   string                         `json:"CreatorName" xml:"CreatorName"`
	ApplicationSecret             string                         `json:"ApplicationSecret" xml:"ApplicationSecret"`
	WhereCondition                string                         `json:"WhereCondition" xml:"WhereCondition"`
	ViewCount                     int64                          `json:"ViewCount" xml:"ViewCount"`
	TaskId                        string                         `json:"TaskId" xml:"TaskId"`
	Content                       string                         `json:"Content" xml:"Content"`
	ProjectName                   string                         `json:"ProjectName" xml:"ProjectName"`
	EnvType                       int                            `json:"EnvType" xml:"EnvType"`
	Message                       string                         `json:"Message" xml:"Message"`
	IsView                        bool                           `json:"IsView" xml:"IsView"`
	FinishTime                    int64                          `json:"FinishTime" xml:"FinishTime"`
	Code                          string                         `json:"Code" xml:"Code"`
	Connection                    string                         `json:"Connection" xml:"Connection"`
	PredictType                   int                            `json:"PredictType" xml:"PredictType"`
	CostTime                      string                         `json:"CostTime" xml:"CostTime"`
	TaskContent                   string                         `json:"TaskContent" xml:"TaskContent"`
	Enabled                       bool                           `json:"Enabled" xml:"Enabled"`
	ExpHour                       int                            `json:"ExpHour" xml:"ExpHour"`
	Status                        string                         `json:"Status" xml:"Status"`
	Name                          string                         `json:"Name" xml:"Name"`
	ModifyTime                    int64                          `json:"ModifyTime" xml:"ModifyTime"`
	ExpTime                       int64                          `json:"ExpTime" xml:"ExpTime"`
	Detail                        string                         `json:"Detail" xml:"Detail"`
	EndCast                       int64                          `json:"EndCast" xml:"EndCast"`
	LastAccessTime                int64                          `json:"LastAccessTime" xml:"LastAccessTime"`
	BusinessId                    int64                          `json:"BusinessId" xml:"BusinessId"`
	BaselineName                  string                         `json:"BaselineName" xml:"BaselineName"`
	Bizdate                       int64                          `json:"Bizdate" xml:"Bizdate"`
	OpenSwitch                    bool                           `json:"OpenSwitch" xml:"OpenSwitch"`
	TotalColumnCount              int64                          `json:"TotalColumnCount" xml:"TotalColumnCount"`
	NextPrimaryKey                string                         `json:"NextPrimaryKey" xml:"NextPrimaryKey"`
	MigrationId                   int64                          `json:"MigrationId" xml:"MigrationId"`
	CommitUser                    string                         `json:"CommitUser" xml:"CommitUser"`
	AlertMarginThreshold          int                            `json:"AlertMarginThreshold" xml:"AlertMarginThreshold"`
	NextToken                     string                         `json:"NextToken" xml:"NextToken"`
	UseFlag                       bool                           `json:"UseFlag" xml:"UseFlag"`
	DqcDescription                string                         `json:"DqcDescription" xml:"DqcDescription"`
	CheckerName                   string                         `json:"CheckerName" xml:"CheckerName"`
	RemindType                    string                         `json:"RemindType" xml:"RemindType"`
	UseType                       string                         `json:"UseType" xml:"UseType"`
	MethodId                      int                            `json:"MethodId" xml:"MethodId"`
	DagType                       string                         `json:"DagType" xml:"DagType"`
	FixCheck                      bool                           `json:"FixCheck" xml:"FixCheck"`
	FileVersion                   int                            `json:"FileVersion" xml:"FileVersion"`
	TestId                        string                         `json:"TestId" xml:"TestId"`
	SlaMinu                       int                            `json:"SlaMinu" xml:"SlaMinu"`
	NodeName                      string                         `json:"NodeName" xml:"NodeName"`
	RunStats                      map[string]interface{}         `json:"RunStats" xml:"RunStats"`
	Gmtdate                       int64                          `json:"Gmtdate" xml:"Gmtdate"`
	AlertEnabled                  bool                           `json:"AlertEnabled" xml:"AlertEnabled"`
	FileContent                   string                         `json:"FileContent" xml:"FileContent"`
	LifeCycle                     int                            `json:"LifeCycle" xml:"LifeCycle"`
	OnDutyAccountName             string                         `json:"OnDutyAccountName" xml:"OnDutyAccountName"`
	Checker                       int                            `json:"Checker" xml:"Checker"`
	Schema                        string                         `json:"Schema" xml:"Schema"`
	Version                       int64                          `json:"Version" xml:"Version"`
	IsCurrentProd                 bool                           `json:"IsCurrentProd" xml:"IsCurrentProd"`
	UpdatedUid                    string                         `json:"UpdatedUid" xml:"UpdatedUid"`
	GmtCreate                     int64                          `json:"GmtCreate" xml:"GmtCreate"`
	ColumnCount                   int                            `json:"ColumnCount" xml:"ColumnCount"`
	Caption                       string                         `json:"Caption" xml:"Caption"`
	AppGuid                       string                         `json:"AppGuid" xml:"AppGuid"`
	NodeId                        int64                          `json:"NodeId" xml:"NodeId"`
	ClusterBizId                  string                         `json:"ClusterBizId" xml:"ClusterBizId"`
	FolderPath                    string                         `json:"FolderPath" xml:"FolderPath"`
	NodeContent                   string                         `json:"NodeContent" xml:"NodeContent"`
	ChangeType                    string                         `json:"ChangeType" xml:"ChangeType"`
	TemplateName                  string                         `json:"TemplateName" xml:"TemplateName"`
	RuleName                      string                         `json:"RuleName" xml:"RuleName"`
	DIJobId                       int64                          `json:"DIJobId" xml:"DIJobId"`
	BeginWaitResTime              int64                          `json:"BeginWaitResTime" xml:"BeginWaitResTime"`
	CreateTime                    int64                          `json:"CreateTime" xml:"CreateTime"`
	RetResult                     string                         `json:"RetResult" xml:"RetResult"`
	FolderId                      string                         `json:"FolderId" xml:"FolderId"`
	BlockType                     int                            `json:"BlockType" xml:"BlockType"`
	OpSeq                         int64                          `json:"OpSeq" xml:"OpSeq"`
	Endpoint                      string                         `json:"Endpoint" xml:"Endpoint"`
	NextTaskId                    string                         `json:"NextTaskId" xml:"NextTaskId"`
	DqcType                       int                            `json:"DqcType" xml:"DqcType"`
	JobName                       string                         `json:"JobName" xml:"JobName"`
	Description                   string                         `json:"Description" xml:"Description"`
	MigrationType                 string                         `json:"MigrationType" xml:"MigrationType"`
	ExpectValue                   string                         `json:"ExpectValue" xml:"ExpectValue"`
	DndEnd                        string                         `json:"DndEnd" xml:"DndEnd"`
	RelatedFlowId                 int64                          `json:"RelatedFlowId" xml:"RelatedFlowId"`
	LastDdlTime                   int64                          `json:"LastDdlTime" xml:"LastDdlTime"`
	Useflag                       bool                           `json:"Useflag" xml:"Useflag"`
	AlertInterval                 int                            `json:"AlertInterval" xml:"AlertInterval"`
	ProcessId                     int64                          `json:"ProcessId" xml:"ProcessId"`
	HourSlaDetail                 string                         `json:"HourSlaDetail" xml:"HourSlaDetail"`
	CycTime                       int64                          `json:"CycTime" xml:"CycTime"`
	Property                      string                         `json:"Property" xml:"Property"`
	BaselineType                  string                         `json:"BaselineType" xml:"BaselineType"`
	CreatedUid                    string                         `json:"CreatedUid" xml:"CreatedUid"`
	IsDefault                     bool                           `json:"IsDefault" xml:"IsDefault"`
	DestinationDataSourceType     string                         `json:"DestinationDataSourceType" xml:"DestinationDataSourceType"`
	Creator                       string                         `json:"Creator" xml:"Creator"`
	RetCode                       int64                          `json:"RetCode" xml:"RetCode"`
	FileName                      string                         `json:"FileName" xml:"FileName"`
	Type                          string                         `json:"Type" xml:"Type"`
	Id                            int64                          `json:"Id" xml:"Id"`
	RuleType                      int                            `json:"RuleType" xml:"RuleType"`
	TableName                     string                         `json:"TableName" xml:"TableName"`
	CategoryId                    int64                          `json:"CategoryId" xml:"CategoryId"`
	Trend                         string                         `json:"Trend" xml:"Trend"`
	InstanceId                    int64                          `json:"InstanceId" xml:"InstanceId"`
	TaskType                      string                         `json:"TaskType" xml:"TaskType"`
	RemindUnit                    string                         `json:"RemindUnit" xml:"RemindUnit"`
	InGroupId                     int                            `json:"InGroupId" xml:"InGroupId"`
	ApplicationId                 int64                          `json:"ApplicationId" xml:"ApplicationId"`
	ReadCount                     int64                          `json:"ReadCount" xml:"ReadCount"`
	FilePropertyContent           string                         `json:"FilePropertyContent" xml:"FilePropertyContent"`
	TemplateId                    int                            `json:"TemplateId" xml:"TemplateId"`
	Repeatability                 bool                           `json:"Repeatability" xml:"Repeatability"`
	TenantId                      int64                          `json:"TenantId" xml:"TenantId"`
	IsPartitionTable              bool                           `json:"IsPartitionTable" xml:"IsPartitionTable"`
	DownloadUrl                   string                         `json:"DownloadUrl" xml:"DownloadUrl"`
	PageSize                      int                            `json:"PageSize" xml:"PageSize"`
	ParamMap                      string                         `json:"ParamMap" xml:"ParamMap"`
	DataSize                      int64                          `json:"DataSize" xml:"DataSize"`
	Comment                       string                         `json:"Comment" xml:"Comment"`
	CreatedTime                   int64                          `json:"CreatedTime" xml:"CreatedTime"`
	RepeatInterval                int64                          `json:"RepeatInterval" xml:"RepeatInterval"`
	Founder                       string                         `json:"Founder" xml:"Founder"`
	CreateUser                    string                         `json:"CreateUser" xml:"CreateUser"`
	JobStatus                     string                         `json:"JobStatus" xml:"JobStatus"`
	ApplicationCode               string                         `json:"ApplicationCode" xml:"ApplicationCode"`
	Priority                      int                            `json:"Priority" xml:"Priority"`
	Buffer                        float64                        `json:"Buffer" xml:"Buffer"`
	Owner                         string                         `json:"Owner" xml:"Owner"`
	PartitionKeys                 string                         `json:"PartitionKeys" xml:"PartitionKeys"`
	FileId                        int64                          `json:"FileId" xml:"FileId"`
	EntityId                      int64                          `json:"EntityId" xml:"EntityId"`
	DebugInfo                     string                         `json:"DebugInfo" xml:"DebugInfo"`
	BeginWaitTimeTime             int64                          `json:"BeginWaitTimeTime" xml:"BeginWaitTimeTime"`
	OwnerId                       string                         `json:"OwnerId" xml:"OwnerId"`
	UpdatedTime                   int64                          `json:"UpdatedTime" xml:"UpdatedTime"`
	SourceDataSourceType          string                         `json:"SourceDataSourceType" xml:"SourceDataSourceType"`
	CriticalThreshold             string                         `json:"CriticalThreshold" xml:"CriticalThreshold"`
	ApplicationKey                string                         `json:"ApplicationKey" xml:"ApplicationKey"`
	DndStart                      string                         `json:"DndStart" xml:"DndStart"`
	BeginRunningTime              int64                          `json:"BeginRunningTime" xml:"BeginRunningTime"`
	Operator                      string                         `json:"Operator" xml:"Operator"`
	StartedTime                   int64                          `json:"StartedTime" xml:"StartedTime"`
	HasNext                       bool                           `json:"HasNext" xml:"HasNext"`
	StartTime                     int64                          `json:"StartTime" xml:"StartTime"`
	ProjectNameCn                 string                         `json:"ProjectNameCn" xml:"ProjectNameCn"`
	TotalCount                    int                            `json:"TotalCount" xml:"TotalCount"`
	HourExpDetail                 string                         `json:"HourExpDetail" xml:"HourExpDetail"`
	DatabaseName                  string                         `json:"DatabaseName" xml:"DatabaseName"`
	ClusterId                     string                         `json:"ClusterId" xml:"ClusterId"`
	TopicId                       int64                          `json:"TopicId" xml:"TopicId"`
	ApplicationName               string                         `json:"ApplicationName" xml:"ApplicationName"`
	OpUser                        string                         `json:"OpUser" xml:"OpUser"`
	OwnerName                     string                         `json:"OwnerName" xml:"OwnerName"`
	BaselineId                    int64                          `json:"BaselineId" xml:"BaselineId"`
	Location                      string                         `json:"Location" xml:"Location"`
	FinishStatus                  string                         `json:"FinishStatus" xml:"FinishStatus"`
	ApiId                         int64                          `json:"ApiId" xml:"ApiId"`
	WarningThreshold              string                         `json:"WarningThreshold" xml:"WarningThreshold"`
	StartedUid                    string                         `json:"StartedUid" xml:"StartedUid"`
	RemindName                    string                         `json:"RemindName" xml:"RemindName"`
	GmtModified                   int64                          `json:"GmtModified" xml:"GmtModified"`
	Config                        string                         `json:"Config" xml:"Config"`
	NodesDebugInfo                string                         `json:"NodesDebugInfo" xml:"NodesDebugInfo"`
	ParamValues                   string                         `json:"ParamValues" xml:"ParamValues"`
	MethodName                    string                         `json:"MethodName" xml:"MethodName"`
	MaxAlertTimes                 int                            `json:"MaxAlertTimes" xml:"MaxAlertTimes"`
	ErrorMessage                  string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	IsVisible                     int                            `json:"IsVisible" xml:"IsVisible"`
	TaskRerunTime                 int                            `json:"TaskRerunTime" xml:"TaskRerunTime"`
	SlaTime                       int64                          `json:"SlaTime" xml:"SlaTime"`
	PageNumber                    int                            `json:"PageNumber" xml:"PageNumber"`
	RemindId                      int64                          `json:"RemindId" xml:"RemindId"`
	TableGuid                     string                         `json:"TableGuid" xml:"TableGuid"`
	ExpMinu                       int                            `json:"ExpMinu" xml:"ExpMinu"`
	OnDuty                        string                         `json:"OnDuty" xml:"OnDuty"`
	DagId                         int64                          `json:"DagId" xml:"DagId"`
	AlertUnit                     string                         `json:"AlertUnit" xml:"AlertUnit"`
	FavoriteCount                 int64                          `json:"FavoriteCount" xml:"FavoriteCount"`
	ModifiedTime                  int64                          `json:"ModifiedTime" xml:"ModifiedTime"`
	Meta                          string                         `json:"Meta" xml:"Meta"`
	CommitTime                    int64                          `json:"CommitTime" xml:"CommitTime"`
	ProjectId                     int64                          `json:"ProjectId" xml:"ProjectId"`
	SlaHour                       int                            `json:"SlaHour" xml:"SlaHour"`
	Webhooks                      []string                       `json:"Webhooks" xml:"Webhooks"`
	NodeIds                       []int64                        `json:"NodeIds" xml:"NodeIds"`
	AlertMethods                  []string                       `json:"AlertMethods" xml:"AlertMethods"`
	CollectionList                []string                       `json:"CollectionList" xml:"CollectionList"`
	AlertTargets                  []string                       `json:"AlertTargets" xml:"AlertTargets"`
	EntityList                    []string                       `json:"EntityList" xml:"EntityList"`
	Deployment                    Deployment                     `json:"Deployment" xml:"Deployment"`
	SolutionInfo                  SolutionInfo                   `json:"SolutionInfo" xml:"SolutionInfo"`
	File                          File                           `json:"File" xml:"File"`
	LastInstance                  LastInstance                   `json:"LastInstance" xml:"LastInstance"`
	ResourceSettings              ResourceSettings               `json:"ResourceSettings" xml:"ResourceSettings"`
	BlockInstance                 BlockInstance                  `json:"BlockInstance" xml:"BlockInstance"`
	SolutionDetail                SolutionDetail                 `json:"SolutionDetail" xml:"SolutionDetail"`
	NodeConfiguration             NodeConfiguration              `json:"NodeConfiguration" xml:"NodeConfiguration"`
	JobSettings                   JobSettings                    `json:"JobSettings" xml:"JobSettings"`
	ColumnList                    []ColumnListItem               `json:"ColumnList" xml:"ColumnList"`
	Apis                          []Api                          `json:"Apis" xml:"Apis"`
	SourceDataSourceSettings      []SourceDataSourceSetting      `json:"SourceDataSourceSettings" xml:"SourceDataSourceSettings"`
	ProjectMemberList             []ProjectMember                `json:"ProjectMemberList" xml:"ProjectMemberList"`
	DataSources                   []DataSourcesItem              `json:"DataSources" xml:"DataSources"`
	TransformationRules           []TransformationRule           `json:"TransformationRules" xml:"TransformationRules"`
	Nodes                         []NodesItem                    `json:"Nodes" xml:"Nodes"`
	FileVersions                  []FileVersion                  `json:"FileVersions" xml:"FileVersions"`
	DestinationDataSourceSettings []DestinationDataSourceSetting `json:"DestinationDataSourceSettings" xml:"DestinationDataSourceSettings"`
	Baselines                     []BaselinesItem                `json:"Baselines" xml:"Baselines"`
	Applications                  []Application                  `json:"Applications" xml:"Applications"`
	Connections                   []ConnectionsItem              `json:"Connections" xml:"Connections"`
	DISyncTasks                   []DISyncTasksItem              `json:"DISyncTasks" xml:"DISyncTasks"`
	Instances                     []Instance                     `json:"Instances" xml:"Instances"`
	DataEntityList                []DataEntityListItem           `json:"DataEntityList" xml:"DataEntityList"`
	BaselineStatuses              []BaselineStatusesItem         `json:"BaselineStatuses" xml:"BaselineStatuses"`
	Projects                      []ProjectsItem                 `json:"Projects" xml:"Projects"`
	Folders                       []FoldersItem                  `json:"Folders" xml:"Folders"`
	Files                         []File                         `json:"Files" xml:"Files"`
	ApiAuthorizedList             []ApiAuthorized                `json:"ApiAuthorizedList" xml:"ApiAuthorizedList"`
	TableMappings                 []TableMapping                 `json:"TableMappings" xml:"TableMappings"`
	Dags                          []Dag                          `json:"Dags" xml:"Dags"`
	ApiAuthorizationList          []ApiAuthorization             `json:"ApiAuthorizationList" xml:"ApiAuthorizationList"`
	OverTimeSettings              []OverTimeSetting              `json:"OverTimeSettings" xml:"OverTimeSettings"`
	Robots                        []RobotsItem                   `json:"Robots" xml:"Robots"`
	AlertSettings                 []AlertSetting                 `json:"AlertSettings" xml:"AlertSettings"`
	BizProcesses                  []BizProcessesItem             `json:"BizProcesses" xml:"BizProcesses"`
	Influences                    []InfluencesItem               `json:"Influences" xml:"Influences"`
	CalcEngines                   []CalcEnginesItem              `json:"CalcEngines" xml:"CalcEngines"`
	Reminds                       []RemindsItem                  `json:"Reminds" xml:"Reminds"`
	DeployedItems                 []DeployedItem                 `json:"DeployedItems" xml:"DeployedItems"`
	Migrations                    []Migration                    `json:"Migrations" xml:"Migrations"`
	Topics                        []TopicsItem                   `json:"Topics" xml:"Topics"`
	Business                      []BusinessItem                 `json:"Business" xml:"Business"`
}