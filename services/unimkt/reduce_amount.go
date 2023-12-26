package unimkt

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

// ReduceAmount invokes the unimkt.ReduceAmount API synchronously
func (client *Client) ReduceAmount(request *ReduceAmountRequest) (response *ReduceAmountResponse, err error) {
	response = CreateReduceAmountResponse()
	err = client.DoAction(request, response)
	return
}

// ReduceAmountWithChan invokes the unimkt.ReduceAmount API asynchronously
func (client *Client) ReduceAmountWithChan(request *ReduceAmountRequest) (<-chan *ReduceAmountResponse, <-chan error) {
	responseChan := make(chan *ReduceAmountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReduceAmount(request)
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

// ReduceAmountWithCallback invokes the unimkt.ReduceAmount API asynchronously
func (client *Client) ReduceAmountWithCallback(request *ReduceAmountRequest, callback func(response *ReduceAmountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReduceAmountResponse
		var err error
		defer close(result)
		response, err = client.ReduceAmount(request)
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

// ReduceAmountRequest is the request struct for api ReduceAmount
type ReduceAmountRequest struct {
	*requests.RpcRequest
	Amount    requests.Integer `position:"Query" name:"Amount"`
	V         string           `position:"Query" name:"V"`
	TaskId    requests.Integer `position:"Query" name:"TaskId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
}

// ReduceAmountResponse is the response struct for api ReduceAmount
type ReduceAmountResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateReduceAmountRequest creates a request to invoke ReduceAmount API
func CreateReduceAmountRequest() (request *ReduceAmountRequest) {
	request = &ReduceAmountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "ReduceAmount", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReduceAmountResponse creates a response to parse from ReduceAmount response
func CreateReduceAmountResponse() (response *ReduceAmountResponse) {
	response = &ReduceAmountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
