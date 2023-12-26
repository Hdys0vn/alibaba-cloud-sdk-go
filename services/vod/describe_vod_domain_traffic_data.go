package vod

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

// DescribeVodDomainTrafficData invokes the vod.DescribeVodDomainTrafficData API synchronously
func (client *Client) DescribeVodDomainTrafficData(request *DescribeVodDomainTrafficDataRequest) (response *DescribeVodDomainTrafficDataResponse, err error) {
	response = CreateDescribeVodDomainTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainTrafficDataWithChan invokes the vod.DescribeVodDomainTrafficData API asynchronously
func (client *Client) DescribeVodDomainTrafficDataWithChan(request *DescribeVodDomainTrafficDataRequest) (<-chan *DescribeVodDomainTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainTrafficData(request)
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

// DescribeVodDomainTrafficDataWithCallback invokes the vod.DescribeVodDomainTrafficData API asynchronously
func (client *Client) DescribeVodDomainTrafficDataWithCallback(request *DescribeVodDomainTrafficDataRequest, callback func(response *DescribeVodDomainTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainTrafficData(request)
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

// DescribeVodDomainTrafficDataRequest is the request struct for api DescribeVodDomainTrafficData
type DescribeVodDomainTrafficDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeVodDomainTrafficDataResponse is the response struct for api DescribeVodDomainTrafficData
type DescribeVodDomainTrafficDataResponse struct {
	*responses.BaseResponse
	EndTime                string                 `json:"EndTime" xml:"EndTime"`
	StartTime              string                 `json:"StartTime" xml:"StartTime"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	DomainName             string                 `json:"DomainName" xml:"DomainName"`
	TotalTraffic           string                 `json:"TotalTraffic" xml:"TotalTraffic"`
	DataInterval           string                 `json:"DataInterval" xml:"DataInterval"`
	TrafficDataPerInterval TrafficDataPerInterval `json:"TrafficDataPerInterval" xml:"TrafficDataPerInterval"`
}

// CreateDescribeVodDomainTrafficDataRequest creates a request to invoke DescribeVodDomainTrafficData API
func CreateDescribeVodDomainTrafficDataRequest() (request *DescribeVodDomainTrafficDataRequest) {
	request = &DescribeVodDomainTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainTrafficData", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodDomainTrafficDataResponse creates a response to parse from DescribeVodDomainTrafficData response
func CreateDescribeVodDomainTrafficDataResponse() (response *DescribeVodDomainTrafficDataResponse) {
	response = &DescribeVodDomainTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
