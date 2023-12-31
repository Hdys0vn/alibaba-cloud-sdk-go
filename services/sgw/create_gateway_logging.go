package sgw

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

// CreateGatewayLogging invokes the sgw.CreateGatewayLogging API synchronously
func (client *Client) CreateGatewayLogging(request *CreateGatewayLoggingRequest) (response *CreateGatewayLoggingResponse, err error) {
	response = CreateCreateGatewayLoggingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGatewayLoggingWithChan invokes the sgw.CreateGatewayLogging API asynchronously
func (client *Client) CreateGatewayLoggingWithChan(request *CreateGatewayLoggingRequest) (<-chan *CreateGatewayLoggingResponse, <-chan error) {
	responseChan := make(chan *CreateGatewayLoggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGatewayLogging(request)
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

// CreateGatewayLoggingWithCallback invokes the sgw.CreateGatewayLogging API asynchronously
func (client *Client) CreateGatewayLoggingWithCallback(request *CreateGatewayLoggingRequest, callback func(response *CreateGatewayLoggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGatewayLoggingResponse
		var err error
		defer close(result)
		response, err = client.CreateGatewayLogging(request)
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

// CreateGatewayLoggingRequest is the request struct for api CreateGatewayLogging
type CreateGatewayLoggingRequest struct {
	*requests.RpcRequest
	SlsLogstore   string `position:"Query" name:"SlsLogstore"`
	SlsProject    string `position:"Query" name:"SlsProject"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// CreateGatewayLoggingResponse is the response struct for api CreateGatewayLogging
type CreateGatewayLoggingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateGatewayLoggingRequest creates a request to invoke CreateGatewayLogging API
func CreateCreateGatewayLoggingRequest() (request *CreateGatewayLoggingRequest) {
	request = &CreateGatewayLoggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "CreateGatewayLogging", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGatewayLoggingResponse creates a response to parse from CreateGatewayLogging response
func CreateCreateGatewayLoggingResponse() (response *CreateGatewayLoggingResponse) {
	response = &CreateGatewayLoggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
