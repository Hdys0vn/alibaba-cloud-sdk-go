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

// DownloadMPCDExcelById invokes the mpaas.DownloadMPCDExcelById API synchronously
func (client *Client) DownloadMPCDExcelById(request *DownloadMPCDExcelByIdRequest) (response *DownloadMPCDExcelByIdResponse, err error) {
	response = CreateDownloadMPCDExcelByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadMPCDExcelByIdWithChan invokes the mpaas.DownloadMPCDExcelById API asynchronously
func (client *Client) DownloadMPCDExcelByIdWithChan(request *DownloadMPCDExcelByIdRequest) (<-chan *DownloadMPCDExcelByIdResponse, <-chan error) {
	responseChan := make(chan *DownloadMPCDExcelByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadMPCDExcelById(request)
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

// DownloadMPCDExcelByIdWithCallback invokes the mpaas.DownloadMPCDExcelById API asynchronously
func (client *Client) DownloadMPCDExcelByIdWithCallback(request *DownloadMPCDExcelByIdRequest, callback func(response *DownloadMPCDExcelByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadMPCDExcelByIdResponse
		var err error
		defer close(result)
		response, err = client.DownloadMPCDExcelById(request)
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

// DownloadMPCDExcelByIdRequest is the request struct for api DownloadMPCDExcelById
type DownloadMPCDExcelByIdRequest struct {
	*requests.RpcRequest
	AppId       string           `position:"Body" name:"AppId"`
	TenantId    string           `position:"Body" name:"TenantId"`
	ScanTaskId  requests.Integer `position:"Body" name:"ScanTaskId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// DownloadMPCDExcelByIdResponse is the response struct for api DownloadMPCDExcelById
type DownloadMPCDExcelByIdResponse struct {
	*responses.BaseResponse
}

// CreateDownloadMPCDExcelByIdRequest creates a request to invoke DownloadMPCDExcelById API
func CreateDownloadMPCDExcelByIdRequest() (request *DownloadMPCDExcelByIdRequest) {
	request = &DownloadMPCDExcelByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DownloadMPCDExcelById", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadMPCDExcelByIdResponse creates a response to parse from DownloadMPCDExcelById response
func CreateDownloadMPCDExcelByIdResponse() (response *DownloadMPCDExcelByIdResponse) {
	response = &DownloadMPCDExcelByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
