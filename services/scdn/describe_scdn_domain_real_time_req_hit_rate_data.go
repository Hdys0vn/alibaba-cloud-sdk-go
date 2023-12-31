package scdn

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

// DescribeScdnDomainRealTimeReqHitRateData invokes the scdn.DescribeScdnDomainRealTimeReqHitRateData API synchronously
func (client *Client) DescribeScdnDomainRealTimeReqHitRateData(request *DescribeScdnDomainRealTimeReqHitRateDataRequest) (response *DescribeScdnDomainRealTimeReqHitRateDataResponse, err error) {
	response = CreateDescribeScdnDomainRealTimeReqHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainRealTimeReqHitRateDataWithChan invokes the scdn.DescribeScdnDomainRealTimeReqHitRateData API asynchronously
func (client *Client) DescribeScdnDomainRealTimeReqHitRateDataWithChan(request *DescribeScdnDomainRealTimeReqHitRateDataRequest) (<-chan *DescribeScdnDomainRealTimeReqHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainRealTimeReqHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainRealTimeReqHitRateData(request)
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

// DescribeScdnDomainRealTimeReqHitRateDataWithCallback invokes the scdn.DescribeScdnDomainRealTimeReqHitRateData API asynchronously
func (client *Client) DescribeScdnDomainRealTimeReqHitRateDataWithCallback(request *DescribeScdnDomainRealTimeReqHitRateDataRequest, callback func(response *DescribeScdnDomainRealTimeReqHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainRealTimeReqHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainRealTimeReqHitRateData(request)
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

// DescribeScdnDomainRealTimeReqHitRateDataRequest is the request struct for api DescribeScdnDomainRealTimeReqHitRateData
type DescribeScdnDomainRealTimeReqHitRateDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeScdnDomainRealTimeReqHitRateDataResponse is the response struct for api DescribeScdnDomainRealTimeReqHitRateData
type DescribeScdnDomainRealTimeReqHitRateDataResponse struct {
	*responses.BaseResponse
	RequestId string                                         `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeScdnDomainRealTimeReqHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeScdnDomainRealTimeReqHitRateDataRequest creates a request to invoke DescribeScdnDomainRealTimeReqHitRateData API
func CreateDescribeScdnDomainRealTimeReqHitRateDataRequest() (request *DescribeScdnDomainRealTimeReqHitRateDataRequest) {
	request = &DescribeScdnDomainRealTimeReqHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainRealTimeReqHitRateData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeScdnDomainRealTimeReqHitRateDataResponse creates a response to parse from DescribeScdnDomainRealTimeReqHitRateData response
func CreateDescribeScdnDomainRealTimeReqHitRateDataResponse() (response *DescribeScdnDomainRealTimeReqHitRateDataResponse) {
	response = &DescribeScdnDomainRealTimeReqHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
