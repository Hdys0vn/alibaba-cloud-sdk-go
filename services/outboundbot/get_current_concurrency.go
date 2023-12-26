package outboundbot

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

// GetCurrentConcurrency invokes the outboundbot.GetCurrentConcurrency API synchronously
func (client *Client) GetCurrentConcurrency(request *GetCurrentConcurrencyRequest) (response *GetCurrentConcurrencyResponse, err error) {
	response = CreateGetCurrentConcurrencyResponse()
	err = client.DoAction(request, response)
	return
}

// GetCurrentConcurrencyWithChan invokes the outboundbot.GetCurrentConcurrency API asynchronously
func (client *Client) GetCurrentConcurrencyWithChan(request *GetCurrentConcurrencyRequest) (<-chan *GetCurrentConcurrencyResponse, <-chan error) {
	responseChan := make(chan *GetCurrentConcurrencyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCurrentConcurrency(request)
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

// GetCurrentConcurrencyWithCallback invokes the outboundbot.GetCurrentConcurrency API asynchronously
func (client *Client) GetCurrentConcurrencyWithCallback(request *GetCurrentConcurrencyRequest, callback func(response *GetCurrentConcurrencyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCurrentConcurrencyResponse
		var err error
		defer close(result)
		response, err = client.GetCurrentConcurrency(request)
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

// GetCurrentConcurrencyRequest is the request struct for api GetCurrentConcurrency
type GetCurrentConcurrencyRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetCurrentConcurrencyResponse is the response struct for api GetCurrentConcurrency
type GetCurrentConcurrencyResponse struct {
	*responses.BaseResponse
	HttpStatusCode            int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	Success                   bool   `json:"Success" xml:"Success"`
	CurrentConcurrency        int    `json:"CurrentConcurrency" xml:"CurrentConcurrency"`
	Code                      string `json:"Code" xml:"Code"`
	Message                   string `json:"Message" xml:"Message"`
	MaxConcurrentConversation int    `json:"MaxConcurrentConversation" xml:"MaxConcurrentConversation"`
	InstanceId                string `json:"InstanceId" xml:"InstanceId"`
}

// CreateGetCurrentConcurrencyRequest creates a request to invoke GetCurrentConcurrency API
func CreateGetCurrentConcurrencyRequest() (request *GetCurrentConcurrencyRequest) {
	request = &GetCurrentConcurrencyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetCurrentConcurrency", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCurrentConcurrencyResponse creates a response to parse from GetCurrentConcurrency response
func CreateGetCurrentConcurrencyResponse() (response *GetCurrentConcurrencyResponse) {
	response = &GetCurrentConcurrencyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
