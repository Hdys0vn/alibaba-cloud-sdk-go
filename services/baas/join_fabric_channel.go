package baas

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

// JoinFabricChannel invokes the baas.JoinFabricChannel API synchronously
// api document: https://help.aliyun.com/api/baas/joinfabricchannel.html
func (client *Client) JoinFabricChannel(request *JoinFabricChannelRequest) (response *JoinFabricChannelResponse, err error) {
	response = CreateJoinFabricChannelResponse()
	err = client.DoAction(request, response)
	return
}

// JoinFabricChannelWithChan invokes the baas.JoinFabricChannel API asynchronously
// api document: https://help.aliyun.com/api/baas/joinfabricchannel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinFabricChannelWithChan(request *JoinFabricChannelRequest) (<-chan *JoinFabricChannelResponse, <-chan error) {
	responseChan := make(chan *JoinFabricChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinFabricChannel(request)
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

// JoinFabricChannelWithCallback invokes the baas.JoinFabricChannel API asynchronously
// api document: https://help.aliyun.com/api/baas/joinfabricchannel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinFabricChannelWithCallback(request *JoinFabricChannelRequest, callback func(response *JoinFabricChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinFabricChannelResponse
		var err error
		defer close(result)
		response, err = client.JoinFabricChannel(request)
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

// JoinFabricChannelRequest is the request struct for api JoinFabricChannel
type JoinFabricChannelRequest struct {
	*requests.RpcRequest
	Do        string `position:"Query" name:"Do"`
	Location  string `position:"Body" name:"Location"`
	ChannelId string `position:"Query" name:"ChannelId"`
}

// JoinFabricChannelResponse is the response struct for api JoinFabricChannel
type JoinFabricChannelResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Success   bool         `json:"Success" xml:"Success"`
	ErrorCode int          `json:"ErrorCode" xml:"ErrorCode"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateJoinFabricChannelRequest creates a request to invoke JoinFabricChannel API
func CreateJoinFabricChannelRequest() (request *JoinFabricChannelRequest) {
	request = &JoinFabricChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "JoinFabricChannel", "baas", "openAPI")
	return
}

// CreateJoinFabricChannelResponse creates a response to parse from JoinFabricChannel response
func CreateJoinFabricChannelResponse() (response *JoinFabricChannelResponse) {
	response = &JoinFabricChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
