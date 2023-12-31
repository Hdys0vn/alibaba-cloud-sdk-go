package cbn

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

// TrafficQosQueue is a nested struct in cbn response
type TrafficQosQueue struct {
	TransitRouterAttachmentId  string `json:"TransitRouterAttachmentId" xml:"TransitRouterAttachmentId"`
	TrafficQosPolicyId         string `json:"TrafficQosPolicyId" xml:"TrafficQosPolicyId"`
	TrafficQosQueueDescription string `json:"TrafficQosQueueDescription" xml:"TrafficQosQueueDescription"`
	QosQueueName               string `json:"QosQueueName" xml:"QosQueueName"`
	TransitRouterId            string `json:"TransitRouterId" xml:"TransitRouterId"`
	TrafficQosQueueName        string `json:"TrafficQosQueueName" xml:"TrafficQosQueueName"`
	TrafficQosQueueId          string `json:"TrafficQosQueueId" xml:"TrafficQosQueueId"`
	Status                     string `json:"Status" xml:"Status"`
	QosQueueId                 string `json:"QosQueueId" xml:"QosQueueId"`
	RemainBandwidthPercent     int    `json:"RemainBandwidthPercent" xml:"RemainBandwidthPercent"`
	QosQueueDescription        string `json:"QosQueueDescription" xml:"QosQueueDescription"`
	Dscps                      []int  `json:"Dscps" xml:"Dscps"`
}
