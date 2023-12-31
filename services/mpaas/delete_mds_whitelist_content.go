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

// DeleteMdsWhitelistContent invokes the mpaas.DeleteMdsWhitelistContent API synchronously
func (client *Client) DeleteMdsWhitelistContent(request *DeleteMdsWhitelistContentRequest) (response *DeleteMdsWhitelistContentResponse, err error) {
	response = CreateDeleteMdsWhitelistContentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMdsWhitelistContentWithChan invokes the mpaas.DeleteMdsWhitelistContent API asynchronously
func (client *Client) DeleteMdsWhitelistContentWithChan(request *DeleteMdsWhitelistContentRequest) (<-chan *DeleteMdsWhitelistContentResponse, <-chan error) {
	responseChan := make(chan *DeleteMdsWhitelistContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMdsWhitelistContent(request)
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

// DeleteMdsWhitelistContentWithCallback invokes the mpaas.DeleteMdsWhitelistContent API asynchronously
func (client *Client) DeleteMdsWhitelistContentWithCallback(request *DeleteMdsWhitelistContentRequest, callback func(response *DeleteMdsWhitelistContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMdsWhitelistContentResponse
		var err error
		defer close(result)
		response, err = client.DeleteMdsWhitelistContent(request)
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

// DeleteMdsWhitelistContentRequest is the request struct for api DeleteMdsWhitelistContent
type DeleteMdsWhitelistContentRequest struct {
	*requests.RpcRequest
	TenantId       string `position:"Body" name:"TenantId"`
	WhitelistValue string `position:"Body" name:"WhitelistValue"`
	AppId          string `position:"Body" name:"AppId"`
	WorkspaceId    string `position:"Body" name:"WorkspaceId"`
	WhitelistId    string `position:"Body" name:"WhitelistId"`
}

// DeleteMdsWhitelistContentResponse is the response struct for api DeleteMdsWhitelistContent
type DeleteMdsWhitelistContentResponse struct {
	*responses.BaseResponse
	ResultMessage string        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContent `json:"ResultContent" xml:"ResultContent"`
}

// CreateDeleteMdsWhitelistContentRequest creates a request to invoke DeleteMdsWhitelistContent API
func CreateDeleteMdsWhitelistContentRequest() (request *DeleteMdsWhitelistContentRequest) {
	request = &DeleteMdsWhitelistContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteMdsWhitelistContent", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMdsWhitelistContentResponse creates a response to parse from DeleteMdsWhitelistContent response
func CreateDeleteMdsWhitelistContentResponse() (response *DeleteMdsWhitelistContentResponse) {
	response = &DeleteMdsWhitelistContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
