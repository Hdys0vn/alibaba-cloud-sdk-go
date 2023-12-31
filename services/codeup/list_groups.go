package codeup

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

// ListGroups invokes the codeup.ListGroups API synchronously
func (client *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
	response = CreateListGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListGroupsWithChan invokes the codeup.ListGroups API asynchronously
func (client *Client) ListGroupsWithChan(request *ListGroupsRequest) (<-chan *ListGroupsResponse, <-chan error) {
	responseChan := make(chan *ListGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGroups(request)
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

// ListGroupsWithCallback invokes the codeup.ListGroups API asynchronously
func (client *Client) ListGroupsWithCallback(request *ListGroupsRequest, callback func(response *ListGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListGroups(request)
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

// ListGroupsRequest is the request struct for api ListGroups
type ListGroupsRequest struct {
	*requests.RoaRequest
	OrganizationId  string           `position:"Query" name:"OrganizationId"`
	IncludePersonal requests.Boolean `position:"Query" name:"IncludePersonal"`
	Search          string           `position:"Query" name:"Search"`
	SubUserId       string           `position:"Query" name:"SubUserId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	AccessToken     string           `position:"Query" name:"AccessToken"`
	Page            requests.Integer `position:"Query" name:"Page"`
}

// ListGroupsResponse is the response struct for api ListGroups
type ListGroupsResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	ErrorCode    string                   `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool                     `json:"Success" xml:"Success"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int64                    `json:"Total" xml:"Total"`
	Result       []ResultItemInListGroups `json:"Result" xml:"Result"`
}

// CreateListGroupsRequest creates a request to invoke ListGroups API
func CreateListGroupsRequest() (request *ListGroupsRequest) {
	request = &ListGroupsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListGroups", "/api/v3/groups/all", "", "")
	request.Method = requests.GET
	return
}

// CreateListGroupsResponse creates a response to parse from ListGroups response
func CreateListGroupsResponse() (response *ListGroupsResponse) {
	response = &ListGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
