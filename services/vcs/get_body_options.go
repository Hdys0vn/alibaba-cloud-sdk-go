package vcs

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

// GetBodyOptions invokes the vcs.GetBodyOptions API synchronously
func (client *Client) GetBodyOptions(request *GetBodyOptionsRequest) (response *GetBodyOptionsResponse, err error) {
	response = CreateGetBodyOptionsResponse()
	err = client.DoAction(request, response)
	return
}

// GetBodyOptionsWithChan invokes the vcs.GetBodyOptions API asynchronously
func (client *Client) GetBodyOptionsWithChan(request *GetBodyOptionsRequest) (<-chan *GetBodyOptionsResponse, <-chan error) {
	responseChan := make(chan *GetBodyOptionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBodyOptions(request)
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

// GetBodyOptionsWithCallback invokes the vcs.GetBodyOptions API asynchronously
func (client *Client) GetBodyOptionsWithCallback(request *GetBodyOptionsRequest, callback func(response *GetBodyOptionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBodyOptionsResponse
		var err error
		defer close(result)
		response, err = client.GetBodyOptions(request)
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

// GetBodyOptionsRequest is the request struct for api GetBodyOptions
type GetBodyOptionsRequest struct {
	*requests.RpcRequest
	CorpId string `position:"Body" name:"CorpId"`
}

// GetBodyOptionsResponse is the response struct for api GetBodyOptions
type GetBodyOptionsResponse struct {
	*responses.BaseResponse
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetBodyOptionsRequest creates a request to invoke GetBodyOptions API
func CreateGetBodyOptionsRequest() (request *GetBodyOptionsRequest) {
	request = &GetBodyOptionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetBodyOptions", "", "")
	request.Method = requests.POST
	return
}

// CreateGetBodyOptionsResponse creates a response to parse from GetBodyOptions response
func CreateGetBodyOptionsResponse() (response *GetBodyOptionsResponse) {
	response = &GetBodyOptionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
