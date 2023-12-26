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

// DescribeDcdnSubList invokes the dcdn.DescribeDcdnSubList API synchronously
func (client *Client) DescribeDcdnSubList(request *DescribeDcdnSubListRequest) (response *DescribeDcdnSubListResponse, err error) {
	response = CreateDescribeDcdnSubListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnSubListWithChan invokes the dcdn.DescribeDcdnSubList API asynchronously
func (client *Client) DescribeDcdnSubListWithChan(request *DescribeDcdnSubListRequest) (<-chan *DescribeDcdnSubListResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnSubListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnSubList(request)
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

// DescribeDcdnSubListWithCallback invokes the dcdn.DescribeDcdnSubList API asynchronously
func (client *Client) DescribeDcdnSubListWithCallback(request *DescribeDcdnSubListRequest, callback func(response *DescribeDcdnSubListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnSubListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnSubList(request)
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

// DescribeDcdnSubListRequest is the request struct for api DescribeDcdnSubList
type DescribeDcdnSubListRequest struct {
	*requests.RpcRequest
}

// DescribeDcdnSubListResponse is the response struct for api DescribeDcdnSubList
type DescribeDcdnSubListResponse struct {
	*responses.BaseResponse
	Content   string `json:"Content" xml:"Content"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeDcdnSubListRequest creates a request to invoke DescribeDcdnSubList API
func CreateDescribeDcdnSubListRequest() (request *DescribeDcdnSubListRequest) {
	request = &DescribeDcdnSubListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnSubList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnSubListResponse creates a response to parse from DescribeDcdnSubList response
func CreateDescribeDcdnSubListResponse() (response *DescribeDcdnSubListResponse) {
	response = &DescribeDcdnSubListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
