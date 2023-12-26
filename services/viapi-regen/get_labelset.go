package viapi_regen

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

// GetLabelset invokes the viapi_regen.GetLabelset API synchronously
func (client *Client) GetLabelset(request *GetLabelsetRequest) (response *GetLabelsetResponse, err error) {
	response = CreateGetLabelsetResponse()
	err = client.DoAction(request, response)
	return
}

// GetLabelsetWithChan invokes the viapi_regen.GetLabelset API asynchronously
func (client *Client) GetLabelsetWithChan(request *GetLabelsetRequest) (<-chan *GetLabelsetResponse, <-chan error) {
	responseChan := make(chan *GetLabelsetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLabelset(request)
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

// GetLabelsetWithCallback invokes the viapi_regen.GetLabelset API asynchronously
func (client *Client) GetLabelsetWithCallback(request *GetLabelsetRequest, callback func(response *GetLabelsetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLabelsetResponse
		var err error
		defer close(result)
		response, err = client.GetLabelset(request)
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

// GetLabelsetRequest is the request struct for api GetLabelset
type GetLabelsetRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// GetLabelsetResponse is the response struct for api GetLabelset
type GetLabelsetResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetLabelsetRequest creates a request to invoke GetLabelset API
func CreateGetLabelsetRequest() (request *GetLabelsetRequest) {
	request = &GetLabelsetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "GetLabelset", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLabelsetResponse creates a response to parse from GetLabelset response
func CreateGetLabelsetResponse() (response *GetLabelsetResponse) {
	response = &GetLabelsetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}