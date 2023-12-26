package vs

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

// StopStream invokes the vs.StopStream API synchronously
func (client *Client) StopStream(request *StopStreamRequest) (response *StopStreamResponse, err error) {
	response = CreateStopStreamResponse()
	err = client.DoAction(request, response)
	return
}

// StopStreamWithChan invokes the vs.StopStream API asynchronously
func (client *Client) StopStreamWithChan(request *StopStreamRequest) (<-chan *StopStreamResponse, <-chan error) {
	responseChan := make(chan *StopStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopStream(request)
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

// StopStreamWithCallback invokes the vs.StopStream API asynchronously
func (client *Client) StopStreamWithCallback(request *StopStreamRequest, callback func(response *StopStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopStreamResponse
		var err error
		defer close(result)
		response, err = client.StopStream(request)
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

// StopStreamRequest is the request struct for api StopStream
type StopStreamRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Id        string           `position:"Query" name:"Id"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	Name      string           `position:"Query" name:"Name"`
}

// StopStreamResponse is the response struct for api StopStream
type StopStreamResponse struct {
	*responses.BaseResponse
	Id        string `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopStreamRequest creates a request to invoke StopStream API
func CreateStopStreamRequest() (request *StopStreamRequest) {
	request = &StopStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "StopStream", "", "")
	request.Method = requests.POST
	return
}

// CreateStopStreamResponse creates a response to parse from StopStream response
func CreateStopStreamResponse() (response *StopStreamResponse) {
	response = &StopStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
