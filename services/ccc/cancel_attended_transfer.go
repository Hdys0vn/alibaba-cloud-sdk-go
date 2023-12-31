package ccc

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

// CancelAttendedTransfer invokes the ccc.CancelAttendedTransfer API synchronously
func (client *Client) CancelAttendedTransfer(request *CancelAttendedTransferRequest) (response *CancelAttendedTransferResponse, err error) {
	response = CreateCancelAttendedTransferResponse()
	err = client.DoAction(request, response)
	return
}

// CancelAttendedTransferWithChan invokes the ccc.CancelAttendedTransfer API asynchronously
func (client *Client) CancelAttendedTransferWithChan(request *CancelAttendedTransferRequest) (<-chan *CancelAttendedTransferResponse, <-chan error) {
	responseChan := make(chan *CancelAttendedTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelAttendedTransfer(request)
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

// CancelAttendedTransferWithCallback invokes the ccc.CancelAttendedTransfer API asynchronously
func (client *Client) CancelAttendedTransferWithCallback(request *CancelAttendedTransferRequest, callback func(response *CancelAttendedTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelAttendedTransferResponse
		var err error
		defer close(result)
		response, err = client.CancelAttendedTransfer(request)
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

// CancelAttendedTransferRequest is the request struct for api CancelAttendedTransfer
type CancelAttendedTransferRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	DeviceId   string `position:"Query" name:"DeviceId"`
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// CancelAttendedTransferResponse is the response struct for api CancelAttendedTransfer
type CancelAttendedTransferResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateCancelAttendedTransferRequest creates a request to invoke CancelAttendedTransfer API
func CreateCancelAttendedTransferRequest() (request *CancelAttendedTransferRequest) {
	request = &CancelAttendedTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "CancelAttendedTransfer", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelAttendedTransferResponse creates a response to parse from CancelAttendedTransfer response
func CreateCancelAttendedTransferResponse() (response *CancelAttendedTransferResponse) {
	response = &CancelAttendedTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
