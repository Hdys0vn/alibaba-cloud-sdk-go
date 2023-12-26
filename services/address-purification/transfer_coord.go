package address_purification

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

// TransferCoord invokes the address_purification.TransferCoord API synchronously
func (client *Client) TransferCoord(request *TransferCoordRequest) (response *TransferCoordResponse, err error) {
	response = CreateTransferCoordResponse()
	err = client.DoAction(request, response)
	return
}

// TransferCoordWithChan invokes the address_purification.TransferCoord API asynchronously
func (client *Client) TransferCoordWithChan(request *TransferCoordRequest) (<-chan *TransferCoordResponse, <-chan error) {
	responseChan := make(chan *TransferCoordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferCoord(request)
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

// TransferCoordWithCallback invokes the address_purification.TransferCoord API asynchronously
func (client *Client) TransferCoordWithCallback(request *TransferCoordRequest, callback func(response *TransferCoordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferCoordResponse
		var err error
		defer close(result)
		response, err = client.TransferCoord(request)
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

// TransferCoordRequest is the request struct for api TransferCoord
type TransferCoordRequest struct {
	*requests.RpcRequest
	DefaultProvince string `position:"Body" name:"DefaultProvince"`
	SrcCoord        string `position:"Body" name:"SrcCoord"`
	DefaultCity     string `position:"Body" name:"DefaultCity"`
	Text            string `position:"Body" name:"Text"`
	ServiceCode     string `position:"Body" name:"ServiceCode"`
	DefaultDistrict string `position:"Body" name:"DefaultDistrict"`
	AppKey          string `position:"Body" name:"AppKey"`
}

// TransferCoordResponse is the response struct for api TransferCoord
type TransferCoordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateTransferCoordRequest creates a request to invoke TransferCoord API
func CreateTransferCoordRequest() (request *TransferCoordRequest) {
	request = &TransferCoordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("address-purification", "2019-11-18", "TransferCoord", "addrp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTransferCoordResponse creates a response to parse from TransferCoord response
func CreateTransferCoordResponse() (response *TransferCoordResponse) {
	response = &TransferCoordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}