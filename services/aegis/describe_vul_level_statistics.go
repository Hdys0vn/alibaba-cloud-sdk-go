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

// DescribeVulLevelStatistics invokes the aegis.DescribeVulLevelStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/describevullevelstatistics.html
func (client *Client) DescribeVulLevelStatistics(request *DescribeVulLevelStatisticsRequest) (response *DescribeVulLevelStatisticsResponse, err error) {
	response = CreateDescribeVulLevelStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulLevelStatisticsWithChan invokes the aegis.DescribeVulLevelStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevullevelstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulLevelStatisticsWithChan(request *DescribeVulLevelStatisticsRequest) (<-chan *DescribeVulLevelStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeVulLevelStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulLevelStatistics(request)
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

// DescribeVulLevelStatisticsWithCallback invokes the aegis.DescribeVulLevelStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevullevelstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulLevelStatisticsWithCallback(request *DescribeVulLevelStatisticsRequest, callback func(response *DescribeVulLevelStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulLevelStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulLevelStatistics(request)
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

// DescribeVulLevelStatisticsRequest is the request struct for api DescribeVulLevelStatistics
type DescribeVulLevelStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	EndTs    requests.Integer `position:"Query" name:"EndTs"`
	StartTs  requests.Integer `position:"Query" name:"StartTs"`
	Uuids    string           `position:"Query" name:"Uuids"`
}

// DescribeVulLevelStatisticsResponse is the response struct for api DescribeVulLevelStatistics
type DescribeVulLevelStatisticsResponse struct {
	*responses.BaseResponse
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	TotalCount      int              `json:"TotalCount" xml:"TotalCount"`
	LevelStatistics []LevelStatistic `json:"LevelStatistics" xml:"LevelStatistics"`
}

// CreateDescribeVulLevelStatisticsRequest creates a request to invoke DescribeVulLevelStatistics API
func CreateDescribeVulLevelStatisticsRequest() (request *DescribeVulLevelStatisticsRequest) {
	request = &DescribeVulLevelStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeVulLevelStatistics", "vipaegis", "openAPI")
	return
}

// CreateDescribeVulLevelStatisticsResponse creates a response to parse from DescribeVulLevelStatistics response
func CreateDescribeVulLevelStatisticsResponse() (response *DescribeVulLevelStatisticsResponse) {
	response = &DescribeVulLevelStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}