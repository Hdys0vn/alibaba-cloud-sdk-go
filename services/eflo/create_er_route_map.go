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

// CreateErRouteMap invokes the eflo.CreateErRouteMap API synchronously
func (client *Client) CreateErRouteMap(request *CreateErRouteMapRequest) (response *CreateErRouteMapResponse, err error) {
	response = CreateCreateErRouteMapResponse()
	err = client.DoAction(request, response)
	return
}

// CreateErRouteMapWithChan invokes the eflo.CreateErRouteMap API asynchronously
func (client *Client) CreateErRouteMapWithChan(request *CreateErRouteMapRequest) (<-chan *CreateErRouteMapResponse, <-chan error) {
	responseChan := make(chan *CreateErRouteMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateErRouteMap(request)
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

// CreateErRouteMapWithCallback invokes the eflo.CreateErRouteMap API asynchronously
func (client *Client) CreateErRouteMapWithCallback(request *CreateErRouteMapRequest, callback func(response *CreateErRouteMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateErRouteMapResponse
		var err error
		defer close(result)
		response, err = client.CreateErRouteMap(request)
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

// CreateErRouteMapRequest is the request struct for api CreateErRouteMap
type CreateErRouteMapRequest struct {
	*requests.RpcRequest
	ReceptionInstanceType     string           `position:"Body" name:"ReceptionInstanceType"`
	Description               string           `position:"Body" name:"Description"`
	ReceptionInstanceId       string           `position:"Body" name:"ReceptionInstanceId"`
	RouteMapAction            string           `position:"Body" name:"RouteMapAction"`
	TransmissionInstanceType  string           `position:"Body" name:"TransmissionInstanceType"`
	DestinationCidrBlock      string           `position:"Body" name:"DestinationCidrBlock"`
	ErId                      string           `position:"Body" name:"ErId"`
	RouteMapNum               requests.Integer `position:"Body" name:"RouteMapNum"`
	ReceptionInstanceOwner    string           `position:"Body" name:"ReceptionInstanceOwner"`
	TransmissionInstanceOwner string           `position:"Body" name:"TransmissionInstanceOwner"`
	TransmissionInstanceId    string           `position:"Body" name:"TransmissionInstanceId"`
}

// CreateErRouteMapResponse is the response struct for api CreateErRouteMap
type CreateErRouteMapResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateCreateErRouteMapRequest creates a request to invoke CreateErRouteMap API
func CreateCreateErRouteMapRequest() (request *CreateErRouteMapRequest) {
	request = &CreateErRouteMapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "CreateErRouteMap", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateErRouteMapResponse creates a response to parse from CreateErRouteMap response
func CreateCreateErRouteMapResponse() (response *CreateErRouteMapResponse) {
	response = &CreateErRouteMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
