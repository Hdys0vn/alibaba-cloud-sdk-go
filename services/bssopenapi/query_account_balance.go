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

// QueryAccountBalance invokes the bssopenapi.QueryAccountBalance API synchronously
func (client *Client) QueryAccountBalance(request *QueryAccountBalanceRequest) (response *QueryAccountBalanceResponse, err error) {
	response = CreateQueryAccountBalanceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAccountBalanceWithChan invokes the bssopenapi.QueryAccountBalance API asynchronously
func (client *Client) QueryAccountBalanceWithChan(request *QueryAccountBalanceRequest) (<-chan *QueryAccountBalanceResponse, <-chan error) {
	responseChan := make(chan *QueryAccountBalanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAccountBalance(request)
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

// QueryAccountBalanceWithCallback invokes the bssopenapi.QueryAccountBalance API asynchronously
func (client *Client) QueryAccountBalanceWithCallback(request *QueryAccountBalanceRequest, callback func(response *QueryAccountBalanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAccountBalanceResponse
		var err error
		defer close(result)
		response, err = client.QueryAccountBalance(request)
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

// QueryAccountBalanceRequest is the request struct for api QueryAccountBalance
type QueryAccountBalanceRequest struct {
	*requests.RpcRequest
}

// QueryAccountBalanceResponse is the response struct for api QueryAccountBalance
type QueryAccountBalanceResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryAccountBalanceRequest creates a request to invoke QueryAccountBalance API
func CreateQueryAccountBalanceRequest() (request *QueryAccountBalanceRequest) {
	request = &QueryAccountBalanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryAccountBalance", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAccountBalanceResponse creates a response to parse from QueryAccountBalance response
func CreateQueryAccountBalanceResponse() (response *QueryAccountBalanceResponse) {
	response = &QueryAccountBalanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
