package vod

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

// DeleteEditingProject invokes the vod.DeleteEditingProject API synchronously
func (client *Client) DeleteEditingProject(request *DeleteEditingProjectRequest) (response *DeleteEditingProjectResponse, err error) {
	response = CreateDeleteEditingProjectResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEditingProjectWithChan invokes the vod.DeleteEditingProject API asynchronously
func (client *Client) DeleteEditingProjectWithChan(request *DeleteEditingProjectRequest) (<-chan *DeleteEditingProjectResponse, <-chan error) {
	responseChan := make(chan *DeleteEditingProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEditingProject(request)
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

// DeleteEditingProjectWithCallback invokes the vod.DeleteEditingProject API asynchronously
func (client *Client) DeleteEditingProjectWithCallback(request *DeleteEditingProjectRequest, callback func(response *DeleteEditingProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEditingProjectResponse
		var err error
		defer close(result)
		response, err = client.DeleteEditingProject(request)
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

// DeleteEditingProjectRequest is the request struct for api DeleteEditingProject
type DeleteEditingProjectRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ProjectIds           string `position:"Query" name:"ProjectIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

// DeleteEditingProjectResponse is the response struct for api DeleteEditingProject
type DeleteEditingProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEditingProjectRequest creates a request to invoke DeleteEditingProject API
func CreateDeleteEditingProjectRequest() (request *DeleteEditingProjectRequest) {
	request = &DeleteEditingProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteEditingProject", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEditingProjectResponse creates a response to parse from DeleteEditingProject response
func CreateDeleteEditingProjectResponse() (response *DeleteEditingProjectResponse) {
	response = &DeleteEditingProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
