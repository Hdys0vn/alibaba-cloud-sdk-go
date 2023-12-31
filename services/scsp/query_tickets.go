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

// QueryTickets invokes the scsp.QueryTickets API synchronously
func (client *Client) QueryTickets(request *QueryTicketsRequest) (response *QueryTicketsResponse, err error) {
	response = CreateQueryTicketsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTicketsWithChan invokes the scsp.QueryTickets API asynchronously
func (client *Client) QueryTicketsWithChan(request *QueryTicketsRequest) (<-chan *QueryTicketsResponse, <-chan error) {
	responseChan := make(chan *QueryTicketsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTickets(request)
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

// QueryTicketsWithCallback invokes the scsp.QueryTickets API asynchronously
func (client *Client) QueryTicketsWithCallback(request *QueryTicketsRequest, callback func(response *QueryTicketsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTicketsResponse
		var err error
		defer close(result)
		response, err = client.QueryTickets(request)
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

// QueryTicketsRequest is the request struct for api QueryTickets
type QueryTicketsRequest struct {
	*requests.RpcRequest
	SrType      requests.Integer `position:"Body"`
	TouchId     requests.Integer `position:"Body"`
	DealId      requests.Integer `position:"Body"`
	CurrentPage requests.Integer `position:"Body"`
	TaskStatus  requests.Integer `position:"Body"`
	InstanceId  string           `position:"Body"`
	AccountName string           `position:"Body"`
	CaseId      requests.Integer `position:"Body"`
	Extra       string           `position:"Body"`
	ChannelType requests.Integer `position:"Body"`
	PageSize    requests.Integer `position:"Body"`
	CaseType    requests.Integer `position:"Body"`
	CaseStatus  requests.Integer `position:"Body"`
	ChannelId   string           `position:"Body"`
}

// QueryTicketsResponse is the response struct for api QueryTickets
type QueryTicketsResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryTicketsRequest creates a request to invoke QueryTickets API
func CreateQueryTicketsRequest() (request *QueryTicketsRequest) {
	request = &QueryTicketsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "QueryTickets", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTicketsResponse creates a response to parse from QueryTickets response
func CreateQueryTicketsResponse() (response *QueryTicketsResponse) {
	response = &QueryTicketsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
