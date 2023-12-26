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

// DeleteTopic invokes the alikafka.DeleteTopic API synchronously
func (client *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	response = CreateDeleteTopicResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTopicWithChan invokes the alikafka.DeleteTopic API asynchronously
func (client *Client) DeleteTopicWithChan(request *DeleteTopicRequest) (<-chan *DeleteTopicResponse, <-chan error) {
	responseChan := make(chan *DeleteTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTopic(request)
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

// DeleteTopicWithCallback invokes the alikafka.DeleteTopic API asynchronously
func (client *Client) DeleteTopicWithCallback(request *DeleteTopicRequest, callback func(response *DeleteTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTopicResponse
		var err error
		defer close(result)
		response, err = client.DeleteTopic(request)
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

// DeleteTopicRequest is the request struct for api DeleteTopic
type DeleteTopicRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// DeleteTopicResponse is the response struct for api DeleteTopic
type DeleteTopicResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteTopicRequest creates a request to invoke DeleteTopic API
func CreateDeleteTopicRequest() (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "DeleteTopic", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTopicResponse creates a response to parse from DeleteTopic response
func CreateDeleteTopicResponse() (response *DeleteTopicResponse) {
	response = &DeleteTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
