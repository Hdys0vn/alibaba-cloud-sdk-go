package mns_open

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

// SetTopicAttributes invokes the mns_open.SetTopicAttributes API synchronously
func (client *Client) SetTopicAttributes(request *SetTopicAttributesRequest) (response *SetTopicAttributesResponse, err error) {
	response = CreateSetTopicAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// SetTopicAttributesWithChan invokes the mns_open.SetTopicAttributes API asynchronously
func (client *Client) SetTopicAttributesWithChan(request *SetTopicAttributesRequest) (<-chan *SetTopicAttributesResponse, <-chan error) {
	responseChan := make(chan *SetTopicAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetTopicAttributes(request)
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

// SetTopicAttributesWithCallback invokes the mns_open.SetTopicAttributes API asynchronously
func (client *Client) SetTopicAttributesWithCallback(request *SetTopicAttributesRequest, callback func(response *SetTopicAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetTopicAttributesResponse
		var err error
		defer close(result)
		response, err = client.SetTopicAttributes(request)
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

// SetTopicAttributesRequest is the request struct for api SetTopicAttributes
type SetTopicAttributesRequest struct {
	*requests.RpcRequest
	TopicName      string           `position:"Query" name:"TopicName"`
	MaxMessageSize requests.Integer `position:"Query" name:"MaxMessageSize"`
	EnableLogging  requests.Boolean `position:"Query" name:"EnableLogging"`
}

// SetTopicAttributesResponse is the response struct for api SetTopicAttributes
type SetTopicAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int64  `json:"Code" xml:"Code"`
	Status    string `json:"Status" xml:"Status"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSetTopicAttributesRequest creates a request to invoke SetTopicAttributes API
func CreateSetTopicAttributesRequest() (request *SetTopicAttributesRequest) {
	request = &SetTopicAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mns-open", "2022-01-19", "SetTopicAttributes", "mns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetTopicAttributesResponse creates a response to parse from SetTopicAttributes response
func CreateSetTopicAttributesResponse() (response *SetTopicAttributesResponse) {
	response = &SetTopicAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
