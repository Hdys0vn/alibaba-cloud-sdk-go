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

// GetLogUrlInMsa invokes the mpaas.GetLogUrlInMsa API synchronously
func (client *Client) GetLogUrlInMsa(request *GetLogUrlInMsaRequest) (response *GetLogUrlInMsaResponse, err error) {
	response = CreateGetLogUrlInMsaResponse()
	err = client.DoAction(request, response)
	return
}

// GetLogUrlInMsaWithChan invokes the mpaas.GetLogUrlInMsa API asynchronously
func (client *Client) GetLogUrlInMsaWithChan(request *GetLogUrlInMsaRequest) (<-chan *GetLogUrlInMsaResponse, <-chan error) {
	responseChan := make(chan *GetLogUrlInMsaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLogUrlInMsa(request)
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

// GetLogUrlInMsaWithCallback invokes the mpaas.GetLogUrlInMsa API asynchronously
func (client *Client) GetLogUrlInMsaWithCallback(request *GetLogUrlInMsaRequest, callback func(response *GetLogUrlInMsaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLogUrlInMsaResponse
		var err error
		defer close(result)
		response, err = client.GetLogUrlInMsa(request)
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

// GetLogUrlInMsaRequest is the request struct for api GetLogUrlInMsa
type GetLogUrlInMsaRequest struct {
	*requests.RpcRequest
	TenantId    string           `position:"Body" name:"TenantId"`
	Id          requests.Integer `position:"Body" name:"Id"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// GetLogUrlInMsaResponse is the response struct for api GetLogUrlInMsa
type GetLogUrlInMsaResponse struct {
	*responses.BaseResponse
	ResultMessage string                        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                        `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInGetLogUrlInMsa `json:"ResultContent" xml:"ResultContent"`
}

// CreateGetLogUrlInMsaRequest creates a request to invoke GetLogUrlInMsa API
func CreateGetLogUrlInMsaRequest() (request *GetLogUrlInMsaRequest) {
	request = &GetLogUrlInMsaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "GetLogUrlInMsa", "", "")
	request.Method = requests.POST
	return
}

// CreateGetLogUrlInMsaResponse creates a response to parse from GetLogUrlInMsa response
func CreateGetLogUrlInMsaResponse() (response *GetLogUrlInMsaResponse) {
	response = &GetLogUrlInMsaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
