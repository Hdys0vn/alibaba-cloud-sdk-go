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

// BaseInspection is a nested struct in das response
type BaseInspection struct {
	EndTime      int64                  `json:"EndTime" xml:"EndTime"`
	StartTime    int64                  `json:"StartTime" xml:"StartTime"`
	Data         map[string]interface{} `json:"Data" xml:"Data"`
	ScoreMap     map[string]interface{} `json:"ScoreMap" xml:"ScoreMap"`
	GmtCreate    int64                  `json:"GmtCreate" xml:"GmtCreate"`
	Score        int                    `json:"Score" xml:"Score"`
	EnableDasPro int                    `json:"EnableDasPro" xml:"EnableDasPro"`
	State        int                    `json:"State" xml:"State"`
	TaskType     int                    `json:"TaskType" xml:"TaskType"`
	Instance     Instance               `json:"Instance" xml:"Instance"`
	AutoFunction AutoFunction           `json:"AutoFunction" xml:"AutoFunction"`
}
