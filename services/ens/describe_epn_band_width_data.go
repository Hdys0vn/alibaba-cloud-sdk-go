package ens

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

// DescribeEpnBandWidthData invokes the ens.DescribeEpnBandWidthData API synchronously
func (client *Client) DescribeEpnBandWidthData(request *DescribeEpnBandWidthDataRequest) (response *DescribeEpnBandWidthDataResponse, err error) {
	response = CreateDescribeEpnBandWidthDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEpnBandWidthDataWithChan invokes the ens.DescribeEpnBandWidthData API asynchronously
func (client *Client) DescribeEpnBandWidthDataWithChan(request *DescribeEpnBandWidthDataRequest) (<-chan *DescribeEpnBandWidthDataResponse, <-chan error) {
	responseChan := make(chan *DescribeEpnBandWidthDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEpnBandWidthData(request)
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

// DescribeEpnBandWidthDataWithCallback invokes the ens.DescribeEpnBandWidthData API asynchronously
func (client *Client) DescribeEpnBandWidthDataWithCallback(request *DescribeEpnBandWidthDataRequest, callback func(response *DescribeEpnBandWidthDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEpnBandWidthDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeEpnBandWidthData(request)
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

// DescribeEpnBandWidthDataRequest is the request struct for api DescribeEpnBandWidthData
type DescribeEpnBandWidthDataRequest struct {
	*requests.RpcRequest
	Isp             string `position:"Query" name:"Isp"`
	StartTime       string `position:"Query" name:"StartTime"`
	EnsRegionId     string `position:"Query" name:"EnsRegionId"`
	EPNInstanceId   string `position:"Query" name:"EPNInstanceId"`
	Period          string `position:"Query" name:"Period"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	NetworkingModel string `position:"Query" name:"NetworkingModel"`
	EndTime         string `position:"Query" name:"EndTime"`
}

// DescribeEpnBandWidthDataResponse is the response struct for api DescribeEpnBandWidthData
type DescribeEpnBandWidthDataResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	MonitorData MonitorData `json:"MonitorData" xml:"MonitorData"`
}

// CreateDescribeEpnBandWidthDataRequest creates a request to invoke DescribeEpnBandWidthData API
func CreateDescribeEpnBandWidthDataRequest() (request *DescribeEpnBandWidthDataRequest) {
	request = &DescribeEpnBandWidthDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEpnBandWidthData", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEpnBandWidthDataResponse creates a response to parse from DescribeEpnBandWidthData response
func CreateDescribeEpnBandWidthDataResponse() (response *DescribeEpnBandWidthDataResponse) {
	response = &DescribeEpnBandWidthDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
