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

// UpdateErRouteMap invokes the eflo.UpdateErRouteMap API synchronously
func (client *Client) UpdateErRouteMap(request *UpdateErRouteMapRequest) (response *UpdateErRouteMapResponse, err error) {
	response = CreateUpdateErRouteMapResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateErRouteMapWithChan invokes the eflo.UpdateErRouteMap API asynchronously
func (client *Client) UpdateErRouteMapWithChan(request *UpdateErRouteMapRequest) (<-chan *UpdateErRouteMapResponse, <-chan error) {
	responseChan := make(chan *UpdateErRouteMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateErRouteMap(request)
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

// UpdateErRouteMapWithCallback invokes the eflo.UpdateErRouteMap API asynchronously
func (client *Client) UpdateErRouteMapWithCallback(request *UpdateErRouteMapRequest, callback func(response *UpdateErRouteMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateErRouteMapResponse
		var err error
		defer close(result)
		response, err = client.UpdateErRouteMap(request)
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

// UpdateErRouteMapRequest is the request struct for api UpdateErRouteMap
type UpdateErRouteMapRequest struct {
	*requests.RpcRequest
	ErId         string `position:"Body" name:"ErId"`
	Description  string `position:"Body" name:"Description"`
	ErRouteMapId string `position:"Body" name:"ErRouteMapId"`
}

// UpdateErRouteMapResponse is the response struct for api UpdateErRouteMap
type UpdateErRouteMapResponse struct {
	*responses.BaseResponse
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Content   map[string]interface{} `json:"Content" xml:"Content"`
}

// CreateUpdateErRouteMapRequest creates a request to invoke UpdateErRouteMap API
func CreateUpdateErRouteMapRequest() (request *UpdateErRouteMapRequest) {
	request = &UpdateErRouteMapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "UpdateErRouteMap", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateErRouteMapResponse creates a response to parse from UpdateErRouteMap response
func CreateUpdateErRouteMapResponse() (response *UpdateErRouteMapResponse) {
	response = &UpdateErRouteMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}