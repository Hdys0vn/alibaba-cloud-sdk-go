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

// RequestPayDemand invokes the domain.RequestPayDemand API synchronously
func (client *Client) RequestPayDemand(request *RequestPayDemandRequest) (response *RequestPayDemandResponse, err error) {
	response = CreateRequestPayDemandResponse()
	err = client.DoAction(request, response)
	return
}

// RequestPayDemandWithChan invokes the domain.RequestPayDemand API asynchronously
func (client *Client) RequestPayDemandWithChan(request *RequestPayDemandRequest) (<-chan *RequestPayDemandResponse, <-chan error) {
	responseChan := make(chan *RequestPayDemandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestPayDemand(request)
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

// RequestPayDemandWithCallback invokes the domain.RequestPayDemand API asynchronously
func (client *Client) RequestPayDemandWithCallback(request *RequestPayDemandRequest, callback func(response *RequestPayDemandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestPayDemandResponse
		var err error
		defer close(result)
		response, err = client.RequestPayDemand(request)
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

// RequestPayDemandRequest is the request struct for api RequestPayDemand
type RequestPayDemandRequest struct {
	*requests.RpcRequest
	Price       requests.Float   `position:"Query" name:"Price"`
	BizId       string           `position:"Query" name:"BizId"`
	DomainName  string           `position:"Query" name:"DomainName"`
	ProduceType requests.Integer `position:"Query" name:"ProduceType"`
	Message     string           `position:"Query" name:"Message"`
}

// RequestPayDemandResponse is the response struct for api RequestPayDemand
type RequestPayDemandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRequestPayDemandRequest creates a request to invoke RequestPayDemand API
func CreateRequestPayDemandRequest() (request *RequestPayDemandRequest) {
	request = &RequestPayDemandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "RequestPayDemand", "", "")
	request.Method = requests.POST
	return
}

// CreateRequestPayDemandResponse creates a response to parse from RequestPayDemand response
func CreateRequestPayDemandResponse() (response *RequestPayDemandResponse) {
	response = &RequestPayDemandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
