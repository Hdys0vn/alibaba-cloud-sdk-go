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

// CreateTicketWithBizData invokes the scsp.CreateTicketWithBizData API synchronously
func (client *Client) CreateTicketWithBizData(request *CreateTicketWithBizDataRequest) (response *CreateTicketWithBizDataResponse, err error) {
	response = CreateCreateTicketWithBizDataResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTicketWithBizDataWithChan invokes the scsp.CreateTicketWithBizData API asynchronously
func (client *Client) CreateTicketWithBizDataWithChan(request *CreateTicketWithBizDataRequest) (<-chan *CreateTicketWithBizDataResponse, <-chan error) {
	responseChan := make(chan *CreateTicketWithBizDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTicketWithBizData(request)
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

// CreateTicketWithBizDataWithCallback invokes the scsp.CreateTicketWithBizData API asynchronously
func (client *Client) CreateTicketWithBizDataWithCallback(request *CreateTicketWithBizDataRequest, callback func(response *CreateTicketWithBizDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTicketWithBizDataResponse
		var err error
		defer close(result)
		response, err = client.CreateTicketWithBizData(request)
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

// CreateTicketWithBizDataRequest is the request struct for api CreateTicketWithBizData
type CreateTicketWithBizDataRequest struct {
	*requests.RpcRequest
	FromInfo    string           `position:"Body"`
	ClientToken string           `position:"Body"`
	CarbonCopy  string           `position:"Body"`
	CreatorId   requests.Integer `position:"Body"`
	BizData     string           `position:"Body"`
	TemplateId  requests.Integer `position:"Body"`
	Priority    requests.Integer `position:"Body"`
	FormData    string           `position:"Body"`
	InstanceId  string           `position:"Body"`
	CreatorType requests.Integer `position:"Body"`
	CreatorName string           `position:"Body"`
	CategoryId  requests.Integer `position:"Body"`
	MemberName  string           `position:"Body"`
	MemberId    requests.Integer `position:"Body"`
}

// CreateTicketWithBizDataResponse is the response struct for api CreateTicketWithBizData
type CreateTicketWithBizDataResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      int64  `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateTicketWithBizDataRequest creates a request to invoke CreateTicketWithBizData API
func CreateCreateTicketWithBizDataRequest() (request *CreateTicketWithBizDataRequest) {
	request = &CreateTicketWithBizDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "CreateTicketWithBizData", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTicketWithBizDataResponse creates a response to parse from CreateTicketWithBizData response
func CreateCreateTicketWithBizDataResponse() (response *CreateTicketWithBizDataResponse) {
	response = &CreateTicketWithBizDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
