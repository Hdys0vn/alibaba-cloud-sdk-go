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

// DeleteRouteEntry invokes the vpc.DeleteRouteEntry API synchronously
func (client *Client) DeleteRouteEntry(request *DeleteRouteEntryRequest) (response *DeleteRouteEntryResponse, err error) {
	response = CreateDeleteRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRouteEntryWithChan invokes the vpc.DeleteRouteEntry API asynchronously
func (client *Client) DeleteRouteEntryWithChan(request *DeleteRouteEntryRequest) (<-chan *DeleteRouteEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRouteEntry(request)
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

// DeleteRouteEntryWithCallback invokes the vpc.DeleteRouteEntry API asynchronously
func (client *Client) DeleteRouteEntryWithCallback(request *DeleteRouteEntryRequest, callback func(response *DeleteRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteRouteEntry(request)
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

// DeleteRouteEntryRequest is the request struct for api DeleteRouteEntry
type DeleteRouteEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	NextHopId            string                         `position:"Query" name:"NextHopId"`
	RouteTableId         string                         `position:"Query" name:"RouteTableId"`
	ResourceOwnerAccount string                         `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock string                         `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                         `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer               `position:"Query" name:"OwnerId"`
	RouteEntryId         string                         `position:"Query" name:"RouteEntryId"`
	NextHopList          *[]DeleteRouteEntryNextHopList `position:"Query" name:"NextHopList"  type:"Repeated"`
}

// DeleteRouteEntryNextHopList is a repeated param struct in DeleteRouteEntryRequest
type DeleteRouteEntryNextHopList struct {
	NextHopId   string `name:"NextHopId"`
	NextHopType string `name:"NextHopType"`
}

// DeleteRouteEntryResponse is the response struct for api DeleteRouteEntry
type DeleteRouteEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRouteEntryRequest creates a request to invoke DeleteRouteEntry API
func CreateDeleteRouteEntryRequest() (request *DeleteRouteEntryRequest) {
	request = &DeleteRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRouteEntryResponse creates a response to parse from DeleteRouteEntry response
func CreateDeleteRouteEntryResponse() (response *DeleteRouteEntryResponse) {
	response = &DeleteRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
