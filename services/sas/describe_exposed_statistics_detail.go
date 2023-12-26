package sas

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

// DescribeExposedStatisticsDetail invokes the sas.DescribeExposedStatisticsDetail API synchronously
func (client *Client) DescribeExposedStatisticsDetail(request *DescribeExposedStatisticsDetailRequest) (response *DescribeExposedStatisticsDetailResponse, err error) {
	response = CreateDescribeExposedStatisticsDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExposedStatisticsDetailWithChan invokes the sas.DescribeExposedStatisticsDetail API asynchronously
func (client *Client) DescribeExposedStatisticsDetailWithChan(request *DescribeExposedStatisticsDetailRequest) (<-chan *DescribeExposedStatisticsDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeExposedStatisticsDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExposedStatisticsDetail(request)
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

// DescribeExposedStatisticsDetailWithCallback invokes the sas.DescribeExposedStatisticsDetail API asynchronously
func (client *Client) DescribeExposedStatisticsDetailWithCallback(request *DescribeExposedStatisticsDetailRequest, callback func(response *DescribeExposedStatisticsDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExposedStatisticsDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeExposedStatisticsDetail(request)
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

// DescribeExposedStatisticsDetailRequest is the request struct for api DescribeExposedStatisticsDetail
type DescribeExposedStatisticsDetailRequest struct {
	*requests.RpcRequest
	StatisticsType              string           `position:"Query" name:"StatisticsType"`
	StatisticsTypeGatewayType   string           `position:"Query" name:"StatisticsTypeGatewayType"`
	CurrentPage                 requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp                    string           `position:"Query" name:"SourceIp"`
	StatisticsTypeInstanceValue string           `position:"Query" name:"StatisticsTypeInstanceValue"`
	PageSize                    requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeExposedStatisticsDetailResponse is the response struct for api DescribeExposedStatisticsDetail
type DescribeExposedStatisticsDetailResponse struct {
	*responses.BaseResponse
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	PageInfo          PageInfo           `json:"PageInfo" xml:"PageInfo"`
	StatisticsDetails []StatisticsDetail `json:"StatisticsDetails" xml:"StatisticsDetails"`
}

// CreateDescribeExposedStatisticsDetailRequest creates a request to invoke DescribeExposedStatisticsDetail API
func CreateDescribeExposedStatisticsDetailRequest() (request *DescribeExposedStatisticsDetailRequest) {
	request = &DescribeExposedStatisticsDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeExposedStatisticsDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeExposedStatisticsDetailResponse creates a response to parse from DescribeExposedStatisticsDetail response
func CreateDescribeExposedStatisticsDetailResponse() (response *DescribeExposedStatisticsDetailResponse) {
	response = &DescribeExposedStatisticsDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
