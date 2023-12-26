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

// ListImsUsers invokes the cloudcallcenter.ListImsUsers API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listimsusers.html
func (client *Client) ListImsUsers(request *ListImsUsersRequest) (response *ListImsUsersResponse, err error) {
	response = CreateListImsUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListImsUsersWithChan invokes the cloudcallcenter.ListImsUsers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listimsusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListImsUsersWithChan(request *ListImsUsersRequest) (<-chan *ListImsUsersResponse, <-chan error) {
	responseChan := make(chan *ListImsUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListImsUsers(request)
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

// ListImsUsersWithCallback invokes the cloudcallcenter.ListImsUsers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listimsusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListImsUsersWithCallback(request *ListImsUsersRequest, callback func(response *ListImsUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListImsUsersResponse
		var err error
		defer close(result)
		response, err = client.ListImsUsers(request)
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

// ListImsUsersRequest is the request struct for api ListImsUsers
type ListImsUsersRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListImsUsersResponse is the response struct for api ListImsUsers
type ListImsUsersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Users          Users  `json:"Users" xml:"Users"`
}

// CreateListImsUsersRequest creates a request to invoke ListImsUsers API
func CreateListImsUsersRequest() (request *ListImsUsersRequest) {
	request = &ListImsUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListImsUsers", "", "")
	request.Method = requests.POST
	return
}

// CreateListImsUsersResponse creates a response to parse from ListImsUsers response
func CreateListImsUsersResponse() (response *ListImsUsersResponse) {
	response = &ListImsUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
