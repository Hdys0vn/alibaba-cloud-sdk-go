package cdn

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

// DescribeDomainMultiUsageData invokes the cdn.DescribeDomainMultiUsageData API synchronously
func (client *Client) DescribeDomainMultiUsageData(request *DescribeDomainMultiUsageDataRequest) (response *DescribeDomainMultiUsageDataResponse, err error) {
	response = CreateDescribeDomainMultiUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainMultiUsageDataWithChan invokes the cdn.DescribeDomainMultiUsageData API asynchronously
func (client *Client) DescribeDomainMultiUsageDataWithChan(request *DescribeDomainMultiUsageDataRequest) (<-chan *DescribeDomainMultiUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainMultiUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainMultiUsageData(request)
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

// DescribeDomainMultiUsageDataWithCallback invokes the cdn.DescribeDomainMultiUsageData API asynchronously
func (client *Client) DescribeDomainMultiUsageDataWithCallback(request *DescribeDomainMultiUsageDataRequest, callback func(response *DescribeDomainMultiUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainMultiUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainMultiUsageData(request)
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

// DescribeDomainMultiUsageDataRequest is the request struct for api DescribeDomainMultiUsageData
type DescribeDomainMultiUsageDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainMultiUsageDataResponse is the response struct for api DescribeDomainMultiUsageData
type DescribeDomainMultiUsageDataResponse struct {
	*responses.BaseResponse
	EndTime            string             `json:"EndTime" xml:"EndTime"`
	StartTime          string             `json:"StartTime" xml:"StartTime"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	RequestPerInterval RequestPerInterval `json:"RequestPerInterval" xml:"RequestPerInterval"`
	TrafficPerInterval TrafficPerInterval `json:"TrafficPerInterval" xml:"TrafficPerInterval"`
}

// CreateDescribeDomainMultiUsageDataRequest creates a request to invoke DescribeDomainMultiUsageData API
func CreateDescribeDomainMultiUsageDataRequest() (request *DescribeDomainMultiUsageDataRequest) {
	request = &DescribeDomainMultiUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainMultiUsageData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainMultiUsageDataResponse creates a response to parse from DescribeDomainMultiUsageData response
func CreateDescribeDomainMultiUsageDataResponse() (response *DescribeDomainMultiUsageDataResponse) {
	response = &DescribeDomainMultiUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
