package sae

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

// ListIngresses invokes the sae.ListIngresses API synchronously
func (client *Client) ListIngresses(request *ListIngressesRequest) (response *ListIngressesResponse, err error) {
	response = CreateListIngressesResponse()
	err = client.DoAction(request, response)
	return
}

// ListIngressesWithChan invokes the sae.ListIngresses API asynchronously
func (client *Client) ListIngressesWithChan(request *ListIngressesRequest) (<-chan *ListIngressesResponse, <-chan error) {
	responseChan := make(chan *ListIngressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIngresses(request)
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

// ListIngressesWithCallback invokes the sae.ListIngresses API asynchronously
func (client *Client) ListIngressesWithCallback(request *ListIngressesRequest, callback func(response *ListIngressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIngressesResponse
		var err error
		defer close(result)
		response, err = client.ListIngresses(request)
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

// ListIngressesRequest is the request struct for api ListIngresses
type ListIngressesRequest struct {
	*requests.RoaRequest
	NamespaceId string `position:"Query" name:"NamespaceId"`
	AppId       string `position:"Query" name:"AppId"`
}

// ListIngressesResponse is the response struct for api ListIngresses
type ListIngressesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListIngressesRequest creates a request to invoke ListIngresses API
func CreateListIngressesRequest() (request *ListIngressesRequest) {
	request = &ListIngressesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "ListIngresses", "/pop/v1/sam/ingress/IngressList", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListIngressesResponse creates a response to parse from ListIngresses response
func CreateListIngressesResponse() (response *ListIngressesResponse) {
	response = &ListIngressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
