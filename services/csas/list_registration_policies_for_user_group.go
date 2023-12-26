package csas

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

// ListRegistrationPoliciesForUserGroup invokes the csas.ListRegistrationPoliciesForUserGroup API synchronously
func (client *Client) ListRegistrationPoliciesForUserGroup(request *ListRegistrationPoliciesForUserGroupRequest) (response *ListRegistrationPoliciesForUserGroupResponse, err error) {
	response = CreateListRegistrationPoliciesForUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListRegistrationPoliciesForUserGroupWithChan invokes the csas.ListRegistrationPoliciesForUserGroup API asynchronously
func (client *Client) ListRegistrationPoliciesForUserGroupWithChan(request *ListRegistrationPoliciesForUserGroupRequest) (<-chan *ListRegistrationPoliciesForUserGroupResponse, <-chan error) {
	responseChan := make(chan *ListRegistrationPoliciesForUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRegistrationPoliciesForUserGroup(request)
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

// ListRegistrationPoliciesForUserGroupWithCallback invokes the csas.ListRegistrationPoliciesForUserGroup API asynchronously
func (client *Client) ListRegistrationPoliciesForUserGroupWithCallback(request *ListRegistrationPoliciesForUserGroupRequest, callback func(response *ListRegistrationPoliciesForUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRegistrationPoliciesForUserGroupResponse
		var err error
		defer close(result)
		response, err = client.ListRegistrationPoliciesForUserGroup(request)
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

// ListRegistrationPoliciesForUserGroupRequest is the request struct for api ListRegistrationPoliciesForUserGroup
type ListRegistrationPoliciesForUserGroupRequest struct {
	*requests.RpcRequest
	UserGroupIds *[]string `position:"Query" name:"UserGroupIds"  type:"Repeated"`
	SourceIp     string    `position:"Query" name:"SourceIp"`
}

// ListRegistrationPoliciesForUserGroupResponse is the response struct for api ListRegistrationPoliciesForUserGroup
type ListRegistrationPoliciesForUserGroupResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	UserGroups []Data `json:"UserGroups" xml:"UserGroups"`
}

// CreateListRegistrationPoliciesForUserGroupRequest creates a request to invoke ListRegistrationPoliciesForUserGroup API
func CreateListRegistrationPoliciesForUserGroupRequest() (request *ListRegistrationPoliciesForUserGroupRequest) {
	request = &ListRegistrationPoliciesForUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListRegistrationPoliciesForUserGroup", "", "")
	request.Method = requests.GET
	return
}

// CreateListRegistrationPoliciesForUserGroupResponse creates a response to parse from ListRegistrationPoliciesForUserGroup response
func CreateListRegistrationPoliciesForUserGroupResponse() (response *ListRegistrationPoliciesForUserGroupResponse) {
	response = &ListRegistrationPoliciesForUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}