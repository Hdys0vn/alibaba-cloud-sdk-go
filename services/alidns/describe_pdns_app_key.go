package alidns

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

// DescribePdnsAppKey invokes the alidns.DescribePdnsAppKey API synchronously
func (client *Client) DescribePdnsAppKey(request *DescribePdnsAppKeyRequest) (response *DescribePdnsAppKeyResponse, err error) {
	response = CreateDescribePdnsAppKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePdnsAppKeyWithChan invokes the alidns.DescribePdnsAppKey API asynchronously
func (client *Client) DescribePdnsAppKeyWithChan(request *DescribePdnsAppKeyRequest) (<-chan *DescribePdnsAppKeyResponse, <-chan error) {
	responseChan := make(chan *DescribePdnsAppKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePdnsAppKey(request)
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

// DescribePdnsAppKeyWithCallback invokes the alidns.DescribePdnsAppKey API asynchronously
func (client *Client) DescribePdnsAppKeyWithCallback(request *DescribePdnsAppKeyRequest, callback func(response *DescribePdnsAppKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePdnsAppKeyResponse
		var err error
		defer close(result)
		response, err = client.DescribePdnsAppKey(request)
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

// DescribePdnsAppKeyRequest is the request struct for api DescribePdnsAppKey
type DescribePdnsAppKeyRequest struct {
	*requests.RpcRequest
	AppKeyId string `position:"Query" name:"AppKeyId"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribePdnsAppKeyResponse is the response struct for api DescribePdnsAppKey
type DescribePdnsAppKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AppKey    AppKey `json:"AppKey" xml:"AppKey"`
}

// CreateDescribePdnsAppKeyRequest creates a request to invoke DescribePdnsAppKey API
func CreateDescribePdnsAppKeyRequest() (request *DescribePdnsAppKeyRequest) {
	request = &DescribePdnsAppKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribePdnsAppKey", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePdnsAppKeyResponse creates a response to parse from DescribePdnsAppKey response
func CreateDescribePdnsAppKeyResponse() (response *DescribePdnsAppKeyResponse) {
	response = &DescribePdnsAppKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
