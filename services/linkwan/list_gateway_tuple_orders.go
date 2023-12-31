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

// ListGatewayTupleOrders invokes the linkwan.ListGatewayTupleOrders API synchronously
func (client *Client) ListGatewayTupleOrders(request *ListGatewayTupleOrdersRequest) (response *ListGatewayTupleOrdersResponse, err error) {
	response = CreateListGatewayTupleOrdersResponse()
	err = client.DoAction(request, response)
	return
}

// ListGatewayTupleOrdersWithChan invokes the linkwan.ListGatewayTupleOrders API asynchronously
func (client *Client) ListGatewayTupleOrdersWithChan(request *ListGatewayTupleOrdersRequest) (<-chan *ListGatewayTupleOrdersResponse, <-chan error) {
	responseChan := make(chan *ListGatewayTupleOrdersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGatewayTupleOrders(request)
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

// ListGatewayTupleOrdersWithCallback invokes the linkwan.ListGatewayTupleOrders API asynchronously
func (client *Client) ListGatewayTupleOrdersWithCallback(request *ListGatewayTupleOrdersRequest, callback func(response *ListGatewayTupleOrdersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGatewayTupleOrdersResponse
		var err error
		defer close(result)
		response, err = client.ListGatewayTupleOrders(request)
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

// ListGatewayTupleOrdersRequest is the request struct for api ListGatewayTupleOrders
type ListGatewayTupleOrdersRequest struct {
	*requests.RpcRequest
	Limit        requests.Integer `position:"Query" name:"Limit"`
	State        *[]string        `position:"Query" name:"State"  type:"Repeated"`
	Offset       requests.Integer `position:"Query" name:"Offset"`
	Ascending    requests.Boolean `position:"Query" name:"Ascending"`
	ApiProduct   string           `position:"Body" name:"ApiProduct"`
	ApiRevision  string           `position:"Body" name:"ApiRevision"`
	SortingField string           `position:"Query" name:"SortingField"`
}

// ListGatewayTupleOrdersResponse is the response struct for api ListGatewayTupleOrders
type ListGatewayTupleOrdersResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Success   bool                         `json:"Success" xml:"Success"`
	Data      DataInListGatewayTupleOrders `json:"Data" xml:"Data"`
}

// CreateListGatewayTupleOrdersRequest creates a request to invoke ListGatewayTupleOrders API
func CreateListGatewayTupleOrdersRequest() (request *ListGatewayTupleOrdersRequest) {
	request = &ListGatewayTupleOrdersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListGatewayTupleOrders", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGatewayTupleOrdersResponse creates a response to parse from ListGatewayTupleOrders response
func CreateListGatewayTupleOrdersResponse() (response *ListGatewayTupleOrdersResponse) {
	response = &ListGatewayTupleOrdersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
