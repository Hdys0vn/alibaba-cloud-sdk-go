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

// EnterStandby invokes the ess.EnterStandby API synchronously
func (client *Client) EnterStandby(request *EnterStandbyRequest) (response *EnterStandbyResponse, err error) {
	response = CreateEnterStandbyResponse()
	err = client.DoAction(request, response)
	return
}

// EnterStandbyWithChan invokes the ess.EnterStandby API asynchronously
func (client *Client) EnterStandbyWithChan(request *EnterStandbyRequest) (<-chan *EnterStandbyResponse, <-chan error) {
	responseChan := make(chan *EnterStandbyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnterStandby(request)
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

// EnterStandbyWithCallback invokes the ess.EnterStandby API asynchronously
func (client *Client) EnterStandbyWithCallback(request *EnterStandbyRequest, callback func(response *EnterStandbyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnterStandbyResponse
		var err error
		defer close(result)
		response, err = client.EnterStandby(request)
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

// EnterStandbyRequest is the request struct for api EnterStandby
type EnterStandbyRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Async                requests.Boolean `position:"Query" name:"Async"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// EnterStandbyResponse is the response struct for api EnterStandby
type EnterStandbyResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
}

// CreateEnterStandbyRequest creates a request to invoke EnterStandby API
func CreateEnterStandbyRequest() (request *EnterStandbyRequest) {
	request = &EnterStandbyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "EnterStandby", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnterStandbyResponse creates a response to parse from EnterStandby response
func CreateEnterStandbyResponse() (response *EnterStandbyResponse) {
	response = &EnterStandbyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
