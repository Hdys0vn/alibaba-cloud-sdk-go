package ccc

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

// ListAttempts invokes the ccc.ListAttempts API synchronously
func (client *Client) ListAttempts(request *ListAttemptsRequest) (response *ListAttemptsResponse, err error) {
	response = CreateListAttemptsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAttemptsWithChan invokes the ccc.ListAttempts API asynchronously
func (client *Client) ListAttemptsWithChan(request *ListAttemptsRequest) (<-chan *ListAttemptsResponse, <-chan error) {
	responseChan := make(chan *ListAttemptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAttempts(request)
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

// ListAttemptsWithCallback invokes the ccc.ListAttempts API asynchronously
func (client *Client) ListAttemptsWithCallback(request *ListAttemptsRequest, callback func(response *ListAttemptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAttemptsResponse
		var err error
		defer close(result)
		response, err = client.ListAttempts(request)
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

// ListAttemptsRequest is the request struct for api ListAttempts
type ListAttemptsRequest struct {
	*requests.RpcRequest
	ContactId  string           `position:"Query" name:"ContactId"`
	CampaignId string           `position:"Query" name:"CampaignId"`
	Criteria   string           `position:"Query" name:"Criteria"`
	Callee     string           `position:"Query" name:"Callee"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	QueueId    string           `position:"Query" name:"QueueId"`
	AgentId    string           `position:"Query" name:"AgentId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	Caller     string           `position:"Query" name:"Caller"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	CaseId     string           `position:"Query" name:"CaseId"`
	AttemptId  string           `position:"Query" name:"AttemptId"`
}

// ListAttemptsResponse is the response struct for api ListAttempts
type ListAttemptsResponse struct {
	*responses.BaseResponse
	Code           string             `json:"Code" xml:"Code"`
	HttpStatusCode int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string             `json:"Message" xml:"Message"`
	RequestId      string             `json:"RequestId" xml:"RequestId"`
	Data           DataInListAttempts `json:"Data" xml:"Data"`
}

// CreateListAttemptsRequest creates a request to invoke ListAttempts API
func CreateListAttemptsRequest() (request *ListAttemptsRequest) {
	request = &ListAttemptsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListAttempts", "CCC", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListAttemptsResponse creates a response to parse from ListAttempts response
func CreateListAttemptsResponse() (response *ListAttemptsResponse) {
	response = &ListAttemptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
