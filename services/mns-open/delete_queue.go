package mns_open

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

// DeleteQueue invokes the mns_open.DeleteQueue API synchronously
func (client *Client) DeleteQueue(request *DeleteQueueRequest) (response *DeleteQueueResponse, err error) {
	response = CreateDeleteQueueResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteQueueWithChan invokes the mns_open.DeleteQueue API asynchronously
func (client *Client) DeleteQueueWithChan(request *DeleteQueueRequest) (<-chan *DeleteQueueResponse, <-chan error) {
	responseChan := make(chan *DeleteQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteQueue(request)
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

// DeleteQueueWithCallback invokes the mns_open.DeleteQueue API asynchronously
func (client *Client) DeleteQueueWithCallback(request *DeleteQueueRequest, callback func(response *DeleteQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteQueueResponse
		var err error
		defer close(result)
		response, err = client.DeleteQueue(request)
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

// DeleteQueueRequest is the request struct for api DeleteQueue
type DeleteQueueRequest struct {
	*requests.RpcRequest
	QueueName string `position:"Query" name:"QueueName"`
}

// DeleteQueueResponse is the response struct for api DeleteQueue
type DeleteQueueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int64  `json:"Code" xml:"Code"`
	Status    string `json:"Status" xml:"Status"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteQueueRequest creates a request to invoke DeleteQueue API
func CreateDeleteQueueRequest() (request *DeleteQueueRequest) {
	request = &DeleteQueueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mns-open", "2022-01-19", "DeleteQueue", "mns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteQueueResponse creates a response to parse from DeleteQueue response
func CreateDeleteQueueResponse() (response *DeleteQueueResponse) {
	response = &DeleteQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}