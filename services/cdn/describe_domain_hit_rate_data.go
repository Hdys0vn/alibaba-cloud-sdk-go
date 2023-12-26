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

// DescribeDomainHitRateData invokes the cdn.DescribeDomainHitRateData API synchronously
func (client *Client) DescribeDomainHitRateData(request *DescribeDomainHitRateDataRequest) (response *DescribeDomainHitRateDataResponse, err error) {
	response = CreateDescribeDomainHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainHitRateDataWithChan invokes the cdn.DescribeDomainHitRateData API asynchronously
func (client *Client) DescribeDomainHitRateDataWithChan(request *DescribeDomainHitRateDataRequest) (<-chan *DescribeDomainHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainHitRateData(request)
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

// DescribeDomainHitRateDataWithCallback invokes the cdn.DescribeDomainHitRateData API asynchronously
func (client *Client) DescribeDomainHitRateDataWithCallback(request *DescribeDomainHitRateDataRequest, callback func(response *DescribeDomainHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainHitRateData(request)
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

// DescribeDomainHitRateDataRequest is the request struct for api DescribeDomainHitRateData
type DescribeDomainHitRateDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	Interval   string `position:"Query" name:"Interval"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainHitRateDataResponse is the response struct for api DescribeDomainHitRateData
type DescribeDomainHitRateDataResponse struct {
	*responses.BaseResponse
	EndTime         string          `json:"EndTime" xml:"EndTime"`
	StartTime       string          `json:"StartTime" xml:"StartTime"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	DomainName      string          `json:"DomainName" xml:"DomainName"`
	DataInterval    string          `json:"DataInterval" xml:"DataInterval"`
	HitRateInterval HitRateInterval `json:"HitRateInterval" xml:"HitRateInterval"`
}

// CreateDescribeDomainHitRateDataRequest creates a request to invoke DescribeDomainHitRateData API
func CreateDescribeDomainHitRateDataRequest() (request *DescribeDomainHitRateDataRequest) {
	request = &DescribeDomainHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainHitRateData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainHitRateDataResponse creates a response to parse from DescribeDomainHitRateData response
func CreateDescribeDomainHitRateDataResponse() (response *DescribeDomainHitRateDataResponse) {
	response = &DescribeDomainHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
