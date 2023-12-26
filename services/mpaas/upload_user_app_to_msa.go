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

// UploadUserAppToMsa invokes the mpaas.UploadUserAppToMsa API synchronously
func (client *Client) UploadUserAppToMsa(request *UploadUserAppToMsaRequest) (response *UploadUserAppToMsaResponse, err error) {
	response = CreateUploadUserAppToMsaResponse()
	err = client.DoAction(request, response)
	return
}

// UploadUserAppToMsaWithChan invokes the mpaas.UploadUserAppToMsa API asynchronously
func (client *Client) UploadUserAppToMsaWithChan(request *UploadUserAppToMsaRequest) (<-chan *UploadUserAppToMsaResponse, <-chan error) {
	responseChan := make(chan *UploadUserAppToMsaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadUserAppToMsa(request)
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

// UploadUserAppToMsaWithCallback invokes the mpaas.UploadUserAppToMsa API asynchronously
func (client *Client) UploadUserAppToMsaWithCallback(request *UploadUserAppToMsaRequest, callback func(response *UploadUserAppToMsaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadUserAppToMsaResponse
		var err error
		defer close(result)
		response, err = client.UploadUserAppToMsa(request)
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

// UploadUserAppToMsaRequest is the request struct for api UploadUserAppToMsa
type UploadUserAppToMsaRequest struct {
	*requests.RpcRequest
	TenantId    string `position:"Body" name:"TenantId"`
	AppId       string `position:"Body" name:"AppId"`
	FileUrl     string `position:"Body" name:"FileUrl"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// UploadUserAppToMsaResponse is the response struct for api UploadUserAppToMsa
type UploadUserAppToMsaResponse struct {
	*responses.BaseResponse
	ResultMessage string                            `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                            `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                            `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInUploadUserAppToMsa `json:"ResultContent" xml:"ResultContent"`
}

// CreateUploadUserAppToMsaRequest creates a request to invoke UploadUserAppToMsa API
func CreateUploadUserAppToMsaRequest() (request *UploadUserAppToMsaRequest) {
	request = &UploadUserAppToMsaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "UploadUserAppToMsa", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadUserAppToMsaResponse creates a response to parse from UploadUserAppToMsa response
func CreateUploadUserAppToMsaResponse() (response *UploadUserAppToMsaResponse) {
	response = &UploadUserAppToMsaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
