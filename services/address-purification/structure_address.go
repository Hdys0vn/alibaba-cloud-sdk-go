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

// StructureAddress invokes the address_purification.StructureAddress API synchronously
func (client *Client) StructureAddress(request *StructureAddressRequest) (response *StructureAddressResponse, err error) {
	response = CreateStructureAddressResponse()
	err = client.DoAction(request, response)
	return
}

// StructureAddressWithChan invokes the address_purification.StructureAddress API asynchronously
func (client *Client) StructureAddressWithChan(request *StructureAddressRequest) (<-chan *StructureAddressResponse, <-chan error) {
	responseChan := make(chan *StructureAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StructureAddress(request)
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

// StructureAddressWithCallback invokes the address_purification.StructureAddress API asynchronously
func (client *Client) StructureAddressWithCallback(request *StructureAddressRequest, callback func(response *StructureAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StructureAddressResponse
		var err error
		defer close(result)
		response, err = client.StructureAddress(request)
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

// StructureAddressRequest is the request struct for api StructureAddress
type StructureAddressRequest struct {
	*requests.RpcRequest
	DefaultProvince string `position:"Body" name:"DefaultProvince"`
	ServiceCode     string `position:"Body" name:"ServiceCode"`
	DefaultCity     string `position:"Body" name:"DefaultCity"`
	DefaultDistrict string `position:"Body" name:"DefaultDistrict"`
	AppKey          string `position:"Body" name:"AppKey"`
	Text            string `position:"Body" name:"Text"`
}

// StructureAddressResponse is the response struct for api StructureAddress
type StructureAddressResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStructureAddressRequest creates a request to invoke StructureAddress API
func CreateStructureAddressRequest() (request *StructureAddressRequest) {
	request = &StructureAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("address-purification", "2019-11-18", "StructureAddress", "addrp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStructureAddressResponse creates a response to parse from StructureAddress response
func CreateStructureAddressResponse() (response *StructureAddressResponse) {
	response = &StructureAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
