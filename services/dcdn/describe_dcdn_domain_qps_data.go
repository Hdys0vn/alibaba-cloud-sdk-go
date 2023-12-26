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

// DescribeDcdnDomainQpsData invokes the dcdn.DescribeDcdnDomainQpsData API synchronously
func (client *Client) DescribeDcdnDomainQpsData(request *DescribeDcdnDomainQpsDataRequest) (response *DescribeDcdnDomainQpsDataResponse, err error) {
	response = CreateDescribeDcdnDomainQpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainQpsDataWithChan invokes the dcdn.DescribeDcdnDomainQpsData API asynchronously
func (client *Client) DescribeDcdnDomainQpsDataWithChan(request *DescribeDcdnDomainQpsDataRequest) (<-chan *DescribeDcdnDomainQpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainQpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainQpsData(request)
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

// DescribeDcdnDomainQpsDataWithCallback invokes the dcdn.DescribeDcdnDomainQpsData API asynchronously
func (client *Client) DescribeDcdnDomainQpsDataWithCallback(request *DescribeDcdnDomainQpsDataRequest, callback func(response *DescribeDcdnDomainQpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainQpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainQpsData(request)
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

// DescribeDcdnDomainQpsDataRequest is the request struct for api DescribeDcdnDomainQpsData
type DescribeDcdnDomainQpsDataRequest struct {
	*requests.RpcRequest
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	Interval       string `position:"Query" name:"Interval"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
}

// DescribeDcdnDomainQpsDataResponse is the response struct for api DescribeDcdnDomainQpsData
type DescribeDcdnDomainQpsDataResponse struct {
	*responses.BaseResponse
	EndTime            string             `json:"EndTime" xml:"EndTime"`
	StartTime          string             `json:"StartTime" xml:"StartTime"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	DomainName         string             `json:"DomainName" xml:"DomainName"`
	DataInterval       string             `json:"DataInterval" xml:"DataInterval"`
	QpsDataPerInterval QpsDataPerInterval `json:"QpsDataPerInterval" xml:"QpsDataPerInterval"`
}

// CreateDescribeDcdnDomainQpsDataRequest creates a request to invoke DescribeDcdnDomainQpsData API
func CreateDescribeDcdnDomainQpsDataRequest() (request *DescribeDcdnDomainQpsDataRequest) {
	request = &DescribeDcdnDomainQpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainQpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainQpsDataResponse creates a response to parse from DescribeDcdnDomainQpsData response
func CreateDescribeDcdnDomainQpsDataResponse() (response *DescribeDcdnDomainQpsDataResponse) {
	response = &DescribeDcdnDomainQpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
