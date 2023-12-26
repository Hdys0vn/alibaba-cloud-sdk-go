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

// UpdateGatewayRouteWafStatus invokes the mse.UpdateGatewayRouteWafStatus API synchronously
func (client *Client) UpdateGatewayRouteWafStatus(request *UpdateGatewayRouteWafStatusRequest) (response *UpdateGatewayRouteWafStatusResponse, err error) {
	response = CreateUpdateGatewayRouteWafStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayRouteWafStatusWithChan invokes the mse.UpdateGatewayRouteWafStatus API asynchronously
func (client *Client) UpdateGatewayRouteWafStatusWithChan(request *UpdateGatewayRouteWafStatusRequest) (<-chan *UpdateGatewayRouteWafStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayRouteWafStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayRouteWafStatus(request)
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

// UpdateGatewayRouteWafStatusWithCallback invokes the mse.UpdateGatewayRouteWafStatus API asynchronously
func (client *Client) UpdateGatewayRouteWafStatusWithCallback(request *UpdateGatewayRouteWafStatusRequest, callback func(response *UpdateGatewayRouteWafStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayRouteWafStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayRouteWafStatus(request)
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

// UpdateGatewayRouteWafStatusRequest is the request struct for api UpdateGatewayRouteWafStatus
type UpdateGatewayRouteWafStatusRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	EnableWaf       requests.Boolean `position:"Query" name:"EnableWaf"`
	RouteId         requests.Integer `position:"Query" name:"RouteId"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewayRouteWafStatusResponse is the response struct for api UpdateGatewayRouteWafStatus
type UpdateGatewayRouteWafStatusResponse struct {
	*responses.BaseResponse
	RequestId      string                            `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                            `json:"Message" xml:"Message"`
	Code           int                               `json:"Code" xml:"Code"`
	Success        bool                              `json:"Success" xml:"Success"`
	Data           DataInUpdateGatewayRouteWafStatus `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayRouteWafStatusRequest creates a request to invoke UpdateGatewayRouteWafStatus API
func CreateUpdateGatewayRouteWafStatusRequest() (request *UpdateGatewayRouteWafStatusRequest) {
	request = &UpdateGatewayRouteWafStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayRouteWafStatus", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayRouteWafStatusResponse creates a response to parse from UpdateGatewayRouteWafStatus response
func CreateUpdateGatewayRouteWafStatusResponse() (response *UpdateGatewayRouteWafStatusResponse) {
	response = &UpdateGatewayRouteWafStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
