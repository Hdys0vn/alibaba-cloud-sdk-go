package alikafka

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

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateTopic invokes the alikafka.CreateTopic API synchronously
func (client *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	response = CreateCreateTopicResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTopicWithChan invokes the alikafka.CreateTopic API asynchronously
func (client *Client) CreateTopicWithChan(request *CreateTopicRequest) (<-chan *CreateTopicResponse, <-chan error) {
	responseChan := make(chan *CreateTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTopic(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateTopicWithCallback invokes the alikafka.CreateTopic API asynchronously
func (client *Client) CreateTopicWithCallback(request *CreateTopicRequest, callback func(response *CreateTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTopicResponse
		var err error
		defer close(result)
		response, err = client.CreateTopic(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateTopicRequest is the request struct for api CreateTopic
type CreateTopicRequest struct {
	*requests.RpcRequest
	Remark            string            `position:"Query" name:"Remark"`
	ReplicationFactor requests.Integer  `position:"Query" name:"ReplicationFactor"`
	MinInsyncReplicas requests.Integer  `position:"Query" name:"MinInsyncReplicas"`
	InstanceId        string            `position:"Query" name:"InstanceId"`
	Topic             string            `position:"Query" name:"Topic"`
	CompactTopic      requests.Boolean  `position:"Query" name:"CompactTopic"`
	Tag               *[]CreateTopicTag `position:"Query" name:"Tag"  type:"Repeated"`
	PartitionNum      string            `position:"Query" name:"PartitionNum"`
	Config            string            `position:"Query" name:"Config"`
	LocalTopic        requests.Boolean  `position:"Query" name:"LocalTopic"`
}

// CreateTopicTag is a repeated param struct in CreateTopicRequest
type CreateTopicTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateTopicResponse is the response struct for api CreateTopic
type CreateTopicResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateTopicRequest creates a request to invoke CreateTopic API
func CreateCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "CreateTopic", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTopicResponse creates a response to parse from CreateTopic response
func CreateCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
