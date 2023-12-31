package yundun_ds

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

// DescribeFlowTotalCount invokes the yundun_ds.DescribeFlowTotalCount API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeflowtotalcount.html
func (client *Client) DescribeFlowTotalCount(request *DescribeFlowTotalCountRequest) (response *DescribeFlowTotalCountResponse, err error) {
	response = CreateDescribeFlowTotalCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowTotalCountWithChan invokes the yundun_ds.DescribeFlowTotalCount API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeflowtotalcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowTotalCountWithChan(request *DescribeFlowTotalCountRequest) (<-chan *DescribeFlowTotalCountResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowTotalCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowTotalCount(request)
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

// DescribeFlowTotalCountWithCallback invokes the yundun_ds.DescribeFlowTotalCount API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeflowtotalcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowTotalCountWithCallback(request *DescribeFlowTotalCountRequest, callback func(response *DescribeFlowTotalCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowTotalCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowTotalCount(request)
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

// DescribeFlowTotalCountRequest is the request struct for api DescribeFlowTotalCount
type DescribeFlowTotalCountRequest struct {
	*requests.RpcRequest
	ProductCode string           `position:"Query" name:"ProductCode"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	DepartId    requests.Integer `position:"Query" name:"DepartId"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DescribeFlowTotalCountResponse is the response struct for api DescribeFlowTotalCount
type DescribeFlowTotalCountResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	FlowCount FlowCount `json:"FlowCount" xml:"FlowCount"`
}

// CreateDescribeFlowTotalCountRequest creates a request to invoke DescribeFlowTotalCount API
func CreateDescribeFlowTotalCountRequest() (request *DescribeFlowTotalCountRequest) {
	request = &DescribeFlowTotalCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeFlowTotalCount", "sddp", "openAPI")
	return
}

// CreateDescribeFlowTotalCountResponse creates a response to parse from DescribeFlowTotalCount response
func CreateDescribeFlowTotalCountResponse() (response *DescribeFlowTotalCountResponse) {
	response = &DescribeFlowTotalCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
