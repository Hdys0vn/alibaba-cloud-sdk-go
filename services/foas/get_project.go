package foas

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

// GetProject invokes the foas.GetProject API synchronously
func (client *Client) GetProject(request *GetProjectRequest) (response *GetProjectResponse, err error) {
	response = CreateGetProjectResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectWithChan invokes the foas.GetProject API asynchronously
func (client *Client) GetProjectWithChan(request *GetProjectRequest) (<-chan *GetProjectResponse, <-chan error) {
	responseChan := make(chan *GetProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProject(request)
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

// GetProjectWithCallback invokes the foas.GetProject API asynchronously
func (client *Client) GetProjectWithCallback(request *GetProjectRequest, callback func(response *GetProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectResponse
		var err error
		defer close(result)
		response, err = client.GetProject(request)
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

// GetProjectRequest is the request struct for api GetProject
type GetProjectRequest struct {
	*requests.RoaRequest
	ProjectName string `position:"Path" name:"projectName"`
}

// GetProjectResponse is the response struct for api GetProject
type GetProjectResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Project   Project `json:"Project" xml:"Project"`
}

// CreateGetProjectRequest creates a request to invoke GetProject API
func CreateGetProjectRequest() (request *GetProjectRequest) {
	request = &GetProjectRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetProject", "/api/v2/projects/[projectName]", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetProjectResponse creates a response to parse from GetProject response
func CreateGetProjectResponse() (response *GetProjectResponse) {
	response = &GetProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
