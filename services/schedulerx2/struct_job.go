package schedulerx2

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

// Job is a nested struct in schedulerx2 response
type Job struct {
	Status          int            `json:"Status" xml:"Status"`
	JarUrl          string         `json:"JarUrl" xml:"JarUrl"`
	MaxAttempt      int            `json:"MaxAttempt" xml:"MaxAttempt"`
	Parameters      string         `json:"Parameters" xml:"Parameters"`
	Description     string         `json:"Description" xml:"Description"`
	JobId           int64          `json:"JobId" xml:"JobId"`
	ExecuteMode     string         `json:"ExecuteMode" xml:"ExecuteMode"`
	MaxConcurrency  string         `json:"MaxConcurrency" xml:"MaxConcurrency"`
	Name            string         `json:"Name" xml:"Name"`
	ClassName       string         `json:"ClassName" xml:"ClassName"`
	Content         string         `json:"Content" xml:"Content"`
	JobType         string         `json:"JobType" xml:"JobType"`
	AttemptInterval int            `json:"AttemptInterval" xml:"AttemptInterval"`
	XAttrs          string         `json:"XAttrs" xml:"XAttrs"`
	MapTaskXAttrs   MapTaskXAttrs  `json:"MapTaskXAttrs" xml:"MapTaskXAttrs"`
	TimeConfig      TimeConfig     `json:"TimeConfig" xml:"TimeConfig"`
	JobMonitorInfo  JobMonitorInfo `json:"JobMonitorInfo" xml:"JobMonitorInfo"`
}
