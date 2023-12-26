package vpc

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

// CreateVcoRouteEntry invokes the vpc.CreateVcoRouteEntry API synchronously
func (client *Client) CreateVcoRouteEntry(request *CreateVcoRouteEntryRequest) (response *CreateVcoRouteEntryResponse, err error) {
	response = CreateCreateVcoRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVcoRouteEntryWithChan invokes the vpc.CreateVcoRouteEntry API asynchronously
func (client *Client) CreateVcoRouteEntryWithChan(request *CreateVcoRouteEntryRequest) (<-chan *CreateVcoRouteEntryResponse, <-chan error) {
	responseChan := make(chan *CreateVcoRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVcoRouteEntry(request)
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

// CreateVcoRouteEntryWithCallback invokes the vpc.CreateVcoRouteEntry API asynchronously
func (client *Client) CreateVcoRouteEntryWithCallback(request *CreateVcoRouteEntryRequest, callback func(response *CreateVcoRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVcoRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateVcoRouteEntry(request)
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

// CreateVcoRouteEntryRequest is the request struct for api CreateVcoRouteEntry
type CreateVcoRouteEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Weight               requests.Integer `position:"Query" name:"Weight"`
	RouteDest            string           `position:"Query" name:"RouteDest"`
	NextHop              string           `position:"Query" name:"NextHop"`
	VpnConnectionId      string           `position:"Query" name:"VpnConnectionId"`
	OverlayMode          string           `position:"Query" name:"OverlayMode"`
}

// CreateVcoRouteEntryResponse is the response struct for api CreateVcoRouteEntry
type CreateVcoRouteEntryResponse struct {
	*responses.BaseResponse
	VpnConnectionId string `json:"VpnConnectionId" xml:"VpnConnectionId"`
	RouteDest       string `json:"RouteDest" xml:"RouteDest"`
	NextHop         string `json:"NextHop" xml:"NextHop"`
	Weight          int    `json:"Weight" xml:"Weight"`
	OverlayMode     string `json:"OverlayMode" xml:"OverlayMode"`
	State           string `json:"State" xml:"State"`
	CreateTime      int64  `json:"CreateTime" xml:"CreateTime"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	Description     string `json:"Description" xml:"Description"`
}

// CreateCreateVcoRouteEntryRequest creates a request to invoke CreateVcoRouteEntry API
func CreateCreateVcoRouteEntryRequest() (request *CreateVcoRouteEntryRequest) {
	request = &CreateVcoRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVcoRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVcoRouteEntryResponse creates a response to parse from CreateVcoRouteEntry response
func CreateCreateVcoRouteEntryResponse() (response *CreateVcoRouteEntryResponse) {
	response = &CreateVcoRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
