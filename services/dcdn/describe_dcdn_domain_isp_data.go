package dcdn

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

// DescribeDcdnDomainIspData invokes the dcdn.DescribeDcdnDomainIspData API synchronously
func (client *Client) DescribeDcdnDomainIspData(request *DescribeDcdnDomainIspDataRequest) (response *DescribeDcdnDomainIspDataResponse, err error) {
	response = CreateDescribeDcdnDomainIspDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainIspDataWithChan invokes the dcdn.DescribeDcdnDomainIspData API asynchronously
func (client *Client) DescribeDcdnDomainIspDataWithChan(request *DescribeDcdnDomainIspDataRequest) (<-chan *DescribeDcdnDomainIspDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainIspDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainIspData(request)
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

// DescribeDcdnDomainIspDataWithCallback invokes the dcdn.DescribeDcdnDomainIspData API asynchronously
func (client *Client) DescribeDcdnDomainIspDataWithCallback(request *DescribeDcdnDomainIspDataRequest, callback func(response *DescribeDcdnDomainIspDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainIspDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainIspData(request)
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

// DescribeDcdnDomainIspDataRequest is the request struct for api DescribeDcdnDomainIspData
type DescribeDcdnDomainIspDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDcdnDomainIspDataResponse is the response struct for api DescribeDcdnDomainIspData
type DescribeDcdnDomainIspDataResponse struct {
	*responses.BaseResponse
	EndTime      string                           `json:"EndTime" xml:"EndTime"`
	StartTime    string                           `json:"StartTime" xml:"StartTime"`
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	DomainName   string                           `json:"DomainName" xml:"DomainName"`
	DataInterval string                           `json:"DataInterval" xml:"DataInterval"`
	Value        ValueInDescribeDcdnDomainIspData `json:"Value" xml:"Value"`
}

// CreateDescribeDcdnDomainIspDataRequest creates a request to invoke DescribeDcdnDomainIspData API
func CreateDescribeDcdnDomainIspDataRequest() (request *DescribeDcdnDomainIspDataRequest) {
	request = &DescribeDcdnDomainIspDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainIspData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainIspDataResponse creates a response to parse from DescribeDcdnDomainIspData response
func CreateDescribeDcdnDomainIspDataResponse() (response *DescribeDcdnDomainIspDataResponse) {
	response = &DescribeDcdnDomainIspDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
