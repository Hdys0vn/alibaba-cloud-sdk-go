package live

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

// DescribeLiveShiftConfigs invokes the live.DescribeLiveShiftConfigs API synchronously
func (client *Client) DescribeLiveShiftConfigs(request *DescribeLiveShiftConfigsRequest) (response *DescribeLiveShiftConfigsResponse, err error) {
	response = CreateDescribeLiveShiftConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveShiftConfigsWithChan invokes the live.DescribeLiveShiftConfigs API asynchronously
func (client *Client) DescribeLiveShiftConfigsWithChan(request *DescribeLiveShiftConfigsRequest) (<-chan *DescribeLiveShiftConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveShiftConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveShiftConfigs(request)
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

// DescribeLiveShiftConfigsWithCallback invokes the live.DescribeLiveShiftConfigs API asynchronously
func (client *Client) DescribeLiveShiftConfigsWithCallback(request *DescribeLiveShiftConfigsRequest, callback func(response *DescribeLiveShiftConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveShiftConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveShiftConfigs(request)
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

// DescribeLiveShiftConfigsRequest is the request struct for api DescribeLiveShiftConfigs
type DescribeLiveShiftConfigsRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveShiftConfigsResponse is the response struct for api DescribeLiveShiftConfigs
type DescribeLiveShiftConfigsResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Content   ContentInDescribeLiveShiftConfigs `json:"Content" xml:"Content"`
}

// CreateDescribeLiveShiftConfigsRequest creates a request to invoke DescribeLiveShiftConfigs API
func CreateDescribeLiveShiftConfigsRequest() (request *DescribeLiveShiftConfigsRequest) {
	request = &DescribeLiveShiftConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveShiftConfigs", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveShiftConfigsResponse creates a response to parse from DescribeLiveShiftConfigs response
func CreateDescribeLiveShiftConfigsResponse() (response *DescribeLiveShiftConfigsResponse) {
	response = &DescribeLiveShiftConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
