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

// DescribeCloudSiemEventDetail invokes the cloud_siem.DescribeCloudSiemEventDetail API synchronously
func (client *Client) DescribeCloudSiemEventDetail(request *DescribeCloudSiemEventDetailRequest) (response *DescribeCloudSiemEventDetailResponse, err error) {
	response = CreateDescribeCloudSiemEventDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudSiemEventDetailWithChan invokes the cloud_siem.DescribeCloudSiemEventDetail API asynchronously
func (client *Client) DescribeCloudSiemEventDetailWithChan(request *DescribeCloudSiemEventDetailRequest) (<-chan *DescribeCloudSiemEventDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudSiemEventDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudSiemEventDetail(request)
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

// DescribeCloudSiemEventDetailWithCallback invokes the cloud_siem.DescribeCloudSiemEventDetail API asynchronously
func (client *Client) DescribeCloudSiemEventDetailWithCallback(request *DescribeCloudSiemEventDetailRequest, callback func(response *DescribeCloudSiemEventDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudSiemEventDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudSiemEventDetail(request)
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

// DescribeCloudSiemEventDetailRequest is the request struct for api DescribeCloudSiemEventDetail
type DescribeCloudSiemEventDetailRequest struct {
	*requests.RpcRequest
	IncidentUuid string `position:"Body" name:"IncidentUuid"`
}

// DescribeCloudSiemEventDetailResponse is the response struct for api DescribeCloudSiemEventDetail
type DescribeCloudSiemEventDetailResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeCloudSiemEventDetailRequest creates a request to invoke DescribeCloudSiemEventDetail API
func CreateDescribeCloudSiemEventDetailRequest() (request *DescribeCloudSiemEventDetailRequest) {
	request = &DescribeCloudSiemEventDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeCloudSiemEventDetail", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudSiemEventDetailResponse creates a response to parse from DescribeCloudSiemEventDetail response
func CreateDescribeCloudSiemEventDetailResponse() (response *DescribeCloudSiemEventDetailResponse) {
	response = &DescribeCloudSiemEventDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
