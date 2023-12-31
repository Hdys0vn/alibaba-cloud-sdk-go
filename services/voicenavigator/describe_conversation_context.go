package voicenavigator

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

// DescribeConversationContext invokes the voicenavigator.DescribeConversationContext API synchronously
func (client *Client) DescribeConversationContext(request *DescribeConversationContextRequest) (response *DescribeConversationContextResponse, err error) {
	response = CreateDescribeConversationContextResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConversationContextWithChan invokes the voicenavigator.DescribeConversationContext API asynchronously
func (client *Client) DescribeConversationContextWithChan(request *DescribeConversationContextRequest) (<-chan *DescribeConversationContextResponse, <-chan error) {
	responseChan := make(chan *DescribeConversationContextResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConversationContext(request)
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

// DescribeConversationContextWithCallback invokes the voicenavigator.DescribeConversationContext API asynchronously
func (client *Client) DescribeConversationContextWithCallback(request *DescribeConversationContextRequest, callback func(response *DescribeConversationContextResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConversationContextResponse
		var err error
		defer close(result)
		response, err = client.DescribeConversationContext(request)
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

// DescribeConversationContextRequest is the request struct for api DescribeConversationContext
type DescribeConversationContextRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// DescribeConversationContextResponse is the response struct for api DescribeConversationContext
type DescribeConversationContextResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	ConversationContext string `json:"ConversationContext" xml:"ConversationContext"`
}

// CreateDescribeConversationContextRequest creates a request to invoke DescribeConversationContext API
func CreateDescribeConversationContextRequest() (request *DescribeConversationContextRequest) {
	request = &DescribeConversationContextRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "DescribeConversationContext", "voicebot", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeConversationContextResponse creates a response to parse from DescribeConversationContext response
func CreateDescribeConversationContextResponse() (response *DescribeConversationContextResponse) {
	response = &DescribeConversationContextResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
