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

// OpenApiCallback invokes the mpaas.OpenApiCallback API synchronously
func (client *Client) OpenApiCallback(request *OpenApiCallbackRequest) (response *OpenApiCallbackResponse, err error) {
	response = CreateOpenApiCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// OpenApiCallbackWithChan invokes the mpaas.OpenApiCallback API asynchronously
func (client *Client) OpenApiCallbackWithChan(request *OpenApiCallbackRequest) (<-chan *OpenApiCallbackResponse, <-chan error) {
	responseChan := make(chan *OpenApiCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenApiCallback(request)
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

// OpenApiCallbackWithCallback invokes the mpaas.OpenApiCallback API asynchronously
func (client *Client) OpenApiCallbackWithCallback(request *OpenApiCallbackRequest, callback func(response *OpenApiCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenApiCallbackResponse
		var err error
		defer close(result)
		response, err = client.OpenApiCallback(request)
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

// OpenApiCallbackRequest is the request struct for api OpenApiCallback
type OpenApiCallbackRequest struct {
	*requests.RpcRequest
	MpaasMqcpOpenApiCallbackRequestJsonStr string `position:"Body" name:"MpaasMqcpOpenApiCallbackRequestJsonStr"`
	TenantId                               string `position:"Body" name:"TenantId"`
	AppId                                  string `position:"Body" name:"AppId"`
	WorkspaceId                            string `position:"Body" name:"WorkspaceId"`
}

// OpenApiCallbackResponse is the response struct for api OpenApiCallback
type OpenApiCallbackResponse struct {
	*responses.BaseResponse
	ResultCode    string `json:"ResultCode" xml:"ResultCode"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ResultContent string `json:"ResultContent" xml:"ResultContent"`
}

// CreateOpenApiCallbackRequest creates a request to invoke OpenApiCallback API
func CreateOpenApiCallbackRequest() (request *OpenApiCallbackRequest) {
	request = &OpenApiCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "OpenApiCallback", "", "")
	request.Method = requests.POST
	return
}

// CreateOpenApiCallbackResponse creates a response to parse from OpenApiCallback response
func CreateOpenApiCallbackResponse() (response *OpenApiCallbackResponse) {
	response = &OpenApiCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
