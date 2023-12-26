package scsp

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

// ExecuteActivity invokes the scsp.ExecuteActivity API synchronously
func (client *Client) ExecuteActivity(request *ExecuteActivityRequest) (response *ExecuteActivityResponse, err error) {
	response = CreateExecuteActivityResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteActivityWithChan invokes the scsp.ExecuteActivity API asynchronously
func (client *Client) ExecuteActivityWithChan(request *ExecuteActivityRequest) (<-chan *ExecuteActivityResponse, <-chan error) {
	responseChan := make(chan *ExecuteActivityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteActivity(request)
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

// ExecuteActivityWithCallback invokes the scsp.ExecuteActivity API asynchronously
func (client *Client) ExecuteActivityWithCallback(request *ExecuteActivityRequest, callback func(response *ExecuteActivityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteActivityResponse
		var err error
		defer close(result)
		response, err = client.ExecuteActivity(request)
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

// ExecuteActivityRequest is the request struct for api ExecuteActivity
type ExecuteActivityRequest struct {
	*requests.RpcRequest
	ClientToken  string           `position:"Body"`
	InstanceId   string           `position:"Body"`
	TicketId     requests.Integer `position:"Body"`
	OperatorId   requests.Integer `position:"Body"`
	ActivityCode string           `position:"Body"`
	ActivityForm string           `position:"Body"`
}

// ExecuteActivityResponse is the response struct for api ExecuteActivity
type ExecuteActivityResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateExecuteActivityRequest creates a request to invoke ExecuteActivity API
func CreateExecuteActivityRequest() (request *ExecuteActivityRequest) {
	request = &ExecuteActivityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "ExecuteActivity", "", "")
	request.Method = requests.POST
	return
}

// CreateExecuteActivityResponse creates a response to parse from ExecuteActivity response
func CreateExecuteActivityResponse() (response *ExecuteActivityResponse) {
	response = &ExecuteActivityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}