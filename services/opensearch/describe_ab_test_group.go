package opensearch

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

// DescribeABTestGroup invokes the opensearch.DescribeABTestGroup API synchronously
func (client *Client) DescribeABTestGroup(request *DescribeABTestGroupRequest) (response *DescribeABTestGroupResponse, err error) {
	response = CreateDescribeABTestGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeABTestGroupWithChan invokes the opensearch.DescribeABTestGroup API asynchronously
func (client *Client) DescribeABTestGroupWithChan(request *DescribeABTestGroupRequest) (<-chan *DescribeABTestGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeABTestGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeABTestGroup(request)
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

// DescribeABTestGroupWithCallback invokes the opensearch.DescribeABTestGroup API asynchronously
func (client *Client) DescribeABTestGroupWithCallback(request *DescribeABTestGroupRequest, callback func(response *DescribeABTestGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeABTestGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeABTestGroup(request)
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

// DescribeABTestGroupRequest is the request struct for api DescribeABTestGroup
type DescribeABTestGroupRequest struct {
	*requests.RoaRequest
	GroupId          requests.Integer `position:"Path" name:"groupId"`
	SceneId          requests.Integer `position:"Path" name:"sceneId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// DescribeABTestGroupResponse is the response struct for api DescribeABTestGroup
type DescribeABTestGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateDescribeABTestGroupRequest creates a request to invoke DescribeABTestGroup API
func CreateDescribeABTestGroupRequest() (request *DescribeABTestGroupRequest) {
	request = &DescribeABTestGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeABTestGroup", "/v4/openapi/app-groups/[appGroupIdentity]/scenes/[sceneId]/groups/[groupId]", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeABTestGroupResponse creates a response to parse from DescribeABTestGroup response
func CreateDescribeABTestGroupResponse() (response *DescribeABTestGroupResponse) {
	response = &DescribeABTestGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
