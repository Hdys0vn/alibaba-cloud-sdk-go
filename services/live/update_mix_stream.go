package live

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

// UpdateMixStream invokes the live.UpdateMixStream API synchronously
func (client *Client) UpdateMixStream(request *UpdateMixStreamRequest) (response *UpdateMixStreamResponse, err error) {
	response = CreateUpdateMixStreamResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMixStreamWithChan invokes the live.UpdateMixStream API asynchronously
func (client *Client) UpdateMixStreamWithChan(request *UpdateMixStreamRequest) (<-chan *UpdateMixStreamResponse, <-chan error) {
	responseChan := make(chan *UpdateMixStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMixStream(request)
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

// UpdateMixStreamWithCallback invokes the live.UpdateMixStream API asynchronously
func (client *Client) UpdateMixStreamWithCallback(request *UpdateMixStreamRequest, callback func(response *UpdateMixStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMixStreamResponse
		var err error
		defer close(result)
		response, err = client.UpdateMixStream(request)
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

// UpdateMixStreamRequest is the request struct for api UpdateMixStream
type UpdateMixStreamRequest struct {
	*requests.RpcRequest
	LayoutId        string           `position:"Query" name:"LayoutId"`
	MixStreamId     string           `position:"Query" name:"MixStreamId"`
	DomainName      string           `position:"Query" name:"DomainName"`
	InputStreamList string           `position:"Query" name:"InputStreamList"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateMixStreamResponse is the response struct for api UpdateMixStream
type UpdateMixStreamResponse struct {
	*responses.BaseResponse
	MixStreamId string `json:"MixStreamId" xml:"MixStreamId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateMixStreamRequest creates a request to invoke UpdateMixStream API
func CreateUpdateMixStreamRequest() (request *UpdateMixStreamRequest) {
	request = &UpdateMixStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateMixStream", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateMixStreamResponse creates a response to parse from UpdateMixStream response
func CreateUpdateMixStreamResponse() (response *UpdateMixStreamResponse) {
	response = &UpdateMixStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
