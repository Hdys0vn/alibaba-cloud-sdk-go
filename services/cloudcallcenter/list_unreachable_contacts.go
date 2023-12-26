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

// ListUnreachableContacts invokes the cloudcallcenter.ListUnreachableContacts API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listunreachablecontacts.html
func (client *Client) ListUnreachableContacts(request *ListUnreachableContactsRequest) (response *ListUnreachableContactsResponse, err error) {
	response = CreateListUnreachableContactsResponse()
	err = client.DoAction(request, response)
	return
}

// ListUnreachableContactsWithChan invokes the cloudcallcenter.ListUnreachableContacts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listunreachablecontacts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUnreachableContactsWithChan(request *ListUnreachableContactsRequest) (<-chan *ListUnreachableContactsResponse, <-chan error) {
	responseChan := make(chan *ListUnreachableContactsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUnreachableContacts(request)
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

// ListUnreachableContactsWithCallback invokes the cloudcallcenter.ListUnreachableContacts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listunreachablecontacts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUnreachableContactsWithCallback(request *ListUnreachableContactsRequest, callback func(response *ListUnreachableContactsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUnreachableContactsResponse
		var err error
		defer close(result)
		response, err = client.ListUnreachableContacts(request)
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

// ListUnreachableContactsRequest is the request struct for api ListUnreachableContacts
type ListUnreachableContactsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	JobGroupId string           `position:"Query" name:"JobGroupId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListUnreachableContactsResponse is the response struct for api ListUnreachableContacts
type ListUnreachableContactsResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	Success             bool                `json:"Success" xml:"Success"`
	Code                string              `json:"Code" xml:"Code"`
	Message             string              `json:"Message" xml:"Message"`
	HttpStatusCode      int                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	UnreachableContacts UnreachableContacts `json:"UnreachableContacts" xml:"UnreachableContacts"`
}

// CreateListUnreachableContactsRequest creates a request to invoke ListUnreachableContacts API
func CreateListUnreachableContactsRequest() (request *ListUnreachableContactsRequest) {
	request = &ListUnreachableContactsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListUnreachableContacts", "", "")
	request.Method = requests.POST
	return
}

// CreateListUnreachableContactsResponse creates a response to parse from ListUnreachableContacts response
func CreateListUnreachableContactsResponse() (response *ListUnreachableContactsResponse) {
	response = &ListUnreachableContactsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
