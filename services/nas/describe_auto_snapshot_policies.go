package nas

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

// DescribeAutoSnapshotPolicies invokes the nas.DescribeAutoSnapshotPolicies API synchronously
func (client *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
	response = CreateDescribeAutoSnapshotPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoSnapshotPoliciesWithChan invokes the nas.DescribeAutoSnapshotPolicies API asynchronously
func (client *Client) DescribeAutoSnapshotPoliciesWithChan(request *DescribeAutoSnapshotPoliciesRequest) (<-chan *DescribeAutoSnapshotPoliciesResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoSnapshotPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoSnapshotPolicies(request)
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

// DescribeAutoSnapshotPoliciesWithCallback invokes the nas.DescribeAutoSnapshotPolicies API asynchronously
func (client *Client) DescribeAutoSnapshotPoliciesWithCallback(request *DescribeAutoSnapshotPoliciesRequest, callback func(response *DescribeAutoSnapshotPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoSnapshotPoliciesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoSnapshotPolicies(request)
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

// DescribeAutoSnapshotPoliciesRequest is the request struct for api DescribeAutoSnapshotPolicies
type DescribeAutoSnapshotPoliciesRequest struct {
	*requests.RpcRequest
	AutoSnapshotPolicyId string           `position:"Query" name:"AutoSnapshotPolicyId"`
	FileSystemType       string           `position:"Query" name:"FileSystemType"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeAutoSnapshotPoliciesResponse is the response struct for api DescribeAutoSnapshotPolicies
type DescribeAutoSnapshotPoliciesResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	AutoSnapshotPolicies AutoSnapshotPolicies `json:"AutoSnapshotPolicies" xml:"AutoSnapshotPolicies"`
}

// CreateDescribeAutoSnapshotPoliciesRequest creates a request to invoke DescribeAutoSnapshotPolicies API
func CreateDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
	request = &DescribeAutoSnapshotPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeAutoSnapshotPolicies", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutoSnapshotPoliciesResponse creates a response to parse from DescribeAutoSnapshotPolicies response
func CreateDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
	response = &DescribeAutoSnapshotPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
