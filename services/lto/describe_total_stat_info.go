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

// DescribeTotalStatInfo invokes the lto.DescribeTotalStatInfo API synchronously
func (client *Client) DescribeTotalStatInfo(request *DescribeTotalStatInfoRequest) (response *DescribeTotalStatInfoResponse, err error) {
	response = CreateDescribeTotalStatInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTotalStatInfoWithChan invokes the lto.DescribeTotalStatInfo API asynchronously
func (client *Client) DescribeTotalStatInfoWithChan(request *DescribeTotalStatInfoRequest) (<-chan *DescribeTotalStatInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeTotalStatInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTotalStatInfo(request)
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

// DescribeTotalStatInfoWithCallback invokes the lto.DescribeTotalStatInfo API asynchronously
func (client *Client) DescribeTotalStatInfoWithCallback(request *DescribeTotalStatInfoRequest, callback func(response *DescribeTotalStatInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTotalStatInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeTotalStatInfo(request)
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

// DescribeTotalStatInfoRequest is the request struct for api DescribeTotalStatInfo
type DescribeTotalStatInfoRequest struct {
	*requests.RpcRequest
}

// DescribeTotalStatInfoResponse is the response struct for api DescribeTotalStatInfo
type DescribeTotalStatInfoResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeTotalStatInfoRequest creates a request to invoke DescribeTotalStatInfo API
func CreateDescribeTotalStatInfoRequest() (request *DescribeTotalStatInfoRequest) {
	request = &DescribeTotalStatInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "DescribeTotalStatInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeTotalStatInfoResponse creates a response to parse from DescribeTotalStatInfo response
func CreateDescribeTotalStatInfoResponse() (response *DescribeTotalStatInfoResponse) {
	response = &DescribeTotalStatInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}