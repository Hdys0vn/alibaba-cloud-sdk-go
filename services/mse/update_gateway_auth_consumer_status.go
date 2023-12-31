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

// UpdateGatewayAuthConsumerStatus invokes the mse.UpdateGatewayAuthConsumerStatus API synchronously
func (client *Client) UpdateGatewayAuthConsumerStatus(request *UpdateGatewayAuthConsumerStatusRequest) (response *UpdateGatewayAuthConsumerStatusResponse, err error) {
	response = CreateUpdateGatewayAuthConsumerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayAuthConsumerStatusWithChan invokes the mse.UpdateGatewayAuthConsumerStatus API asynchronously
func (client *Client) UpdateGatewayAuthConsumerStatusWithChan(request *UpdateGatewayAuthConsumerStatusRequest) (<-chan *UpdateGatewayAuthConsumerStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayAuthConsumerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayAuthConsumerStatus(request)
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

// UpdateGatewayAuthConsumerStatusWithCallback invokes the mse.UpdateGatewayAuthConsumerStatus API asynchronously
func (client *Client) UpdateGatewayAuthConsumerStatusWithCallback(request *UpdateGatewayAuthConsumerStatusRequest, callback func(response *UpdateGatewayAuthConsumerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayAuthConsumerStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayAuthConsumerStatus(request)
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

// UpdateGatewayAuthConsumerStatusRequest is the request struct for api UpdateGatewayAuthConsumerStatus
type UpdateGatewayAuthConsumerStatusRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	Id              requests.Integer `position:"Query" name:"Id"`
	ConsumerStatus  requests.Boolean `position:"Query" name:"ConsumerStatus"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewayAuthConsumerStatusResponse is the response struct for api UpdateGatewayAuthConsumerStatus
type UpdateGatewayAuthConsumerStatusResponse struct {
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

// CreateUpdateGatewayAuthConsumerStatusRequest creates a request to invoke UpdateGatewayAuthConsumerStatus API
func CreateUpdateGatewayAuthConsumerStatusRequest() (request *UpdateGatewayAuthConsumerStatusRequest) {
	request = &UpdateGatewayAuthConsumerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayAuthConsumerStatus", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayAuthConsumerStatusResponse creates a response to parse from UpdateGatewayAuthConsumerStatus response
func CreateUpdateGatewayAuthConsumerStatusResponse() (response *UpdateGatewayAuthConsumerStatusResponse) {
	response = &UpdateGatewayAuthConsumerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
