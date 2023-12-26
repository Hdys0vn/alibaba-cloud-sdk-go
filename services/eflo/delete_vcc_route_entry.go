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

// DeleteVccRouteEntry invokes the eflo.DeleteVccRouteEntry API synchronously
func (client *Client) DeleteVccRouteEntry(request *DeleteVccRouteEntryRequest) (response *DeleteVccRouteEntryResponse, err error) {
	response = CreateDeleteVccRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVccRouteEntryWithChan invokes the eflo.DeleteVccRouteEntry API asynchronously
func (client *Client) DeleteVccRouteEntryWithChan(request *DeleteVccRouteEntryRequest) (<-chan *DeleteVccRouteEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteVccRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVccRouteEntry(request)
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

// DeleteVccRouteEntryWithCallback invokes the eflo.DeleteVccRouteEntry API asynchronously
func (client *Client) DeleteVccRouteEntryWithCallback(request *DeleteVccRouteEntryRequest, callback func(response *DeleteVccRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVccRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteVccRouteEntry(request)
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

// DeleteVccRouteEntryRequest is the request struct for api DeleteVccRouteEntry
type DeleteVccRouteEntryRequest struct {
	*requests.RpcRequest
	DestinationCidrBlock string `position:"Body" name:"DestinationCidrBlock"`
	VccId                string `position:"Body" name:"VccId"`
	VccRouteEntryId      string `position:"Body" name:"VccRouteEntryId"`
}

// DeleteVccRouteEntryResponse is the response struct for api DeleteVccRouteEntry
type DeleteVccRouteEntryResponse struct {
	*responses.BaseResponse
}

// CreateDeleteVccRouteEntryRequest creates a request to invoke DeleteVccRouteEntry API
func CreateDeleteVccRouteEntryRequest() (request *DeleteVccRouteEntryRequest) {
	request = &DeleteVccRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "DeleteVccRouteEntry", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVccRouteEntryResponse creates a response to parse from DeleteVccRouteEntry response
func CreateDeleteVccRouteEntryResponse() (response *DeleteVccRouteEntryResponse) {
	response = &DeleteVccRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}