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

// DescribeDomainReqHitRateData invokes the cdn.DescribeDomainReqHitRateData API synchronously
func (client *Client) DescribeDomainReqHitRateData(request *DescribeDomainReqHitRateDataRequest) (response *DescribeDomainReqHitRateDataResponse, err error) {
	response = CreateDescribeDomainReqHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainReqHitRateDataWithChan invokes the cdn.DescribeDomainReqHitRateData API asynchronously
func (client *Client) DescribeDomainReqHitRateDataWithChan(request *DescribeDomainReqHitRateDataRequest) (<-chan *DescribeDomainReqHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainReqHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainReqHitRateData(request)
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

// DescribeDomainReqHitRateDataWithCallback invokes the cdn.DescribeDomainReqHitRateData API asynchronously
func (client *Client) DescribeDomainReqHitRateDataWithCallback(request *DescribeDomainReqHitRateDataRequest, callback func(response *DescribeDomainReqHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainReqHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainReqHitRateData(request)
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

// DescribeDomainReqHitRateDataRequest is the request struct for api DescribeDomainReqHitRateData
type DescribeDomainReqHitRateDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	Interval   string `position:"Query" name:"Interval"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainReqHitRateDataResponse is the response struct for api DescribeDomainReqHitRateData
type DescribeDomainReqHitRateDataResponse struct {
	*responses.BaseResponse
	EndTime            string             `json:"EndTime" xml:"EndTime"`
	StartTime          string             `json:"StartTime" xml:"StartTime"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	DomainName         string             `json:"DomainName" xml:"DomainName"`
	DataInterval       string             `json:"DataInterval" xml:"DataInterval"`
	ReqHitRateInterval ReqHitRateInterval `json:"ReqHitRateInterval" xml:"ReqHitRateInterval"`
}

// CreateDescribeDomainReqHitRateDataRequest creates a request to invoke DescribeDomainReqHitRateData API
func CreateDescribeDomainReqHitRateDataRequest() (request *DescribeDomainReqHitRateDataRequest) {
	request = &DescribeDomainReqHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainReqHitRateData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainReqHitRateDataResponse creates a response to parse from DescribeDomainReqHitRateData response
func CreateDescribeDomainReqHitRateDataResponse() (response *DescribeDomainReqHitRateDataResponse) {
	response = &DescribeDomainReqHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
