package ess

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

// ExitStandby invokes the ess.ExitStandby API synchronously
func (client *Client) ExitStandby(request *ExitStandbyRequest) (response *ExitStandbyResponse, err error) {
	response = CreateExitStandbyResponse()
	err = client.DoAction(request, response)
	return
}

// ExitStandbyWithChan invokes the ess.ExitStandby API asynchronously
func (client *Client) ExitStandbyWithChan(request *ExitStandbyRequest) (<-chan *ExitStandbyResponse, <-chan error) {
	responseChan := make(chan *ExitStandbyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExitStandby(request)
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

// ExitStandbyWithCallback invokes the ess.ExitStandby API asynchronously
func (client *Client) ExitStandbyWithCallback(request *ExitStandbyRequest, callback func(response *ExitStandbyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExitStandbyResponse
		var err error
		defer close(result)
		response, err = client.ExitStandby(request)
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

// ExitStandbyRequest is the request struct for api ExitStandby
type ExitStandbyRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Async                requests.Boolean `position:"Query" name:"Async"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// ExitStandbyResponse is the response struct for api ExitStandby
type ExitStandbyResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
}

// CreateExitStandbyRequest creates a request to invoke ExitStandby API
func CreateExitStandbyRequest() (request *ExitStandbyRequest) {
	request = &ExitStandbyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ExitStandby", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExitStandbyResponse creates a response to parse from ExitStandby response
func CreateExitStandbyResponse() (response *ExitStandbyResponse) {
	response = &ExitStandbyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
