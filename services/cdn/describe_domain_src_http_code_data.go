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

// DescribeDomainSrcHttpCodeData invokes the cdn.DescribeDomainSrcHttpCodeData API synchronously
func (client *Client) DescribeDomainSrcHttpCodeData(request *DescribeDomainSrcHttpCodeDataRequest) (response *DescribeDomainSrcHttpCodeDataResponse, err error) {
	response = CreateDescribeDomainSrcHttpCodeDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainSrcHttpCodeDataWithChan invokes the cdn.DescribeDomainSrcHttpCodeData API asynchronously
func (client *Client) DescribeDomainSrcHttpCodeDataWithChan(request *DescribeDomainSrcHttpCodeDataRequest) (<-chan *DescribeDomainSrcHttpCodeDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSrcHttpCodeDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSrcHttpCodeData(request)
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

// DescribeDomainSrcHttpCodeDataWithCallback invokes the cdn.DescribeDomainSrcHttpCodeData API asynchronously
func (client *Client) DescribeDomainSrcHttpCodeDataWithCallback(request *DescribeDomainSrcHttpCodeDataRequest, callback func(response *DescribeDomainSrcHttpCodeDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSrcHttpCodeDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSrcHttpCodeData(request)
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

// DescribeDomainSrcHttpCodeDataRequest is the request struct for api DescribeDomainSrcHttpCodeData
type DescribeDomainSrcHttpCodeDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	Interval   string `position:"Query" name:"Interval"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainSrcHttpCodeDataResponse is the response struct for api DescribeDomainSrcHttpCodeData
type DescribeDomainSrcHttpCodeDataResponse struct {
	*responses.BaseResponse
	EndTime      string                                      `json:"EndTime" xml:"EndTime"`
	StartTime    string                                      `json:"StartTime" xml:"StartTime"`
	RequestId    string                                      `json:"RequestId" xml:"RequestId"`
	DomainName   string                                      `json:"DomainName" xml:"DomainName"`
	DataInterval string                                      `json:"DataInterval" xml:"DataInterval"`
	HttpCodeData HttpCodeDataInDescribeDomainSrcHttpCodeData `json:"HttpCodeData" xml:"HttpCodeData"`
}

// CreateDescribeDomainSrcHttpCodeDataRequest creates a request to invoke DescribeDomainSrcHttpCodeData API
func CreateDescribeDomainSrcHttpCodeDataRequest() (request *DescribeDomainSrcHttpCodeDataRequest) {
	request = &DescribeDomainSrcHttpCodeDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainSrcHttpCodeData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainSrcHttpCodeDataResponse creates a response to parse from DescribeDomainSrcHttpCodeData response
func CreateDescribeDomainSrcHttpCodeDataResponse() (response *DescribeDomainSrcHttpCodeDataResponse) {
	response = &DescribeDomainSrcHttpCodeDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
