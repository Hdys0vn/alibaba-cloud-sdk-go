package oceanbasepro

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

// StartProject invokes the oceanbasepro.StartProject API synchronously
func (client *Client) StartProject(request *StartProjectRequest) (response *StartProjectResponse, err error) {
	response = CreateStartProjectResponse()
	err = client.DoAction(request, response)
	return
}

// StartProjectWithChan invokes the oceanbasepro.StartProject API asynchronously
func (client *Client) StartProjectWithChan(request *StartProjectRequest) (<-chan *StartProjectResponse, <-chan error) {
	responseChan := make(chan *StartProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartProject(request)
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

// StartProjectWithCallback invokes the oceanbasepro.StartProject API asynchronously
func (client *Client) StartProjectWithCallback(request *StartProjectRequest, callback func(response *StartProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartProjectResponse
		var err error
		defer close(result)
		response, err = client.StartProject(request)
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

// StartProjectRequest is the request struct for api StartProject
type StartProjectRequest struct {
	*requests.RpcRequest
	Id string `position:"Body" name:"Id"`
}

// StartProjectResponse is the response struct for api StartProject
type StartProjectResponse struct {
	*responses.BaseResponse
}

// CreateStartProjectRequest creates a request to invoke StartProject API
func CreateStartProjectRequest() (request *StartProjectRequest) {
	request = &StartProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "StartProject", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartProjectResponse creates a response to parse from StartProject response
func CreateStartProjectResponse() (response *StartProjectResponse) {
	response = &StartProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
