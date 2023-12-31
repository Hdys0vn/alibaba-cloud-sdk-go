package ons

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

// PublishInfoDo is a nested struct in ons response
type PublishInfoDo struct {
	MessageType       int                `json:"MessageType" xml:"MessageType"`
	RelationName      string             `json:"RelationName" xml:"RelationName"`
	Owner             string             `json:"Owner" xml:"Owner"`
	IndependentNaming bool               `json:"IndependentNaming" xml:"IndependentNaming"`
	Remark            string             `json:"Remark" xml:"Remark"`
	Relation          int                `json:"Relation" xml:"Relation"`
	CreateTime        int64              `json:"CreateTime" xml:"CreateTime"`
	Topic             string             `json:"Topic" xml:"Topic"`
	InstanceId        string             `json:"InstanceId" xml:"InstanceId"`
	ServiceStatus     int                `json:"ServiceStatus" xml:"ServiceStatus"`
	Tags              TagsInOnsTopicList `json:"Tags" xml:"Tags"`
}
