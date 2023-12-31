package scsp

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

// GenerateWebSocketSign invokes the scsp.GenerateWebSocketSign API synchronously
func (client *Client) GenerateWebSocketSign(request *GenerateWebSocketSignRequest) (response *GenerateWebSocketSignResponse, err error) {
	response = CreateGenerateWebSocketSignResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateWebSocketSignWithChan invokes the scsp.GenerateWebSocketSign API asynchronously
func (client *Client) GenerateWebSocketSignWithChan(request *GenerateWebSocketSignRequest) (<-chan *GenerateWebSocketSignResponse, <-chan error) {
	responseChan := make(chan *GenerateWebSocketSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateWebSocketSign(request)
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

// GenerateWebSocketSignWithCallback invokes the scsp.GenerateWebSocketSign API asynchronously
func (client *Client) GenerateWebSocketSignWithCallback(request *GenerateWebSocketSignRequest, callback func(response *GenerateWebSocketSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateWebSocketSignResponse
		var err error
		defer close(result)
		response, err = client.GenerateWebSocketSign(request)
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

// GenerateWebSocketSignRequest is the request struct for api GenerateWebSocketSign
type GenerateWebSocketSignRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Body"`
	InstanceId  string `position:"Body"`
	AccountName string `position:"Body"`
}

// GenerateWebSocketSignResponse is the response struct for api GenerateWebSocketSign
type GenerateWebSocketSignResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           string `json:"Data" xml:"Data"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateGenerateWebSocketSignRequest creates a request to invoke GenerateWebSocketSign API
func CreateGenerateWebSocketSignRequest() (request *GenerateWebSocketSignRequest) {
	request = &GenerateWebSocketSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GenerateWebSocketSign", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateWebSocketSignResponse creates a response to parse from GenerateWebSocketSign response
func CreateGenerateWebSocketSignResponse() (response *GenerateWebSocketSignResponse) {
	response = &GenerateWebSocketSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
