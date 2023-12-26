package xtrace

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

// SearchTraces invokes the xtrace.SearchTraces API synchronously
func (client *Client) SearchTraces(request *SearchTracesRequest) (response *SearchTracesResponse, err error) {
	response = CreateSearchTracesResponse()
	err = client.DoAction(request, response)
	return
}

// SearchTracesWithChan invokes the xtrace.SearchTraces API asynchronously
func (client *Client) SearchTracesWithChan(request *SearchTracesRequest) (<-chan *SearchTracesResponse, <-chan error) {
	responseChan := make(chan *SearchTracesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchTraces(request)
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

// SearchTracesWithCallback invokes the xtrace.SearchTraces API asynchronously
func (client *Client) SearchTracesWithCallback(request *SearchTracesRequest, callback func(response *SearchTracesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchTracesResponse
		var err error
		defer close(result)
		response, err = client.SearchTraces(request)
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

// SearchTracesRequest is the request struct for api SearchTraces
type SearchTracesRequest struct {
	*requests.RpcRequest
	AppType       string             `position:"Query" name:"AppType"`
	EndTime       requests.Integer   `position:"Query" name:"EndTime"`
	StartTime     requests.Integer   `position:"Query" name:"StartTime"`
	Reverse       requests.Boolean   `position:"Query" name:"Reverse"`
	MinDuration   requests.Integer   `position:"Query" name:"MinDuration"`
	PageNumber    requests.Integer   `position:"Query" name:"PageNumber"`
	ServiceIp     string             `position:"Query" name:"ServiceIp"`
	OperationName string             `position:"Query" name:"OperationName"`
	PageSize      requests.Integer   `position:"Query" name:"PageSize"`
	ServiceName   string             `position:"Query" name:"ServiceName"`
	Tag           *[]SearchTracesTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// SearchTracesTag is a repeated param struct in SearchTracesRequest
type SearchTracesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// SearchTracesResponse is the response struct for api SearchTraces
type SearchTracesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateSearchTracesRequest creates a request to invoke SearchTraces API
func CreateSearchTracesRequest() (request *SearchTracesRequest) {
	request = &SearchTracesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("xtrace", "2019-08-08", "SearchTraces", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchTracesResponse creates a response to parse from SearchTraces response
func CreateSearchTracesResponse() (response *SearchTracesResponse) {
	response = &SearchTracesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
