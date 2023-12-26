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

// GetMcubeFileToken invokes the mpaas.GetMcubeFileToken API synchronously
func (client *Client) GetMcubeFileToken(request *GetMcubeFileTokenRequest) (response *GetMcubeFileTokenResponse, err error) {
	response = CreateGetMcubeFileTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetMcubeFileTokenWithChan invokes the mpaas.GetMcubeFileToken API asynchronously
func (client *Client) GetMcubeFileTokenWithChan(request *GetMcubeFileTokenRequest) (<-chan *GetMcubeFileTokenResponse, <-chan error) {
	responseChan := make(chan *GetMcubeFileTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMcubeFileToken(request)
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

// GetMcubeFileTokenWithCallback invokes the mpaas.GetMcubeFileToken API asynchronously
func (client *Client) GetMcubeFileTokenWithCallback(request *GetMcubeFileTokenRequest, callback func(response *GetMcubeFileTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMcubeFileTokenResponse
		var err error
		defer close(result)
		response, err = client.GetMcubeFileToken(request)
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

// GetMcubeFileTokenRequest is the request struct for api GetMcubeFileToken
type GetMcubeFileTokenRequest struct {
	*requests.RpcRequest
	OnexFlag    requests.Boolean `position:"Body" name:"OnexFlag"`
	TenantId    string           `position:"Body" name:"TenantId"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// GetMcubeFileTokenResponse is the response struct for api GetMcubeFileToken
type GetMcubeFileTokenResponse struct {
	*responses.BaseResponse
	ResultMessage      string             `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode         string             `json:"ResultCode" xml:"ResultCode"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	GetFileTokenResult GetFileTokenResult `json:"GetFileTokenResult" xml:"GetFileTokenResult"`
}

// CreateGetMcubeFileTokenRequest creates a request to invoke GetMcubeFileToken API
func CreateGetMcubeFileTokenRequest() (request *GetMcubeFileTokenRequest) {
	request = &GetMcubeFileTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "GetMcubeFileToken", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMcubeFileTokenResponse creates a response to parse from GetMcubeFileToken response
func CreateGetMcubeFileTokenResponse() (response *GetMcubeFileTokenResponse) {
	response = &GetMcubeFileTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}