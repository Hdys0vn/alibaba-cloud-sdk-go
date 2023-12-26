package ddoscoo

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

// EmptyAutoCcWhitelist invokes the ddoscoo.EmptyAutoCcWhitelist API synchronously
func (client *Client) EmptyAutoCcWhitelist(request *EmptyAutoCcWhitelistRequest) (response *EmptyAutoCcWhitelistResponse, err error) {
	response = CreateEmptyAutoCcWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// EmptyAutoCcWhitelistWithChan invokes the ddoscoo.EmptyAutoCcWhitelist API asynchronously
func (client *Client) EmptyAutoCcWhitelistWithChan(request *EmptyAutoCcWhitelistRequest) (<-chan *EmptyAutoCcWhitelistResponse, <-chan error) {
	responseChan := make(chan *EmptyAutoCcWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EmptyAutoCcWhitelist(request)
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

// EmptyAutoCcWhitelistWithCallback invokes the ddoscoo.EmptyAutoCcWhitelist API asynchronously
func (client *Client) EmptyAutoCcWhitelistWithCallback(request *EmptyAutoCcWhitelistRequest, callback func(response *EmptyAutoCcWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EmptyAutoCcWhitelistResponse
		var err error
		defer close(result)
		response, err = client.EmptyAutoCcWhitelist(request)
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

// EmptyAutoCcWhitelistRequest is the request struct for api EmptyAutoCcWhitelist
type EmptyAutoCcWhitelistRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
}

// EmptyAutoCcWhitelistResponse is the response struct for api EmptyAutoCcWhitelist
type EmptyAutoCcWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEmptyAutoCcWhitelistRequest creates a request to invoke EmptyAutoCcWhitelist API
func CreateEmptyAutoCcWhitelistRequest() (request *EmptyAutoCcWhitelistRequest) {
	request = &EmptyAutoCcWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "EmptyAutoCcWhitelist", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEmptyAutoCcWhitelistResponse creates a response to parse from EmptyAutoCcWhitelist response
func CreateEmptyAutoCcWhitelistResponse() (response *EmptyAutoCcWhitelistResponse) {
	response = &EmptyAutoCcWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
