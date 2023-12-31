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

// CreateRouteEntry invokes the vpc.CreateRouteEntry API synchronously
func (client *Client) CreateRouteEntry(request *CreateRouteEntryRequest) (response *CreateRouteEntryResponse, err error) {
	response = CreateCreateRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRouteEntryWithChan invokes the vpc.CreateRouteEntry API asynchronously
func (client *Client) CreateRouteEntryWithChan(request *CreateRouteEntryRequest) (<-chan *CreateRouteEntryResponse, <-chan error) {
	responseChan := make(chan *CreateRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRouteEntry(request)
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

// CreateRouteEntryWithCallback invokes the vpc.CreateRouteEntry API asynchronously
func (client *Client) CreateRouteEntryWithCallback(request *CreateRouteEntryRequest, callback func(response *CreateRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateRouteEntry(request)
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

// CreateRouteEntryRequest is the request struct for api CreateRouteEntry
type CreateRouteEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	RouteEntryName       string                         `position:"Query" name:"RouteEntryName"`
	ClientToken          string                         `position:"Query" name:"ClientToken"`
	Description          string                         `position:"Query" name:"Description"`
	NextHopId            string                         `position:"Query" name:"NextHopId"`
	NextHopType          string                         `position:"Query" name:"NextHopType"`
	RouteTableId         string                         `position:"Query" name:"RouteTableId"`
	ResourceOwnerAccount string                         `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock string                         `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                         `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer               `position:"Query" name:"OwnerId"`
	PrivateIpAddress     string                         `position:"Query" name:"PrivateIpAddress"`
	NextHopList          *[]CreateRouteEntryNextHopList `position:"Query" name:"NextHopList"  type:"Repeated"`
}

// CreateRouteEntryNextHopList is a repeated param struct in CreateRouteEntryRequest
type CreateRouteEntryNextHopList struct {
	Weight      string `name:"Weight"`
	NextHopId   string `name:"NextHopId"`
	NextHopType string `name:"NextHopType"`
}

// CreateRouteEntryResponse is the response struct for api CreateRouteEntry
type CreateRouteEntryResponse struct {
	*responses.BaseResponse
	RouteEntryId string `json:"RouteEntryId" xml:"RouteEntryId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateRouteEntryRequest creates a request to invoke CreateRouteEntry API
func CreateCreateRouteEntryRequest() (request *CreateRouteEntryRequest) {
	request = &CreateRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRouteEntryResponse creates a response to parse from CreateRouteEntry response
func CreateCreateRouteEntryResponse() (response *CreateRouteEntryResponse) {
	response = &CreateRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
