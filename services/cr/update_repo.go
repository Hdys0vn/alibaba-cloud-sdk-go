package cr

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

// UpdateRepo invokes the cr.UpdateRepo API synchronously
func (client *Client) UpdateRepo(request *UpdateRepoRequest) (response *UpdateRepoResponse, err error) {
	response = CreateUpdateRepoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRepoWithChan invokes the cr.UpdateRepo API asynchronously
func (client *Client) UpdateRepoWithChan(request *UpdateRepoRequest) (<-chan *UpdateRepoResponse, <-chan error) {
	responseChan := make(chan *UpdateRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRepo(request)
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

// UpdateRepoWithCallback invokes the cr.UpdateRepo API asynchronously
func (client *Client) UpdateRepoWithCallback(request *UpdateRepoRequest, callback func(response *UpdateRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRepoResponse
		var err error
		defer close(result)
		response, err = client.UpdateRepo(request)
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

// UpdateRepoRequest is the request struct for api UpdateRepo
type UpdateRepoRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// UpdateRepoResponse is the response struct for api UpdateRepo
type UpdateRepoResponse struct {
	*responses.BaseResponse
}

// CreateUpdateRepoRequest creates a request to invoke UpdateRepo API
func CreateUpdateRepoRequest() (request *UpdateRepoRequest) {
	request = &UpdateRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "UpdateRepo", "/repos/[RepoNamespace]/[RepoName]", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRepoResponse creates a response to parse from UpdateRepo response
func CreateUpdateRepoResponse() (response *UpdateRepoResponse) {
	response = &UpdateRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
