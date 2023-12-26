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

// RecognizeCarNum invokes the iot.RecognizeCarNum API synchronously
func (client *Client) RecognizeCarNum(request *RecognizeCarNumRequest) (response *RecognizeCarNumResponse, err error) {
	response = CreateRecognizeCarNumResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeCarNumWithChan invokes the iot.RecognizeCarNum API asynchronously
func (client *Client) RecognizeCarNumWithChan(request *RecognizeCarNumRequest) (<-chan *RecognizeCarNumResponse, <-chan error) {
	responseChan := make(chan *RecognizeCarNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeCarNum(request)
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

// RecognizeCarNumWithCallback invokes the iot.RecognizeCarNum API asynchronously
func (client *Client) RecognizeCarNumWithCallback(request *RecognizeCarNumRequest, callback func(response *RecognizeCarNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeCarNumResponse
		var err error
		defer close(result)
		response, err = client.RecognizeCarNum(request)
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

// RecognizeCarNumRequest is the request struct for api RecognizeCarNum
type RecognizeCarNumRequest struct {
	*requests.RpcRequest
	Url         string `position:"Query" name:"Url"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// RecognizeCarNumResponse is the response struct for api RecognizeCarNum
type RecognizeCarNumResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         string `json:"Data" xml:"Data"`
}

// CreateRecognizeCarNumRequest creates a request to invoke RecognizeCarNum API
func CreateRecognizeCarNumRequest() (request *RecognizeCarNumRequest) {
	request = &RecognizeCarNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "RecognizeCarNum", "", "")
	request.Method = requests.POST
	return
}

// CreateRecognizeCarNumResponse creates a response to parse from RecognizeCarNum response
func CreateRecognizeCarNumResponse() (response *RecognizeCarNumResponse) {
	response = &RecognizeCarNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
