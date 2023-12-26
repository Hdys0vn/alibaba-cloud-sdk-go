package quickbi_public

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

// Data is a nested struct in quickbi_public response
type Data struct {
	WorksName                    string                       `json:"WorksName" xml:"WorksName"`
	IdentifiedPath               string                       `json:"IdentifiedPath" xml:"IdentifiedPath"`
	UsergroupDesc                string                       `json:"UsergroupDesc" xml:"UsergroupDesc"`
	ModifiedTime                 string                       `json:"ModifiedTime" xml:"ModifiedTime"`
	CubeName                     string                       `json:"CubeName" xml:"CubeName"`
	TagName                      string                       `json:"TagName" xml:"TagName"`
	QueryOverTenSecPercent       string                       `json:"QueryOverTenSecPercent" xml:"QueryOverTenSecPercent"`
	AuthPoint                    int                          `json:"AuthPoint" xml:"AuthPoint"`
	CostTimeAvg                  string                       `json:"CostTimeAvg" xml:"CostTimeAvg"`
	FavoriteId                   int                          `json:"FavoriteId" xml:"FavoriteId"`
	WorkType                     string                       `json:"WorkType" xml:"WorkType"`
	Id                           string                       `json:"Id" xml:"Id"`
	QueryTimeoutCountPercent     string                       `json:"QueryTimeoutCountPercent" xml:"QueryTimeoutCountPercent"`
	ShareToId                    string                       `json:"ShareToId" xml:"ShareToId"`
	WorksType                    string                       `json:"WorksType" xml:"WorksType"`
	RepeatSqlQueryCount          int                          `json:"RepeatSqlQueryCount" xml:"RepeatSqlQueryCount"`
	ModifyUser                   string                       `json:"ModifyUser" xml:"ModifyUser"`
	OwnerName                    string                       `json:"OwnerName" xml:"OwnerName"`
	UserGroupId                  string                       `json:"UserGroupId" xml:"UserGroupId"`
	WorkspaceName                string                       `json:"WorkspaceName" xml:"WorkspaceName"`
	TagId                        string                       `json:"TagId" xml:"TagId"`
	WorksId                      string                       `json:"WorksId" xml:"WorksId"`
	AdviceType                   string                       `json:"AdviceType" xml:"AdviceType"`
	UsergroupName                string                       `json:"UsergroupName" xml:"UsergroupName"`
	QuickIndexCostTimeAvg        string                       `json:"QuickIndexCostTimeAvg" xml:"QuickIndexCostTimeAvg"`
	MenuId                       string                       `json:"MenuId" xml:"MenuId"`
	ThirdPartAuthFlag            int                          `json:"ThirdPartAuthFlag" xml:"ThirdPartAuthFlag"`
	SecurityLevel                string                       `json:"SecurityLevel" xml:"SecurityLevel"`
	QueryOverTenSecPercentNum    string                       `json:"QueryOverTenSecPercentNum" xml:"QueryOverTenSecPercentNum"`
	WorkspaceId                  string                       `json:"WorkspaceId" xml:"WorkspaceId"`
	ReportType                   string                       `json:"ReportType" xml:"ReportType"`
	RepeatQueryPercent           string                       `json:"RepeatQueryPercent" xml:"RepeatQueryPercent"`
	RepeatSqlQueryPercent        string                       `json:"RepeatSqlQueryPercent" xml:"RepeatSqlQueryPercent"`
	Name                         string                       `json:"Name" xml:"Name"`
	ModifyTime                   string                       `json:"ModifyTime" xml:"ModifyTime"`
	ShareId                      string                       `json:"ShareId" xml:"ShareId"`
	CubeId                       string                       `json:"CubeId" xml:"CubeId"`
	ReportId                     string                       `json:"ReportId" xml:"ReportId"`
	CreateTime                   string                       `json:"CreateTime" xml:"CreateTime"`
	ShowOnlyWithAccess           bool                         `json:"ShowOnlyWithAccess" xml:"ShowOnlyWithAccess"`
	ParentUserGroupName          string                       `json:"ParentUserGroupName" xml:"ParentUserGroupName"`
	UsergroupId                  string                       `json:"UsergroupId" xml:"UsergroupId"`
	WorkName                     string                       `json:"WorkName" xml:"WorkName"`
	UserGroupName                string                       `json:"UserGroupName" xml:"UserGroupName"`
	CreateUser                   string                       `json:"CreateUser" xml:"CreateUser"`
	QueryCountAvg                string                       `json:"QueryCountAvg" xml:"QueryCountAvg"`
	TagValue                     string                       `json:"TagValue" xml:"TagValue"`
	QueryCount                   int                          `json:"QueryCount" xml:"QueryCount"`
	Description                  string                       `json:"Description" xml:"Description"`
	CacheQueryCount              int                          `json:"CacheQueryCount" xml:"CacheQueryCount"`
	IsUserGroup                  bool                         `json:"IsUserGroup" xml:"IsUserGroup"`
	ExpireDate                   int64                        `json:"ExpireDate" xml:"ExpireDate"`
	CacheCostTimeAvg             string                       `json:"CacheCostTimeAvg" xml:"CacheCostTimeAvg"`
	Status                       int                          `json:"Status" xml:"Status"`
	ParentUserGroupId            string                       `json:"ParentUserGroupId" xml:"ParentUserGroupId"`
	QueryTimeoutCount            int                          `json:"QueryTimeoutCount" xml:"QueryTimeoutCount"`
	RepeatQueryPercentNum        string                       `json:"RepeatQueryPercentNum" xml:"RepeatQueryPercentNum"`
	ModifyName                   string                       `json:"ModifyName" xml:"ModifyName"`
	ComponentName                string                       `json:"ComponentName" xml:"ComponentName"`
	QueryOverFivePercentNum      string                       `json:"QueryOverFivePercentNum" xml:"QueryOverFivePercentNum"`
	ComponentQueryCountAvg       string                       `json:"ComponentQueryCountAvg" xml:"ComponentQueryCountAvg"`
	QueryOverFiveSecPercent      string                       `json:"QueryOverFiveSecPercent" xml:"QueryOverFiveSecPercent"`
	ShareToName                  string                       `json:"ShareToName" xml:"ShareToName"`
	ShareToType                  int                          `json:"ShareToType" xml:"ShareToType"`
	QuickIndexQueryCount         int                          `json:"QuickIndexQueryCount" xml:"QuickIndexQueryCount"`
	ParentUsergroupId            string                       `json:"ParentUsergroupId" xml:"ParentUsergroupId"`
	ComponentId                  string                       `json:"ComponentId" xml:"ComponentId"`
	ComponentQueryCount          int                          `json:"ComponentQueryCount" xml:"ComponentQueryCount"`
	OwnerId                      string                       `json:"OwnerId" xml:"OwnerId"`
	UserGroupDescription         string                       `json:"UserGroupDescription" xml:"UserGroupDescription"`
	ReportName                   string                       `json:"ReportName" xml:"ReportName"`
	ShareType                    string                       `json:"ShareType" xml:"ShareType"`
	CubePerformanceDiagnoseModel CubePerformanceDiagnoseModel `json:"CubePerformanceDiagnoseModel" xml:"CubePerformanceDiagnoseModel"`
	Directory                    Directory                    `json:"Directory" xml:"Directory"`
	Receivers                    []ReceiversItem              `json:"Receivers" xml:"Receivers"`
}
