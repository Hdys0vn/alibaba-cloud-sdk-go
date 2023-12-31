package ecd

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

// SuspEvent is a nested struct in ecd response
type SuspEvent struct {
	DataSource            string   `json:"DataSource" xml:"DataSource"`
	EventSubType          string   `json:"EventSubType" xml:"EventSubType"`
	DesktopName           string   `json:"DesktopName" xml:"DesktopName"`
	DesktopId             string   `json:"DesktopId" xml:"DesktopId"`
	OccurrenceTime        string   `json:"OccurrenceTime" xml:"OccurrenceTime"`
	AlarmEventTypeDisplay string   `json:"AlarmEventTypeDisplay" xml:"AlarmEventTypeDisplay"`
	AlarmUniqueInfo       string   `json:"AlarmUniqueInfo" xml:"AlarmUniqueInfo"`
	Desc                  string   `json:"Desc" xml:"Desc"`
	AlarmEventNameDisplay string   `json:"AlarmEventNameDisplay" xml:"AlarmEventNameDisplay"`
	CanCancelFault        bool     `json:"CanCancelFault" xml:"CanCancelFault"`
	LastTime              string   `json:"LastTime" xml:"LastTime"`
	OperateMsg            string   `json:"OperateMsg" xml:"OperateMsg"`
	CanBeDealOnLine       string   `json:"CanBeDealOnLine" xml:"CanBeDealOnLine"`
	AlarmEventType        string   `json:"AlarmEventType" xml:"AlarmEventType"`
	EventStatus           int      `json:"EventStatus" xml:"EventStatus"`
	OperateErrorCode      string   `json:"OperateErrorCode" xml:"OperateErrorCode"`
	AlarmEventName        string   `json:"AlarmEventName" xml:"AlarmEventName"`
	Name                  string   `json:"Name" xml:"Name"`
	UniqueInfo            string   `json:"UniqueInfo" xml:"UniqueInfo"`
	Level                 string   `json:"Level" xml:"Level"`
	Id                    int64    `json:"Id" xml:"Id"`
	Details               []Detail `json:"Details" xml:"Details"`
}
