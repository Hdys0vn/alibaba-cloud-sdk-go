package ecd

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

// ListOfficeSiteUsers invokes the ecd.ListOfficeSiteUsers API synchronously
func (client *Client) ListOfficeSiteUsers(request *ListOfficeSiteUsersRequest) (response *ListOfficeSiteUsersResponse, err error) {
	response = CreateListOfficeSiteUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListOfficeSiteUsersWithChan invokes the ecd.ListOfficeSiteUsers API asynchronously
func (client *Client) ListOfficeSiteUsersWithChan(request *ListOfficeSiteUsersRequest) (<-chan *ListOfficeSiteUsersResponse, <-chan error) {
	responseChan := make(chan *ListOfficeSiteUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOfficeSiteUsers(request)
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

// ListOfficeSiteUsersWithCallback invokes the ecd.ListOfficeSiteUsers API asynchronously
func (client *Client) ListOfficeSiteUsersWithCallback(request *ListOfficeSiteUsersRequest, callback func(response *ListOfficeSiteUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOfficeSiteUsersResponse
		var err error
		defer close(result)
		response, err = client.ListOfficeSiteUsers(request)
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

// ListOfficeSiteUsersRequest is the request struct for api ListOfficeSiteUsers
type ListOfficeSiteUsersRequest struct {
	*requests.RpcRequest
	OfficeSiteId string           `position:"Query" name:"OfficeSiteId"`
	OUPath       string           `position:"Query" name:"OUPath"`
	Filter       string           `position:"Query" name:"Filter"`
	NextToken    string           `position:"Query" name:"NextToken"`
	MaxResults   requests.Integer `position:"Query" name:"MaxResults"`
}

// ListOfficeSiteUsersResponse is the response struct for api ListOfficeSiteUsers
type ListOfficeSiteUsersResponse struct {
	*responses.BaseResponse
	NextToken string `json:"NextToken" xml:"NextToken"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Users     []User `json:"Users" xml:"Users"`
}

// CreateListOfficeSiteUsersRequest creates a request to invoke ListOfficeSiteUsers API
func CreateListOfficeSiteUsersRequest() (request *ListOfficeSiteUsersRequest) {
	request = &ListOfficeSiteUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ListOfficeSiteUsers", "", "")
	request.Method = requests.POST
	return
}

// CreateListOfficeSiteUsersResponse creates a response to parse from ListOfficeSiteUsers response
func CreateListOfficeSiteUsersResponse() (response *ListOfficeSiteUsersResponse) {
	response = &ListOfficeSiteUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}