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

// TestEventPattern invokes the eventbridge.TestEventPattern API synchronously
func (client *Client) TestEventPattern(request *TestEventPatternRequest) (response *TestEventPatternResponse, err error) {
	response = CreateTestEventPatternResponse()
	err = client.DoAction(request, response)
	return
}

// TestEventPatternWithChan invokes the eventbridge.TestEventPattern API asynchronously
func (client *Client) TestEventPatternWithChan(request *TestEventPatternRequest) (<-chan *TestEventPatternResponse, <-chan error) {
	responseChan := make(chan *TestEventPatternResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestEventPattern(request)
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

// TestEventPatternWithCallback invokes the eventbridge.TestEventPattern API asynchronously
func (client *Client) TestEventPatternWithCallback(request *TestEventPatternRequest, callback func(response *TestEventPatternResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestEventPatternResponse
		var err error
		defer close(result)
		response, err = client.TestEventPattern(request)
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

// TestEventPatternRequest is the request struct for api TestEventPattern
type TestEventPatternRequest struct {
	*requests.RpcRequest
	EventPattern string `position:"Body" name:"EventPattern"`
	Event        string `position:"Body" name:"Event"`
}

// TestEventPatternResponse is the response struct for api TestEventPattern
type TestEventPatternResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTestEventPatternRequest creates a request to invoke TestEventPattern API
func CreateTestEventPatternRequest() (request *TestEventPatternRequest) {
	request = &TestEventPatternRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "TestEventPattern", "", "")
	request.Method = requests.POST
	return
}

// CreateTestEventPatternResponse creates a response to parse from TestEventPattern response
func CreateTestEventPatternResponse() (response *TestEventPatternResponse) {
	response = &TestEventPatternResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
