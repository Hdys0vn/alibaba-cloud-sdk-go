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

// DescribeScdnDomainRealTimeSrcBpsData invokes the scdn.DescribeScdnDomainRealTimeSrcBpsData API synchronously
func (client *Client) DescribeScdnDomainRealTimeSrcBpsData(request *DescribeScdnDomainRealTimeSrcBpsDataRequest) (response *DescribeScdnDomainRealTimeSrcBpsDataResponse, err error) {
	response = CreateDescribeScdnDomainRealTimeSrcBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainRealTimeSrcBpsDataWithChan invokes the scdn.DescribeScdnDomainRealTimeSrcBpsData API asynchronously
func (client *Client) DescribeScdnDomainRealTimeSrcBpsDataWithChan(request *DescribeScdnDomainRealTimeSrcBpsDataRequest) (<-chan *DescribeScdnDomainRealTimeSrcBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainRealTimeSrcBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainRealTimeSrcBpsData(request)
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

// DescribeScdnDomainRealTimeSrcBpsDataWithCallback invokes the scdn.DescribeScdnDomainRealTimeSrcBpsData API asynchronously
func (client *Client) DescribeScdnDomainRealTimeSrcBpsDataWithCallback(request *DescribeScdnDomainRealTimeSrcBpsDataRequest, callback func(response *DescribeScdnDomainRealTimeSrcBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainRealTimeSrcBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainRealTimeSrcBpsData(request)
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

// DescribeScdnDomainRealTimeSrcBpsDataRequest is the request struct for api DescribeScdnDomainRealTimeSrcBpsData
type DescribeScdnDomainRealTimeSrcBpsDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeScdnDomainRealTimeSrcBpsDataResponse is the response struct for api DescribeScdnDomainRealTimeSrcBpsData
type DescribeScdnDomainRealTimeSrcBpsDataResponse struct {
	*responses.BaseResponse
	EndTime                       string                        `json:"EndTime" xml:"EndTime"`
	StartTime                     string                        `json:"StartTime" xml:"StartTime"`
	RequestId                     string                        `json:"RequestId" xml:"RequestId"`
	DomainName                    string                        `json:"DomainName" xml:"DomainName"`
	DataInterval                  string                        `json:"DataInterval" xml:"DataInterval"`
	RealTimeSrcBpsDataPerInterval RealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval" xml:"RealTimeSrcBpsDataPerInterval"`
}

// CreateDescribeScdnDomainRealTimeSrcBpsDataRequest creates a request to invoke DescribeScdnDomainRealTimeSrcBpsData API
func CreateDescribeScdnDomainRealTimeSrcBpsDataRequest() (request *DescribeScdnDomainRealTimeSrcBpsDataRequest) {
	request = &DescribeScdnDomainRealTimeSrcBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainRealTimeSrcBpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainRealTimeSrcBpsDataResponse creates a response to parse from DescribeScdnDomainRealTimeSrcBpsData response
func CreateDescribeScdnDomainRealTimeSrcBpsDataResponse() (response *DescribeScdnDomainRealTimeSrcBpsDataResponse) {
	response = &DescribeScdnDomainRealTimeSrcBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
