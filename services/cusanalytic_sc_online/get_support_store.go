package cusanalytic_sc_online

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

// GetSupportStore invokes the cusanalytic_sc_online.GetSupportStore API synchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/getsupportstore.html
func (client *Client) GetSupportStore(request *GetSupportStoreRequest) (response *GetSupportStoreResponse, err error) {
	response = CreateGetSupportStoreResponse()
	err = client.DoAction(request, response)
	return
}

// GetSupportStoreWithChan invokes the cusanalytic_sc_online.GetSupportStore API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/getsupportstore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSupportStoreWithChan(request *GetSupportStoreRequest) (<-chan *GetSupportStoreResponse, <-chan error) {
	responseChan := make(chan *GetSupportStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSupportStore(request)
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

// GetSupportStoreWithCallback invokes the cusanalytic_sc_online.GetSupportStore API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/getsupportstore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSupportStoreWithCallback(request *GetSupportStoreRequest, callback func(response *GetSupportStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSupportStoreResponse
		var err error
		defer close(result)
		response, err = client.GetSupportStore(request)
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

// GetSupportStoreRequest is the request struct for api GetSupportStore
type GetSupportStoreRequest struct {
	*requests.RpcRequest
}

// GetSupportStoreResponse is the response struct for api GetSupportStore
type GetSupportStoreResponse struct {
	*responses.BaseResponse
	StorePopDTOs StorePopDTOs `json:"StorePopDTOs" xml:"StorePopDTOs"`
}

// CreateGetSupportStoreRequest creates a request to invoke GetSupportStore API
func CreateGetSupportStoreRequest() (request *GetSupportStoreRequest) {
	request = &GetSupportStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cusanalytic_sc_online", "2019-05-24", "GetSupportStore", "", "")
	return
}

// CreateGetSupportStoreResponse creates a response to parse from GetSupportStore response
func CreateGetSupportStoreResponse() (response *GetSupportStoreResponse) {
	response = &GetSupportStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}