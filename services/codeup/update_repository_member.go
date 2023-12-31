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

// UpdateRepositoryMember invokes the codeup.UpdateRepositoryMember API synchronously
func (client *Client) UpdateRepositoryMember(request *UpdateRepositoryMemberRequest) (response *UpdateRepositoryMemberResponse, err error) {
	response = CreateUpdateRepositoryMemberResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRepositoryMemberWithChan invokes the codeup.UpdateRepositoryMember API asynchronously
func (client *Client) UpdateRepositoryMemberWithChan(request *UpdateRepositoryMemberRequest) (<-chan *UpdateRepositoryMemberResponse, <-chan error) {
	responseChan := make(chan *UpdateRepositoryMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRepositoryMember(request)
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

// UpdateRepositoryMemberWithCallback invokes the codeup.UpdateRepositoryMember API asynchronously
func (client *Client) UpdateRepositoryMemberWithCallback(request *UpdateRepositoryMemberRequest, callback func(response *UpdateRepositoryMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRepositoryMemberResponse
		var err error
		defer close(result)
		response, err = client.UpdateRepositoryMember(request)
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

// UpdateRepositoryMemberRequest is the request struct for api UpdateRepositoryMember
type UpdateRepositoryMemberRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
	UserId         requests.Integer `position:"Path" name:"UserId"`
}

// UpdateRepositoryMemberResponse is the response struct for api UpdateRepositoryMember
type UpdateRepositoryMemberResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateUpdateRepositoryMemberRequest creates a request to invoke UpdateRepositoryMember API
func CreateUpdateRepositoryMemberRequest() (request *UpdateRepositoryMemberRequest) {
	request = &UpdateRepositoryMemberRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "UpdateRepositoryMember", "/api/v3/projects/[ProjectId]/members/[UserId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateRepositoryMemberResponse creates a response to parse from UpdateRepositoryMember response
func CreateUpdateRepositoryMemberResponse() (response *UpdateRepositoryMemberResponse) {
	response = &UpdateRepositoryMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
