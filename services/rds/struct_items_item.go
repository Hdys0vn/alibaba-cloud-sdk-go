package rds

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

// ItemsItem is a nested struct in rds response
type ItemsItem struct {
	Region                string                `json:"Region" xml:"Region"`
	ResultInfo            string                `json:"ResultInfo" xml:"ResultInfo"`
	InsName               string                `json:"InsName" xml:"InsName"`
	TaskTypeEn            string                `json:"TaskTypeEn" xml:"TaskTypeEn"`
	Impact                string                `json:"Impact" xml:"Impact"`
	TaskParams            string                `json:"TaskParams" xml:"TaskParams"`
	InsComment            string                `json:"InsComment" xml:"InsComment"`
	InstanceType          string                `json:"InstanceType" xml:"InstanceType"`
	ReasonCode            string                `json:"ReasonCode" xml:"ReasonCode"`
	Deadline              string                `json:"Deadline" xml:"Deadline"`
	ImpactZh              string                `json:"ImpactZh" xml:"ImpactZh"`
	DbType                string                `json:"DbType" xml:"DbType"`
	ModifiedTime          string                `json:"ModifiedTime" xml:"ModifiedTime"`
	DbVersion             string                `json:"DbVersion" xml:"DbVersion"`
	AllowCancel           string                `json:"AllowCancel" xml:"AllowCancel"`
	CreatedTime           string                `json:"CreatedTime" xml:"CreatedTime"`
	TaskId                string                `json:"TaskId" xml:"TaskId"`
	RegionId              string                `json:"RegionId" xml:"RegionId"`
	InstanceId            string                `json:"InstanceId" xml:"InstanceId"`
	SwitchTime            string                `json:"SwitchTime" xml:"SwitchTime"`
	Progress              float64               `json:"Progress" xml:"Progress"`
	DBInstanceId          string                `json:"DBInstanceId" xml:"DBInstanceId"`
	AllowChange           string                `json:"AllowChange" xml:"AllowChange"`
	TaskDetail            string                `json:"TaskDetail" xml:"TaskDetail"`
	ChangeLevel           string                `json:"ChangeLevel" xml:"ChangeLevel"`
	Uid                   string                `json:"Uid" xml:"Uid"`
	Id                    int                   `json:"Id" xml:"Id"`
	CallerSource          string                `json:"CallerSource" xml:"CallerSource"`
	Status                int                   `json:"Status" xml:"Status"`
	Product               string                `json:"Product" xml:"Product"`
	CallerUid             string                `json:"CallerUid" xml:"CallerUid"`
	InstanceName          string                `json:"InstanceName" xml:"InstanceName"`
	ActionInfo            string                `json:"ActionInfo" xml:"ActionInfo"`
	ImpactEn              string                `json:"ImpactEn" xml:"ImpactEn"`
	ChangeLevelEn         string                `json:"ChangeLevelEn" xml:"ChangeLevelEn"`
	PrepareInterval       string                `json:"PrepareInterval" xml:"PrepareInterval"`
	CurrentAVZ            string                `json:"CurrentAVZ" xml:"CurrentAVZ"`
	StartTime             string                `json:"StartTime" xml:"StartTime"`
	EndTime               string                `json:"EndTime" xml:"EndTime"`
	RemainTime            int                   `json:"RemainTime" xml:"RemainTime"`
	CurrentStepName       string                `json:"CurrentStepName" xml:"CurrentStepName"`
	ChangeLevelZh         string                `json:"ChangeLevelZh" xml:"ChangeLevelZh"`
	TaskType              string                `json:"TaskType" xml:"TaskType"`
	TaskTypeZh            string                `json:"TaskTypeZh" xml:"TaskTypeZh"`
	ReadDBInstanceNames   []string              `json:"ReadDBInstanceNames" xml:"ReadDBInstanceNames"`
	ReadDelayTimes        []string              `json:"ReadDelayTimes" xml:"ReadDelayTimes"`
	SubInsNames           []string              `json:"SubInsNames" xml:"SubInsNames"`
	ReadonlyInstanceDelay ReadonlyInstanceDelay `json:"ReadonlyInstanceDelay" xml:"ReadonlyInstanceDelay"`
}
