package cloudcallcenter

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

// GetProvider invokes the cloudcallcenter.GetProvider API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getprovider.html
func (client *Client) GetProvider(request *GetProviderRequest) (response *GetProviderResponse, err error) {
	response = CreateGetProviderResponse()
	err = client.DoAction(request, response)
	return
}

// GetProviderWithChan invokes the cloudcallcenter.GetProvider API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getprovider.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProviderWithChan(request *GetProviderRequest) (<-chan *GetProviderResponse, <-chan error) {
	responseChan := make(chan *GetProviderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProvider(request)
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

// GetProviderWithCallback invokes the cloudcallcenter.GetProvider API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getprovider.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProviderWithCallback(request *GetProviderRequest, callback func(response *GetProviderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProviderResponse
		var err error
		defer close(result)
		response, err = client.GetProvider(request)
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

// GetProviderRequest is the request struct for api GetProvider
type GetProviderRequest struct {
	*requests.RpcRequest
}

// GetProviderResponse is the response struct for api GetProvider
type GetProviderResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Provider       string `json:"Provider" xml:"Provider"`
}

// CreateGetProviderRequest creates a request to invoke GetProvider API
func CreateGetProviderRequest() (request *GetProviderRequest) {
	request = &GetProviderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetProvider", "", "")
	request.Method = requests.POST
	return
}

// CreateGetProviderResponse creates a response to parse from GetProvider response
func CreateGetProviderResponse() (response *GetProviderResponse) {
	response = &GetProviderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}