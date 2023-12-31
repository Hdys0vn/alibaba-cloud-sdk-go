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

// DescribeDashboardApiInfo invokes the lto.DescribeDashboardApiInfo API synchronously
func (client *Client) DescribeDashboardApiInfo(request *DescribeDashboardApiInfoRequest) (response *DescribeDashboardApiInfoResponse, err error) {
	response = CreateDescribeDashboardApiInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDashboardApiInfoWithChan invokes the lto.DescribeDashboardApiInfo API asynchronously
func (client *Client) DescribeDashboardApiInfoWithChan(request *DescribeDashboardApiInfoRequest) (<-chan *DescribeDashboardApiInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDashboardApiInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDashboardApiInfo(request)
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

// DescribeDashboardApiInfoWithCallback invokes the lto.DescribeDashboardApiInfo API asynchronously
func (client *Client) DescribeDashboardApiInfoWithCallback(request *DescribeDashboardApiInfoRequest, callback func(response *DescribeDashboardApiInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDashboardApiInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDashboardApiInfo(request)
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

// DescribeDashboardApiInfoRequest is the request struct for api DescribeDashboardApiInfo
type DescribeDashboardApiInfoRequest struct {
	*requests.RpcRequest
}

// DescribeDashboardApiInfoResponse is the response struct for api DescribeDashboardApiInfo
type DescribeDashboardApiInfoResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDashboardApiInfoRequest creates a request to invoke DescribeDashboardApiInfo API
func CreateDescribeDashboardApiInfoRequest() (request *DescribeDashboardApiInfoRequest) {
	request = &DescribeDashboardApiInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "DescribeDashboardApiInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDashboardApiInfoResponse creates a response to parse from DescribeDashboardApiInfo response
func CreateDescribeDashboardApiInfoResponse() (response *DescribeDashboardApiInfoResponse) {
	response = &DescribeDashboardApiInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
