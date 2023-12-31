package mts

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

// RefreshCdnDomainConfigsCache invokes the mts.RefreshCdnDomainConfigsCache API synchronously
func (client *Client) RefreshCdnDomainConfigsCache(request *RefreshCdnDomainConfigsCacheRequest) (response *RefreshCdnDomainConfigsCacheResponse, err error) {
	response = CreateRefreshCdnDomainConfigsCacheResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshCdnDomainConfigsCacheWithChan invokes the mts.RefreshCdnDomainConfigsCache API asynchronously
func (client *Client) RefreshCdnDomainConfigsCacheWithChan(request *RefreshCdnDomainConfigsCacheRequest) (<-chan *RefreshCdnDomainConfigsCacheResponse, <-chan error) {
	responseChan := make(chan *RefreshCdnDomainConfigsCacheResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshCdnDomainConfigsCache(request)
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

// RefreshCdnDomainConfigsCacheWithCallback invokes the mts.RefreshCdnDomainConfigsCache API asynchronously
func (client *Client) RefreshCdnDomainConfigsCacheWithCallback(request *RefreshCdnDomainConfigsCacheRequest, callback func(response *RefreshCdnDomainConfigsCacheResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshCdnDomainConfigsCacheResponse
		var err error
		defer close(result)
		response, err = client.RefreshCdnDomainConfigsCache(request)
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

// RefreshCdnDomainConfigsCacheRequest is the request struct for api RefreshCdnDomainConfigsCache
type RefreshCdnDomainConfigsCacheRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Domains              string `position:"Query" name:"Domains"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

// RefreshCdnDomainConfigsCacheResponse is the response struct for api RefreshCdnDomainConfigsCache
type RefreshCdnDomainConfigsCacheResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	SucessDomains SucessDomains `json:"SucessDomains" xml:"SucessDomains"`
	FailedDomains FailedDomains `json:"FailedDomains" xml:"FailedDomains"`
}

// CreateRefreshCdnDomainConfigsCacheRequest creates a request to invoke RefreshCdnDomainConfigsCache API
func CreateRefreshCdnDomainConfigsCacheRequest() (request *RefreshCdnDomainConfigsCacheRequest) {
	request = &RefreshCdnDomainConfigsCacheRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "RefreshCdnDomainConfigsCache", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRefreshCdnDomainConfigsCacheResponse creates a response to parse from RefreshCdnDomainConfigsCache response
func CreateRefreshCdnDomainConfigsCacheResponse() (response *RefreshCdnDomainConfigsCacheResponse) {
	response = &RefreshCdnDomainConfigsCacheResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
