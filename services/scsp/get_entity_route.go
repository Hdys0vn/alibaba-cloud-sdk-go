package scsp

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

// GetEntityRoute invokes the scsp.GetEntityRoute API synchronously
func (client *Client) GetEntityRoute(request *GetEntityRouteRequest) (response *GetEntityRouteResponse, err error) {
	response = CreateGetEntityRouteResponse()
	err = client.DoAction(request, response)
	return
}

// GetEntityRouteWithChan invokes the scsp.GetEntityRoute API asynchronously
func (client *Client) GetEntityRouteWithChan(request *GetEntityRouteRequest) (<-chan *GetEntityRouteResponse, <-chan error) {
	responseChan := make(chan *GetEntityRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEntityRoute(request)
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

// GetEntityRouteWithCallback invokes the scsp.GetEntityRoute API asynchronously
func (client *Client) GetEntityRouteWithCallback(request *GetEntityRouteRequest, callback func(response *GetEntityRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEntityRouteResponse
		var err error
		defer close(result)
		response, err = client.GetEntityRoute(request)
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

// GetEntityRouteRequest is the request struct for api GetEntityRoute
type GetEntityRouteRequest struct {
	*requests.RpcRequest
	EntityBizCode        string           `position:"Body"`
	InstanceId           string           `position:"Body"`
	EntityName           string           `position:"Body"`
	EntityId             string           `position:"Body"`
	EntityRelationNumber string           `position:"Body"`
	UniqueId             requests.Integer `position:"Body"`
}

// GetEntityRouteResponse is the response struct for api GetEntityRoute
type GetEntityRouteResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetEntityRouteRequest creates a request to invoke GetEntityRoute API
func CreateGetEntityRouteRequest() (request *GetEntityRouteRequest) {
	request = &GetEntityRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GetEntityRoute", "", "")
	request.Method = requests.POST
	return
}

// CreateGetEntityRouteResponse creates a response to parse from GetEntityRoute response
func CreateGetEntityRouteResponse() (response *GetEntityRouteResponse) {
	response = &GetEntityRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
