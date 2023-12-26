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

// ResumeVsStream invokes the vs.ResumeVsStream API synchronously
func (client *Client) ResumeVsStream(request *ResumeVsStreamRequest) (response *ResumeVsStreamResponse, err error) {
	response = CreateResumeVsStreamResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeVsStreamWithChan invokes the vs.ResumeVsStream API asynchronously
func (client *Client) ResumeVsStreamWithChan(request *ResumeVsStreamRequest) (<-chan *ResumeVsStreamResponse, <-chan error) {
	responseChan := make(chan *ResumeVsStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeVsStream(request)
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

// ResumeVsStreamWithCallback invokes the vs.ResumeVsStream API asynchronously
func (client *Client) ResumeVsStreamWithCallback(request *ResumeVsStreamRequest, callback func(response *ResumeVsStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeVsStreamResponse
		var err error
		defer close(result)
		response, err = client.ResumeVsStream(request)
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

// ResumeVsStreamRequest is the request struct for api ResumeVsStream
type ResumeVsStreamRequest struct {
	*requests.RpcRequest
	AppName             string           `position:"Query" name:"AppName"`
	StreamName          string           `position:"Query" name:"StreamName"`
	ShowLog             string           `position:"Query" name:"ShowLog"`
	ControlStreamAction string           `position:"Query" name:"ControlStreamAction"`
	LiveStreamType      string           `position:"Query" name:"LiveStreamType"`
	DomainName          string           `position:"Query" name:"DomainName"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
}

// ResumeVsStreamResponse is the response struct for api ResumeVsStream
type ResumeVsStreamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResumeVsStreamRequest creates a request to invoke ResumeVsStream API
func CreateResumeVsStreamRequest() (request *ResumeVsStreamRequest) {
	request = &ResumeVsStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ResumeVsStream", "", "")
	request.Method = requests.POST
	return
}

// CreateResumeVsStreamResponse creates a response to parse from ResumeVsStream response
func CreateResumeVsStreamResponse() (response *ResumeVsStreamResponse) {
	response = &ResumeVsStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
