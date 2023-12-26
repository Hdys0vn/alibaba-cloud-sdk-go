package companyreg

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

// TransferIntentionOwner invokes the companyreg.TransferIntentionOwner API synchronously
func (client *Client) TransferIntentionOwner(request *TransferIntentionOwnerRequest) (response *TransferIntentionOwnerResponse, err error) {
	response = CreateTransferIntentionOwnerResponse()
	err = client.DoAction(request, response)
	return
}

// TransferIntentionOwnerWithChan invokes the companyreg.TransferIntentionOwner API asynchronously
func (client *Client) TransferIntentionOwnerWithChan(request *TransferIntentionOwnerRequest) (<-chan *TransferIntentionOwnerResponse, <-chan error) {
	responseChan := make(chan *TransferIntentionOwnerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferIntentionOwner(request)
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

// TransferIntentionOwnerWithCallback invokes the companyreg.TransferIntentionOwner API asynchronously
func (client *Client) TransferIntentionOwnerWithCallback(request *TransferIntentionOwnerRequest, callback func(response *TransferIntentionOwnerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferIntentionOwnerResponse
		var err error
		defer close(result)
		response, err = client.TransferIntentionOwner(request)
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

// TransferIntentionOwnerRequest is the request struct for api TransferIntentionOwner
type TransferIntentionOwnerRequest struct {
	*requests.RpcRequest
	BizType  string           `position:"Query" name:"BizType"`
	BizId    string           `position:"Query" name:"BizId"`
	PersonId requests.Integer `position:"Query" name:"PersonId"`
	Remark   string           `position:"Query" name:"Remark"`
}

// TransferIntentionOwnerResponse is the response struct for api TransferIntentionOwner
type TransferIntentionOwnerResponse struct {
	*responses.BaseResponse
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateTransferIntentionOwnerRequest creates a request to invoke TransferIntentionOwner API
func CreateTransferIntentionOwnerRequest() (request *TransferIntentionOwnerRequest) {
	request = &TransferIntentionOwnerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "TransferIntentionOwner", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTransferIntentionOwnerResponse creates a response to parse from TransferIntentionOwner response
func CreateTransferIntentionOwnerResponse() (response *TransferIntentionOwnerResponse) {
	response = &TransferIntentionOwnerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
