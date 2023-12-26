package aegis

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

// DescribeGroupList invokes the aegis.DescribeGroupList API synchronously
// api document: https://help.aliyun.com/api/aegis/describegrouplist.html
func (client *Client) DescribeGroupList(request *DescribeGroupListRequest) (response *DescribeGroupListResponse, err error) {
	response = CreateDescribeGroupListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGroupListWithChan invokes the aegis.DescribeGroupList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describegrouplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGroupListWithChan(request *DescribeGroupListRequest) (<-chan *DescribeGroupListResponse, <-chan error) {
	responseChan := make(chan *DescribeGroupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGroupList(request)
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

// DescribeGroupListWithCallback invokes the aegis.DescribeGroupList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describegrouplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGroupListWithCallback(request *DescribeGroupListRequest, callback func(response *DescribeGroupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGroupListResponse
		var err error
		defer close(result)
		response, err = client.DescribeGroupList(request)
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

// DescribeGroupListRequest is the request struct for api DescribeGroupList
type DescribeGroupListRequest struct {
	*requests.RpcRequest
	WarnLevel   string           `position:"Query" name:"WarnLevel"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	RuleGroupId requests.Integer `position:"Query" name:"RuleGroupId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	ExGroupId   requests.Integer `position:"Query" name:"ExGroupId"`
}

// DescribeGroupListResponse is the response struct for api DescribeGroupList
type DescribeGroupListResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Success    bool             `json:"Success" xml:"Success"`
	PageInfo   PageInfo         `json:"PageInfo" xml:"PageInfo"`
	RuleGroups []RuleGroupsItem `json:"RuleGroups" xml:"RuleGroups"`
}

// CreateDescribeGroupListRequest creates a request to invoke DescribeGroupList API
func CreateDescribeGroupListRequest() (request *DescribeGroupListRequest) {
	request = &DescribeGroupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeGroupList", "vipaegis", "openAPI")
	return
}

// CreateDescribeGroupListResponse creates a response to parse from DescribeGroupList response
func CreateDescribeGroupListResponse() (response *DescribeGroupListResponse) {
	response = &DescribeGroupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
