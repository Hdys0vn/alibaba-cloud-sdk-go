package swas_open

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

// StopInstances invokes the swas_open.StopInstances API synchronously
func (client *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
	response = CreateStopInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// StopInstancesWithChan invokes the swas_open.StopInstances API asynchronously
func (client *Client) StopInstancesWithChan(request *StopInstancesRequest) (<-chan *StopInstancesResponse, <-chan error) {
	responseChan := make(chan *StopInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopInstances(request)
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

// StopInstancesWithCallback invokes the swas_open.StopInstances API asynchronously
func (client *Client) StopInstancesWithCallback(request *StopInstancesRequest, callback func(response *StopInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopInstancesResponse
		var err error
		defer close(result)
		response, err = client.StopInstances(request)
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

// StopInstancesRequest is the request struct for api StopInstances
type StopInstancesRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query" name:"ClientToken"`
	ForceStop   requests.Boolean `position:"Query" name:"ForceStop"`
	InstanceIds string           `position:"Query" name:"InstanceIds"`
}

// StopInstancesResponse is the response struct for api StopInstances
type StopInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopInstancesRequest creates a request to invoke StopInstances API
func CreateStopInstancesRequest() (request *StopInstancesRequest) {
	request = &StopInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "StopInstances", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopInstancesResponse creates a response to parse from StopInstances response
func CreateStopInstancesResponse() (response *StopInstancesResponse) {
	response = &StopInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}