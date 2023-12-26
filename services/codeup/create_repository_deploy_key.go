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

// CreateRepositoryDeployKey invokes the codeup.CreateRepositoryDeployKey API synchronously
func (client *Client) CreateRepositoryDeployKey(request *CreateRepositoryDeployKeyRequest) (response *CreateRepositoryDeployKeyResponse, err error) {
	response = CreateCreateRepositoryDeployKeyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRepositoryDeployKeyWithChan invokes the codeup.CreateRepositoryDeployKey API asynchronously
func (client *Client) CreateRepositoryDeployKeyWithChan(request *CreateRepositoryDeployKeyRequest) (<-chan *CreateRepositoryDeployKeyResponse, <-chan error) {
	responseChan := make(chan *CreateRepositoryDeployKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepositoryDeployKey(request)
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

// CreateRepositoryDeployKeyWithCallback invokes the codeup.CreateRepositoryDeployKey API asynchronously
func (client *Client) CreateRepositoryDeployKeyWithCallback(request *CreateRepositoryDeployKeyRequest, callback func(response *CreateRepositoryDeployKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepositoryDeployKeyResponse
		var err error
		defer close(result)
		response, err = client.CreateRepositoryDeployKey(request)
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

// CreateRepositoryDeployKeyRequest is the request struct for api CreateRepositoryDeployKey
type CreateRepositoryDeployKeyRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// CreateRepositoryDeployKeyResponse is the response struct for api CreateRepositoryDeployKey
type CreateRepositoryDeployKeyResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateCreateRepositoryDeployKeyRequest creates a request to invoke CreateRepositoryDeployKey API
func CreateCreateRepositoryDeployKeyRequest() (request *CreateRepositoryDeployKeyRequest) {
	request = &CreateRepositoryDeployKeyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "CreateRepositoryDeployKey", "/api/v3/projects/[ProjectId]/keys", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRepositoryDeployKeyResponse creates a response to parse from CreateRepositoryDeployKey response
func CreateCreateRepositoryDeployKeyResponse() (response *CreateRepositoryDeployKeyResponse) {
	response = &CreateRepositoryDeployKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
