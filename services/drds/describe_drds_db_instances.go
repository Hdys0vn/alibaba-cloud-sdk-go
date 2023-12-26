package drds

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

// DescribeDrdsDbInstances invokes the drds.DescribeDrdsDbInstances API synchronously
func (client *Client) DescribeDrdsDbInstances(request *DescribeDrdsDbInstancesRequest) (response *DescribeDrdsDbInstancesResponse, err error) {
	response = CreateDescribeDrdsDbInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDbInstancesWithChan invokes the drds.DescribeDrdsDbInstances API asynchronously
func (client *Client) DescribeDrdsDbInstancesWithChan(request *DescribeDrdsDbInstancesRequest) (<-chan *DescribeDrdsDbInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDbInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDbInstances(request)
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

// DescribeDrdsDbInstancesWithCallback invokes the drds.DescribeDrdsDbInstances API asynchronously
func (client *Client) DescribeDrdsDbInstancesWithCallback(request *DescribeDrdsDbInstancesRequest, callback func(response *DescribeDrdsDbInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDbInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDbInstances(request)
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

// DescribeDrdsDbInstancesRequest is the request struct for api DescribeDrdsDbInstances
type DescribeDrdsDbInstancesRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string           `position:"Query" name:"DrdsInstanceId"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	DbName         string           `position:"Query" name:"DbName"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeDrdsDbInstancesResponse is the response struct for api DescribeDrdsDbInstances
type DescribeDrdsDbInstancesResponse struct {
	*responses.BaseResponse
	RequestId   string                               `json:"RequestId" xml:"RequestId"`
	Success     bool                                 `json:"Success" xml:"Success"`
	PageNumber  string                               `json:"PageNumber" xml:"PageNumber"`
	PageSize    string                               `json:"PageSize" xml:"PageSize"`
	Total       string                               `json:"Total" xml:"Total"`
	DbInstances DbInstancesInDescribeDrdsDbInstances `json:"DbInstances" xml:"DbInstances"`
}

// CreateDescribeDrdsDbInstancesRequest creates a request to invoke DescribeDrdsDbInstances API
func CreateDescribeDrdsDbInstancesRequest() (request *DescribeDrdsDbInstancesRequest) {
	request = &DescribeDrdsDbInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDbInstances", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDrdsDbInstancesResponse creates a response to parse from DescribeDrdsDbInstances response
func CreateDescribeDrdsDbInstancesResponse() (response *DescribeDrdsDbInstancesResponse) {
	response = &DescribeDrdsDbInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
