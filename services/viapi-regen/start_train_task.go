package viapi_regen

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

// StartTrainTask invokes the viapi_regen.StartTrainTask API synchronously
func (client *Client) StartTrainTask(request *StartTrainTaskRequest) (response *StartTrainTaskResponse, err error) {
	response = CreateStartTrainTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StartTrainTaskWithChan invokes the viapi_regen.StartTrainTask API asynchronously
func (client *Client) StartTrainTaskWithChan(request *StartTrainTaskRequest) (<-chan *StartTrainTaskResponse, <-chan error) {
	responseChan := make(chan *StartTrainTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartTrainTask(request)
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

// StartTrainTaskWithCallback invokes the viapi_regen.StartTrainTask API asynchronously
func (client *Client) StartTrainTaskWithCallback(request *StartTrainTaskRequest, callback func(response *StartTrainTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartTrainTaskResponse
		var err error
		defer close(result)
		response, err = client.StartTrainTask(request)
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

// StartTrainTaskRequest is the request struct for api StartTrainTask
type StartTrainTaskRequest struct {
	*requests.RpcRequest
	RelyOnTaskId   requests.Integer `position:"Body" name:"RelyOnTaskId"`
	ForceStartFlag requests.Boolean `position:"Body" name:"ForceStartFlag"`
	Id             requests.Integer `position:"Body" name:"Id"`
}

// StartTrainTaskResponse is the response struct for api StartTrainTask
type StartTrainTaskResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateStartTrainTaskRequest creates a request to invoke StartTrainTask API
func CreateStartTrainTaskRequest() (request *StartTrainTaskRequest) {
	request = &StartTrainTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "StartTrainTask", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartTrainTaskResponse creates a response to parse from StartTrainTask response
func CreateStartTrainTaskResponse() (response *StartTrainTaskResponse) {
	response = &StartTrainTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}