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

// ListGatewayService invokes the mse.ListGatewayService API synchronously
func (client *Client) ListGatewayService(request *ListGatewayServiceRequest) (response *ListGatewayServiceResponse, err error) {
	response = CreateListGatewayServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ListGatewayServiceWithChan invokes the mse.ListGatewayService API asynchronously
func (client *Client) ListGatewayServiceWithChan(request *ListGatewayServiceRequest) (<-chan *ListGatewayServiceResponse, <-chan error) {
	responseChan := make(chan *ListGatewayServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGatewayService(request)
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

// ListGatewayServiceWithCallback invokes the mse.ListGatewayService API asynchronously
func (client *Client) ListGatewayServiceWithCallback(request *ListGatewayServiceRequest, callback func(response *ListGatewayServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGatewayServiceResponse
		var err error
		defer close(result)
		response, err = client.ListGatewayService(request)
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

// ListGatewayServiceRequest is the request struct for api ListGatewayService
type ListGatewayServiceRequest struct {
	*requests.RpcRequest
	MseSessionId   string                         `position:"Query" name:"MseSessionId"`
	PageNumber     requests.Integer               `position:"Query" name:"PageNumber"`
	OrderItem      string                         `position:"Query" name:"OrderItem"`
	PageSize       requests.Integer               `position:"Query" name:"PageSize"`
	DescSort       requests.Boolean               `position:"Query" name:"DescSort"`
	FilterParams   ListGatewayServiceFilterParams `position:"Query" name:"FilterParams"  type:"Struct"`
	AcceptLanguage string                         `position:"Query" name:"AcceptLanguage"`
}

// ListGatewayServiceFilterParams is a repeated param struct in ListGatewayServiceRequest
type ListGatewayServiceFilterParams struct {
	GatewayUniqueId string `name:"GatewayUniqueId"`
	ServiceProtocol string `name:"ServiceProtocol"`
	Name            string `name:"Name"`
	Namespace       string `name:"Namespace"`
	SourceType      string `name:"SourceType"`
	GroupName       string `name:"GroupName"`
}

// ListGatewayServiceResponse is the response struct for api ListGatewayService
type ListGatewayServiceResponse struct {
	*responses.BaseResponse
	RequestId      string                   `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                   `json:"Message" xml:"Message"`
	Code           int                      `json:"Code" xml:"Code"`
	Success        bool                     `json:"Success" xml:"Success"`
	Data           DataInListGatewayService `json:"Data" xml:"Data"`
}

// CreateListGatewayServiceRequest creates a request to invoke ListGatewayService API
func CreateListGatewayServiceRequest() (request *ListGatewayServiceRequest) {
	request = &ListGatewayServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListGatewayService", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGatewayServiceResponse creates a response to parse from ListGatewayService response
func CreateListGatewayServiceResponse() (response *ListGatewayServiceResponse) {
	response = &ListGatewayServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
