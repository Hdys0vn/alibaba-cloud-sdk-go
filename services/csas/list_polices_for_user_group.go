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

// ListPolicesForUserGroup invokes the csas.ListPolicesForUserGroup API synchronously
func (client *Client) ListPolicesForUserGroup(request *ListPolicesForUserGroupRequest) (response *ListPolicesForUserGroupResponse, err error) {
	response = CreateListPolicesForUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListPolicesForUserGroupWithChan invokes the csas.ListPolicesForUserGroup API asynchronously
func (client *Client) ListPolicesForUserGroupWithChan(request *ListPolicesForUserGroupRequest) (<-chan *ListPolicesForUserGroupResponse, <-chan error) {
	responseChan := make(chan *ListPolicesForUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPolicesForUserGroup(request)
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

// ListPolicesForUserGroupWithCallback invokes the csas.ListPolicesForUserGroup API asynchronously
func (client *Client) ListPolicesForUserGroupWithCallback(request *ListPolicesForUserGroupRequest, callback func(response *ListPolicesForUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPolicesForUserGroupResponse
		var err error
		defer close(result)
		response, err = client.ListPolicesForUserGroup(request)
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

// ListPolicesForUserGroupRequest is the request struct for api ListPolicesForUserGroup
type ListPolicesForUserGroupRequest struct {
	*requests.RpcRequest
	UserGroupIds *[]string `position:"Query" name:"UserGroupIds"  type:"Repeated"`
	SourceIp     string    `position:"Query" name:"SourceIp"`
}

// ListPolicesForUserGroupResponse is the response struct for api ListPolicesForUserGroup
type ListPolicesForUserGroupResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	UserGroups []UserGroup `json:"UserGroups" xml:"UserGroups"`
}

// CreateListPolicesForUserGroupRequest creates a request to invoke ListPolicesForUserGroup API
func CreateListPolicesForUserGroupRequest() (request *ListPolicesForUserGroupRequest) {
	request = &ListPolicesForUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListPolicesForUserGroup", "", "")
	request.Method = requests.GET
	return
}

// CreateListPolicesForUserGroupResponse creates a response to parse from ListPolicesForUserGroup response
func CreateListPolicesForUserGroupResponse() (response *ListPolicesForUserGroupResponse) {
	response = &ListPolicesForUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
