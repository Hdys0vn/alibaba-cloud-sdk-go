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

// GetErRouteEntry invokes the eflo.GetErRouteEntry API synchronously
func (client *Client) GetErRouteEntry(request *GetErRouteEntryRequest) (response *GetErRouteEntryResponse, err error) {
	response = CreateGetErRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// GetErRouteEntryWithChan invokes the eflo.GetErRouteEntry API asynchronously
func (client *Client) GetErRouteEntryWithChan(request *GetErRouteEntryRequest) (<-chan *GetErRouteEntryResponse, <-chan error) {
	responseChan := make(chan *GetErRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetErRouteEntry(request)
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

// GetErRouteEntryWithCallback invokes the eflo.GetErRouteEntry API asynchronously
func (client *Client) GetErRouteEntryWithCallback(request *GetErRouteEntryRequest, callback func(response *GetErRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetErRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.GetErRouteEntry(request)
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

// GetErRouteEntryRequest is the request struct for api GetErRouteEntry
type GetErRouteEntryRequest struct {
	*requests.RpcRequest
	ErId           string `position:"Body" name:"ErId"`
	ErRouteEntryId string `position:"Body" name:"ErRouteEntryId"`
}

// GetErRouteEntryResponse is the response struct for api GetErRouteEntry
type GetErRouteEntryResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetErRouteEntryRequest creates a request to invoke GetErRouteEntry API
func CreateGetErRouteEntryRequest() (request *GetErRouteEntryRequest) {
	request = &GetErRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetErRouteEntry", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetErRouteEntryResponse creates a response to parse from GetErRouteEntry response
func CreateGetErRouteEntryResponse() (response *GetErRouteEntryResponse) {
	response = &GetErRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
