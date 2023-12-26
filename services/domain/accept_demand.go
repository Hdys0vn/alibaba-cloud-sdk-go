package domain

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

// AcceptDemand invokes the domain.AcceptDemand API synchronously
func (client *Client) AcceptDemand(request *AcceptDemandRequest) (response *AcceptDemandResponse, err error) {
	response = CreateAcceptDemandResponse()
	err = client.DoAction(request, response)
	return
}

// AcceptDemandWithChan invokes the domain.AcceptDemand API asynchronously
func (client *Client) AcceptDemandWithChan(request *AcceptDemandRequest) (<-chan *AcceptDemandResponse, <-chan error) {
	responseChan := make(chan *AcceptDemandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AcceptDemand(request)
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

// AcceptDemandWithCallback invokes the domain.AcceptDemand API asynchronously
func (client *Client) AcceptDemandWithCallback(request *AcceptDemandRequest, callback func(response *AcceptDemandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AcceptDemandResponse
		var err error
		defer close(result)
		response, err = client.AcceptDemand(request)
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

// AcceptDemandRequest is the request struct for api AcceptDemand
type AcceptDemandRequest struct {
	*requests.RpcRequest
	BizId   string `position:"Query" name:"BizId"`
	Message string `position:"Query" name:"Message"`
}

// AcceptDemandResponse is the response struct for api AcceptDemand
type AcceptDemandResponse struct {
	*responses.BaseResponse
	BindUrl   string `json:"BindUrl" xml:"BindUrl"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAcceptDemandRequest creates a request to invoke AcceptDemand API
func CreateAcceptDemandRequest() (request *AcceptDemandRequest) {
	request = &AcceptDemandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "AcceptDemand", "", "")
	request.Method = requests.POST
	return
}

// CreateAcceptDemandResponse creates a response to parse from AcceptDemand response
func CreateAcceptDemandResponse() (response *AcceptDemandResponse) {
	response = &AcceptDemandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
