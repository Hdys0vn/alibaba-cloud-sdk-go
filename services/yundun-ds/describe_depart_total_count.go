package yundun_ds

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

// DescribeDepartTotalCount invokes the yundun_ds.DescribeDepartTotalCount API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedeparttotalcount.html
func (client *Client) DescribeDepartTotalCount(request *DescribeDepartTotalCountRequest) (response *DescribeDepartTotalCountResponse, err error) {
	response = CreateDescribeDepartTotalCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDepartTotalCountWithChan invokes the yundun_ds.DescribeDepartTotalCount API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedeparttotalcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDepartTotalCountWithChan(request *DescribeDepartTotalCountRequest) (<-chan *DescribeDepartTotalCountResponse, <-chan error) {
	responseChan := make(chan *DescribeDepartTotalCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDepartTotalCount(request)
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

// DescribeDepartTotalCountWithCallback invokes the yundun_ds.DescribeDepartTotalCount API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedeparttotalcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDepartTotalCountWithCallback(request *DescribeDepartTotalCountRequest, callback func(response *DescribeDepartTotalCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDepartTotalCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeDepartTotalCount(request)
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

// DescribeDepartTotalCountRequest is the request struct for api DescribeDepartTotalCount
type DescribeDepartTotalCountRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DescribeDepartTotalCountResponse is the response struct for api DescribeDepartTotalCount
type DescribeDepartTotalCountResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	DepartCount DepartCount `json:"DepartCount" xml:"DepartCount"`
}

// CreateDescribeDepartTotalCountRequest creates a request to invoke DescribeDepartTotalCount API
func CreateDescribeDepartTotalCountRequest() (request *DescribeDepartTotalCountRequest) {
	request = &DescribeDepartTotalCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeDepartTotalCount", "sddp", "openAPI")
	return
}

// CreateDescribeDepartTotalCountResponse creates a response to parse from DescribeDepartTotalCount response
func CreateDescribeDepartTotalCountResponse() (response *DescribeDepartTotalCountResponse) {
	response = &DescribeDepartTotalCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
