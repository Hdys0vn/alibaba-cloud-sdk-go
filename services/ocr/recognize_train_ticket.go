package ocr

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

// RecognizeTrainTicket invokes the ocr.RecognizeTrainTicket API synchronously
func (client *Client) RecognizeTrainTicket(request *RecognizeTrainTicketRequest) (response *RecognizeTrainTicketResponse, err error) {
	response = CreateRecognizeTrainTicketResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeTrainTicketWithChan invokes the ocr.RecognizeTrainTicket API asynchronously
func (client *Client) RecognizeTrainTicketWithChan(request *RecognizeTrainTicketRequest) (<-chan *RecognizeTrainTicketResponse, <-chan error) {
	responseChan := make(chan *RecognizeTrainTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeTrainTicket(request)
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

// RecognizeTrainTicketWithCallback invokes the ocr.RecognizeTrainTicket API asynchronously
func (client *Client) RecognizeTrainTicketWithCallback(request *RecognizeTrainTicketRequest, callback func(response *RecognizeTrainTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeTrainTicketResponse
		var err error
		defer close(result)
		response, err = client.RecognizeTrainTicket(request)
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

// RecognizeTrainTicketRequest is the request struct for api RecognizeTrainTicket
type RecognizeTrainTicketRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Body" name:"ImageType"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// RecognizeTrainTicketResponse is the response struct for api RecognizeTrainTicket
type RecognizeTrainTicketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeTrainTicketRequest creates a request to invoke RecognizeTrainTicket API
func CreateRecognizeTrainTicketRequest() (request *RecognizeTrainTicketRequest) {
	request = &RecognizeTrainTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeTrainTicket", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeTrainTicketResponse creates a response to parse from RecognizeTrainTicket response
func CreateRecognizeTrainTicketResponse() (response *RecognizeTrainTicketResponse) {
	response = &RecognizeTrainTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}