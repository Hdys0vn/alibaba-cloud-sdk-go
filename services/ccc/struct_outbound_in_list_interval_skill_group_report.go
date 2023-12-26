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

// OutboundInListIntervalSkillGroupReport is a nested struct in ccc response
type OutboundInListIntervalSkillGroupReport struct {
	AverageRingTime              float64 `json:"AverageRingTime" xml:"AverageRingTime"`
	CallsDialed                  int64   `json:"CallsDialed" xml:"CallsDialed"`
	CallsAnswered                int64   `json:"CallsAnswered" xml:"CallsAnswered"`
	TotalWorkTime                int64   `json:"TotalWorkTime" xml:"TotalWorkTime"`
	CallsAttendedTransferOut     int64   `json:"CallsAttendedTransferOut" xml:"CallsAttendedTransferOut"`
	MaxWorkTime                  int64   `json:"MaxWorkTime" xml:"MaxWorkTime"`
	TotalDialingTime             int64   `json:"TotalDialingTime" xml:"TotalDialingTime"`
	TotalHoldTime                int64   `json:"TotalHoldTime" xml:"TotalHoldTime"`
	AverageWorkTime              float64 `json:"AverageWorkTime" xml:"AverageWorkTime"`
	CallsBlindTransferIn         int64   `json:"CallsBlindTransferIn" xml:"CallsBlindTransferIn"`
	SatisfactionIndex            float64 `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	CallsRinged                  int64   `json:"CallsRinged" xml:"CallsRinged"`
	CallsAttendedTransferIn      int64   `json:"CallsAttendedTransferIn" xml:"CallsAttendedTransferIn"`
	CallsBlindTransferOut        int64   `json:"CallsBlindTransferOut" xml:"CallsBlindTransferOut"`
	TotalRingTime                int64   `json:"TotalRingTime" xml:"TotalRingTime"`
	MaxTalkTime                  int64   `json:"MaxTalkTime" xml:"MaxTalkTime"`
	MaxRingTime                  int64   `json:"MaxRingTime" xml:"MaxRingTime"`
	TotalTalkTime                int64   `json:"TotalTalkTime" xml:"TotalTalkTime"`
	MaxDialingTime               int64   `json:"MaxDialingTime" xml:"MaxDialingTime"`
	AnswerRate                   float64 `json:"AnswerRate" xml:"AnswerRate"`
	MaxHoldTime                  int64   `json:"MaxHoldTime" xml:"MaxHoldTime"`
	AverageTalkTime              float64 `json:"AverageTalkTime" xml:"AverageTalkTime"`
	SatisfactionRate             float64 `json:"SatisfactionRate" xml:"SatisfactionRate"`
	CallsHold                    int64   `json:"CallsHold" xml:"CallsHold"`
	SatisfactionSurveysOffered   int64   `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	SatisfactionSurveysResponded int64   `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	AverageHoldTime              float64 `json:"AverageHoldTime" xml:"AverageHoldTime"`
	AverageDialingTime           float64 `json:"AverageDialingTime" xml:"AverageDialingTime"`
}
