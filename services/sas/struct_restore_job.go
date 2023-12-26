package sas

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

// RestoreJob is a nested struct in sas response
type RestoreJob struct {
	Status          string `json:"Status" xml:"Status"`
	SnapshotHash    string `json:"SnapshotHash" xml:"SnapshotHash"`
	SourceClientId  string `json:"SourceClientId" xml:"SourceClientId"`
	ErrorFileUrl    string `json:"ErrorFileUrl" xml:"ErrorFileUrl"`
	Includes        string `json:"Includes" xml:"Includes"`
	RestoreName     string `json:"RestoreName" xml:"RestoreName"`
	InternetIp      string `json:"InternetIp" xml:"InternetIp"`
	VaultId         string `json:"VaultId" xml:"VaultId"`
	ActualBytes     int64  `json:"ActualBytes" xml:"ActualBytes"`
	Message         string `json:"Message" xml:"Message"`
	Percentage      int    `json:"Percentage" xml:"Percentage"`
	GmtModified     string `json:"GmtModified" xml:"GmtModified"`
	RestoreType     string `json:"RestoreType" xml:"RestoreType"`
	ExitCode        string `json:"ExitCode" xml:"ExitCode"`
	ClientId        string `json:"ClientId" xml:"ClientId"`
	ItemsDone       int64  `json:"ItemsDone" xml:"ItemsDone"`
	BytesTotal      int64  `json:"BytesTotal" xml:"BytesTotal"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	InstanceName    string `json:"InstanceName" xml:"InstanceName"`
	CompleteTime    int64  `json:"CompleteTime" xml:"CompleteTime"`
	ErrorType       string `json:"ErrorType" xml:"ErrorType"`
	SnapshotVersion string `json:"SnapshotVersion" xml:"SnapshotVersion"`
	Target          string `json:"Target" xml:"Target"`
	CreatedTime     int64  `json:"CreatedTime" xml:"CreatedTime"`
	InstanceId      string `json:"InstanceId" xml:"InstanceId"`
	Source          string `json:"Source" xml:"Source"`
	IntranetIp      string `json:"IntranetIp" xml:"IntranetIp"`
	ErrorFile       string `json:"ErrorFile" xml:"ErrorFile"`
	Uuid            string `json:"Uuid" xml:"Uuid"`
	Excludes        string `json:"Excludes" xml:"Excludes"`
	Speed           int64  `json:"Speed" xml:"Speed"`
	SnapshotId      string `json:"SnapshotId" xml:"SnapshotId"`
	UpdatedTime     int64  `json:"UpdatedTime" xml:"UpdatedTime"`
	RestoreId       string `json:"RestoreId" xml:"RestoreId"`
	GmtCreate       string `json:"GmtCreate" xml:"GmtCreate"`
	Eta             int64  `json:"Eta" xml:"Eta"`
	Duration        int64  `json:"Duration" xml:"Duration"`
	ErrorCount      int64  `json:"ErrorCount" xml:"ErrorCount"`
	ItemsTotal      int64  `json:"ItemsTotal" xml:"ItemsTotal"`
	BytesDone       int64  `json:"BytesDone" xml:"BytesDone"`
}