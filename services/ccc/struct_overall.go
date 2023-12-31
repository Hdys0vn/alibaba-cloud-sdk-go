package ccc

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

// Overall is a nested struct in ccc response
type Overall struct {
	TotalLoggedInTime              int64             `json:"TotalLoggedInTime" xml:"TotalLoggedInTime"`
	AverageHoldTime                float64           `json:"AverageHoldTime" xml:"AverageHoldTime"`
	TotalOutboundScenarioTime      int64             `json:"TotalOutboundScenarioTime" xml:"TotalOutboundScenarioTime"`
	MaxReadyTime                   int64             `json:"MaxReadyTime" xml:"MaxReadyTime"`
	AverageTalkTime                float64           `json:"AverageTalkTime" xml:"AverageTalkTime"`
	MaxWorkTime                    int64             `json:"MaxWorkTime" xml:"MaxWorkTime"`
	TotalReadyTime                 int64             `json:"TotalReadyTime" xml:"TotalReadyTime"`
	TotalWorkTime                  int64             `json:"TotalWorkTime" xml:"TotalWorkTime"`
	FirstCheckInTime               int64             `json:"FirstCheckInTime" xml:"FirstCheckInTime"`
	SatisfactionSurveysOffered     int64             `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	TotalBreakTime                 int64             `json:"TotalBreakTime" xml:"TotalBreakTime"`
	MaxBreakTime                   int64             `json:"MaxBreakTime" xml:"MaxBreakTime"`
	LastCheckOutTime               int64             `json:"LastCheckOutTime" xml:"LastCheckOutTime"`
	TotalHoldTime                  int64             `json:"TotalHoldTime" xml:"TotalHoldTime"`
	SatisfactionRate               float64           `json:"SatisfactionRate" xml:"SatisfactionRate"`
	MaxHoldTime                    int64             `json:"MaxHoldTime" xml:"MaxHoldTime"`
	SatisfactionSurveysResponded   int64             `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	AverageBreakTime               float64           `json:"AverageBreakTime" xml:"AverageBreakTime"`
	TotalOffSiteOnlineTime         int64             `json:"TotalOffSiteOnlineTime" xml:"TotalOffSiteOnlineTime"`
	SatisfactionIndex              float64           `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	TotalOfficePhoneOnlineTime     int64             `json:"TotalOfficePhoneOnlineTime" xml:"TotalOfficePhoneOnlineTime"`
	LastCheckoutTime               int64             `json:"LastCheckoutTime" xml:"LastCheckoutTime"`
	TotalOutboundScenarioReadyTime int64             `json:"TotalOutboundScenarioReadyTime" xml:"TotalOutboundScenarioReadyTime"`
	MaxTalkTime                    int64             `json:"MaxTalkTime" xml:"MaxTalkTime"`
	TotalTalkTime                  int64             `json:"TotalTalkTime" xml:"TotalTalkTime"`
	AverageReadyTime               float64           `json:"AverageReadyTime" xml:"AverageReadyTime"`
	TotalCalls                     int64             `json:"TotalCalls" xml:"TotalCalls"`
	TotalOnSiteOnlineTime          int64             `json:"TotalOnSiteOnlineTime" xml:"TotalOnSiteOnlineTime"`
	OccupancyRate                  float64           `json:"OccupancyRate" xml:"OccupancyRate"`
	AverageWorkTime                float64           `json:"AverageWorkTime" xml:"AverageWorkTime"`
	BreakCodeDetailList            []BreakCodeDetail `json:"BreakCodeDetailList" xml:"BreakCodeDetailList"`
}
