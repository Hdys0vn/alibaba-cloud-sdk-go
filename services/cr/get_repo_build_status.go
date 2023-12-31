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

// GetRepoBuildStatus invokes the cr.GetRepoBuildStatus API synchronously
func (client *Client) GetRepoBuildStatus(request *GetRepoBuildStatusRequest) (response *GetRepoBuildStatusResponse, err error) {
	response = CreateGetRepoBuildStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoBuildStatusWithChan invokes the cr.GetRepoBuildStatus API asynchronously
func (client *Client) GetRepoBuildStatusWithChan(request *GetRepoBuildStatusRequest) (<-chan *GetRepoBuildStatusResponse, <-chan error) {
	responseChan := make(chan *GetRepoBuildStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoBuildStatus(request)
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

// GetRepoBuildStatusWithCallback invokes the cr.GetRepoBuildStatus API asynchronously
func (client *Client) GetRepoBuildStatusWithCallback(request *GetRepoBuildStatusRequest, callback func(response *GetRepoBuildStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoBuildStatusResponse
		var err error
		defer close(result)
		response, err = client.GetRepoBuildStatus(request)
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

// GetRepoBuildStatusRequest is the request struct for api GetRepoBuildStatus
type GetRepoBuildStatusRequest struct {
	*requests.RoaRequest
	BuildId       string `position:"Path" name:"BuildId"`
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// GetRepoBuildStatusResponse is the response struct for api GetRepoBuildStatus
type GetRepoBuildStatusResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoBuildStatusRequest creates a request to invoke GetRepoBuildStatus API
func CreateGetRepoBuildStatusRequest() (request *GetRepoBuildStatusRequest) {
	request = &GetRepoBuildStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildStatus", "/repos/[RepoNamespace]/[RepoName]/build/[BuildId]/status", "acr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRepoBuildStatusResponse creates a response to parse from GetRepoBuildStatus response
func CreateGetRepoBuildStatusResponse() (response *GetRepoBuildStatusResponse) {
	response = &GetRepoBuildStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
