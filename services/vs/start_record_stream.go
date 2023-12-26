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

// StartRecordStream invokes the vs.StartRecordStream API synchronously
func (client *Client) StartRecordStream(request *StartRecordStreamRequest) (response *StartRecordStreamResponse, err error) {
	response = CreateStartRecordStreamResponse()
	err = client.DoAction(request, response)
	return
}

// StartRecordStreamWithChan invokes the vs.StartRecordStream API asynchronously
func (client *Client) StartRecordStreamWithChan(request *StartRecordStreamRequest) (<-chan *StartRecordStreamResponse, <-chan error) {
	responseChan := make(chan *StartRecordStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartRecordStream(request)
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

// StartRecordStreamWithCallback invokes the vs.StartRecordStream API asynchronously
func (client *Client) StartRecordStreamWithCallback(request *StartRecordStreamRequest, callback func(response *StartRecordStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartRecordStreamResponse
		var err error
		defer close(result)
		response, err = client.StartRecordStream(request)
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

// StartRecordStreamRequest is the request struct for api StartRecordStream
type StartRecordStreamRequest struct {
	*requests.RpcRequest
	Id         string           `position:"Query" name:"Id"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	PlayDomain string           `position:"Query" name:"PlayDomain"`
	App        string           `position:"Query" name:"App"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Name       string           `position:"Query" name:"Name"`
}

// StartRecordStreamResponse is the response struct for api StartRecordStream
type StartRecordStreamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartRecordStreamRequest creates a request to invoke StartRecordStream API
func CreateStartRecordStreamRequest() (request *StartRecordStreamRequest) {
	request = &StartRecordStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "StartRecordStream", "", "")
	request.Method = requests.POST
	return
}

// CreateStartRecordStreamResponse creates a response to parse from StartRecordStream response
func CreateStartRecordStreamResponse() (response *StartRecordStreamResponse) {
	response = &StartRecordStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}