package arms

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

// IMRobots is a nested struct in arms response
type IMRobots struct {
	RobotId       float64        `json:"RobotId" xml:"RobotId"`
	Type          string         `json:"Type" xml:"Type"`
	RobotName     string         `json:"RobotName" xml:"RobotName"`
	RobotAddr     string         `json:"RobotAddr" xml:"RobotAddr"`
	DailyNoc      bool           `json:"DailyNoc" xml:"DailyNoc"`
	DailyNocTime  string         `json:"DailyNocTime" xml:"DailyNocTime"`
	DingSignKey   string         `json:"DingSignKey" xml:"DingSignKey"`
	CreateTime    string         `json:"CreateTime" xml:"CreateTime"`
	DispatchRules []DispatchRule `json:"DispatchRules" xml:"DispatchRules"`
}
