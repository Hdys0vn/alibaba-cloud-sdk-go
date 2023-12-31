package rdc

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

// SyncUserToRdc invokes the rdc.SyncUserToRdc API synchronously
// api document: https://help.aliyun.com/api/rdc/syncusertordc.html
func (client *Client) SyncUserToRdc(request *SyncUserToRdcRequest) (response *SyncUserToRdcResponse, err error) {
	response = CreateSyncUserToRdcResponse()
	err = client.DoAction(request, response)
	return
}

// SyncUserToRdcWithChan invokes the rdc.SyncUserToRdc API asynchronously
// api document: https://help.aliyun.com/api/rdc/syncusertordc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SyncUserToRdcWithChan(request *SyncUserToRdcRequest) (<-chan *SyncUserToRdcResponse, <-chan error) {
	responseChan := make(chan *SyncUserToRdcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncUserToRdc(request)
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

// SyncUserToRdcWithCallback invokes the rdc.SyncUserToRdc API asynchronously
// api document: https://help.aliyun.com/api/rdc/syncusertordc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SyncUserToRdcWithCallback(request *SyncUserToRdcRequest, callback func(response *SyncUserToRdcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncUserToRdcResponse
		var err error
		defer close(result)
		response, err = client.SyncUserToRdc(request)
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

// SyncUserToRdcRequest is the request struct for api SyncUserToRdc
type SyncUserToRdcRequest struct {
	*requests.RpcRequest
	LoginTicket string `position:"Body" name:"LoginTicket"`
}

// SyncUserToRdcResponse is the response struct for api SyncUserToRdc
type SyncUserToRdcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSyncUserToRdcRequest creates a request to invoke SyncUserToRdc API
func CreateSyncUserToRdcRequest() (request *SyncUserToRdcRequest) {
	request = &SyncUserToRdcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "SyncUserToRdc", "rdc", "openAPI")
	return
}

// CreateSyncUserToRdcResponse creates a response to parse from SyncUserToRdc response
func CreateSyncUserToRdcResponse() (response *SyncUserToRdcResponse) {
	response = &SyncUserToRdcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
