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

// DownloadMPCDById invokes the mpaas.DownloadMPCDById API synchronously
func (client *Client) DownloadMPCDById(request *DownloadMPCDByIdRequest) (response *DownloadMPCDByIdResponse, err error) {
	response = CreateDownloadMPCDByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadMPCDByIdWithChan invokes the mpaas.DownloadMPCDById API asynchronously
func (client *Client) DownloadMPCDByIdWithChan(request *DownloadMPCDByIdRequest) (<-chan *DownloadMPCDByIdResponse, <-chan error) {
	responseChan := make(chan *DownloadMPCDByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadMPCDById(request)
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

// DownloadMPCDByIdWithCallback invokes the mpaas.DownloadMPCDById API asynchronously
func (client *Client) DownloadMPCDByIdWithCallback(request *DownloadMPCDByIdRequest, callback func(response *DownloadMPCDByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadMPCDByIdResponse
		var err error
		defer close(result)
		response, err = client.DownloadMPCDById(request)
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

// DownloadMPCDByIdRequest is the request struct for api DownloadMPCDById
type DownloadMPCDByIdRequest struct {
	*requests.RpcRequest
	AppId       string           `position:"Body" name:"AppId"`
	TenantId    string           `position:"Body" name:"TenantId"`
	ScanTaskId  requests.Integer `position:"Body" name:"ScanTaskId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// DownloadMPCDByIdResponse is the response struct for api DownloadMPCDById
type DownloadMPCDByIdResponse struct {
	*responses.BaseResponse
}

// CreateDownloadMPCDByIdRequest creates a request to invoke DownloadMPCDById API
func CreateDownloadMPCDByIdRequest() (request *DownloadMPCDByIdRequest) {
	request = &DownloadMPCDByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DownloadMPCDById", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadMPCDByIdResponse creates a response to parse from DownloadMPCDById response
func CreateDownloadMPCDByIdResponse() (response *DownloadMPCDByIdResponse) {
	response = &DownloadMPCDByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}