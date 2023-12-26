package mpaas

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

// UploadApp invokes the mpaas.UploadApp API synchronously
func (client *Client) UploadApp(request *UploadAppRequest) (response *UploadAppResponse, err error) {
	response = CreateUploadAppResponse()
	err = client.DoAction(request, response)
	return
}

// UploadAppWithChan invokes the mpaas.UploadApp API asynchronously
func (client *Client) UploadAppWithChan(request *UploadAppRequest) (<-chan *UploadAppResponse, <-chan error) {
	responseChan := make(chan *UploadAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadApp(request)
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

// UploadAppWithCallback invokes the mpaas.UploadApp API asynchronously
func (client *Client) UploadAppWithCallback(request *UploadAppRequest, callback func(response *UploadAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadAppResponse
		var err error
		defer close(result)
		response, err = client.UploadApp(request)
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

// UploadAppRequest is the request struct for api UploadApp
type UploadAppRequest struct {
	*requests.RpcRequest
	AppFileUrl  string `position:"Body" name:"AppFileUrl"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// UploadAppResponse is the response struct for api UploadApp
type UploadAppResponse struct {
	*responses.BaseResponse
}

// CreateUploadAppRequest creates a request to invoke UploadApp API
func CreateUploadAppRequest() (request *UploadAppRequest) {
	request = &UploadAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "UploadApp", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadAppResponse creates a response to parse from UploadApp response
func CreateUploadAppResponse() (response *UploadAppResponse) {
	response = &UploadAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
