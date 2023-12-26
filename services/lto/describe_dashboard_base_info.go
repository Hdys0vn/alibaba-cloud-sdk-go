package lto

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

// DescribeDashboardBaseInfo invokes the lto.DescribeDashboardBaseInfo API synchronously
func (client *Client) DescribeDashboardBaseInfo(request *DescribeDashboardBaseInfoRequest) (response *DescribeDashboardBaseInfoResponse, err error) {
	response = CreateDescribeDashboardBaseInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDashboardBaseInfoWithChan invokes the lto.DescribeDashboardBaseInfo API asynchronously
func (client *Client) DescribeDashboardBaseInfoWithChan(request *DescribeDashboardBaseInfoRequest) (<-chan *DescribeDashboardBaseInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDashboardBaseInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDashboardBaseInfo(request)
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

// DescribeDashboardBaseInfoWithCallback invokes the lto.DescribeDashboardBaseInfo API asynchronously
func (client *Client) DescribeDashboardBaseInfoWithCallback(request *DescribeDashboardBaseInfoRequest, callback func(response *DescribeDashboardBaseInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDashboardBaseInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDashboardBaseInfo(request)
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

// DescribeDashboardBaseInfoRequest is the request struct for api DescribeDashboardBaseInfo
type DescribeDashboardBaseInfoRequest struct {
	*requests.RpcRequest
}

// DescribeDashboardBaseInfoResponse is the response struct for api DescribeDashboardBaseInfo
type DescribeDashboardBaseInfoResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDashboardBaseInfoRequest creates a request to invoke DescribeDashboardBaseInfo API
func CreateDescribeDashboardBaseInfoRequest() (request *DescribeDashboardBaseInfoRequest) {
	request = &DescribeDashboardBaseInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "DescribeDashboardBaseInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDashboardBaseInfoResponse creates a response to parse from DescribeDashboardBaseInfo response
func CreateDescribeDashboardBaseInfoResponse() (response *DescribeDashboardBaseInfoResponse) {
	response = &DescribeDashboardBaseInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}