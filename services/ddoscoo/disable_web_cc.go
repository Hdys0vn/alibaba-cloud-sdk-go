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

// DisableWebCC invokes the ddoscoo.DisableWebCC API synchronously
func (client *Client) DisableWebCC(request *DisableWebCCRequest) (response *DisableWebCCResponse, err error) {
	response = CreateDisableWebCCResponse()
	err = client.DoAction(request, response)
	return
}

// DisableWebCCWithChan invokes the ddoscoo.DisableWebCC API asynchronously
func (client *Client) DisableWebCCWithChan(request *DisableWebCCRequest) (<-chan *DisableWebCCResponse, <-chan error) {
	responseChan := make(chan *DisableWebCCResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableWebCC(request)
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

// DisableWebCCWithCallback invokes the ddoscoo.DisableWebCC API asynchronously
func (client *Client) DisableWebCCWithCallback(request *DisableWebCCRequest, callback func(response *DisableWebCCResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableWebCCResponse
		var err error
		defer close(result)
		response, err = client.DisableWebCC(request)
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

// DisableWebCCRequest is the request struct for api DisableWebCC
type DisableWebCCRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// DisableWebCCResponse is the response struct for api DisableWebCC
type DisableWebCCResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableWebCCRequest creates a request to invoke DisableWebCC API
func CreateDisableWebCCRequest() (request *DisableWebCCRequest) {
	request = &DisableWebCCRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DisableWebCC", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableWebCCResponse creates a response to parse from DisableWebCC response
func CreateDisableWebCCResponse() (response *DisableWebCCResponse) {
	response = &DisableWebCCResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
