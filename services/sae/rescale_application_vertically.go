package sae

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

// RescaleApplicationVertically invokes the sae.RescaleApplicationVertically API synchronously
func (client *Client) RescaleApplicationVertically(request *RescaleApplicationVerticallyRequest) (response *RescaleApplicationVerticallyResponse, err error) {
	response = CreateRescaleApplicationVerticallyResponse()
	err = client.DoAction(request, response)
	return
}

// RescaleApplicationVerticallyWithChan invokes the sae.RescaleApplicationVertically API asynchronously
func (client *Client) RescaleApplicationVerticallyWithChan(request *RescaleApplicationVerticallyRequest) (<-chan *RescaleApplicationVerticallyResponse, <-chan error) {
	responseChan := make(chan *RescaleApplicationVerticallyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RescaleApplicationVertically(request)
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

// RescaleApplicationVerticallyWithCallback invokes the sae.RescaleApplicationVertically API asynchronously
func (client *Client) RescaleApplicationVerticallyWithCallback(request *RescaleApplicationVerticallyRequest, callback func(response *RescaleApplicationVerticallyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RescaleApplicationVerticallyResponse
		var err error
		defer close(result)
		response, err = client.RescaleApplicationVertically(request)
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

// RescaleApplicationVerticallyRequest is the request struct for api RescaleApplicationVertically
type RescaleApplicationVerticallyRequest struct {
	*requests.RoaRequest
	Memory string `position:"Query" name:"Memory"`
	AppId  string `position:"Query" name:"AppId"`
	Cpu    string `position:"Query" name:"Cpu"`
}

// RescaleApplicationVerticallyResponse is the response struct for api RescaleApplicationVertically
type RescaleApplicationVerticallyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRescaleApplicationVerticallyRequest creates a request to invoke RescaleApplicationVertically API
func CreateRescaleApplicationVerticallyRequest() (request *RescaleApplicationVerticallyRequest) {
	request = &RescaleApplicationVerticallyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "RescaleApplicationVertically", "/pop/v1/sam/app/rescaleApplicationVertically", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRescaleApplicationVerticallyResponse creates a response to parse from RescaleApplicationVertically response
func CreateRescaleApplicationVerticallyResponse() (response *RescaleApplicationVerticallyResponse) {
	response = &RescaleApplicationVerticallyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
