package dms_enterprise

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

// AddLhMembers invokes the dms_enterprise.AddLhMembers API synchronously
func (client *Client) AddLhMembers(request *AddLhMembersRequest) (response *AddLhMembersResponse, err error) {
	response = CreateAddLhMembersResponse()
	err = client.DoAction(request, response)
	return
}

// AddLhMembersWithChan invokes the dms_enterprise.AddLhMembers API asynchronously
func (client *Client) AddLhMembersWithChan(request *AddLhMembersRequest) (<-chan *AddLhMembersResponse, <-chan error) {
	responseChan := make(chan *AddLhMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLhMembers(request)
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

// AddLhMembersWithCallback invokes the dms_enterprise.AddLhMembers API asynchronously
func (client *Client) AddLhMembersWithCallback(request *AddLhMembersRequest, callback func(response *AddLhMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLhMembersResponse
		var err error
		defer close(result)
		response, err = client.AddLhMembers(request)
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

// AddLhMembersRequest is the request struct for api AddLhMembers
type AddLhMembersRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer       `position:"Query" name:"Tid"`
	Members    *[]AddLhMembersMembers `position:"Query" name:"Members"  type:"Json"`
	ObjectType requests.Integer       `position:"Query" name:"ObjectType"`
	ObjectId   requests.Integer       `position:"Query" name:"ObjectId"`
}

// AddLhMembersMembers is a repeated param struct in AddLhMembersRequest
type AddLhMembersMembers struct {
	Roles  *[]string `name:"Roles" type:"Repeated"`
	UserId string    `name:"UserId"`
}

// AddLhMembersResponse is the response struct for api AddLhMembers
type AddLhMembersResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateAddLhMembersRequest creates a request to invoke AddLhMembers API
func CreateAddLhMembersRequest() (request *AddLhMembersRequest) {
	request = &AddLhMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "AddLhMembers", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLhMembersResponse creates a response to parse from AddLhMembers response
func CreateAddLhMembersResponse() (response *AddLhMembersResponse) {
	response = &AddLhMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
