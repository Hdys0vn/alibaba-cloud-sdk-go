package eventbridge

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

// QueryTracedEvents invokes the eventbridge.QueryTracedEvents API synchronously
func (client *Client) QueryTracedEvents(request *QueryTracedEventsRequest) (response *QueryTracedEventsResponse, err error) {
	response = CreateQueryTracedEventsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTracedEventsWithChan invokes the eventbridge.QueryTracedEvents API asynchronously
func (client *Client) QueryTracedEventsWithChan(request *QueryTracedEventsRequest) (<-chan *QueryTracedEventsResponse, <-chan error) {
	responseChan := make(chan *QueryTracedEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTracedEvents(request)
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

// QueryTracedEventsWithCallback invokes the eventbridge.QueryTracedEvents API asynchronously
func (client *Client) QueryTracedEventsWithCallback(request *QueryTracedEventsRequest, callback func(response *QueryTracedEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTracedEventsResponse
		var err error
		defer close(result)
		response, err = client.QueryTracedEvents(request)
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

// QueryTracedEventsRequest is the request struct for api QueryTracedEvents
type QueryTracedEventsRequest struct {
	*requests.RpcRequest
	MatchedRule  string           `position:"Query" name:"MatchedRule"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	EventBusName string           `position:"Query" name:"EventBusName"`
	EventSource  string           `position:"Query" name:"EventSource"`
	NextToken    string           `position:"Query" name:"NextToken"`
	Limit        requests.Integer `position:"Query" name:"Limit"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	EventType    string           `position:"Query" name:"EventType"`
}

// QueryTracedEventsResponse is the response struct for api QueryTracedEvents
type QueryTracedEventsResponse struct {
	*responses.BaseResponse
	Message   string                  `json:"Message" xml:"Message"`
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Code      string                  `json:"Code" xml:"Code"`
	Success   bool                    `json:"Success" xml:"Success"`
	Data      DataInQueryTracedEvents `json:"Data" xml:"Data"`
}

// CreateQueryTracedEventsRequest creates a request to invoke QueryTracedEvents API
func CreateQueryTracedEventsRequest() (request *QueryTracedEventsRequest) {
	request = &QueryTracedEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "QueryTracedEvents", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTracedEventsResponse creates a response to parse from QueryTracedEvents response
func CreateQueryTracedEventsResponse() (response *QueryTracedEventsResponse) {
	response = &QueryTracedEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
