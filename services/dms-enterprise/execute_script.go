package dms_enterprise

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

// ExecuteScript invokes the dms_enterprise.ExecuteScript API synchronously
func (client *Client) ExecuteScript(request *ExecuteScriptRequest) (response *ExecuteScriptResponse, err error) {
	response = CreateExecuteScriptResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteScriptWithChan invokes the dms_enterprise.ExecuteScript API asynchronously
func (client *Client) ExecuteScriptWithChan(request *ExecuteScriptRequest) (<-chan *ExecuteScriptResponse, <-chan error) {
	responseChan := make(chan *ExecuteScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteScript(request)
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

// ExecuteScriptWithCallback invokes the dms_enterprise.ExecuteScript API asynchronously
func (client *Client) ExecuteScriptWithCallback(request *ExecuteScriptRequest, callback func(response *ExecuteScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteScriptResponse
		var err error
		defer close(result)
		response, err = client.ExecuteScript(request)
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

// ExecuteScriptRequest is the request struct for api ExecuteScript
type ExecuteScriptRequest struct {
	*requests.RpcRequest
	Tid    requests.Integer `position:"Query" name:"Tid"`
	Script string           `position:"Query" name:"Script"`
	DbId   requests.Integer `position:"Query" name:"DbId"`
	Logic  requests.Boolean `position:"Query" name:"Logic"`
}

// ExecuteScriptResponse is the response struct for api ExecuteScript
type ExecuteScriptResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	ErrorCode    string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool     `json:"Success" xml:"Success"`
	Results      []Result `json:"Results" xml:"Results"`
}

// CreateExecuteScriptRequest creates a request to invoke ExecuteScript API
func CreateExecuteScriptRequest() (request *ExecuteScriptRequest) {
	request = &ExecuteScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ExecuteScript", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExecuteScriptResponse creates a response to parse from ExecuteScript response
func CreateExecuteScriptResponse() (response *ExecuteScriptResponse) {
	response = &ExecuteScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}