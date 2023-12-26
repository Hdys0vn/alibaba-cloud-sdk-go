package oceanbasepro

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

// DescribeTenants invokes the oceanbasepro.DescribeTenants API synchronously
func (client *Client) DescribeTenants(request *DescribeTenantsRequest) (response *DescribeTenantsResponse, err error) {
	response = CreateDescribeTenantsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTenantsWithChan invokes the oceanbasepro.DescribeTenants API asynchronously
func (client *Client) DescribeTenantsWithChan(request *DescribeTenantsRequest) (<-chan *DescribeTenantsResponse, <-chan error) {
	responseChan := make(chan *DescribeTenantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTenants(request)
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

// DescribeTenantsWithCallback invokes the oceanbasepro.DescribeTenants API asynchronously
func (client *Client) DescribeTenantsWithCallback(request *DescribeTenantsRequest, callback func(response *DescribeTenantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTenantsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTenants(request)
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

// DescribeTenantsRequest is the request struct for api DescribeTenants
type DescribeTenantsRequest struct {
	*requests.RpcRequest
	SearchKey  string           `position:"Body" name:"SearchKey"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	InstanceId string           `position:"Body" name:"InstanceId"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	TenantId   string           `position:"Body" name:"TenantId"`
	TenantName string           `position:"Body" name:"TenantName"`
}

// DescribeTenantsResponse is the response struct for api DescribeTenants
type DescribeTenantsResponse struct {
	*responses.BaseResponse
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Tenants    []Data `json:"Tenants" xml:"Tenants"`
}

// CreateDescribeTenantsRequest creates a request to invoke DescribeTenants API
func CreateDescribeTenantsRequest() (request *DescribeTenantsRequest) {
	request = &DescribeTenantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeTenants", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTenantsResponse creates a response to parse from DescribeTenants response
func CreateDescribeTenantsResponse() (response *DescribeTenantsResponse) {
	response = &DescribeTenantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}