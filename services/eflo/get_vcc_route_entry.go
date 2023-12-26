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

// GetVccRouteEntry invokes the eflo.GetVccRouteEntry API synchronously
func (client *Client) GetVccRouteEntry(request *GetVccRouteEntryRequest) (response *GetVccRouteEntryResponse, err error) {
	response = CreateGetVccRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// GetVccRouteEntryWithChan invokes the eflo.GetVccRouteEntry API asynchronously
func (client *Client) GetVccRouteEntryWithChan(request *GetVccRouteEntryRequest) (<-chan *GetVccRouteEntryResponse, <-chan error) {
	responseChan := make(chan *GetVccRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVccRouteEntry(request)
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

// GetVccRouteEntryWithCallback invokes the eflo.GetVccRouteEntry API asynchronously
func (client *Client) GetVccRouteEntryWithCallback(request *GetVccRouteEntryRequest, callback func(response *GetVccRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVccRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.GetVccRouteEntry(request)
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

// GetVccRouteEntryRequest is the request struct for api GetVccRouteEntry
type GetVccRouteEntryRequest struct {
	*requests.RpcRequest
	VccId           string `position:"Body" name:"VccId"`
	VccRouteEntryId string `position:"Body" name:"VccRouteEntryId"`
}

// GetVccRouteEntryResponse is the response struct for api GetVccRouteEntry
type GetVccRouteEntryResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetVccRouteEntryRequest creates a request to invoke GetVccRouteEntry API
func CreateGetVccRouteEntryRequest() (request *GetVccRouteEntryRequest) {
	request = &GetVccRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetVccRouteEntry", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVccRouteEntryResponse creates a response to parse from GetVccRouteEntry response
func CreateGetVccRouteEntryResponse() (response *GetVccRouteEntryResponse) {
	response = &GetVccRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
