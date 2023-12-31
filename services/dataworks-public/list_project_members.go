package dataworks_public

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

// ListProjectMembers invokes the dataworks_public.ListProjectMembers API synchronously
func (client *Client) ListProjectMembers(request *ListProjectMembersRequest) (response *ListProjectMembersResponse, err error) {
	response = CreateListProjectMembersResponse()
	err = client.DoAction(request, response)
	return
}

// ListProjectMembersWithChan invokes the dataworks_public.ListProjectMembers API asynchronously
func (client *Client) ListProjectMembersWithChan(request *ListProjectMembersRequest) (<-chan *ListProjectMembersResponse, <-chan error) {
	responseChan := make(chan *ListProjectMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProjectMembers(request)
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

// ListProjectMembersWithCallback invokes the dataworks_public.ListProjectMembers API asynchronously
func (client *Client) ListProjectMembersWithCallback(request *ListProjectMembersRequest, callback func(response *ListProjectMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProjectMembersResponse
		var err error
		defer close(result)
		response, err = client.ListProjectMembers(request)
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

// ListProjectMembersRequest is the request struct for api ListProjectMembers
type ListProjectMembersRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	ProjectId  requests.Integer `position:"Query" name:"ProjectId"`
}

// ListProjectMembersResponse is the response struct for api ListProjectMembers
type ListProjectMembersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListProjectMembersRequest creates a request to invoke ListProjectMembers API
func CreateListProjectMembersRequest() (request *ListProjectMembersRequest) {
	request = &ListProjectMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListProjectMembers", "", "")
	request.Method = requests.POST
	return
}

// CreateListProjectMembersResponse creates a response to parse from ListProjectMembers response
func CreateListProjectMembersResponse() (response *ListProjectMembersResponse) {
	response = &ListProjectMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
