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

// DeleteGatewayAuthConsumerResource invokes the mse.DeleteGatewayAuthConsumerResource API synchronously
func (client *Client) DeleteGatewayAuthConsumerResource(request *DeleteGatewayAuthConsumerResourceRequest) (response *DeleteGatewayAuthConsumerResourceResponse, err error) {
	response = CreateDeleteGatewayAuthConsumerResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayAuthConsumerResourceWithChan invokes the mse.DeleteGatewayAuthConsumerResource API asynchronously
func (client *Client) DeleteGatewayAuthConsumerResourceWithChan(request *DeleteGatewayAuthConsumerResourceRequest) (<-chan *DeleteGatewayAuthConsumerResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayAuthConsumerResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGatewayAuthConsumerResource(request)
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

// DeleteGatewayAuthConsumerResourceWithCallback invokes the mse.DeleteGatewayAuthConsumerResource API asynchronously
func (client *Client) DeleteGatewayAuthConsumerResourceWithCallback(request *DeleteGatewayAuthConsumerResourceRequest, callback func(response *DeleteGatewayAuthConsumerResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayAuthConsumerResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteGatewayAuthConsumerResource(request)
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

// DeleteGatewayAuthConsumerResourceRequest is the request struct for api DeleteGatewayAuthConsumerResource
type DeleteGatewayAuthConsumerResourceRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	ConsumerId      requests.Integer `position:"Query" name:"ConsumerId"`
	IdList          string           `position:"Query" name:"IdList"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// DeleteGatewayAuthConsumerResourceResponse is the response struct for api DeleteGatewayAuthConsumerResource
type DeleteGatewayAuthConsumerResourceResponse struct {
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

// CreateDeleteGatewayAuthConsumerResourceRequest creates a request to invoke DeleteGatewayAuthConsumerResource API
func CreateDeleteGatewayAuthConsumerResourceRequest() (request *DeleteGatewayAuthConsumerResourceRequest) {
	request = &DeleteGatewayAuthConsumerResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteGatewayAuthConsumerResource", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayAuthConsumerResourceResponse creates a response to parse from DeleteGatewayAuthConsumerResource response
func CreateDeleteGatewayAuthConsumerResourceResponse() (response *DeleteGatewayAuthConsumerResourceResponse) {
	response = &DeleteGatewayAuthConsumerResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
