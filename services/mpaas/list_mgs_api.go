package mpaas

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

// ListMgsApi invokes the mpaas.ListMgsApi API synchronously
func (client *Client) ListMgsApi(request *ListMgsApiRequest) (response *ListMgsApiResponse, err error) {
	response = CreateListMgsApiResponse()
	err = client.DoAction(request, response)
	return
}

// ListMgsApiWithChan invokes the mpaas.ListMgsApi API asynchronously
func (client *Client) ListMgsApiWithChan(request *ListMgsApiRequest) (<-chan *ListMgsApiResponse, <-chan error) {
	responseChan := make(chan *ListMgsApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMgsApi(request)
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

// ListMgsApiWithCallback invokes the mpaas.ListMgsApi API asynchronously
func (client *Client) ListMgsApiWithCallback(request *ListMgsApiRequest, callback func(response *ListMgsApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMgsApiResponse
		var err error
		defer close(result)
		response, err = client.ListMgsApi(request)
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

// ListMgsApiRequest is the request struct for api ListMgsApi
type ListMgsApiRequest struct {
	*requests.RpcRequest
	NeedEtag      string           `position:"Body" name:"NeedEtag"`
	ApiType       string           `position:"Body" name:"ApiType"`
	OptFuzzy      string           `position:"Body" name:"OptFuzzy"`
	Host          string           `position:"Body" name:"Host"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	TenantId      string           `position:"Body" name:"TenantId"`
	PageIndex     requests.Integer `position:"Body" name:"PageIndex"`
	ApiStatus     string           `position:"Body" name:"ApiStatus"`
	SysId         requests.Integer `position:"Body" name:"SysId"`
	Format        string           `position:"Body" name:"Format"`
	NeedEncrypt   string           `position:"Body" name:"NeedEncrypt"`
	OperationType string           `position:"Body" name:"OperationType"`
	NeedSign      string           `position:"Body" name:"NeedSign"`
	AppId         string           `position:"Body" name:"AppId"`
	SysName       string           `position:"Body" name:"SysName"`
	WorkspaceId   string           `position:"Body" name:"WorkspaceId"`
}

// ListMgsApiResponse is the response struct for api ListMgsApi
type ListMgsApiResponse struct {
	*responses.BaseResponse
	ResultMessage string                    `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                    `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                    `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInListMgsApi `json:"ResultContent" xml:"ResultContent"`
}

// CreateListMgsApiRequest creates a request to invoke ListMgsApi API
func CreateListMgsApiRequest() (request *ListMgsApiRequest) {
	request = &ListMgsApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "ListMgsApi", "", "")
	request.Method = requests.POST
	return
}

// CreateListMgsApiResponse creates a response to parse from ListMgsApi response
func CreateListMgsApiResponse() (response *ListMgsApiResponse) {
	response = &ListMgsApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
