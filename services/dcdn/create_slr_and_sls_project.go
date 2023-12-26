package dcdn

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

// CreateSlrAndSlsProject invokes the dcdn.CreateSlrAndSlsProject API synchronously
func (client *Client) CreateSlrAndSlsProject(request *CreateSlrAndSlsProjectRequest) (response *CreateSlrAndSlsProjectResponse, err error) {
	response = CreateCreateSlrAndSlsProjectResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSlrAndSlsProjectWithChan invokes the dcdn.CreateSlrAndSlsProject API asynchronously
func (client *Client) CreateSlrAndSlsProjectWithChan(request *CreateSlrAndSlsProjectRequest) (<-chan *CreateSlrAndSlsProjectResponse, <-chan error) {
	responseChan := make(chan *CreateSlrAndSlsProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSlrAndSlsProject(request)
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

// CreateSlrAndSlsProjectWithCallback invokes the dcdn.CreateSlrAndSlsProject API asynchronously
func (client *Client) CreateSlrAndSlsProjectWithCallback(request *CreateSlrAndSlsProjectRequest, callback func(response *CreateSlrAndSlsProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSlrAndSlsProjectResponse
		var err error
		defer close(result)
		response, err = client.CreateSlrAndSlsProject(request)
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

// CreateSlrAndSlsProjectRequest is the request struct for api CreateSlrAndSlsProject
type CreateSlrAndSlsProjectRequest struct {
	*requests.RpcRequest
	Region       string `position:"Body" name:"Region"`
	BusinessType string `position:"Body" name:"BusinessType"`
}

// CreateSlrAndSlsProjectResponse is the response struct for api CreateSlrAndSlsProject
type CreateSlrAndSlsProjectResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	SlsInfo   SlsInfo `json:"SlsInfo" xml:"SlsInfo"`
}

// CreateCreateSlrAndSlsProjectRequest creates a request to invoke CreateSlrAndSlsProject API
func CreateCreateSlrAndSlsProjectRequest() (request *CreateSlrAndSlsProjectRequest) {
	request = &CreateSlrAndSlsProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "CreateSlrAndSlsProject", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSlrAndSlsProjectResponse creates a response to parse from CreateSlrAndSlsProject response
func CreateCreateSlrAndSlsProjectResponse() (response *CreateSlrAndSlsProjectResponse) {
	response = &CreateSlrAndSlsProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}