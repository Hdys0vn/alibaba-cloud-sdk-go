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

// DescribeIspFlushCacheRemainQuota invokes the alidns.DescribeIspFlushCacheRemainQuota API synchronously
func (client *Client) DescribeIspFlushCacheRemainQuota(request *DescribeIspFlushCacheRemainQuotaRequest) (response *DescribeIspFlushCacheRemainQuotaResponse, err error) {
	response = CreateDescribeIspFlushCacheRemainQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIspFlushCacheRemainQuotaWithChan invokes the alidns.DescribeIspFlushCacheRemainQuota API asynchronously
func (client *Client) DescribeIspFlushCacheRemainQuotaWithChan(request *DescribeIspFlushCacheRemainQuotaRequest) (<-chan *DescribeIspFlushCacheRemainQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeIspFlushCacheRemainQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIspFlushCacheRemainQuota(request)
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

// DescribeIspFlushCacheRemainQuotaWithCallback invokes the alidns.DescribeIspFlushCacheRemainQuota API asynchronously
func (client *Client) DescribeIspFlushCacheRemainQuotaWithCallback(request *DescribeIspFlushCacheRemainQuotaRequest, callback func(response *DescribeIspFlushCacheRemainQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIspFlushCacheRemainQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeIspFlushCacheRemainQuota(request)
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

// DescribeIspFlushCacheRemainQuotaRequest is the request struct for api DescribeIspFlushCacheRemainQuota
type DescribeIspFlushCacheRemainQuotaRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeIspFlushCacheRemainQuotaResponse is the response struct for api DescribeIspFlushCacheRemainQuota
type DescribeIspFlushCacheRemainQuotaResponse struct {
	*responses.BaseResponse
	TelecomRemainQuota int    `json:"TelecomRemainQuota" xml:"TelecomRemainQuota"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeIspFlushCacheRemainQuotaRequest creates a request to invoke DescribeIspFlushCacheRemainQuota API
func CreateDescribeIspFlushCacheRemainQuotaRequest() (request *DescribeIspFlushCacheRemainQuotaRequest) {
	request = &DescribeIspFlushCacheRemainQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeIspFlushCacheRemainQuota", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIspFlushCacheRemainQuotaResponse creates a response to parse from DescribeIspFlushCacheRemainQuota response
func CreateDescribeIspFlushCacheRemainQuotaResponse() (response *DescribeIspFlushCacheRemainQuotaResponse) {
	response = &DescribeIspFlushCacheRemainQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}