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

// GetAppMessageQueueRoute invokes the mse.GetAppMessageQueueRoute API synchronously
func (client *Client) GetAppMessageQueueRoute(request *GetAppMessageQueueRouteRequest) (response *GetAppMessageQueueRouteResponse, err error) {
	response = CreateGetAppMessageQueueRouteResponse()
	err = client.DoAction(request, response)
	return
}

// GetAppMessageQueueRouteWithChan invokes the mse.GetAppMessageQueueRoute API asynchronously
func (client *Client) GetAppMessageQueueRouteWithChan(request *GetAppMessageQueueRouteRequest) (<-chan *GetAppMessageQueueRouteResponse, <-chan error) {
	responseChan := make(chan *GetAppMessageQueueRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAppMessageQueueRoute(request)
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

// GetAppMessageQueueRouteWithCallback invokes the mse.GetAppMessageQueueRoute API asynchronously
func (client *Client) GetAppMessageQueueRouteWithCallback(request *GetAppMessageQueueRouteRequest, callback func(response *GetAppMessageQueueRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAppMessageQueueRouteResponse
		var err error
		defer close(result)
		response, err = client.GetAppMessageQueueRoute(request)
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

// GetAppMessageQueueRouteRequest is the request struct for api GetAppMessageQueueRoute
type GetAppMessageQueueRouteRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	AppName        string `position:"Query" name:"AppName"`
	AppId          string `position:"Query" name:"AppId"`
	Namespace      string `position:"Query" name:"Namespace"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	Region         string `position:"Query" name:"Region"`
}

// GetAppMessageQueueRouteResponse is the response struct for api GetAppMessageQueueRoute
type GetAppMessageQueueRouteResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetAppMessageQueueRouteRequest creates a request to invoke GetAppMessageQueueRoute API
func CreateGetAppMessageQueueRouteRequest() (request *GetAppMessageQueueRouteRequest) {
	request = &GetAppMessageQueueRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetAppMessageQueueRoute", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAppMessageQueueRouteResponse creates a response to parse from GetAppMessageQueueRoute response
func CreateGetAppMessageQueueRouteResponse() (response *GetAppMessageQueueRouteResponse) {
	response = &GetAppMessageQueueRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
