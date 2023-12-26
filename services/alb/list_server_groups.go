package alb

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

// ListServerGroups invokes the alb.ListServerGroups API synchronously
func (client *Client) ListServerGroups(request *ListServerGroupsRequest) (response *ListServerGroupsResponse, err error) {
	response = CreateListServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListServerGroupsWithChan invokes the alb.ListServerGroups API asynchronously
func (client *Client) ListServerGroupsWithChan(request *ListServerGroupsRequest) (<-chan *ListServerGroupsResponse, <-chan error) {
	responseChan := make(chan *ListServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListServerGroups(request)
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

// ListServerGroupsWithCallback invokes the alb.ListServerGroups API asynchronously
func (client *Client) ListServerGroupsWithCallback(request *ListServerGroupsRequest, callback func(response *ListServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListServerGroups(request)
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

// ListServerGroupsRequest is the request struct for api ListServerGroups
type ListServerGroupsRequest struct {
	*requests.RpcRequest
	ServerGroupNames    *[]string              `position:"Query" name:"ServerGroupNames"  type:"Repeated"`
	ResourceGroupId     string                 `position:"Query" name:"ResourceGroupId"`
	NextToken           string                 `position:"Query" name:"NextToken"`
	Tag                 *[]ListServerGroupsTag `position:"Query" name:"Tag"  type:"Repeated"`
	ShowRelationEnabled requests.Boolean       `position:"Query" name:"ShowRelationEnabled"`
	ServerGroupIds      *[]string              `position:"Query" name:"ServerGroupIds"  type:"Repeated"`
	ServerGroupType     string                 `position:"Query" name:"ServerGroupType"`
	VpcId               string                 `position:"Query" name:"VpcId"`
	MaxResults          requests.Integer       `position:"Query" name:"MaxResults"`
}

// ListServerGroupsTag is a repeated param struct in ListServerGroupsRequest
type ListServerGroupsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListServerGroupsResponse is the response struct for api ListServerGroups
type ListServerGroupsResponse struct {
	*responses.BaseResponse
	MaxResults   int           `json:"MaxResults" xml:"MaxResults"`
	NextToken    string        `json:"NextToken" xml:"NextToken"`
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	TotalCount   int           `json:"TotalCount" xml:"TotalCount"`
	ServerGroups []ServerGroup `json:"ServerGroups" xml:"ServerGroups"`
}

// CreateListServerGroupsRequest creates a request to invoke ListServerGroups API
func CreateListServerGroupsRequest() (request *ListServerGroupsRequest) {
	request = &ListServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "ListServerGroups", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListServerGroupsResponse creates a response to parse from ListServerGroups response
func CreateListServerGroupsResponse() (response *ListServerGroupsResponse) {
	response = &ListServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
