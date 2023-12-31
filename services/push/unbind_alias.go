package push

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

// UnbindAlias invokes the push.UnbindAlias API synchronously
func (client *Client) UnbindAlias(request *UnbindAliasRequest) (response *UnbindAliasResponse, err error) {
	response = CreateUnbindAliasResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindAliasWithChan invokes the push.UnbindAlias API asynchronously
func (client *Client) UnbindAliasWithChan(request *UnbindAliasRequest) (<-chan *UnbindAliasResponse, <-chan error) {
	responseChan := make(chan *UnbindAliasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindAlias(request)
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

// UnbindAliasWithCallback invokes the push.UnbindAlias API asynchronously
func (client *Client) UnbindAliasWithCallback(request *UnbindAliasRequest, callback func(response *UnbindAliasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindAliasResponse
		var err error
		defer close(result)
		response, err = client.UnbindAlias(request)
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

// UnbindAliasRequest is the request struct for api UnbindAlias
type UnbindAliasRequest struct {
	*requests.RpcRequest
	DeviceId  string           `position:"Query" name:"DeviceId"`
	AliasName string           `position:"Query" name:"AliasName"`
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	UnbindAll requests.Boolean `position:"Query" name:"UnbindAll"`
}

// UnbindAliasResponse is the response struct for api UnbindAlias
type UnbindAliasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindAliasRequest creates a request to invoke UnbindAlias API
func CreateUnbindAliasRequest() (request *UnbindAliasRequest) {
	request = &UnbindAliasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "UnbindAlias", "", "")
	request.Method = requests.POST
	return
}

// CreateUnbindAliasResponse creates a response to parse from UnbindAlias response
func CreateUnbindAliasResponse() (response *UnbindAliasResponse) {
	response = &UnbindAliasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
