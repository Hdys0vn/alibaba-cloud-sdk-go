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

// OpenApiUpdateActiveScene invokes the mpaas.OpenApiUpdateActiveScene API synchronously
func (client *Client) OpenApiUpdateActiveScene(request *OpenApiUpdateActiveSceneRequest) (response *OpenApiUpdateActiveSceneResponse, err error) {
	response = CreateOpenApiUpdateActiveSceneResponse()
	err = client.DoAction(request, response)
	return
}

// OpenApiUpdateActiveSceneWithChan invokes the mpaas.OpenApiUpdateActiveScene API asynchronously
func (client *Client) OpenApiUpdateActiveSceneWithChan(request *OpenApiUpdateActiveSceneRequest) (<-chan *OpenApiUpdateActiveSceneResponse, <-chan error) {
	responseChan := make(chan *OpenApiUpdateActiveSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenApiUpdateActiveScene(request)
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

// OpenApiUpdateActiveSceneWithCallback invokes the mpaas.OpenApiUpdateActiveScene API asynchronously
func (client *Client) OpenApiUpdateActiveSceneWithCallback(request *OpenApiUpdateActiveSceneRequest, callback func(response *OpenApiUpdateActiveSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenApiUpdateActiveSceneResponse
		var err error
		defer close(result)
		response, err = client.OpenApiUpdateActiveScene(request)
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

// OpenApiUpdateActiveSceneRequest is the request struct for api OpenApiUpdateActiveScene
type OpenApiUpdateActiveSceneRequest struct {
	*requests.RpcRequest
	MpaasMqcpOpenApiUpdateActiveSceneReqJsonStr string `position:"Body" name:"MpaasMqcpOpenApiUpdateActiveSceneReqJsonStr"`
	TenantId                                    string `position:"Body" name:"TenantId"`
	AppId                                       string `position:"Body" name:"AppId"`
	WorkspaceId                                 string `position:"Body" name:"WorkspaceId"`
}

// OpenApiUpdateActiveSceneResponse is the response struct for api OpenApiUpdateActiveScene
type OpenApiUpdateActiveSceneResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ResultCode    string `json:"ResultCode" xml:"ResultCode"`
	ResultContent string `json:"ResultContent" xml:"ResultContent"`
}

// CreateOpenApiUpdateActiveSceneRequest creates a request to invoke OpenApiUpdateActiveScene API
func CreateOpenApiUpdateActiveSceneRequest() (request *OpenApiUpdateActiveSceneRequest) {
	request = &OpenApiUpdateActiveSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "OpenApiUpdateActiveScene", "", "")
	request.Method = requests.POST
	return
}

// CreateOpenApiUpdateActiveSceneResponse creates a response to parse from OpenApiUpdateActiveScene response
func CreateOpenApiUpdateActiveSceneResponse() (response *OpenApiUpdateActiveSceneResponse) {
	response = &OpenApiUpdateActiveSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
