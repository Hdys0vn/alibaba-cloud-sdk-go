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

// RestartCaster invokes the live.RestartCaster API synchronously
func (client *Client) RestartCaster(request *RestartCasterRequest) (response *RestartCasterResponse, err error) {
	response = CreateRestartCasterResponse()
	err = client.DoAction(request, response)
	return
}

// RestartCasterWithChan invokes the live.RestartCaster API asynchronously
func (client *Client) RestartCasterWithChan(request *RestartCasterRequest) (<-chan *RestartCasterResponse, <-chan error) {
	responseChan := make(chan *RestartCasterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartCaster(request)
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

// RestartCasterWithCallback invokes the live.RestartCaster API asynchronously
func (client *Client) RestartCasterWithCallback(request *RestartCasterRequest, callback func(response *RestartCasterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartCasterResponse
		var err error
		defer close(result)
		response, err = client.RestartCaster(request)
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

// RestartCasterRequest is the request struct for api RestartCaster
type RestartCasterRequest struct {
	*requests.RpcRequest
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// RestartCasterResponse is the response struct for api RestartCaster
type RestartCasterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestartCasterRequest creates a request to invoke RestartCaster API
func CreateRestartCasterRequest() (request *RestartCasterRequest) {
	request = &RestartCasterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "RestartCaster", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestartCasterResponse creates a response to parse from RestartCaster response
func CreateRestartCasterResponse() (response *RestartCasterResponse) {
	response = &RestartCasterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
