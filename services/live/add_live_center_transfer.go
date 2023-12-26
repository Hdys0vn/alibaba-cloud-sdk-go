package live

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

// AddLiveCenterTransfer invokes the live.AddLiveCenterTransfer API synchronously
func (client *Client) AddLiveCenterTransfer(request *AddLiveCenterTransferRequest) (response *AddLiveCenterTransferResponse, err error) {
	response = CreateAddLiveCenterTransferResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveCenterTransferWithChan invokes the live.AddLiveCenterTransfer API asynchronously
func (client *Client) AddLiveCenterTransferWithChan(request *AddLiveCenterTransferRequest) (<-chan *AddLiveCenterTransferResponse, <-chan error) {
	responseChan := make(chan *AddLiveCenterTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveCenterTransfer(request)
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

// AddLiveCenterTransferWithCallback invokes the live.AddLiveCenterTransfer API asynchronously
func (client *Client) AddLiveCenterTransferWithCallback(request *AddLiveCenterTransferRequest, callback func(response *AddLiveCenterTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveCenterTransferResponse
		var err error
		defer close(result)
		response, err = client.AddLiveCenterTransfer(request)
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

// AddLiveCenterTransferRequest is the request struct for api AddLiveCenterTransfer
type AddLiveCenterTransferRequest struct {
	*requests.RpcRequest
	TransferArgs string           `position:"Query" name:"TransferArgs"`
	StartTime    string           `position:"Query" name:"StartTime"`
	AppName      string           `position:"Query" name:"AppName"`
	StreamName   string           `position:"Query" name:"StreamName"`
	DstUrl       string           `position:"Query" name:"DstUrl"`
	DomainName   string           `position:"Query" name:"DomainName"`
	EndTime      string           `position:"Query" name:"EndTime"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// AddLiveCenterTransferResponse is the response struct for api AddLiveCenterTransfer
type AddLiveCenterTransferResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveCenterTransferRequest creates a request to invoke AddLiveCenterTransfer API
func CreateAddLiveCenterTransferRequest() (request *AddLiveCenterTransferRequest) {
	request = &AddLiveCenterTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLiveCenterTransfer", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLiveCenterTransferResponse creates a response to parse from AddLiveCenterTransfer response
func CreateAddLiveCenterTransferResponse() (response *AddLiveCenterTransferResponse) {
	response = &AddLiveCenterTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}