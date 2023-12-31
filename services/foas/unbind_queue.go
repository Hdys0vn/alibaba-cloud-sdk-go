package foas

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

// UnbindQueue invokes the foas.UnbindQueue API synchronously
func (client *Client) UnbindQueue(request *UnbindQueueRequest) (response *UnbindQueueResponse, err error) {
	response = CreateUnbindQueueResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindQueueWithChan invokes the foas.UnbindQueue API asynchronously
func (client *Client) UnbindQueueWithChan(request *UnbindQueueRequest) (<-chan *UnbindQueueResponse, <-chan error) {
	responseChan := make(chan *UnbindQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindQueue(request)
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

// UnbindQueueWithCallback invokes the foas.UnbindQueue API asynchronously
func (client *Client) UnbindQueueWithCallback(request *UnbindQueueRequest, callback func(response *UnbindQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindQueueResponse
		var err error
		defer close(result)
		response, err = client.UnbindQueue(request)
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

// UnbindQueueRequest is the request struct for api UnbindQueue
type UnbindQueueRequest struct {
	*requests.RoaRequest
	QueueName   string `position:"Query" name:"queueName"`
	ProjectName string `position:"Path" name:"projectName"`
	ClusterId   string `position:"Query" name:"clusterId"`
}

// UnbindQueueResponse is the response struct for api UnbindQueue
type UnbindQueueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindQueueRequest creates a request to invoke UnbindQueue API
func CreateUnbindQueueRequest() (request *UnbindQueueRequest) {
	request = &UnbindQueueRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "UnbindQueue", "/api/v2/projects/[projectName]/queue", "foas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateUnbindQueueResponse creates a response to parse from UnbindQueue response
func CreateUnbindQueueResponse() (response *UnbindQueueResponse) {
	response = &UnbindQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
