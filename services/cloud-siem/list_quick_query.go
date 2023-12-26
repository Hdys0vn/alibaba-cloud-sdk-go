package cloud_siem

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

// ListQuickQuery invokes the cloud_siem.ListQuickQuery API synchronously
func (client *Client) ListQuickQuery(request *ListQuickQueryRequest) (response *ListQuickQueryResponse, err error) {
	response = CreateListQuickQueryResponse()
	err = client.DoAction(request, response)
	return
}

// ListQuickQueryWithChan invokes the cloud_siem.ListQuickQuery API asynchronously
func (client *Client) ListQuickQueryWithChan(request *ListQuickQueryRequest) (<-chan *ListQuickQueryResponse, <-chan error) {
	responseChan := make(chan *ListQuickQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQuickQuery(request)
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

// ListQuickQueryWithCallback invokes the cloud_siem.ListQuickQuery API asynchronously
func (client *Client) ListQuickQueryWithCallback(request *ListQuickQueryRequest, callback func(response *ListQuickQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQuickQueryResponse
		var err error
		defer close(result)
		response, err = client.ListQuickQuery(request)
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

// ListQuickQueryRequest is the request struct for api ListQuickQuery
type ListQuickQueryRequest struct {
	*requests.RpcRequest
	PageSize requests.Integer `position:"Body" name:"PageSize"`
	Offset   requests.Integer `position:"Body" name:"Offset"`
}

// ListQuickQueryResponse is the response struct for api ListQuickQuery
type ListQuickQueryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListQuickQueryRequest creates a request to invoke ListQuickQuery API
func CreateListQuickQueryRequest() (request *ListQuickQueryRequest) {
	request = &ListQuickQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListQuickQuery", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListQuickQueryResponse creates a response to parse from ListQuickQuery response
func CreateListQuickQueryResponse() (response *ListQuickQueryResponse) {
	response = &ListQuickQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
