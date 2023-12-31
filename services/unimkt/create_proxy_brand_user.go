package unimkt

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

// CreateProxyBrandUser invokes the unimkt.CreateProxyBrandUser API synchronously
func (client *Client) CreateProxyBrandUser(request *CreateProxyBrandUserRequest) (response *CreateProxyBrandUserResponse, err error) {
	response = CreateCreateProxyBrandUserResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProxyBrandUserWithChan invokes the unimkt.CreateProxyBrandUser API asynchronously
func (client *Client) CreateProxyBrandUserWithChan(request *CreateProxyBrandUserRequest) (<-chan *CreateProxyBrandUserResponse, <-chan error) {
	responseChan := make(chan *CreateProxyBrandUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProxyBrandUser(request)
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

// CreateProxyBrandUserWithCallback invokes the unimkt.CreateProxyBrandUser API asynchronously
func (client *Client) CreateProxyBrandUserWithCallback(request *CreateProxyBrandUserRequest, callback func(response *CreateProxyBrandUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProxyBrandUserResponse
		var err error
		defer close(result)
		response, err = client.CreateProxyBrandUser(request)
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

// CreateProxyBrandUserRequest is the request struct for api CreateProxyBrandUser
type CreateProxyBrandUserRequest struct {
	*requests.RpcRequest
	ContactName   string           `position:"Query" name:"ContactName"`
	ClientToken   string           `position:"Query" name:"ClientToken"`
	Company       string           `position:"Query" name:"Company"`
	Industry      string           `position:"Query" name:"Industry"`
	BrandUserNick string           `position:"Query" name:"BrandUserNick"`
	ContactPhone  string           `position:"Query" name:"ContactPhone"`
	ProxyUserId   requests.Integer `position:"Query" name:"ProxyUserId"`
	ChannelId     string           `position:"Query" name:"ChannelId"`
}

// CreateProxyBrandUserResponse is the response struct for api CreateProxyBrandUser
type CreateProxyBrandUserResponse struct {
	*responses.BaseResponse
	Code      int                      `json:"Code" xml:"Code"`
	Success   bool                     `json:"Success" xml:"Success"`
	ErrorMsg  string                   `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Data      []map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateCreateProxyBrandUserRequest creates a request to invoke CreateProxyBrandUser API
func CreateCreateProxyBrandUserRequest() (request *CreateProxyBrandUserRequest) {
	request = &CreateProxyBrandUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "CreateProxyBrandUser", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProxyBrandUserResponse creates a response to parse from CreateProxyBrandUser response
func CreateCreateProxyBrandUserResponse() (response *CreateProxyBrandUserResponse) {
	response = &CreateProxyBrandUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
