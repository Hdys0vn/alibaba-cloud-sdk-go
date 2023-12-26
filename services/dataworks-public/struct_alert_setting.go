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

// AlertSetting is a nested struct in dataworks_public response
type AlertSetting struct {
	AlertType            string      `json:"AlertType" xml:"AlertType"`
	SilenceStartTime     string      `json:"SilenceStartTime" xml:"SilenceStartTime"`
	SilenceEndTime       string      `json:"SilenceEndTime" xml:"SilenceEndTime"`
	AlertInterval        int         `json:"AlertInterval" xml:"AlertInterval"`
	AlertMaximum         int         `json:"AlertMaximum" xml:"AlertMaximum"`
	BaselineAlertEnabled bool        `json:"BaselineAlertEnabled" xml:"BaselineAlertEnabled"`
	AlertRecipientType   string      `json:"AlertRecipientType" xml:"AlertRecipientType"`
	AlertRecipient       string      `json:"AlertRecipient" xml:"AlertRecipient"`
	AlertMethods         []string    `json:"AlertMethods" xml:"AlertMethods"`
	Webhooks             []string    `json:"Webhooks" xml:"Webhooks"`
	TopicTypes           []string    `json:"TopicTypes" xml:"TopicTypes"`
	DingRobots           []DingRobot `json:"DingRobots" xml:"DingRobots"`
}