package ccc

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

// GetNumberLocation invokes the ccc.GetNumberLocation API synchronously
func (client *Client) GetNumberLocation(request *GetNumberLocationRequest) (response *GetNumberLocationResponse, err error) {
	response = CreateGetNumberLocationResponse()
	err = client.DoAction(request, response)
	return
}

// GetNumberLocationWithChan invokes the ccc.GetNumberLocation API asynchronously
func (client *Client) GetNumberLocationWithChan(request *GetNumberLocationRequest) (<-chan *GetNumberLocationResponse, <-chan error) {
	responseChan := make(chan *GetNumberLocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNumberLocation(request)
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

// GetNumberLocationWithCallback invokes the ccc.GetNumberLocation API asynchronously
func (client *Client) GetNumberLocationWithCallback(request *GetNumberLocationRequest, callback func(response *GetNumberLocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNumberLocationResponse
		var err error
		defer close(result)
		response, err = client.GetNumberLocation(request)
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

// GetNumberLocationRequest is the request struct for api GetNumberLocation
type GetNumberLocationRequest struct {
	*requests.RpcRequest
	Number     string `position:"Query" name:"Number"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetNumberLocationResponse is the response struct for api GetNumberLocation
type GetNumberLocationResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetNumberLocationRequest creates a request to invoke GetNumberLocation API
func CreateGetNumberLocationRequest() (request *GetNumberLocationRequest) {
	request = &GetNumberLocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetNumberLocation", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNumberLocationResponse creates a response to parse from GetNumberLocation response
func CreateGetNumberLocationResponse() (response *GetNumberLocationResponse) {
	response = &GetNumberLocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
