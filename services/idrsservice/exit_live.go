package idrsservice

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

// ExitLive invokes the idrsservice.ExitLive API synchronously
func (client *Client) ExitLive(request *ExitLiveRequest) (response *ExitLiveResponse, err error) {
	response = CreateExitLiveResponse()
	err = client.DoAction(request, response)
	return
}

// ExitLiveWithChan invokes the idrsservice.ExitLive API asynchronously
func (client *Client) ExitLiveWithChan(request *ExitLiveRequest) (<-chan *ExitLiveResponse, <-chan error) {
	responseChan := make(chan *ExitLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExitLive(request)
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

// ExitLiveWithCallback invokes the idrsservice.ExitLive API asynchronously
func (client *Client) ExitLiveWithCallback(request *ExitLiveRequest, callback func(response *ExitLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExitLiveResponse
		var err error
		defer close(result)
		response, err = client.ExitLive(request)
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

// ExitLiveRequest is the request struct for api ExitLive
type ExitLiveRequest struct {
	*requests.RpcRequest
	Channel string `position:"Query" name:"Channel"`
	UserId  string `position:"Query" name:"UserId"`
	RtcCode string `position:"Query" name:"RtcCode"`
}

// ExitLiveResponse is the response struct for api ExitLive
type ExitLiveResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateExitLiveRequest creates a request to invoke ExitLive API
func CreateExitLiveRequest() (request *ExitLiveRequest) {
	request = &ExitLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "ExitLive", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExitLiveResponse creates a response to parse from ExitLive response
func CreateExitLiveResponse() (response *ExitLiveResponse) {
	response = &ExitLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}