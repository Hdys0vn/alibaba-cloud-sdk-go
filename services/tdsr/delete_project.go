package tdsr

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

// DeleteProject invokes the tdsr.DeleteProject API synchronously
// api document: https://help.aliyun.com/api/tdsr/deleteproject.html
func (client *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
	response = CreateDeleteProjectResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteProjectWithChan invokes the tdsr.DeleteProject API asynchronously
// api document: https://help.aliyun.com/api/tdsr/deleteproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProjectWithChan(request *DeleteProjectRequest) (<-chan *DeleteProjectResponse, <-chan error) {
	responseChan := make(chan *DeleteProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteProject(request)
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

// DeleteProjectWithCallback invokes the tdsr.DeleteProject API asynchronously
// api document: https://help.aliyun.com/api/tdsr/deleteproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProjectWithCallback(request *DeleteProjectRequest, callback func(response *DeleteProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteProjectResponse
		var err error
		defer close(result)
		response, err = client.DeleteProject(request)
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

// DeleteProjectRequest is the request struct for api DeleteProject
type DeleteProjectRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DeleteProjectResponse is the response struct for api DeleteProject
type DeleteProjectResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateDeleteProjectRequest creates a request to invoke DeleteProject API
func CreateDeleteProjectRequest() (request *DeleteProjectRequest) {
	request = &DeleteProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("tdsr", "2020-01-01", "DeleteProject", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteProjectResponse creates a response to parse from DeleteProject response
func CreateDeleteProjectResponse() (response *DeleteProjectResponse) {
	response = &DeleteProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
