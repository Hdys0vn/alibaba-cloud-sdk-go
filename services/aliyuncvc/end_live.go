package aliyuncvc

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

// EndLive invokes the aliyuncvc.EndLive API synchronously
func (client *Client) EndLive(request *EndLiveRequest) (response *EndLiveResponse, err error) {
	response = CreateEndLiveResponse()
	err = client.DoAction(request, response)
	return
}

// EndLiveWithChan invokes the aliyuncvc.EndLive API asynchronously
func (client *Client) EndLiveWithChan(request *EndLiveRequest) (<-chan *EndLiveResponse, <-chan error) {
	responseChan := make(chan *EndLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EndLive(request)
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

// EndLiveWithCallback invokes the aliyuncvc.EndLive API asynchronously
func (client *Client) EndLiveWithCallback(request *EndLiveRequest, callback func(response *EndLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EndLiveResponse
		var err error
		defer close(result)
		response, err = client.EndLive(request)
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

// EndLiveRequest is the request struct for api EndLive
type EndLiveRequest struct {
	*requests.RpcRequest
	LiveUUID string `position:"Body" name:"LiveUUID"`
	UserId   string `position:"Body" name:"UserId"`
}

// EndLiveResponse is the response struct for api EndLive
type EndLiveResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEndLiveRequest creates a request to invoke EndLive API
func CreateEndLiveRequest() (request *EndLiveRequest) {
	request = &EndLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "EndLive", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEndLiveResponse creates a response to parse from EndLive response
func CreateEndLiveResponse() (response *EndLiveResponse) {
	response = &EndLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}