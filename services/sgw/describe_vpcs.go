package sgw

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

// DescribeVpcs invokes the sgw.DescribeVpcs API synchronously
func (client *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	response = CreateDescribeVpcsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpcsWithChan invokes the sgw.DescribeVpcs API asynchronously
func (client *Client) DescribeVpcsWithChan(request *DescribeVpcsRequest) (<-chan *DescribeVpcsResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcs(request)
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

// DescribeVpcsWithCallback invokes the sgw.DescribeVpcs API asynchronously
func (client *Client) DescribeVpcsWithCallback(request *DescribeVpcsRequest, callback func(response *DescribeVpcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcs(request)
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

// DescribeVpcsRequest is the request struct for api DescribeVpcs
type DescribeVpcsRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	VpcId           string           `position:"Query" name:"VpcId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	StorageBundleId string           `position:"Query" name:"StorageBundleId"`
}

// DescribeVpcsResponse is the response struct for api DescribeVpcs
type DescribeVpcsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Vpcs       Vpcs   `json:"Vpcs" xml:"Vpcs"`
}

// CreateDescribeVpcsRequest creates a request to invoke DescribeVpcs API
func CreateDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeVpcs", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpcsResponse creates a response to parse from DescribeVpcs response
func CreateDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
