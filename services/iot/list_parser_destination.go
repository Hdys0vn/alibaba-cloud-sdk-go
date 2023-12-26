package iot

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

// ListParserDestination invokes the iot.ListParserDestination API synchronously
func (client *Client) ListParserDestination(request *ListParserDestinationRequest) (response *ListParserDestinationResponse, err error) {
	response = CreateListParserDestinationResponse()
	err = client.DoAction(request, response)
	return
}

// ListParserDestinationWithChan invokes the iot.ListParserDestination API asynchronously
func (client *Client) ListParserDestinationWithChan(request *ListParserDestinationRequest) (<-chan *ListParserDestinationResponse, <-chan error) {
	responseChan := make(chan *ListParserDestinationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListParserDestination(request)
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

// ListParserDestinationWithCallback invokes the iot.ListParserDestination API asynchronously
func (client *Client) ListParserDestinationWithCallback(request *ListParserDestinationRequest, callback func(response *ListParserDestinationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListParserDestinationResponse
		var err error
		defer close(result)
		response, err = client.ListParserDestination(request)
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

// ListParserDestinationRequest is the request struct for api ListParserDestination
type ListParserDestinationRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ParserId      requests.Integer `position:"Query" name:"ParserId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	IsFailover    requests.Boolean `position:"Query" name:"IsFailover"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// ListParserDestinationResponse is the response struct for api ListParserDestination
type ListParserDestinationResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	Success      bool                        `json:"Success" xml:"Success"`
	Code         string                      `json:"Code" xml:"Code"`
	ErrorMessage string                      `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInListParserDestination `json:"Data" xml:"Data"`
}

// CreateListParserDestinationRequest creates a request to invoke ListParserDestination API
func CreateListParserDestinationRequest() (request *ListParserDestinationRequest) {
	request = &ListParserDestinationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListParserDestination", "", "")
	request.Method = requests.POST
	return
}

// CreateListParserDestinationResponse creates a response to parse from ListParserDestination response
func CreateListParserDestinationResponse() (response *ListParserDestinationResponse) {
	response = &ListParserDestinationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
