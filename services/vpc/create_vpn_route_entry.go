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

// CreateVpnRouteEntry invokes the vpc.CreateVpnRouteEntry API synchronously
func (client *Client) CreateVpnRouteEntry(request *CreateVpnRouteEntryRequest) (response *CreateVpnRouteEntryResponse, err error) {
	response = CreateCreateVpnRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpnRouteEntryWithChan invokes the vpc.CreateVpnRouteEntry API asynchronously
func (client *Client) CreateVpnRouteEntryWithChan(request *CreateVpnRouteEntryRequest) (<-chan *CreateVpnRouteEntryResponse, <-chan error) {
	responseChan := make(chan *CreateVpnRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpnRouteEntry(request)
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

// CreateVpnRouteEntryWithCallback invokes the vpc.CreateVpnRouteEntry API asynchronously
func (client *Client) CreateVpnRouteEntryWithCallback(request *CreateVpnRouteEntryRequest, callback func(response *CreateVpnRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpnRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateVpnRouteEntry(request)
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

// CreateVpnRouteEntryRequest is the request struct for api CreateVpnRouteEntry
type CreateVpnRouteEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	PublishVpc           requests.Boolean `position:"Query" name:"PublishVpc"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Weight               requests.Integer `position:"Query" name:"Weight"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouteDest            string           `position:"Query" name:"RouteDest"`
	NextHop              string           `position:"Query" name:"NextHop"`
	OverlayMode          string           `position:"Query" name:"OverlayMode"`
}

// CreateVpnRouteEntryResponse is the response struct for api CreateVpnRouteEntry
type CreateVpnRouteEntryResponse struct {
	*responses.BaseResponse
	NextHop       string `json:"NextHop" xml:"NextHop"`
	Weight        int    `json:"Weight" xml:"Weight"`
	RouteDest     string `json:"RouteDest" xml:"RouteDest"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Description   string `json:"Description" xml:"Description"`
	State         string `json:"State" xml:"State"`
	CreateTime    int64  `json:"CreateTime" xml:"CreateTime"`
	OverlayMode   string `json:"OverlayMode" xml:"OverlayMode"`
	VpnInstanceId string `json:"VpnInstanceId" xml:"VpnInstanceId"`
}

// CreateCreateVpnRouteEntryRequest creates a request to invoke CreateVpnRouteEntry API
func CreateCreateVpnRouteEntryRequest() (request *CreateVpnRouteEntryRequest) {
	request = &CreateVpnRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVpnRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpnRouteEntryResponse creates a response to parse from CreateVpnRouteEntry response
func CreateCreateVpnRouteEntryResponse() (response *CreateVpnRouteEntryResponse) {
	response = &CreateVpnRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
