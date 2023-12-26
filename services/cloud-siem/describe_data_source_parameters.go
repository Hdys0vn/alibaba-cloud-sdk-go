package cloud_siem

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

// DescribeDataSourceParameters invokes the cloud_siem.DescribeDataSourceParameters API synchronously
func (client *Client) DescribeDataSourceParameters(request *DescribeDataSourceParametersRequest) (response *DescribeDataSourceParametersResponse, err error) {
	response = CreateDescribeDataSourceParametersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataSourceParametersWithChan invokes the cloud_siem.DescribeDataSourceParameters API asynchronously
func (client *Client) DescribeDataSourceParametersWithChan(request *DescribeDataSourceParametersRequest) (<-chan *DescribeDataSourceParametersResponse, <-chan error) {
	responseChan := make(chan *DescribeDataSourceParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataSourceParameters(request)
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

// DescribeDataSourceParametersWithCallback invokes the cloud_siem.DescribeDataSourceParameters API asynchronously
func (client *Client) DescribeDataSourceParametersWithCallback(request *DescribeDataSourceParametersRequest, callback func(response *DescribeDataSourceParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataSourceParametersResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataSourceParameters(request)
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

// DescribeDataSourceParametersRequest is the request struct for api DescribeDataSourceParameters
type DescribeDataSourceParametersRequest struct {
	*requests.RpcRequest
	DataSourceType string `position:"Body" name:"DataSourceType"`
	CloudCode      string `position:"Body" name:"CloudCode"`
}

// DescribeDataSourceParametersResponse is the response struct for api DescribeDataSourceParameters
type DescribeDataSourceParametersResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeDataSourceParametersRequest creates a request to invoke DescribeDataSourceParameters API
func CreateDescribeDataSourceParametersRequest() (request *DescribeDataSourceParametersRequest) {
	request = &DescribeDataSourceParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeDataSourceParameters", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataSourceParametersResponse creates a response to parse from DescribeDataSourceParameters response
func CreateDescribeDataSourceParametersResponse() (response *DescribeDataSourceParametersResponse) {
	response = &DescribeDataSourceParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
