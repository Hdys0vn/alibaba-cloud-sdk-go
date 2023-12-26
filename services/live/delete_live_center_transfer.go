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

// DeleteLiveCenterTransfer invokes the live.DeleteLiveCenterTransfer API synchronously
func (client *Client) DeleteLiveCenterTransfer(request *DeleteLiveCenterTransferRequest) (response *DeleteLiveCenterTransferResponse, err error) {
	response = CreateDeleteLiveCenterTransferResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveCenterTransferWithChan invokes the live.DeleteLiveCenterTransfer API asynchronously
func (client *Client) DeleteLiveCenterTransferWithChan(request *DeleteLiveCenterTransferRequest) (<-chan *DeleteLiveCenterTransferResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveCenterTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveCenterTransfer(request)
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

// DeleteLiveCenterTransferWithCallback invokes the live.DeleteLiveCenterTransfer API asynchronously
func (client *Client) DeleteLiveCenterTransferWithCallback(request *DeleteLiveCenterTransferRequest, callback func(response *DeleteLiveCenterTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveCenterTransferResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveCenterTransfer(request)
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

// DeleteLiveCenterTransferRequest is the request struct for api DeleteLiveCenterTransfer
type DeleteLiveCenterTransferRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DstUrl     string           `position:"Query" name:"DstUrl"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveCenterTransferResponse is the response struct for api DeleteLiveCenterTransfer
type DeleteLiveCenterTransferResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveCenterTransferRequest creates a request to invoke DeleteLiveCenterTransfer API
func CreateDeleteLiveCenterTransferRequest() (request *DeleteLiveCenterTransferRequest) {
	request = &DeleteLiveCenterTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveCenterTransfer", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveCenterTransferResponse creates a response to parse from DeleteLiveCenterTransfer response
func CreateDeleteLiveCenterTransferResponse() (response *DeleteLiveCenterTransferResponse) {
	response = &DeleteLiveCenterTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
