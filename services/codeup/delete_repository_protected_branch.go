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

// DeleteRepositoryProtectedBranch invokes the codeup.DeleteRepositoryProtectedBranch API synchronously
func (client *Client) DeleteRepositoryProtectedBranch(request *DeleteRepositoryProtectedBranchRequest) (response *DeleteRepositoryProtectedBranchResponse, err error) {
	response = CreateDeleteRepositoryProtectedBranchResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepositoryProtectedBranchWithChan invokes the codeup.DeleteRepositoryProtectedBranch API asynchronously
func (client *Client) DeleteRepositoryProtectedBranchWithChan(request *DeleteRepositoryProtectedBranchRequest) (<-chan *DeleteRepositoryProtectedBranchResponse, <-chan error) {
	responseChan := make(chan *DeleteRepositoryProtectedBranchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepositoryProtectedBranch(request)
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

// DeleteRepositoryProtectedBranchWithCallback invokes the codeup.DeleteRepositoryProtectedBranch API asynchronously
func (client *Client) DeleteRepositoryProtectedBranchWithCallback(request *DeleteRepositoryProtectedBranchRequest, callback func(response *DeleteRepositoryProtectedBranchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepositoryProtectedBranchResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepositoryProtectedBranch(request)
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

// DeleteRepositoryProtectedBranchRequest is the request struct for api DeleteRepositoryProtectedBranch
type DeleteRepositoryProtectedBranchRequest struct {
	*requests.RoaRequest
	OrganizationId    string           `position:"Query" name:"OrganizationId"`
	ProtectedBranchId requests.Integer `position:"Path" name:"ProtectedBranchId"`
	AccessToken       string           `position:"Query" name:"AccessToken"`
	ProjectId         requests.Integer `position:"Path" name:"ProjectId"`
}

// DeleteRepositoryProtectedBranchResponse is the response struct for api DeleteRepositoryProtectedBranch
type DeleteRepositoryProtectedBranchResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateDeleteRepositoryProtectedBranchRequest creates a request to invoke DeleteRepositoryProtectedBranch API
func CreateDeleteRepositoryProtectedBranchRequest() (request *DeleteRepositoryProtectedBranchRequest) {
	request = &DeleteRepositoryProtectedBranchRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "DeleteRepositoryProtectedBranch", "/api/v4/projects/[ProjectId]/repository/protect_branches/[ProtectedBranchId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteRepositoryProtectedBranchResponse creates a response to parse from DeleteRepositoryProtectedBranch response
func CreateDeleteRepositoryProtectedBranchResponse() (response *DeleteRepositoryProtectedBranchResponse) {
	response = &DeleteRepositoryProtectedBranchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
