package cr_ee

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

// ListRepoTag invokes the cr.ListRepoTag API synchronously
// api document: https://help.aliyun.com/api/cr/listrepotag.html
func (client *Client) ListRepoTag(request *ListRepoTagRequest) (response *ListRepoTagResponse, err error) {
	response = CreateListRepoTagResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepoTagWithChan invokes the cr.ListRepoTag API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTagWithChan(request *ListRepoTagRequest) (<-chan *ListRepoTagResponse, <-chan error) {
	responseChan := make(chan *ListRepoTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepoTag(request)
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

// ListRepoTagWithCallback invokes the cr.ListRepoTag API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTagWithCallback(request *ListRepoTagRequest, callback func(response *ListRepoTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepoTagResponse
		var err error
		defer close(result)
		response, err = client.ListRepoTag(request)
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

// ListRepoTagRequest is the request struct for api ListRepoTag
type ListRepoTagRequest struct {
	*requests.RpcRequest
	RepoId     string           `position:"Query" name:"RepoId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageNo     requests.Integer `position:"Query" name:"PageNo"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListRepoTagResponse is the response struct for api ListRepoTag
type ListRepoTagResponse struct {
	*responses.BaseResponse
	ListRepoTagIsSuccess bool         `json:"IsSuccess" xml:"IsSuccess"`
	Code                 string       `json:"Code" xml:"Code"`
	RequestId            string       `json:"RequestId" xml:"RequestId"`
	PageNo               int          `json:"PageNo" xml:"PageNo"`
	PageSize             int          `json:"PageSize" xml:"PageSize"`
	TotalCount           string       `json:"TotalCount" xml:"TotalCount"`
	Images               []ImagesItem `json:"Images" xml:"Images"`
}

// CreateListRepoTagRequest creates a request to invoke ListRepoTag API
func CreateListRepoTagRequest() (request *ListRepoTagRequest) {
	request = &ListRepoTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListRepoTag", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRepoTagResponse creates a response to parse from ListRepoTag response
func CreateListRepoTagResponse() (response *ListRepoTagResponse) {
	response = &ListRepoTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
