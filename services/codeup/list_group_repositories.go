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

// ListGroupRepositories invokes the codeup.ListGroupRepositories API synchronously
func (client *Client) ListGroupRepositories(request *ListGroupRepositoriesRequest) (response *ListGroupRepositoriesResponse, err error) {
	response = CreateListGroupRepositoriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListGroupRepositoriesWithChan invokes the codeup.ListGroupRepositories API asynchronously
func (client *Client) ListGroupRepositoriesWithChan(request *ListGroupRepositoriesRequest) (<-chan *ListGroupRepositoriesResponse, <-chan error) {
	responseChan := make(chan *ListGroupRepositoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGroupRepositories(request)
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

// ListGroupRepositoriesWithCallback invokes the codeup.ListGroupRepositories API asynchronously
func (client *Client) ListGroupRepositoriesWithCallback(request *ListGroupRepositoriesRequest, callback func(response *ListGroupRepositoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGroupRepositoriesResponse
		var err error
		defer close(result)
		response, err = client.ListGroupRepositories(request)
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

// ListGroupRepositoriesRequest is the request struct for api ListGroupRepositories
type ListGroupRepositoriesRequest struct {
	*requests.RoaRequest
	AccessToken    string           `position:"Query" name:"AccessToken"`
	IsMember       requests.Boolean `position:"Query" name:"IsMember"`
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	Search         string           `position:"Query" name:"Search"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	Identity       string           `position:"Path" name:"Identity"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Page           requests.Integer `position:"Query" name:"Page"`
}

// ListGroupRepositoriesResponse is the response struct for api ListGroupRepositories
type ListGroupRepositoriesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool         `json:"Success" xml:"Success"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int64        `json:"Total" xml:"Total"`
	Result       []ResultItem `json:"Result" xml:"Result"`
}

// CreateListGroupRepositoriesRequest creates a request to invoke ListGroupRepositories API
func CreateListGroupRepositoriesRequest() (request *ListGroupRepositoriesRequest) {
	request = &ListGroupRepositoriesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListGroupRepositories", "/api/v3/groups/[Identity]/projects", "", "")
	request.Method = requests.GET
	return
}

// CreateListGroupRepositoriesResponse creates a response to parse from ListGroupRepositories response
func CreateListGroupRepositoriesResponse() (response *ListGroupRepositoriesResponse) {
	response = &ListGroupRepositoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
