package arms

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

// ListInsightsEvents invokes the arms.ListInsightsEvents API synchronously
func (client *Client) ListInsightsEvents(request *ListInsightsEventsRequest) (response *ListInsightsEventsResponse, err error) {
	response = CreateListInsightsEventsResponse()
	err = client.DoAction(request, response)
	return
}

// ListInsightsEventsWithChan invokes the arms.ListInsightsEvents API asynchronously
func (client *Client) ListInsightsEventsWithChan(request *ListInsightsEventsRequest) (<-chan *ListInsightsEventsResponse, <-chan error) {
	responseChan := make(chan *ListInsightsEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInsightsEvents(request)
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

// ListInsightsEventsWithCallback invokes the arms.ListInsightsEvents API asynchronously
func (client *Client) ListInsightsEventsWithCallback(request *ListInsightsEventsRequest, callback func(response *ListInsightsEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInsightsEventsResponse
		var err error
		defer close(result)
		response, err = client.ListInsightsEvents(request)
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

// ListInsightsEventsRequest is the request struct for api ListInsightsEvents
type ListInsightsEventsRequest struct {
	*requests.RpcRequest
	InsightsTypes string `position:"Query" name:"InsightsTypes"`
	EndTime       string `position:"Query" name:"EndTime"`
	Pid           string `position:"Query" name:"Pid"`
	StartTime     string `position:"Query" name:"StartTime"`
}

// ListInsightsEventsResponse is the response struct for api ListInsightsEvents
type ListInsightsEventsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	InsightsEvents []ProblemInfos `json:"InsightsEvents" xml:"InsightsEvents"`
}

// CreateListInsightsEventsRequest creates a request to invoke ListInsightsEvents API
func CreateListInsightsEventsRequest() (request *ListInsightsEventsRequest) {
	request = &ListInsightsEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListInsightsEvents", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInsightsEventsResponse creates a response to parse from ListInsightsEvents response
func CreateListInsightsEventsResponse() (response *ListInsightsEventsResponse) {
	response = &ListInsightsEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}