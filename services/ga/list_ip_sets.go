package ga

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

// ListIpSets invokes the ga.ListIpSets API synchronously
func (client *Client) ListIpSets(request *ListIpSetsRequest) (response *ListIpSetsResponse, err error) {
	response = CreateListIpSetsResponse()
	err = client.DoAction(request, response)
	return
}

// ListIpSetsWithChan invokes the ga.ListIpSets API asynchronously
func (client *Client) ListIpSetsWithChan(request *ListIpSetsRequest) (<-chan *ListIpSetsResponse, <-chan error) {
	responseChan := make(chan *ListIpSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIpSets(request)
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

// ListIpSetsWithCallback invokes the ga.ListIpSets API asynchronously
func (client *Client) ListIpSetsWithCallback(request *ListIpSetsRequest, callback func(response *ListIpSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIpSetsResponse
		var err error
		defer close(result)
		response, err = client.ListIpSets(request)
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

// ListIpSetsRequest is the request struct for api ListIpSets
type ListIpSetsRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	AcceleratorId string           `position:"Query" name:"AcceleratorId"`
}

// ListIpSetsResponse is the response struct for api ListIpSets
type ListIpSetsResponse struct {
	*responses.BaseResponse
	TotalCount int          `json:"TotalCount" xml:"TotalCount"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	PageSize   int          `json:"PageSize" xml:"PageSize"`
	PageNumber int          `json:"PageNumber" xml:"PageNumber"`
	IpSets     []IpSetsItem `json:"IpSets" xml:"IpSets"`
}

// CreateListIpSetsRequest creates a request to invoke ListIpSets API
func CreateListIpSetsRequest() (request *ListIpSetsRequest) {
	request = &ListIpSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ListIpSets", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListIpSetsResponse creates a response to parse from ListIpSets response
func CreateListIpSetsResponse() (response *ListIpSetsResponse) {
	response = &ListIpSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
