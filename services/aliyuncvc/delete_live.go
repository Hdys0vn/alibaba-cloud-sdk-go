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

// DeleteLive invokes the aliyuncvc.DeleteLive API synchronously
func (client *Client) DeleteLive(request *DeleteLiveRequest) (response *DeleteLiveResponse, err error) {
	response = CreateDeleteLiveResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveWithChan invokes the aliyuncvc.DeleteLive API asynchronously
func (client *Client) DeleteLiveWithChan(request *DeleteLiveRequest) (<-chan *DeleteLiveResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLive(request)
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

// DeleteLiveWithCallback invokes the aliyuncvc.DeleteLive API asynchronously
func (client *Client) DeleteLiveWithCallback(request *DeleteLiveRequest, callback func(response *DeleteLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveResponse
		var err error
		defer close(result)
		response, err = client.DeleteLive(request)
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

// DeleteLiveRequest is the request struct for api DeleteLive
type DeleteLiveRequest struct {
	*requests.RpcRequest
	LiveUUID string `position:"Body" name:"LiveUUID"`
	UserId   string `position:"Body" name:"UserId"`
}

// DeleteLiveResponse is the response struct for api DeleteLive
type DeleteLiveResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveRequest creates a request to invoke DeleteLive API
func CreateDeleteLiveRequest() (request *DeleteLiveRequest) {
	request = &DeleteLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "DeleteLive", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveResponse creates a response to parse from DeleteLive response
func CreateDeleteLiveResponse() (response *DeleteLiveResponse) {
	response = &DeleteLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
