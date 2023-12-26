package qualitycheck

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

// RestartAsrTask invokes the qualitycheck.RestartAsrTask API synchronously
func (client *Client) RestartAsrTask(request *RestartAsrTaskRequest) (response *RestartAsrTaskResponse, err error) {
	response = CreateRestartAsrTaskResponse()
	err = client.DoAction(request, response)
	return
}

// RestartAsrTaskWithChan invokes the qualitycheck.RestartAsrTask API asynchronously
func (client *Client) RestartAsrTaskWithChan(request *RestartAsrTaskRequest) (<-chan *RestartAsrTaskResponse, <-chan error) {
	responseChan := make(chan *RestartAsrTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartAsrTask(request)
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

// RestartAsrTaskWithCallback invokes the qualitycheck.RestartAsrTask API asynchronously
func (client *Client) RestartAsrTaskWithCallback(request *RestartAsrTaskRequest, callback func(response *RestartAsrTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartAsrTaskResponse
		var err error
		defer close(result)
		response, err = client.RestartAsrTask(request)
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

// RestartAsrTaskRequest is the request struct for api RestartAsrTask
type RestartAsrTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// RestartAsrTaskResponse is the response struct for api RestartAsrTask
type RestartAsrTaskResponse struct {
	*responses.BaseResponse
	Code      string               `json:"Code" xml:"Code"`
	Message   string               `json:"Message" xml:"Message"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Success   bool                 `json:"Success" xml:"Success"`
	Data      DataInRestartAsrTask `json:"Data" xml:"Data"`
}

// CreateRestartAsrTaskRequest creates a request to invoke RestartAsrTask API
func CreateRestartAsrTaskRequest() (request *RestartAsrTaskRequest) {
	request = &RestartAsrTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "RestartAsrTask", "", "")
	request.Method = requests.POST
	return
}

// CreateRestartAsrTaskResponse creates a response to parse from RestartAsrTask response
func CreateRestartAsrTaskResponse() (response *RestartAsrTaskResponse) {
	response = &RestartAsrTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
