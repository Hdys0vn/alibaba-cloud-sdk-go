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

// DeleteCubecardWhitelistContent invokes the mpaas.DeleteCubecardWhitelistContent API synchronously
func (client *Client) DeleteCubecardWhitelistContent(request *DeleteCubecardWhitelistContentRequest) (response *DeleteCubecardWhitelistContentResponse, err error) {
	response = CreateDeleteCubecardWhitelistContentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCubecardWhitelistContentWithChan invokes the mpaas.DeleteCubecardWhitelistContent API asynchronously
func (client *Client) DeleteCubecardWhitelistContentWithChan(request *DeleteCubecardWhitelistContentRequest) (<-chan *DeleteCubecardWhitelistContentResponse, <-chan error) {
	responseChan := make(chan *DeleteCubecardWhitelistContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCubecardWhitelistContent(request)
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

// DeleteCubecardWhitelistContentWithCallback invokes the mpaas.DeleteCubecardWhitelistContent API asynchronously
func (client *Client) DeleteCubecardWhitelistContentWithCallback(request *DeleteCubecardWhitelistContentRequest, callback func(response *DeleteCubecardWhitelistContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCubecardWhitelistContentResponse
		var err error
		defer close(result)
		response, err = client.DeleteCubecardWhitelistContent(request)
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

// DeleteCubecardWhitelistContentRequest is the request struct for api DeleteCubecardWhitelistContent
type DeleteCubecardWhitelistContentRequest struct {
	*requests.RpcRequest
	TenantId       string `position:"Body" name:"TenantId"`
	WhitelistValue string `position:"Body" name:"WhitelistValue"`
	AppId          string `position:"Body" name:"AppId"`
	WorkspaceId    string `position:"Body" name:"WorkspaceId"`
	WhitelistId    string `position:"Body" name:"WhitelistId"`
}

// DeleteCubecardWhitelistContentResponse is the response struct for api DeleteCubecardWhitelistContent
type DeleteCubecardWhitelistContentResponse struct {
	*responses.BaseResponse
	ResultMessage string        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContent `json:"ResultContent" xml:"ResultContent"`
}

// CreateDeleteCubecardWhitelistContentRequest creates a request to invoke DeleteCubecardWhitelistContent API
func CreateDeleteCubecardWhitelistContentRequest() (request *DeleteCubecardWhitelistContentRequest) {
	request = &DeleteCubecardWhitelistContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteCubecardWhitelistContent", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCubecardWhitelistContentResponse creates a response to parse from DeleteCubecardWhitelistContent response
func CreateDeleteCubecardWhitelistContentResponse() (response *DeleteCubecardWhitelistContentResponse) {
	response = &DeleteCubecardWhitelistContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}