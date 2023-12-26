package devops_rdc

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

// InsertProjectMembers invokes the devops_rdc.InsertProjectMembers API synchronously
func (client *Client) InsertProjectMembers(request *InsertProjectMembersRequest) (response *InsertProjectMembersResponse, err error) {
	response = CreateInsertProjectMembersResponse()
	err = client.DoAction(request, response)
	return
}

// InsertProjectMembersWithChan invokes the devops_rdc.InsertProjectMembers API asynchronously
func (client *Client) InsertProjectMembersWithChan(request *InsertProjectMembersRequest) (<-chan *InsertProjectMembersResponse, <-chan error) {
	responseChan := make(chan *InsertProjectMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertProjectMembers(request)
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

// InsertProjectMembersWithCallback invokes the devops_rdc.InsertProjectMembers API asynchronously
func (client *Client) InsertProjectMembersWithCallback(request *InsertProjectMembersRequest, callback func(response *InsertProjectMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertProjectMembersResponse
		var err error
		defer close(result)
		response, err = client.InsertProjectMembers(request)
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

// InsertProjectMembersRequest is the request struct for api InsertProjectMembers
type InsertProjectMembersRequest struct {
	*requests.RpcRequest
	Members   string `position:"Body" name:"Members"`
	ProjectId string `position:"Body" name:"ProjectId"`
	OrgId     string `position:"Body" name:"OrgId"`
}

// InsertProjectMembersResponse is the response struct for api InsertProjectMembers
type InsertProjectMembersResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	Object     bool   `json:"Object" xml:"Object"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateInsertProjectMembersRequest creates a request to invoke InsertProjectMembers API
func CreateInsertProjectMembersRequest() (request *InsertProjectMembersRequest) {
	request = &InsertProjectMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "InsertProjectMembers", "", "")
	request.Method = requests.POST
	return
}

// CreateInsertProjectMembersResponse creates a response to parse from InsertProjectMembers response
func CreateInsertProjectMembersResponse() (response *InsertProjectMembersResponse) {
	response = &InsertProjectMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
