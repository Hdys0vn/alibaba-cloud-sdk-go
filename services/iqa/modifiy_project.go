package iqa

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

// ModifiyProject invokes the iqa.ModifiyProject API synchronously
// api document: https://help.aliyun.com/api/iqa/modifiyproject.html
func (client *Client) ModifiyProject(request *ModifiyProjectRequest) (response *ModifiyProjectResponse, err error) {
	response = CreateModifiyProjectResponse()
	err = client.DoAction(request, response)
	return
}

// ModifiyProjectWithChan invokes the iqa.ModifiyProject API asynchronously
// api document: https://help.aliyun.com/api/iqa/modifiyproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifiyProjectWithChan(request *ModifiyProjectRequest) (<-chan *ModifiyProjectResponse, <-chan error) {
	responseChan := make(chan *ModifiyProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifiyProject(request)
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

// ModifiyProjectWithCallback invokes the iqa.ModifiyProject API asynchronously
// api document: https://help.aliyun.com/api/iqa/modifiyproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifiyProjectWithCallback(request *ModifiyProjectRequest, callback func(response *ModifiyProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifiyProjectResponse
		var err error
		defer close(result)
		response, err = client.ModifiyProject(request)
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

// ModifiyProjectRequest is the request struct for api ModifiyProject
type ModifiyProjectRequest struct {
	*requests.RpcRequest
	ProjectName string `position:"Body" name:"ProjectName"`
	ModelId     string `position:"Body" name:"ModelId"`
	ProjectId   string `position:"Body" name:"ProjectId"`
}

// ModifiyProjectResponse is the response struct for api ModifiyProject
type ModifiyProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ProjectId string `json:"ProjectId" xml:"ProjectId"`
}

// CreateModifiyProjectRequest creates a request to invoke ModifiyProject API
func CreateModifiyProjectRequest() (request *ModifiyProjectRequest) {
	request = &ModifiyProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("iqa", "2019-08-13", "ModifiyProject", "iqa", "openAPI")
	return
}

// CreateModifiyProjectResponse creates a response to parse from ModifiyProject response
func CreateModifiyProjectResponse() (response *ModifiyProjectResponse) {
	response = &ModifiyProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
