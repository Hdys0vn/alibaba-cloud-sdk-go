package emr

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

// ModifyFlowProject invokes the emr.ModifyFlowProject API synchronously
func (client *Client) ModifyFlowProject(request *ModifyFlowProjectRequest) (response *ModifyFlowProjectResponse, err error) {
	response = CreateModifyFlowProjectResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFlowProjectWithChan invokes the emr.ModifyFlowProject API asynchronously
func (client *Client) ModifyFlowProjectWithChan(request *ModifyFlowProjectRequest) (<-chan *ModifyFlowProjectResponse, <-chan error) {
	responseChan := make(chan *ModifyFlowProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFlowProject(request)
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

// ModifyFlowProjectWithCallback invokes the emr.ModifyFlowProject API asynchronously
func (client *Client) ModifyFlowProjectWithCallback(request *ModifyFlowProjectRequest, callback func(response *ModifyFlowProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFlowProjectResponse
		var err error
		defer close(result)
		response, err = client.ModifyFlowProject(request)
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

// ModifyFlowProjectRequest is the request struct for api ModifyFlowProject
type ModifyFlowProjectRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	Name        string `position:"Query" name:"Name"`
	ProjectId   string `position:"Query" name:"ProjectId"`
}

// ModifyFlowProjectResponse is the response struct for api ModifyFlowProject
type ModifyFlowProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateModifyFlowProjectRequest creates a request to invoke ModifyFlowProject API
func CreateModifyFlowProjectRequest() (request *ModifyFlowProjectRequest) {
	request = &ModifyFlowProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlowProject", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyFlowProjectResponse creates a response to parse from ModifyFlowProject response
func CreateModifyFlowProjectResponse() (response *ModifyFlowProjectResponse) {
	response = &ModifyFlowProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
