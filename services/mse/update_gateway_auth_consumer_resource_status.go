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

// UpdateGatewayAuthConsumerResourceStatus invokes the mse.UpdateGatewayAuthConsumerResourceStatus API synchronously
func (client *Client) UpdateGatewayAuthConsumerResourceStatus(request *UpdateGatewayAuthConsumerResourceStatusRequest) (response *UpdateGatewayAuthConsumerResourceStatusResponse, err error) {
	response = CreateUpdateGatewayAuthConsumerResourceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayAuthConsumerResourceStatusWithChan invokes the mse.UpdateGatewayAuthConsumerResourceStatus API asynchronously
func (client *Client) UpdateGatewayAuthConsumerResourceStatusWithChan(request *UpdateGatewayAuthConsumerResourceStatusRequest) (<-chan *UpdateGatewayAuthConsumerResourceStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayAuthConsumerResourceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayAuthConsumerResourceStatus(request)
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

// UpdateGatewayAuthConsumerResourceStatusWithCallback invokes the mse.UpdateGatewayAuthConsumerResourceStatus API asynchronously
func (client *Client) UpdateGatewayAuthConsumerResourceStatusWithCallback(request *UpdateGatewayAuthConsumerResourceStatusRequest, callback func(response *UpdateGatewayAuthConsumerResourceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayAuthConsumerResourceStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayAuthConsumerResourceStatus(request)
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

// UpdateGatewayAuthConsumerResourceStatusRequest is the request struct for api UpdateGatewayAuthConsumerResourceStatus
type UpdateGatewayAuthConsumerResourceStatusRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	ConsumerId      requests.Integer `position:"Query" name:"ConsumerId"`
	IdList          string           `position:"Query" name:"IdList"`
	ResourceStatus  requests.Boolean `position:"Query" name:"ResourceStatus"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewayAuthConsumerResourceStatusResponse is the response struct for api UpdateGatewayAuthConsumerResourceStatus
type UpdateGatewayAuthConsumerResourceStatusResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           bool   `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayAuthConsumerResourceStatusRequest creates a request to invoke UpdateGatewayAuthConsumerResourceStatus API
func CreateUpdateGatewayAuthConsumerResourceStatusRequest() (request *UpdateGatewayAuthConsumerResourceStatusRequest) {
	request = &UpdateGatewayAuthConsumerResourceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayAuthConsumerResourceStatus", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayAuthConsumerResourceStatusResponse creates a response to parse from UpdateGatewayAuthConsumerResourceStatus response
func CreateUpdateGatewayAuthConsumerResourceStatusResponse() (response *UpdateGatewayAuthConsumerResourceStatusResponse) {
	response = &UpdateGatewayAuthConsumerResourceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}