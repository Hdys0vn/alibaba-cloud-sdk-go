package mpaas

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

// ApkInfo is a nested struct in mpaas response
type ApkInfo struct {
	VersionName         string               `json:"VersionName" xml:"VersionName"`
	BeforeMd5           string               `json:"BeforeMd5" xml:"BeforeMd5"`
	AfterMd5            string               `json:"AfterMd5" xml:"AfterMd5"`
	Label               string               `json:"Label" xml:"Label"`
	Status              int64                `json:"Status" xml:"Status"`
	Id                  int64                `json:"Id" xml:"Id"`
	VersionCode         string               `json:"VersionCode" xml:"VersionCode"`
	AfterSize           int64                `json:"AfterSize" xml:"AfterSize"`
	BeforeSize          int64                `json:"BeforeSize" xml:"BeforeSize"`
	Progress            int64                `json:"Progress" xml:"Progress"`
	ClassForest         string               `json:"ClassForest" xml:"ClassForest"`
	AppCode             string               `json:"AppCode" xml:"AppCode"`
	TaskType            string               `json:"TaskType" xml:"TaskType"`
	AppPackage          string               `json:"AppPackage" xml:"AppPackage"`
	EnhancedClasses     []string             `json:"EnhancedClasses" xml:"EnhancedClasses"`
	SoFileList          []string             `json:"SoFileList" xml:"SoFileList"`
	EnhancedSoFiles     []string             `json:"EnhancedSoFiles" xml:"EnhancedSoFiles"`
	AssetsFileList      []string             `json:"AssetsFileList" xml:"AssetsFileList"`
	EnhanceRules        []string             `json:"EnhanceRules" xml:"EnhanceRules"`
	EnhancedAssetsFiles []string             `json:"EnhancedAssetsFiles" xml:"EnhancedAssetsFiles"`
	EnhanceMapping      []EnhanceMappingItem `json:"EnhanceMapping" xml:"EnhanceMapping"`
}