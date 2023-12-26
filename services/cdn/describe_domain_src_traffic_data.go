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

// DescribeDomainSrcTrafficData invokes the cdn.DescribeDomainSrcTrafficData API synchronously
func (client *Client) DescribeDomainSrcTrafficData(request *DescribeDomainSrcTrafficDataRequest) (response *DescribeDomainSrcTrafficDataResponse, err error) {
	response = CreateDescribeDomainSrcTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainSrcTrafficDataWithChan invokes the cdn.DescribeDomainSrcTrafficData API asynchronously
func (client *Client) DescribeDomainSrcTrafficDataWithChan(request *DescribeDomainSrcTrafficDataRequest) (<-chan *DescribeDomainSrcTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSrcTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSrcTrafficData(request)
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

// DescribeDomainSrcTrafficDataWithCallback invokes the cdn.DescribeDomainSrcTrafficData API asynchronously
func (client *Client) DescribeDomainSrcTrafficDataWithCallback(request *DescribeDomainSrcTrafficDataRequest, callback func(response *DescribeDomainSrcTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSrcTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSrcTrafficData(request)
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

// DescribeDomainSrcTrafficDataRequest is the request struct for api DescribeDomainSrcTrafficData
type DescribeDomainSrcTrafficDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	Interval   string `position:"Query" name:"Interval"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainSrcTrafficDataResponse is the response struct for api DescribeDomainSrcTrafficData
type DescribeDomainSrcTrafficDataResponse struct {
	*responses.BaseResponse
	EndTime                   string                    `json:"EndTime" xml:"EndTime"`
	StartTime                 string                    `json:"StartTime" xml:"StartTime"`
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	DomainName                string                    `json:"DomainName" xml:"DomainName"`
	DataInterval              string                    `json:"DataInterval" xml:"DataInterval"`
	SrcTrafficDataPerInterval SrcTrafficDataPerInterval `json:"SrcTrafficDataPerInterval" xml:"SrcTrafficDataPerInterval"`
}

// CreateDescribeDomainSrcTrafficDataRequest creates a request to invoke DescribeDomainSrcTrafficData API
func CreateDescribeDomainSrcTrafficDataRequest() (request *DescribeDomainSrcTrafficDataRequest) {
	request = &DescribeDomainSrcTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainSrcTrafficData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainSrcTrafficDataResponse creates a response to parse from DescribeDomainSrcTrafficData response
func CreateDescribeDomainSrcTrafficDataResponse() (response *DescribeDomainSrcTrafficDataResponse) {
	response = &DescribeDomainSrcTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
