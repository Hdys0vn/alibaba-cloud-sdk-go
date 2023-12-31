package outboundbot

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

// DeleteGlobalQuestion invokes the outboundbot.DeleteGlobalQuestion API synchronously
func (client *Client) DeleteGlobalQuestion(request *DeleteGlobalQuestionRequest) (response *DeleteGlobalQuestionResponse, err error) {
	response = CreateDeleteGlobalQuestionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGlobalQuestionWithChan invokes the outboundbot.DeleteGlobalQuestion API asynchronously
func (client *Client) DeleteGlobalQuestionWithChan(request *DeleteGlobalQuestionRequest) (<-chan *DeleteGlobalQuestionResponse, <-chan error) {
	responseChan := make(chan *DeleteGlobalQuestionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGlobalQuestion(request)
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

// DeleteGlobalQuestionWithCallback invokes the outboundbot.DeleteGlobalQuestion API asynchronously
func (client *Client) DeleteGlobalQuestionWithCallback(request *DeleteGlobalQuestionRequest, callback func(response *DeleteGlobalQuestionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGlobalQuestionResponse
		var err error
		defer close(result)
		response, err = client.DeleteGlobalQuestion(request)
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

// DeleteGlobalQuestionRequest is the request struct for api DeleteGlobalQuestion
type DeleteGlobalQuestionRequest struct {
	*requests.RpcRequest
	GlobalQuestionId string `position:"Query" name:"GlobalQuestionId"`
	ScriptId         string `position:"Query" name:"ScriptId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
}

// DeleteGlobalQuestionResponse is the response struct for api DeleteGlobalQuestion
type DeleteGlobalQuestionResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteGlobalQuestionRequest creates a request to invoke DeleteGlobalQuestion API
func CreateDeleteGlobalQuestionRequest() (request *DeleteGlobalQuestionRequest) {
	request = &DeleteGlobalQuestionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DeleteGlobalQuestion", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteGlobalQuestionResponse creates a response to parse from DeleteGlobalQuestion response
func CreateDeleteGlobalQuestionResponse() (response *DeleteGlobalQuestionResponse) {
	response = &DeleteGlobalQuestionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
