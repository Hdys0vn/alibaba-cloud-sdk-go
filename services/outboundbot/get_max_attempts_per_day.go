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

// GetMaxAttemptsPerDay invokes the outboundbot.GetMaxAttemptsPerDay API synchronously
func (client *Client) GetMaxAttemptsPerDay(request *GetMaxAttemptsPerDayRequest) (response *GetMaxAttemptsPerDayResponse, err error) {
	response = CreateGetMaxAttemptsPerDayResponse()
	err = client.DoAction(request, response)
	return
}

// GetMaxAttemptsPerDayWithChan invokes the outboundbot.GetMaxAttemptsPerDay API asynchronously
func (client *Client) GetMaxAttemptsPerDayWithChan(request *GetMaxAttemptsPerDayRequest) (<-chan *GetMaxAttemptsPerDayResponse, <-chan error) {
	responseChan := make(chan *GetMaxAttemptsPerDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMaxAttemptsPerDay(request)
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

// GetMaxAttemptsPerDayWithCallback invokes the outboundbot.GetMaxAttemptsPerDay API asynchronously
func (client *Client) GetMaxAttemptsPerDayWithCallback(request *GetMaxAttemptsPerDayRequest, callback func(response *GetMaxAttemptsPerDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMaxAttemptsPerDayResponse
		var err error
		defer close(result)
		response, err = client.GetMaxAttemptsPerDay(request)
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

// GetMaxAttemptsPerDayRequest is the request struct for api GetMaxAttemptsPerDay
type GetMaxAttemptsPerDayRequest struct {
	*requests.RpcRequest
	StrategyLevel requests.Integer `position:"Query" name:"StrategyLevel"`
	EntryId       string           `position:"Query" name:"EntryId"`
}

// GetMaxAttemptsPerDayResponse is the response struct for api GetMaxAttemptsPerDay
type GetMaxAttemptsPerDayResponse struct {
	*responses.BaseResponse
	HttpStatusCode    int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Success           bool   `json:"Success" xml:"Success"`
	MaxAttemptsPerDay int    `json:"MaxAttemptsPerDay" xml:"MaxAttemptsPerDay"`
	Code              string `json:"Code" xml:"Code"`
	Message           string `json:"Message" xml:"Message"`
}

// CreateGetMaxAttemptsPerDayRequest creates a request to invoke GetMaxAttemptsPerDay API
func CreateGetMaxAttemptsPerDayRequest() (request *GetMaxAttemptsPerDayRequest) {
	request = &GetMaxAttemptsPerDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetMaxAttemptsPerDay", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMaxAttemptsPerDayResponse creates a response to parse from GetMaxAttemptsPerDay response
func CreateGetMaxAttemptsPerDayResponse() (response *GetMaxAttemptsPerDayResponse) {
	response = &GetMaxAttemptsPerDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
