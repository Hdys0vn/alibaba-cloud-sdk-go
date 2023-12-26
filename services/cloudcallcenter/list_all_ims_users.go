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

// ListAllImsUsers invokes the cloudcallcenter.ListAllImsUsers API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listallimsusers.html
func (client *Client) ListAllImsUsers(request *ListAllImsUsersRequest) (response *ListAllImsUsersResponse, err error) {
	response = CreateListAllImsUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListAllImsUsersWithChan invokes the cloudcallcenter.ListAllImsUsers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listallimsusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAllImsUsersWithChan(request *ListAllImsUsersRequest) (<-chan *ListAllImsUsersResponse, <-chan error) {
	responseChan := make(chan *ListAllImsUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAllImsUsers(request)
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

// ListAllImsUsersWithCallback invokes the cloudcallcenter.ListAllImsUsers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listallimsusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAllImsUsersWithCallback(request *ListAllImsUsersRequest, callback func(response *ListAllImsUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAllImsUsersResponse
		var err error
		defer close(result)
		response, err = client.ListAllImsUsers(request)
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

// ListAllImsUsersRequest is the request struct for api ListAllImsUsers
type ListAllImsUsersRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListAllImsUsersResponse is the response struct for api ListAllImsUsers
type ListAllImsUsersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Users          Users  `json:"Users" xml:"Users"`
}

// CreateListAllImsUsersRequest creates a request to invoke ListAllImsUsers API
func CreateListAllImsUsersRequest() (request *ListAllImsUsersRequest) {
	request = &ListAllImsUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListAllImsUsers", "", "")
	request.Method = requests.POST
	return
}

// CreateListAllImsUsersResponse creates a response to parse from ListAllImsUsers response
func CreateListAllImsUsersResponse() (response *ListAllImsUsersResponse) {
	response = &ListAllImsUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}