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

// GetRepositoryCommit invokes the codeup.GetRepositoryCommit API synchronously
func (client *Client) GetRepositoryCommit(request *GetRepositoryCommitRequest) (response *GetRepositoryCommitResponse, err error) {
	response = CreateGetRepositoryCommitResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepositoryCommitWithChan invokes the codeup.GetRepositoryCommit API asynchronously
func (client *Client) GetRepositoryCommitWithChan(request *GetRepositoryCommitRequest) (<-chan *GetRepositoryCommitResponse, <-chan error) {
	responseChan := make(chan *GetRepositoryCommitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepositoryCommit(request)
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

// GetRepositoryCommitWithCallback invokes the codeup.GetRepositoryCommit API asynchronously
func (client *Client) GetRepositoryCommitWithCallback(request *GetRepositoryCommitRequest, callback func(response *GetRepositoryCommitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepositoryCommitResponse
		var err error
		defer close(result)
		response, err = client.GetRepositoryCommit(request)
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

// GetRepositoryCommitRequest is the request struct for api GetRepositoryCommit
type GetRepositoryCommitRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
	Sha            string           `position:"Path" name:"Sha"`
}

// GetRepositoryCommitResponse is the response struct for api GetRepositoryCommit
type GetRepositoryCommitResponse struct {
	*responses.BaseResponse
	ErrorCode    string                      `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                      `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	Success      bool                        `json:"Success" xml:"Success"`
	Result       ResultInGetRepositoryCommit `json:"Result" xml:"Result"`
}

// CreateGetRepositoryCommitRequest creates a request to invoke GetRepositoryCommit API
func CreateGetRepositoryCommitRequest() (request *GetRepositoryCommitRequest) {
	request = &GetRepositoryCommitRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "GetRepositoryCommit", "/api/v4/projects/[ProjectId]/repository/commits/[Sha]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetRepositoryCommitResponse creates a response to parse from GetRepositoryCommit response
func CreateGetRepositoryCommitResponse() (response *GetRepositoryCommitResponse) {
	response = &GetRepositoryCommitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
