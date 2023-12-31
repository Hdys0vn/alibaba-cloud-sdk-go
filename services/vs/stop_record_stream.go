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

// StopRecordStream invokes the vs.StopRecordStream API synchronously
func (client *Client) StopRecordStream(request *StopRecordStreamRequest) (response *StopRecordStreamResponse, err error) {
	response = CreateStopRecordStreamResponse()
	err = client.DoAction(request, response)
	return
}

// StopRecordStreamWithChan invokes the vs.StopRecordStream API asynchronously
func (client *Client) StopRecordStreamWithChan(request *StopRecordStreamRequest) (<-chan *StopRecordStreamResponse, <-chan error) {
	responseChan := make(chan *StopRecordStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopRecordStream(request)
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

// StopRecordStreamWithCallback invokes the vs.StopRecordStream API asynchronously
func (client *Client) StopRecordStreamWithCallback(request *StopRecordStreamRequest, callback func(response *StopRecordStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopRecordStreamResponse
		var err error
		defer close(result)
		response, err = client.StopRecordStream(request)
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

// StopRecordStreamRequest is the request struct for api StopRecordStream
type StopRecordStreamRequest struct {
	*requests.RpcRequest
	Id         string           `position:"Query" name:"Id"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	PlayDomain string           `position:"Query" name:"PlayDomain"`
	App        string           `position:"Query" name:"App"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Name       string           `position:"Query" name:"Name"`
}

// StopRecordStreamResponse is the response struct for api StopRecordStream
type StopRecordStreamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopRecordStreamRequest creates a request to invoke StopRecordStream API
func CreateStopRecordStreamRequest() (request *StopRecordStreamRequest) {
	request = &StopRecordStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "StopRecordStream", "", "")
	request.Method = requests.POST
	return
}

// CreateStopRecordStreamResponse creates a response to parse from StopRecordStream response
func CreateStopRecordStreamResponse() (response *StopRecordStreamResponse) {
	response = &StopRecordStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
