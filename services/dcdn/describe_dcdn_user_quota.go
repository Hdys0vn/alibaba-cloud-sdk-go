package dcdn

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

// DescribeDcdnUserQuota invokes the dcdn.DescribeDcdnUserQuota API synchronously
func (client *Client) DescribeDcdnUserQuota(request *DescribeDcdnUserQuotaRequest) (response *DescribeDcdnUserQuotaResponse, err error) {
	response = CreateDescribeDcdnUserQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnUserQuotaWithChan invokes the dcdn.DescribeDcdnUserQuota API asynchronously
func (client *Client) DescribeDcdnUserQuotaWithChan(request *DescribeDcdnUserQuotaRequest) (<-chan *DescribeDcdnUserQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnUserQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnUserQuota(request)
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

// DescribeDcdnUserQuotaWithCallback invokes the dcdn.DescribeDcdnUserQuota API asynchronously
func (client *Client) DescribeDcdnUserQuotaWithCallback(request *DescribeDcdnUserQuotaRequest, callback func(response *DescribeDcdnUserQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnUserQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnUserQuota(request)
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

// DescribeDcdnUserQuotaRequest is the request struct for api DescribeDcdnUserQuota
type DescribeDcdnUserQuotaRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnUserQuotaResponse is the response struct for api DescribeDcdnUserQuota
type DescribeDcdnUserQuotaResponse struct {
	*responses.BaseResponse
	BlockQuota         int    `json:"BlockQuota" xml:"BlockQuota"`
	RefreshUrlRemain   int    `json:"RefreshUrlRemain" xml:"RefreshUrlRemain"`
	DomainQuota        int    `json:"DomainQuota" xml:"DomainQuota"`
	BlockRemain        int    `json:"BlockRemain" xml:"BlockRemain"`
	PreloadRemain      int    `json:"PreloadRemain" xml:"PreloadRemain"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	RefreshUrlQuota    int    `json:"RefreshUrlQuota" xml:"RefreshUrlQuota"`
	PreloadQuota       int    `json:"PreloadQuota" xml:"PreloadQuota"`
	RefreshDirQuota    int    `json:"RefreshDirQuota" xml:"RefreshDirQuota"`
	RefreshDirRemain   int    `json:"RefreshDirRemain" xml:"RefreshDirRemain"`
	IgnoreParamsQuota  int    `json:"IgnoreParamsQuota" xml:"IgnoreParamsQuota"`
	IgnoreParamsRemain int    `json:"IgnoreParamsRemain" xml:"IgnoreParamsRemain"`
}

// CreateDescribeDcdnUserQuotaRequest creates a request to invoke DescribeDcdnUserQuota API
func CreateDescribeDcdnUserQuotaRequest() (request *DescribeDcdnUserQuotaRequest) {
	request = &DescribeDcdnUserQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserQuota", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnUserQuotaResponse creates a response to parse from DescribeDcdnUserQuota response
func CreateDescribeDcdnUserQuotaResponse() (response *DescribeDcdnUserQuotaResponse) {
	response = &DescribeDcdnUserQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}