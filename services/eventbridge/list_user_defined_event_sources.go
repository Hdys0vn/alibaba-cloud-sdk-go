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

// ListUserDefinedEventSources invokes the eventbridge.ListUserDefinedEventSources API synchronously
func (client *Client) ListUserDefinedEventSources(request *ListUserDefinedEventSourcesRequest) (response *ListUserDefinedEventSourcesResponse, err error) {
	response = CreateListUserDefinedEventSourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserDefinedEventSourcesWithChan invokes the eventbridge.ListUserDefinedEventSources API asynchronously
func (client *Client) ListUserDefinedEventSourcesWithChan(request *ListUserDefinedEventSourcesRequest) (<-chan *ListUserDefinedEventSourcesResponse, <-chan error) {
	responseChan := make(chan *ListUserDefinedEventSourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserDefinedEventSources(request)
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

// ListUserDefinedEventSourcesWithCallback invokes the eventbridge.ListUserDefinedEventSources API asynchronously
func (client *Client) ListUserDefinedEventSourcesWithCallback(request *ListUserDefinedEventSourcesRequest, callback func(response *ListUserDefinedEventSourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserDefinedEventSourcesResponse
		var err error
		defer close(result)
		response, err = client.ListUserDefinedEventSources(request)
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

// ListUserDefinedEventSourcesRequest is the request struct for api ListUserDefinedEventSources
type ListUserDefinedEventSourcesRequest struct {
	*requests.RpcRequest
	Type         string           `position:"Query" name:"Type"`
	EventBusName string           `position:"Query" name:"EventBusName"`
	NextToken    string           `position:"Query" name:"NextToken"`
	Limit        requests.Integer `position:"Query" name:"Limit"`
	NamePrefix   string           `position:"Query" name:"NamePrefix"`
}

// ListUserDefinedEventSourcesResponse is the response struct for api ListUserDefinedEventSources
type ListUserDefinedEventSourcesResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListUserDefinedEventSourcesRequest creates a request to invoke ListUserDefinedEventSources API
func CreateListUserDefinedEventSourcesRequest() (request *ListUserDefinedEventSourcesRequest) {
	request = &ListUserDefinedEventSourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "ListUserDefinedEventSources", "", "")
	request.Method = requests.POST
	return
}

// CreateListUserDefinedEventSourcesResponse creates a response to parse from ListUserDefinedEventSources response
func CreateListUserDefinedEventSourcesResponse() (response *ListUserDefinedEventSourcesResponse) {
	response = &ListUserDefinedEventSourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
