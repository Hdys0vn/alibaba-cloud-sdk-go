package bssopenapi

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

// SetAllExpirationDay invokes the bssopenapi.SetAllExpirationDay API synchronously
func (client *Client) SetAllExpirationDay(request *SetAllExpirationDayRequest) (response *SetAllExpirationDayResponse, err error) {
	response = CreateSetAllExpirationDayResponse()
	err = client.DoAction(request, response)
	return
}

// SetAllExpirationDayWithChan invokes the bssopenapi.SetAllExpirationDay API asynchronously
func (client *Client) SetAllExpirationDayWithChan(request *SetAllExpirationDayRequest) (<-chan *SetAllExpirationDayResponse, <-chan error) {
	responseChan := make(chan *SetAllExpirationDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAllExpirationDay(request)
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

// SetAllExpirationDayWithCallback invokes the bssopenapi.SetAllExpirationDay API asynchronously
func (client *Client) SetAllExpirationDayWithCallback(request *SetAllExpirationDayRequest, callback func(response *SetAllExpirationDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAllExpirationDayResponse
		var err error
		defer close(result)
		response, err = client.SetAllExpirationDay(request)
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

// SetAllExpirationDayRequest is the request struct for api SetAllExpirationDay
type SetAllExpirationDayRequest struct {
	*requests.RpcRequest
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	UnifyExpireDay string           `position:"Query" name:"UnifyExpireDay"`
}

// SetAllExpirationDayResponse is the response struct for api SetAllExpirationDay
type SetAllExpirationDayResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSetAllExpirationDayRequest creates a request to invoke SetAllExpirationDay API
func CreateSetAllExpirationDayRequest() (request *SetAllExpirationDayRequest) {
	request = &SetAllExpirationDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "SetAllExpirationDay", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetAllExpirationDayResponse creates a response to parse from SetAllExpirationDay response
func CreateSetAllExpirationDayResponse() (response *SetAllExpirationDayResponse) {
	response = &SetAllExpirationDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
