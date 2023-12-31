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

// ListPredefinedPrivileges invokes the cloudcallcenter.ListPredefinedPrivileges API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listpredefinedprivileges.html
func (client *Client) ListPredefinedPrivileges(request *ListPredefinedPrivilegesRequest) (response *ListPredefinedPrivilegesResponse, err error) {
	response = CreateListPredefinedPrivilegesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPredefinedPrivilegesWithChan invokes the cloudcallcenter.ListPredefinedPrivileges API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listpredefinedprivileges.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPredefinedPrivilegesWithChan(request *ListPredefinedPrivilegesRequest) (<-chan *ListPredefinedPrivilegesResponse, <-chan error) {
	responseChan := make(chan *ListPredefinedPrivilegesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPredefinedPrivileges(request)
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

// ListPredefinedPrivilegesWithCallback invokes the cloudcallcenter.ListPredefinedPrivileges API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listpredefinedprivileges.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPredefinedPrivilegesWithCallback(request *ListPredefinedPrivilegesRequest, callback func(response *ListPredefinedPrivilegesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPredefinedPrivilegesResponse
		var err error
		defer close(result)
		response, err = client.ListPredefinedPrivileges(request)
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

// ListPredefinedPrivilegesRequest is the request struct for api ListPredefinedPrivileges
type ListPredefinedPrivilegesRequest struct {
	*requests.RpcRequest
}

// ListPredefinedPrivilegesResponse is the response struct for api ListPredefinedPrivileges
type ListPredefinedPrivilegesResponse struct {
	*responses.BaseResponse
	RequestId      string                               `json:"RequestId" xml:"RequestId"`
	Success        bool                                 `json:"Success" xml:"Success"`
	Code           string                               `json:"Code" xml:"Code"`
	Message        string                               `json:"Message" xml:"Message"`
	HttpStatusCode int                                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Privileges     PrivilegesInListPredefinedPrivileges `json:"Privileges" xml:"Privileges"`
}

// CreateListPredefinedPrivilegesRequest creates a request to invoke ListPredefinedPrivileges API
func CreateListPredefinedPrivilegesRequest() (request *ListPredefinedPrivilegesRequest) {
	request = &ListPredefinedPrivilegesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListPredefinedPrivileges", "", "")
	request.Method = requests.POST
	return
}

// CreateListPredefinedPrivilegesResponse creates a response to parse from ListPredefinedPrivileges response
func CreateListPredefinedPrivilegesResponse() (response *ListPredefinedPrivilegesResponse) {
	response = &ListPredefinedPrivilegesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
