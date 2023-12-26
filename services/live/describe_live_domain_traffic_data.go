package live

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

// DescribeLiveDomainTrafficData invokes the live.DescribeLiveDomainTrafficData API synchronously
func (client *Client) DescribeLiveDomainTrafficData(request *DescribeLiveDomainTrafficDataRequest) (response *DescribeLiveDomainTrafficDataResponse, err error) {
	response = CreateDescribeLiveDomainTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainTrafficDataWithChan invokes the live.DescribeLiveDomainTrafficData API asynchronously
func (client *Client) DescribeLiveDomainTrafficDataWithChan(request *DescribeLiveDomainTrafficDataRequest) (<-chan *DescribeLiveDomainTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainTrafficData(request)
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

// DescribeLiveDomainTrafficDataWithCallback invokes the live.DescribeLiveDomainTrafficData API asynchronously
func (client *Client) DescribeLiveDomainTrafficDataWithCallback(request *DescribeLiveDomainTrafficDataRequest, callback func(response *DescribeLiveDomainTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainTrafficData(request)
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

// DescribeLiveDomainTrafficDataRequest is the request struct for api DescribeLiveDomainTrafficData
type DescribeLiveDomainTrafficDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeLiveDomainTrafficDataResponse is the response struct for api DescribeLiveDomainTrafficData
type DescribeLiveDomainTrafficDataResponse struct {
	*responses.BaseResponse
	EndTime                string                                                `json:"EndTime" xml:"EndTime"`
	StartTime              string                                                `json:"StartTime" xml:"StartTime"`
	RequestId              string                                                `json:"RequestId" xml:"RequestId"`
	DomainName             string                                                `json:"DomainName" xml:"DomainName"`
	DataInterval           string                                                `json:"DataInterval" xml:"DataInterval"`
	TrafficDataPerInterval TrafficDataPerIntervalInDescribeLiveDomainTrafficData `json:"TrafficDataPerInterval" xml:"TrafficDataPerInterval"`
}

// CreateDescribeLiveDomainTrafficDataRequest creates a request to invoke DescribeLiveDomainTrafficData API
func CreateDescribeLiveDomainTrafficDataRequest() (request *DescribeLiveDomainTrafficDataRequest) {
	request = &DescribeLiveDomainTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainTrafficData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainTrafficDataResponse creates a response to parse from DescribeLiveDomainTrafficData response
func CreateDescribeLiveDomainTrafficDataResponse() (response *DescribeLiveDomainTrafficDataResponse) {
	response = &DescribeLiveDomainTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
