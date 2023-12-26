package drds

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

// ReleaseInstanceInternetAddress invokes the drds.ReleaseInstanceInternetAddress API synchronously
func (client *Client) ReleaseInstanceInternetAddress(request *ReleaseInstanceInternetAddressRequest) (response *ReleaseInstanceInternetAddressResponse, err error) {
	response = CreateReleaseInstanceInternetAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseInstanceInternetAddressWithChan invokes the drds.ReleaseInstanceInternetAddress API asynchronously
func (client *Client) ReleaseInstanceInternetAddressWithChan(request *ReleaseInstanceInternetAddressRequest) (<-chan *ReleaseInstanceInternetAddressResponse, <-chan error) {
	responseChan := make(chan *ReleaseInstanceInternetAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseInstanceInternetAddress(request)
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

// ReleaseInstanceInternetAddressWithCallback invokes the drds.ReleaseInstanceInternetAddress API asynchronously
func (client *Client) ReleaseInstanceInternetAddressWithCallback(request *ReleaseInstanceInternetAddressRequest, callback func(response *ReleaseInstanceInternetAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseInstanceInternetAddressResponse
		var err error
		defer close(result)
		response, err = client.ReleaseInstanceInternetAddress(request)
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

// ReleaseInstanceInternetAddressRequest is the request struct for api ReleaseInstanceInternetAddress
type ReleaseInstanceInternetAddressRequest struct {
	*requests.RpcRequest
	DrdsPassword   string `position:"Query" name:"DrdsPassword"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// ReleaseInstanceInternetAddressResponse is the response struct for api ReleaseInstanceInternetAddress
type ReleaseInstanceInternetAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateReleaseInstanceInternetAddressRequest creates a request to invoke ReleaseInstanceInternetAddress API
func CreateReleaseInstanceInternetAddressRequest() (request *ReleaseInstanceInternetAddressRequest) {
	request = &ReleaseInstanceInternetAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "ReleaseInstanceInternetAddress", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleaseInstanceInternetAddressResponse creates a response to parse from ReleaseInstanceInternetAddress response
func CreateReleaseInstanceInternetAddressResponse() (response *ReleaseInstanceInternetAddressResponse) {
	response = &ReleaseInstanceInternetAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
