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

// DescribeDcdnDeliverList invokes the dcdn.DescribeDcdnDeliverList API synchronously
func (client *Client) DescribeDcdnDeliverList(request *DescribeDcdnDeliverListRequest) (response *DescribeDcdnDeliverListResponse, err error) {
	response = CreateDescribeDcdnDeliverListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDeliverListWithChan invokes the dcdn.DescribeDcdnDeliverList API asynchronously
func (client *Client) DescribeDcdnDeliverListWithChan(request *DescribeDcdnDeliverListRequest) (<-chan *DescribeDcdnDeliverListResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDeliverListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDeliverList(request)
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

// DescribeDcdnDeliverListWithCallback invokes the dcdn.DescribeDcdnDeliverList API asynchronously
func (client *Client) DescribeDcdnDeliverListWithCallback(request *DescribeDcdnDeliverListRequest, callback func(response *DescribeDcdnDeliverListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDeliverListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDeliverList(request)
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

// DescribeDcdnDeliverListRequest is the request struct for api DescribeDcdnDeliverList
type DescribeDcdnDeliverListRequest struct {
	*requests.RpcRequest
	DeliverId requests.Integer `position:"Query" name:"DeliverId"`
}

// DescribeDcdnDeliverListResponse is the response struct for api DescribeDcdnDeliverList
type DescribeDcdnDeliverListResponse struct {
	*responses.BaseResponse
	Content   string `json:"Content" xml:"Content"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeDcdnDeliverListRequest creates a request to invoke DescribeDcdnDeliverList API
func CreateDescribeDcdnDeliverListRequest() (request *DescribeDcdnDeliverListRequest) {
	request = &DescribeDcdnDeliverListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDeliverList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDeliverListResponse creates a response to parse from DescribeDcdnDeliverList response
func CreateDescribeDcdnDeliverListResponse() (response *DescribeDcdnDeliverListResponse) {
	response = &DescribeDcdnDeliverListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
