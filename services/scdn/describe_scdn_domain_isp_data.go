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

// DescribeScdnDomainIspData invokes the scdn.DescribeScdnDomainIspData API synchronously
func (client *Client) DescribeScdnDomainIspData(request *DescribeScdnDomainIspDataRequest) (response *DescribeScdnDomainIspDataResponse, err error) {
	response = CreateDescribeScdnDomainIspDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainIspDataWithChan invokes the scdn.DescribeScdnDomainIspData API asynchronously
func (client *Client) DescribeScdnDomainIspDataWithChan(request *DescribeScdnDomainIspDataRequest) (<-chan *DescribeScdnDomainIspDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainIspDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainIspData(request)
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

// DescribeScdnDomainIspDataWithCallback invokes the scdn.DescribeScdnDomainIspData API asynchronously
func (client *Client) DescribeScdnDomainIspDataWithCallback(request *DescribeScdnDomainIspDataRequest, callback func(response *DescribeScdnDomainIspDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainIspDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainIspData(request)
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

// DescribeScdnDomainIspDataRequest is the request struct for api DescribeScdnDomainIspData
type DescribeScdnDomainIspDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeScdnDomainIspDataResponse is the response struct for api DescribeScdnDomainIspData
type DescribeScdnDomainIspDataResponse struct {
	*responses.BaseResponse
	EndTime      string                           `json:"EndTime" xml:"EndTime"`
	StartTime    string                           `json:"StartTime" xml:"StartTime"`
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	DomainName   string                           `json:"DomainName" xml:"DomainName"`
	DataInterval string                           `json:"DataInterval" xml:"DataInterval"`
	Value        ValueInDescribeScdnDomainIspData `json:"Value" xml:"Value"`
}

// CreateDescribeScdnDomainIspDataRequest creates a request to invoke DescribeScdnDomainIspData API
func CreateDescribeScdnDomainIspDataRequest() (request *DescribeScdnDomainIspDataRequest) {
	request = &DescribeScdnDomainIspDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainIspData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainIspDataResponse creates a response to parse from DescribeScdnDomainIspData response
func CreateDescribeScdnDomainIspDataResponse() (response *DescribeScdnDomainIspDataResponse) {
	response = &DescribeScdnDomainIspDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
