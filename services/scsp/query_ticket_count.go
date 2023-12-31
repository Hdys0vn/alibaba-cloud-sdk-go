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

// QueryTicketCount invokes the scsp.QueryTicketCount API synchronously
func (client *Client) QueryTicketCount(request *QueryTicketCountRequest) (response *QueryTicketCountResponse, err error) {
	response = CreateQueryTicketCountResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTicketCountWithChan invokes the scsp.QueryTicketCount API asynchronously
func (client *Client) QueryTicketCountWithChan(request *QueryTicketCountRequest) (<-chan *QueryTicketCountResponse, <-chan error) {
	responseChan := make(chan *QueryTicketCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTicketCount(request)
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

// QueryTicketCountWithCallback invokes the scsp.QueryTicketCount API asynchronously
func (client *Client) QueryTicketCountWithCallback(request *QueryTicketCountRequest, callback func(response *QueryTicketCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTicketCountResponse
		var err error
		defer close(result)
		response, err = client.QueryTicketCount(request)
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

// QueryTicketCountRequest is the request struct for api QueryTicketCount
type QueryTicketCountRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query"`
	InstanceId  string           `position:"Query"`
	OperatorId  requests.Integer `position:"Query"`
}

// QueryTicketCountResponse is the response struct for api QueryTicketCount
type QueryTicketCountResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryTicketCountRequest creates a request to invoke QueryTicketCount API
func CreateQueryTicketCountRequest() (request *QueryTicketCountRequest) {
	request = &QueryTicketCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "QueryTicketCount", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTicketCountResponse creates a response to parse from QueryTicketCount response
func CreateQueryTicketCountResponse() (response *QueryTicketCountResponse) {
	response = &QueryTicketCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
