package mse

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

// ListGatewayRoute invokes the mse.ListGatewayRoute API synchronously
func (client *Client) ListGatewayRoute(request *ListGatewayRouteRequest) (response *ListGatewayRouteResponse, err error) {
	response = CreateListGatewayRouteResponse()
	err = client.DoAction(request, response)
	return
}

// ListGatewayRouteWithChan invokes the mse.ListGatewayRoute API asynchronously
func (client *Client) ListGatewayRouteWithChan(request *ListGatewayRouteRequest) (<-chan *ListGatewayRouteResponse, <-chan error) {
	responseChan := make(chan *ListGatewayRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGatewayRoute(request)
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

// ListGatewayRouteWithCallback invokes the mse.ListGatewayRoute API asynchronously
func (client *Client) ListGatewayRouteWithCallback(request *ListGatewayRouteRequest, callback func(response *ListGatewayRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGatewayRouteResponse
		var err error
		defer close(result)
		response, err = client.ListGatewayRoute(request)
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

// ListGatewayRouteRequest is the request struct for api ListGatewayRoute
type ListGatewayRouteRequest struct {
	*requests.RpcRequest
	MseSessionId   string                       `position:"Query" name:"MseSessionId"`
	PageNumber     requests.Integer             `position:"Query" name:"PageNumber"`
	OrderItem      string                       `position:"Query" name:"OrderItem"`
	PageSize       requests.Integer             `position:"Query" name:"PageSize"`
	DescSort       requests.Boolean             `position:"Query" name:"DescSort"`
	FilterParams   ListGatewayRouteFilterParams `position:"Query" name:"FilterParams"  type:"Struct"`
	AcceptLanguage string                       `position:"Query" name:"AcceptLanguage"`
}

// ListGatewayRouteFilterParams is a repeated param struct in ListGatewayRouteRequest
type ListGatewayRouteFilterParams struct {
	DefaultServiceId string `name:"DefaultServiceId"`
	RouteOrder       string `name:"RouteOrder"`
	GatewayUniqueId  string `name:"GatewayUniqueId"`
	Name             string `name:"Name"`
	DomainName       string `name:"DomainName"`
	GatewayId        string `name:"GatewayId"`
	DomainId         string `name:"DomainId"`
	Status           string `name:"Status"`
}

// ListGatewayRouteResponse is the response struct for api ListGatewayRoute
type ListGatewayRouteResponse struct {
	*responses.BaseResponse
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                 `json:"Message" xml:"Message"`
	Code           int                    `json:"Code" xml:"Code"`
	Success        bool                   `json:"Success" xml:"Success"`
	Data           DataInListGatewayRoute `json:"Data" xml:"Data"`
}

// CreateListGatewayRouteRequest creates a request to invoke ListGatewayRoute API
func CreateListGatewayRouteRequest() (request *ListGatewayRouteRequest) {
	request = &ListGatewayRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListGatewayRoute", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGatewayRouteResponse creates a response to parse from ListGatewayRoute response
func CreateListGatewayRouteResponse() (response *ListGatewayRouteResponse) {
	response = &ListGatewayRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
