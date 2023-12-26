package oos

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

// ContinueDeployApplicationGroup invokes the oos.ContinueDeployApplicationGroup API synchronously
func (client *Client) ContinueDeployApplicationGroup(request *ContinueDeployApplicationGroupRequest) (response *ContinueDeployApplicationGroupResponse, err error) {
	response = CreateContinueDeployApplicationGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ContinueDeployApplicationGroupWithChan invokes the oos.ContinueDeployApplicationGroup API asynchronously
func (client *Client) ContinueDeployApplicationGroupWithChan(request *ContinueDeployApplicationGroupRequest) (<-chan *ContinueDeployApplicationGroupResponse, <-chan error) {
	responseChan := make(chan *ContinueDeployApplicationGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ContinueDeployApplicationGroup(request)
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

// ContinueDeployApplicationGroupWithCallback invokes the oos.ContinueDeployApplicationGroup API asynchronously
func (client *Client) ContinueDeployApplicationGroupWithCallback(request *ContinueDeployApplicationGroupRequest, callback func(response *ContinueDeployApplicationGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ContinueDeployApplicationGroupResponse
		var err error
		defer close(result)
		response, err = client.ContinueDeployApplicationGroup(request)
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

// ContinueDeployApplicationGroupRequest is the request struct for api ContinueDeployApplicationGroup
type ContinueDeployApplicationGroupRequest struct {
	*requests.RpcRequest
	DeployParameters string `position:"Query" name:"DeployParameters"`
	ApplicationName  string `position:"Query" name:"ApplicationName"`
	Name             string `position:"Query" name:"Name"`
}

// ContinueDeployApplicationGroupResponse is the response struct for api ContinueDeployApplicationGroup
type ContinueDeployApplicationGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateContinueDeployApplicationGroupRequest creates a request to invoke ContinueDeployApplicationGroup API
func CreateContinueDeployApplicationGroupRequest() (request *ContinueDeployApplicationGroupRequest) {
	request = &ContinueDeployApplicationGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ContinueDeployApplicationGroup", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateContinueDeployApplicationGroupResponse creates a response to parse from ContinueDeployApplicationGroup response
func CreateContinueDeployApplicationGroupResponse() (response *ContinueDeployApplicationGroupResponse) {
	response = &ContinueDeployApplicationGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
