package scsp

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

// TransferCallToPhone invokes the scsp.TransferCallToPhone API synchronously
func (client *Client) TransferCallToPhone(request *TransferCallToPhoneRequest) (response *TransferCallToPhoneResponse, err error) {
	response = CreateTransferCallToPhoneResponse()
	err = client.DoAction(request, response)
	return
}

// TransferCallToPhoneWithChan invokes the scsp.TransferCallToPhone API asynchronously
func (client *Client) TransferCallToPhoneWithChan(request *TransferCallToPhoneRequest) (<-chan *TransferCallToPhoneResponse, <-chan error) {
	responseChan := make(chan *TransferCallToPhoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferCallToPhone(request)
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

// TransferCallToPhoneWithCallback invokes the scsp.TransferCallToPhone API asynchronously
func (client *Client) TransferCallToPhoneWithCallback(request *TransferCallToPhoneRequest, callback func(response *TransferCallToPhoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferCallToPhoneResponse
		var err error
		defer close(result)
		response, err = client.TransferCallToPhone(request)
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

// TransferCallToPhoneRequest is the request struct for api TransferCallToPhone
type TransferCallToPhoneRequest struct {
	*requests.RpcRequest
	ClientToken      string           `position:"Body"`
	InstanceId       string           `position:"Body"`
	AccountName      string           `position:"Body"`
	Caller           string           `position:"Body"`
	Callee           string           `position:"Body"`
	CallId           string           `position:"Body"`
	JobId            string           `position:"Body"`
	ConnectionId     string           `position:"Body"`
	HoldConnectionId string           `position:"Body"`
	Type             requests.Integer `position:"Body"`
	IsSingleTransfer requests.Boolean `position:"Body"`
	CallerPhone      string           `position:"Body"`
	CalleePhone      string           `position:"Body"`
}

// TransferCallToPhoneResponse is the response struct for api TransferCallToPhone
type TransferCallToPhoneResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateTransferCallToPhoneRequest creates a request to invoke TransferCallToPhone API
func CreateTransferCallToPhoneRequest() (request *TransferCallToPhoneRequest) {
	request = &TransferCallToPhoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "TransferCallToPhone", "", "")
	request.Method = requests.POST
	return
}

// CreateTransferCallToPhoneResponse creates a response to parse from TransferCallToPhone response
func CreateTransferCallToPhoneResponse() (response *TransferCallToPhoneResponse) {
	response = &TransferCallToPhoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
