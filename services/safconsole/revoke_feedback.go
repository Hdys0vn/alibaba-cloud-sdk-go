package safconsole

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

// RevokeFeedback invokes the safconsole.RevokeFeedback API synchronously
func (client *Client) RevokeFeedback(request *RevokeFeedbackRequest) (response *RevokeFeedbackResponse, err error) {
	response = CreateRevokeFeedbackResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeFeedbackWithChan invokes the safconsole.RevokeFeedback API asynchronously
func (client *Client) RevokeFeedbackWithChan(request *RevokeFeedbackRequest) (<-chan *RevokeFeedbackResponse, <-chan error) {
	responseChan := make(chan *RevokeFeedbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeFeedback(request)
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

// RevokeFeedbackWithCallback invokes the safconsole.RevokeFeedback API asynchronously
func (client *Client) RevokeFeedbackWithCallback(request *RevokeFeedbackRequest, callback func(response *RevokeFeedbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeFeedbackResponse
		var err error
		defer close(result)
		response, err = client.RevokeFeedback(request)
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

// RevokeFeedbackRequest is the request struct for api RevokeFeedback
type RevokeFeedbackRequest struct {
	*requests.RpcRequest
	SampleType string `position:"Body" name:"SampleType"`
	Value      string `position:"Body" name:"Value"`
}

// RevokeFeedbackResponse is the response struct for api RevokeFeedback
type RevokeFeedbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateRevokeFeedbackRequest creates a request to invoke RevokeFeedback API
func CreateRevokeFeedbackRequest() (request *RevokeFeedbackRequest) {
	request = &RevokeFeedbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("safconsole", "2021-01-12", "RevokeFeedback", "saf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRevokeFeedbackResponse creates a response to parse from RevokeFeedback response
func CreateRevokeFeedbackResponse() (response *RevokeFeedbackResponse) {
	response = &RevokeFeedbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
