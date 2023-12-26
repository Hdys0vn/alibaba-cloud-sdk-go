package pairecservice

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

// FeatureConsistencyCheckConfigsItem is a nested struct in pairecservice response
type FeatureConsistencyCheckConfigsItem struct {
	FeatureConsistencyCheckJobConfigId string `json:"FeatureConsistencyCheckJobConfigId" xml:"FeatureConsistencyCheckJobConfigId"`
	Name                               string `json:"Name" xml:"Name"`
	SceneId                            string `json:"SceneId" xml:"SceneId"`
	SceneName                          string `json:"SceneName" xml:"SceneName"`
	Status                             string `json:"Status" xml:"Status"`
	CompareFeature                     bool   `json:"CompareFeature" xml:"CompareFeature"`
	LatestJobId                        string `json:"LatestJobId" xml:"LatestJobId"`
	LatestJobGmtSamplingStartTime      string `json:"LatestJobGmtSamplingStartTime" xml:"LatestJobGmtSamplingStartTime"`
	LatestJobGmtSamplingEndTime        string `json:"LatestJobGmtSamplingEndTime" xml:"LatestJobGmtSamplingEndTime"`
	SampleRate                         string `json:"SampleRate" xml:"SampleRate"`
	FeatureLandingResourceId           string `json:"FeatureLandingResourceId" xml:"FeatureLandingResourceId"`
	FeatureLandingResourceUri          string `json:"FeatureLandingResourceUri" xml:"FeatureLandingResourceUri"`
	EasServiceName                     string `json:"EasServiceName" xml:"EasServiceName"`
	FgJsonFileName                     string `json:"FgJsonFileName" xml:"FgJsonFileName"`
	UserTable                          string `json:"UserTable" xml:"UserTable"`
	UserIdField                        string `json:"UserIdField" xml:"UserIdField"`
	UserTablePartitionField            string `json:"UserTablePartitionField" xml:"UserTablePartitionField"`
	UserTablePartitionFieldFormat      string `json:"UserTablePartitionFieldFormat" xml:"UserTablePartitionFieldFormat"`
	ItemTable                          string `json:"ItemTable" xml:"ItemTable"`
	ItemIdField                        string `json:"ItemIdField" xml:"ItemIdField"`
	ItemTablePartitionField            string `json:"ItemTablePartitionField" xml:"ItemTablePartitionField"`
	ItemTablePartitionFieldFormat      string `json:"ItemTablePartitionFieldFormat" xml:"ItemTablePartitionFieldFormat"`
	GenerateZip                        bool   `json:"GenerateZip" xml:"GenerateZip"`
	ServiceId                          string `json:"ServiceId" xml:"ServiceId"`
	ServiceName                        string `json:"ServiceName" xml:"ServiceName"`
	WorkflowName                       string `json:"WorkflowName" xml:"WorkflowName"`
	OssResourceId                      string `json:"OssResourceId" xml:"OssResourceId"`
	OssBucket                          string `json:"OssBucket" xml:"OssBucket"`
	EasyRecVersion                     string `json:"EasyRecVersion" xml:"EasyRecVersion"`
	EasyRecPackagePath                 string `json:"EasyRecPackagePath" xml:"EasyRecPackagePath"`
	FgJarVersion                       string `json:"FgJarVersion" xml:"FgJarVersion"`
	FeaturePriority                    string `json:"FeaturePriority" xml:"FeaturePriority"`
	FeatureDisplayExclude              string `json:"FeatureDisplayExclude" xml:"FeatureDisplayExclude"`
	GmtCreateTime                      string `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime                    string `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
}
