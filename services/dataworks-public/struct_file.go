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

// File is a nested struct in dataworks_public response
type File struct {
	AutoParsing        bool   `json:"AutoParsing" xml:"AutoParsing"`
	LastEditTime       int64  `json:"LastEditTime" xml:"LastEditTime"`
	Owner              string `json:"Owner" xml:"Owner"`
	ConnectionName     string `json:"ConnectionName" xml:"ConnectionName"`
	CreateTime         int64  `json:"CreateTime" xml:"CreateTime"`
	FileFolderId       string `json:"FileFolderId" xml:"FileFolderId"`
	DeletedStatus      string `json:"DeletedStatus" xml:"DeletedStatus"`
	FileDescription    string `json:"FileDescription" xml:"FileDescription"`
	FileId             int64  `json:"FileId" xml:"FileId"`
	CurrentVersion     int    `json:"CurrentVersion" xml:"CurrentVersion"`
	BizId              int64  `json:"BizId" xml:"BizId"`
	AdvancedSettings   string `json:"AdvancedSettings" xml:"AdvancedSettings"`
	FileType           int    `json:"FileType" xml:"FileType"`
	Content            string `json:"Content" xml:"Content"`
	NodeId             int64  `json:"NodeId" xml:"NodeId"`
	AbsoluteFolderPath string `json:"AbsoluteFolderPath" xml:"AbsoluteFolderPath"`
	CreateUser         string `json:"CreateUser" xml:"CreateUser"`
	BusinessId         int64  `json:"BusinessId" xml:"BusinessId"`
	LastEditUser       string `json:"LastEditUser" xml:"LastEditUser"`
	FileName           string `json:"FileName" xml:"FileName"`
	UseType            string `json:"UseType" xml:"UseType"`
	CommitStatus       int    `json:"CommitStatus" xml:"CommitStatus"`
	ParentId           int64  `json:"ParentId" xml:"ParentId"`
	IsMaxCompute       bool   `json:"IsMaxCompute" xml:"IsMaxCompute"`
}
