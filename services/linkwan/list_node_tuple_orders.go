package linkwan

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

// ListNodeTupleOrders invokes the linkwan.ListNodeTupleOrders API synchronously
func (client *Client) ListNodeTupleOrders(request *ListNodeTupleOrdersRequest) (response *ListNodeTupleOrdersResponse, err error) {
	response = CreateListNodeTupleOrdersResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeTupleOrdersWithChan invokes the linkwan.ListNodeTupleOrders API asynchronously
func (client *Client) ListNodeTupleOrdersWithChan(request *ListNodeTupleOrdersRequest) (<-chan *ListNodeTupleOrdersResponse, <-chan error) {
	responseChan := make(chan *ListNodeTupleOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeTupleOrders(request)
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

// ListNodeTupleOrdersWithCallback invokes the linkwan.ListNodeTupleOrders API asynchronously
func (client *Client) ListNodeTupleOrdersWithCallback(request *ListNodeTupleOrdersRequest, callback func(response *ListNodeTupleOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeTupleOrdersResponse
		var err error
		defer close(result)
		response, err = client.ListNodeTupleOrders(request)
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

// ListNodeTupleOrdersRequest is the request struct for api ListNodeTupleOrders
type ListNodeTupleOrdersRequest struct {
	*requests.RpcRequest
	IsKpm        requests.Boolean `position:"Query" name:"IsKpm"`
	Limit        requests.Integer `position:"Query" name:"Limit"`
	State        *[]string        `position:"Query" name:"State"  type:"Repeated"`
	Offset       requests.Integer `position:"Query" name:"Offset"`
	Ascending    requests.Boolean `position:"Query" name:"Ascending"`
	ApiProduct   string           `position:"Body" name:"ApiProduct"`
	ApiRevision  string           `position:"Body" name:"ApiRevision"`
	SortingField string           `position:"Query" name:"SortingField"`
}

// ListNodeTupleOrdersResponse is the response struct for api ListNodeTupleOrders
type ListNodeTupleOrdersResponse struct {
	*responses.BaseResponse
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Success   bool                      `json:"Success" xml:"Success"`
	Data      DataInListNodeTupleOrders `json:"Data" xml:"Data"`
}

// CreateListNodeTupleOrdersRequest creates a request to invoke ListNodeTupleOrders API
func CreateListNodeTupleOrdersRequest() (request *ListNodeTupleOrdersRequest) {
	request = &ListNodeTupleOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListNodeTupleOrders", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodeTupleOrdersResponse creates a response to parse from ListNodeTupleOrders response
func CreateListNodeTupleOrdersResponse() (response *ListNodeTupleOrdersResponse) {
	response = &ListNodeTupleOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
