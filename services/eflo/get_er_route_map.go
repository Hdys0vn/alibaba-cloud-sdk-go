package eflo

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

// GetErRouteMap invokes the eflo.GetErRouteMap API synchronously
func (client *Client) GetErRouteMap(request *GetErRouteMapRequest) (response *GetErRouteMapResponse, err error) {
	response = CreateGetErRouteMapResponse()
	err = client.DoAction(request, response)
	return
}

// GetErRouteMapWithChan invokes the eflo.GetErRouteMap API asynchronously
func (client *Client) GetErRouteMapWithChan(request *GetErRouteMapRequest) (<-chan *GetErRouteMapResponse, <-chan error) {
	responseChan := make(chan *GetErRouteMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetErRouteMap(request)
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

// GetErRouteMapWithCallback invokes the eflo.GetErRouteMap API asynchronously
func (client *Client) GetErRouteMapWithCallback(request *GetErRouteMapRequest, callback func(response *GetErRouteMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetErRouteMapResponse
		var err error
		defer close(result)
		response, err = client.GetErRouteMap(request)
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

// GetErRouteMapRequest is the request struct for api GetErRouteMap
type GetErRouteMapRequest struct {
	*requests.RpcRequest
	ErId         string `position:"Body" name:"ErId"`
	ErRouteMapId string `position:"Body" name:"ErRouteMapId"`
}

// GetErRouteMapResponse is the response struct for api GetErRouteMap
type GetErRouteMapResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetErRouteMapRequest creates a request to invoke GetErRouteMap API
func CreateGetErRouteMapRequest() (request *GetErRouteMapRequest) {
	request = &GetErRouteMapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetErRouteMap", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetErRouteMapResponse creates a response to parse from GetErRouteMap response
func CreateGetErRouteMapResponse() (response *GetErRouteMapResponse) {
	response = &GetErRouteMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
