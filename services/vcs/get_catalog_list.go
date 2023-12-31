package vcs

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

// GetCatalogList invokes the vcs.GetCatalogList API synchronously
func (client *Client) GetCatalogList(request *GetCatalogListRequest) (response *GetCatalogListResponse, err error) {
	response = CreateGetCatalogListResponse()
	err = client.DoAction(request, response)
	return
}

// GetCatalogListWithChan invokes the vcs.GetCatalogList API asynchronously
func (client *Client) GetCatalogListWithChan(request *GetCatalogListRequest) (<-chan *GetCatalogListResponse, <-chan error) {
	responseChan := make(chan *GetCatalogListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCatalogList(request)
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

// GetCatalogListWithCallback invokes the vcs.GetCatalogList API asynchronously
func (client *Client) GetCatalogListWithCallback(request *GetCatalogListRequest, callback func(response *GetCatalogListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCatalogListResponse
		var err error
		defer close(result)
		response, err = client.GetCatalogList(request)
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

// GetCatalogListRequest is the request struct for api GetCatalogList
type GetCatalogListRequest struct {
	*requests.RpcRequest
	IsvSubId string `position:"Query" name:"IsvSubId"`
	CorpId   string `position:"Query" name:"CorpId"`
}

// GetCatalogListResponse is the response struct for api GetCatalogList
type GetCatalogListResponse struct {
	*responses.BaseResponse
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetCatalogListRequest creates a request to invoke GetCatalogList API
func CreateGetCatalogListRequest() (request *GetCatalogListRequest) {
	request = &GetCatalogListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetCatalogList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCatalogListResponse creates a response to parse from GetCatalogList response
func CreateGetCatalogListResponse() (response *GetCatalogListResponse) {
	response = &GetCatalogListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
