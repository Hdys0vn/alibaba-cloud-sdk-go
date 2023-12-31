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

// AttachVServerGroups invokes the ess.AttachVServerGroups API synchronously
func (client *Client) AttachVServerGroups(request *AttachVServerGroupsRequest) (response *AttachVServerGroupsResponse, err error) {
	response = CreateAttachVServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// AttachVServerGroupsWithChan invokes the ess.AttachVServerGroups API asynchronously
func (client *Client) AttachVServerGroupsWithChan(request *AttachVServerGroupsRequest) (<-chan *AttachVServerGroupsResponse, <-chan error) {
	responseChan := make(chan *AttachVServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachVServerGroups(request)
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

// AttachVServerGroupsWithCallback invokes the ess.AttachVServerGroups API asynchronously
func (client *Client) AttachVServerGroupsWithCallback(request *AttachVServerGroupsRequest, callback func(response *AttachVServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachVServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.AttachVServerGroups(request)
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

// AttachVServerGroupsRequest is the request struct for api AttachVServerGroups
type AttachVServerGroupsRequest struct {
	*requests.RpcRequest
	ClientToken          string                             `position:"Query" name:"ClientToken"`
	ScalingGroupId       string                             `position:"Query" name:"ScalingGroupId"`
	ForceAttach          requests.Boolean                   `position:"Query" name:"ForceAttach"`
	ResourceOwnerAccount string                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer                   `position:"Query" name:"OwnerId"`
	VServerGroup         *[]AttachVServerGroupsVServerGroup `position:"Query" name:"VServerGroup"  type:"Repeated"`
}

// AttachVServerGroupsVServerGroup is a repeated param struct in AttachVServerGroupsRequest
type AttachVServerGroupsVServerGroup struct {
	LoadBalancerId        string                                                  `name:"LoadBalancerId"`
	VServerGroupAttribute *[]AttachVServerGroupsVServerGroupVServerGroupAttribute `name:"VServerGroupAttribute" type:"Repeated"`
}

// AttachVServerGroupsVServerGroupVServerGroupAttribute is a repeated param struct in AttachVServerGroupsRequest
type AttachVServerGroupsVServerGroupVServerGroupAttribute struct {
	VServerGroupId string `name:"VServerGroupId"`
	Port           string `name:"Port"`
	Weight         string `name:"Weight"`
}

// AttachVServerGroupsResponse is the response struct for api AttachVServerGroups
type AttachVServerGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachVServerGroupsRequest creates a request to invoke AttachVServerGroups API
func CreateAttachVServerGroupsRequest() (request *AttachVServerGroupsRequest) {
	request = &AttachVServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "AttachVServerGroups", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachVServerGroupsResponse creates a response to parse from AttachVServerGroups response
func CreateAttachVServerGroupsResponse() (response *AttachVServerGroupsResponse) {
	response = &AttachVServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
