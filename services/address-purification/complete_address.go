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

// CompleteAddress invokes the address_purification.CompleteAddress API synchronously
func (client *Client) CompleteAddress(request *CompleteAddressRequest) (response *CompleteAddressResponse, err error) {
	response = CreateCompleteAddressResponse()
	err = client.DoAction(request, response)
	return
}

// CompleteAddressWithChan invokes the address_purification.CompleteAddress API asynchronously
func (client *Client) CompleteAddressWithChan(request *CompleteAddressRequest) (<-chan *CompleteAddressResponse, <-chan error) {
	responseChan := make(chan *CompleteAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompleteAddress(request)
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

// CompleteAddressWithCallback invokes the address_purification.CompleteAddress API asynchronously
func (client *Client) CompleteAddressWithCallback(request *CompleteAddressRequest, callback func(response *CompleteAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompleteAddressResponse
		var err error
		defer close(result)
		response, err = client.CompleteAddress(request)
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

// CompleteAddressRequest is the request struct for api CompleteAddress
type CompleteAddressRequest struct {
	*requests.RpcRequest
	DefaultProvince string `position:"Body" name:"DefaultProvince"`
	ServiceCode     string `position:"Body" name:"ServiceCode"`
	DefaultCity     string `position:"Body" name:"DefaultCity"`
	DefaultDistrict string `position:"Body" name:"DefaultDistrict"`
	AppKey          string `position:"Body" name:"AppKey"`
	Text            string `position:"Body" name:"Text"`
}

// CompleteAddressResponse is the response struct for api CompleteAddress
type CompleteAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCompleteAddressRequest creates a request to invoke CompleteAddress API
func CreateCompleteAddressRequest() (request *CompleteAddressRequest) {
	request = &CompleteAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("address-purification", "2019-11-18", "CompleteAddress", "addrp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCompleteAddressResponse creates a response to parse from CompleteAddress response
func CreateCompleteAddressResponse() (response *CompleteAddressResponse) {
	response = &CompleteAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
