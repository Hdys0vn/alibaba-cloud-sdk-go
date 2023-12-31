package amqp_open

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

// Data is a nested struct in amqp_open response
type Data struct {
	MasterUId           int64             `json:"MasterUId" xml:"MasterUId"`
	CurrentVirtualHosts int               `json:"CurrentVirtualHosts" xml:"CurrentVirtualHosts"`
	MaxQueues           int               `json:"MaxQueues" xml:"MaxQueues"`
	MaxVirtualHosts     int               `json:"MaxVirtualHosts" xml:"MaxVirtualHosts"`
	MaxExchanges        int               `json:"MaxExchanges" xml:"MaxExchanges"`
	Password            string            `json:"Password" xml:"Password"`
	CreateTimeStamp     int64             `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
	NextToken           string            `json:"NextToken" xml:"NextToken"`
	InstanceId          string            `json:"InstanceId" xml:"InstanceId"`
	CurrentExchanges    int               `json:"CurrentExchanges" xml:"CurrentExchanges"`
	UserName            string            `json:"UserName" xml:"UserName"`
	CurrentQueues       int               `json:"CurrentQueues" xml:"CurrentQueues"`
	AccessKey           string            `json:"AccessKey" xml:"AccessKey"`
	MaxResults          int               `json:"MaxResults" xml:"MaxResults"`
	Exchanges           []ExchangeVO      `json:"Exchanges" xml:"Exchanges"`
	Queues              []QueueVO         `json:"Queues" xml:"Queues"`
	Bindings            []BindingDO       `json:"Bindings" xml:"Bindings"`
	Consumers           []QueueConsumerVO `json:"Consumers" xml:"Consumers"`
	VirtualHosts        []VhostVO         `json:"VirtualHosts" xml:"VirtualHosts"`
	Instances           []InstanceVO      `json:"Instances" xml:"Instances"`
}
