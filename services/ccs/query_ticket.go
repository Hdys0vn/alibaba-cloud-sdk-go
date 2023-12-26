package ccs

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

// QueryTicket invokes the ccs.QueryTicket API synchronously
// api document: https://help.aliyun.com/api/ccs/queryticket.html
func (client *Client) QueryTicket(request *QueryTicketRequest) (response *QueryTicketResponse, err error) {
	response = CreateQueryTicketResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTicketWithChan invokes the ccs.QueryTicket API asynchronously
// api document: https://help.aliyun.com/api/ccs/queryticket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTicketWithChan(request *QueryTicketRequest) (<-chan *QueryTicketResponse, <-chan error) {
	responseChan := make(chan *QueryTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTicket(request)
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

// QueryTicketWithCallback invokes the ccs.QueryTicket API asynchronously
// api document: https://help.aliyun.com/api/ccs/queryticket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTicketWithCallback(request *QueryTicketRequest, callback func(response *QueryTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTicketResponse
		var err error
		defer close(result)
		response, err = client.QueryTicket(request)
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

// QueryTicketRequest is the request struct for api QueryTicket
type QueryTicketRequest struct {
	*requests.RpcRequest
	Stage         string           `position:"Query" name:"Stage"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CreatorId     string           `position:"Query" name:"CreatorId"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	Type          string           `position:"Query" name:"Type"`
	CcsInstanceId string           `position:"Query" name:"CcsInstanceId"`
}

// QueryTicketResponse is the response struct for api QueryTicket
type QueryTicketResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNum    int     `json:"PageNum" xml:"PageNum"`
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	Tickets    Tickets `json:"Tickets" xml:"Tickets"`
}

// CreateQueryTicketRequest creates a request to invoke QueryTicket API
func CreateQueryTicketRequest() (request *QueryTicketRequest) {
	request = &QueryTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ccs", "2017-10-01", "QueryTicket", "ccs", "openAPI")
	return
}

// CreateQueryTicketResponse creates a response to parse from QueryTicket response
func CreateQueryTicketResponse() (response *QueryTicketResponse) {
	response = &QueryTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
