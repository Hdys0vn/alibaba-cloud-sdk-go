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

// GetGateway invokes the mse.GetGateway API synchronously
func (client *Client) GetGateway(request *GetGatewayRequest) (response *GetGatewayResponse, err error) {
	response = CreateGetGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// GetGatewayWithChan invokes the mse.GetGateway API asynchronously
func (client *Client) GetGatewayWithChan(request *GetGatewayRequest) (<-chan *GetGatewayResponse, <-chan error) {
	responseChan := make(chan *GetGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGateway(request)
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

// GetGatewayWithCallback invokes the mse.GetGateway API asynchronously
func (client *Client) GetGatewayWithCallback(request *GetGatewayRequest, callback func(response *GetGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGatewayResponse
		var err error
		defer close(result)
		response, err = client.GetGateway(request)
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

// GetGatewayRequest is the request struct for api GetGateway
type GetGatewayRequest struct {
	*requests.RpcRequest
	MseSessionId    string `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string `position:"Query" name:"GatewayUniqueId"`
	AcceptLanguage  string `position:"Query" name:"AcceptLanguage"`
}

// GetGatewayResponse is the response struct for api GetGateway
type GetGatewayResponse struct {
	*responses.BaseResponse
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string           `json:"Message" xml:"Message"`
	Code           int              `json:"Code" xml:"Code"`
	Success        bool             `json:"Success" xml:"Success"`
	Data           DataInGetGateway `json:"Data" xml:"Data"`
}

// CreateGetGatewayRequest creates a request to invoke GetGateway API
func CreateGetGatewayRequest() (request *GetGatewayRequest) {
	request = &GetGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetGateway", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetGatewayResponse creates a response to parse from GetGateway response
func CreateGetGatewayResponse() (response *GetGatewayResponse) {
	response = &GetGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
