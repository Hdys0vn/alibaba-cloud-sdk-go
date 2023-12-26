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

// DisableScalingGroup invokes the ess.DisableScalingGroup API synchronously
func (client *Client) DisableScalingGroup(request *DisableScalingGroupRequest) (response *DisableScalingGroupResponse, err error) {
	response = CreateDisableScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DisableScalingGroupWithChan invokes the ess.DisableScalingGroup API asynchronously
func (client *Client) DisableScalingGroupWithChan(request *DisableScalingGroupRequest) (<-chan *DisableScalingGroupResponse, <-chan error) {
	responseChan := make(chan *DisableScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableScalingGroup(request)
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

// DisableScalingGroupWithCallback invokes the ess.DisableScalingGroup API asynchronously
func (client *Client) DisableScalingGroupWithCallback(request *DisableScalingGroupRequest, callback func(response *DisableScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.DisableScalingGroup(request)
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

// DisableScalingGroupRequest is the request struct for api DisableScalingGroup
type DisableScalingGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DisableScalingGroupResponse is the response struct for api DisableScalingGroup
type DisableScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableScalingGroupRequest creates a request to invoke DisableScalingGroup API
func CreateDisableScalingGroupRequest() (request *DisableScalingGroupRequest) {
	request = &DisableScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DisableScalingGroup", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableScalingGroupResponse creates a response to parse from DisableScalingGroup response
func CreateDisableScalingGroupResponse() (response *DisableScalingGroupResponse) {
	response = &DisableScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
