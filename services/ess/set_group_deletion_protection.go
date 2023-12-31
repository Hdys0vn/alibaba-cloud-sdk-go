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

// SetGroupDeletionProtection invokes the ess.SetGroupDeletionProtection API synchronously
func (client *Client) SetGroupDeletionProtection(request *SetGroupDeletionProtectionRequest) (response *SetGroupDeletionProtectionResponse, err error) {
	response = CreateSetGroupDeletionProtectionResponse()
	err = client.DoAction(request, response)
	return
}

// SetGroupDeletionProtectionWithChan invokes the ess.SetGroupDeletionProtection API asynchronously
func (client *Client) SetGroupDeletionProtectionWithChan(request *SetGroupDeletionProtectionRequest) (<-chan *SetGroupDeletionProtectionResponse, <-chan error) {
	responseChan := make(chan *SetGroupDeletionProtectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetGroupDeletionProtection(request)
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

// SetGroupDeletionProtectionWithCallback invokes the ess.SetGroupDeletionProtection API asynchronously
func (client *Client) SetGroupDeletionProtectionWithCallback(request *SetGroupDeletionProtectionRequest, callback func(response *SetGroupDeletionProtectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetGroupDeletionProtectionResponse
		var err error
		defer close(result)
		response, err = client.SetGroupDeletionProtection(request)
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

// SetGroupDeletionProtectionRequest is the request struct for api SetGroupDeletionProtection
type SetGroupDeletionProtectionRequest struct {
	*requests.RpcRequest
	ScalingGroupId          string           `position:"Query" name:"ScalingGroupId"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	GroupDeletionProtection requests.Boolean `position:"Query" name:"GroupDeletionProtection"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
}

// SetGroupDeletionProtectionResponse is the response struct for api SetGroupDeletionProtection
type SetGroupDeletionProtectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetGroupDeletionProtectionRequest creates a request to invoke SetGroupDeletionProtection API
func CreateSetGroupDeletionProtectionRequest() (request *SetGroupDeletionProtectionRequest) {
	request = &SetGroupDeletionProtectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "SetGroupDeletionProtection", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetGroupDeletionProtectionResponse creates a response to parse from SetGroupDeletionProtection response
func CreateSetGroupDeletionProtectionResponse() (response *SetGroupDeletionProtectionResponse) {
	response = &SetGroupDeletionProtectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
