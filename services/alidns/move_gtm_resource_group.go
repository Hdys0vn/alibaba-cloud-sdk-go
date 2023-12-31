package alidns

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

// MoveGtmResourceGroup invokes the alidns.MoveGtmResourceGroup API synchronously
func (client *Client) MoveGtmResourceGroup(request *MoveGtmResourceGroupRequest) (response *MoveGtmResourceGroupResponse, err error) {
	response = CreateMoveGtmResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// MoveGtmResourceGroupWithChan invokes the alidns.MoveGtmResourceGroup API asynchronously
func (client *Client) MoveGtmResourceGroupWithChan(request *MoveGtmResourceGroupRequest) (<-chan *MoveGtmResourceGroupResponse, <-chan error) {
	responseChan := make(chan *MoveGtmResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoveGtmResourceGroup(request)
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

// MoveGtmResourceGroupWithCallback invokes the alidns.MoveGtmResourceGroup API asynchronously
func (client *Client) MoveGtmResourceGroupWithCallback(request *MoveGtmResourceGroupRequest, callback func(response *MoveGtmResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoveGtmResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.MoveGtmResourceGroup(request)
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

// MoveGtmResourceGroupRequest is the request struct for api MoveGtmResourceGroup
type MoveGtmResourceGroupRequest struct {
	*requests.RpcRequest
	ResourceId         string `position:"Query" name:"ResourceId"`
	NewResourceGroupId string `position:"Query" name:"NewResourceGroupId"`
	UserClientIp       string `position:"Query" name:"UserClientIp"`
	Lang               string `position:"Query" name:"Lang"`
}

// MoveGtmResourceGroupResponse is the response struct for api MoveGtmResourceGroup
type MoveGtmResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMoveGtmResourceGroupRequest creates a request to invoke MoveGtmResourceGroup API
func CreateMoveGtmResourceGroupRequest() (request *MoveGtmResourceGroupRequest) {
	request = &MoveGtmResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "MoveGtmResourceGroup", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMoveGtmResourceGroupResponse creates a response to parse from MoveGtmResourceGroup response
func CreateMoveGtmResourceGroupResponse() (response *MoveGtmResourceGroupResponse) {
	response = &MoveGtmResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
