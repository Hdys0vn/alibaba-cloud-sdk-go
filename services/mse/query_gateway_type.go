package mse

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

// QueryGatewayType invokes the mse.QueryGatewayType API synchronously
func (client *Client) QueryGatewayType(request *QueryGatewayTypeRequest) (response *QueryGatewayTypeResponse, err error) {
	response = CreateQueryGatewayTypeResponse()
	err = client.DoAction(request, response)
	return
}

// QueryGatewayTypeWithChan invokes the mse.QueryGatewayType API asynchronously
func (client *Client) QueryGatewayTypeWithChan(request *QueryGatewayTypeRequest) (<-chan *QueryGatewayTypeResponse, <-chan error) {
	responseChan := make(chan *QueryGatewayTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryGatewayType(request)
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

// QueryGatewayTypeWithCallback invokes the mse.QueryGatewayType API asynchronously
func (client *Client) QueryGatewayTypeWithCallback(request *QueryGatewayTypeRequest, callback func(response *QueryGatewayTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryGatewayTypeResponse
		var err error
		defer close(result)
		response, err = client.QueryGatewayType(request)
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

// QueryGatewayTypeRequest is the request struct for api QueryGatewayType
type QueryGatewayTypeRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// QueryGatewayTypeResponse is the response struct for api QueryGatewayType
type QueryGatewayTypeResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	Code           int      `json:"Code" xml:"Code"`
	Success        bool     `json:"Success" xml:"Success"`
	Data           []string `json:"Data" xml:"Data"`
}

// CreateQueryGatewayTypeRequest creates a request to invoke QueryGatewayType API
func CreateQueryGatewayTypeRequest() (request *QueryGatewayTypeRequest) {
	request = &QueryGatewayTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryGatewayType", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryGatewayTypeResponse creates a response to parse from QueryGatewayType response
func CreateQueryGatewayTypeResponse() (response *QueryGatewayTypeResponse) {
	response = &QueryGatewayTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
