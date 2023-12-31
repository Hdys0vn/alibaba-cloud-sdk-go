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

// DeleteMcubeNebulaApp invokes the mpaas.DeleteMcubeNebulaApp API synchronously
func (client *Client) DeleteMcubeNebulaApp(request *DeleteMcubeNebulaAppRequest) (response *DeleteMcubeNebulaAppResponse, err error) {
	response = CreateDeleteMcubeNebulaAppResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMcubeNebulaAppWithChan invokes the mpaas.DeleteMcubeNebulaApp API asynchronously
func (client *Client) DeleteMcubeNebulaAppWithChan(request *DeleteMcubeNebulaAppRequest) (<-chan *DeleteMcubeNebulaAppResponse, <-chan error) {
	responseChan := make(chan *DeleteMcubeNebulaAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMcubeNebulaApp(request)
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

// DeleteMcubeNebulaAppWithCallback invokes the mpaas.DeleteMcubeNebulaApp API asynchronously
func (client *Client) DeleteMcubeNebulaAppWithCallback(request *DeleteMcubeNebulaAppRequest, callback func(response *DeleteMcubeNebulaAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMcubeNebulaAppResponse
		var err error
		defer close(result)
		response, err = client.DeleteMcubeNebulaApp(request)
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

// DeleteMcubeNebulaAppRequest is the request struct for api DeleteMcubeNebulaApp
type DeleteMcubeNebulaAppRequest struct {
	*requests.RpcRequest
	H5Id        string `position:"Body" name:"H5Id"`
	TenantId    string `position:"Body" name:"TenantId"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// DeleteMcubeNebulaAppResponse is the response struct for api DeleteMcubeNebulaApp
type DeleteMcubeNebulaAppResponse struct {
	*responses.BaseResponse
	ResultMessage              string                     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode                 string                     `json:"ResultCode" xml:"ResultCode"`
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	DeleteMcubeNebulaAppResult DeleteMcubeNebulaAppResult `json:"DeleteMcubeNebulaAppResult" xml:"DeleteMcubeNebulaAppResult"`
}

// CreateDeleteMcubeNebulaAppRequest creates a request to invoke DeleteMcubeNebulaApp API
func CreateDeleteMcubeNebulaAppRequest() (request *DeleteMcubeNebulaAppRequest) {
	request = &DeleteMcubeNebulaAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteMcubeNebulaApp", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMcubeNebulaAppResponse creates a response to parse from DeleteMcubeNebulaApp response
func CreateDeleteMcubeNebulaAppResponse() (response *DeleteMcubeNebulaAppResponse) {
	response = &DeleteMcubeNebulaAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
