package das

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

// DataInGetSqlOptimizeAdvice is a nested struct in das response
type DataInGetSqlOptimizeAdvice struct {
	Status      string `json:"Status" xml:"Status"`
	StatusCode  string `json:"StatusCode" xml:"StatusCode"`
	DownloadUrl string `json:"DownloadUrl" xml:"DownloadUrl"`
	CreateTime  string `json:"CreateTime" xml:"CreateTime"`
	ExpireTime  string `json:"ExpireTime" xml:"ExpireTime"`
	TaskId      string `json:"TaskId" xml:"TaskId"`
}