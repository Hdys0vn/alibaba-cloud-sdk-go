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

// DescribeLiveDomainPushBpsData invokes the live.DescribeLiveDomainPushBpsData API synchronously
func (client *Client) DescribeLiveDomainPushBpsData(request *DescribeLiveDomainPushBpsDataRequest) (response *DescribeLiveDomainPushBpsDataResponse, err error) {
	response = CreateDescribeLiveDomainPushBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainPushBpsDataWithChan invokes the live.DescribeLiveDomainPushBpsData API asynchronously
func (client *Client) DescribeLiveDomainPushBpsDataWithChan(request *DescribeLiveDomainPushBpsDataRequest) (<-chan *DescribeLiveDomainPushBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainPushBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainPushBpsData(request)
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

// DescribeLiveDomainPushBpsDataWithCallback invokes the live.DescribeLiveDomainPushBpsData API asynchronously
func (client *Client) DescribeLiveDomainPushBpsDataWithCallback(request *DescribeLiveDomainPushBpsDataRequest, callback func(response *DescribeLiveDomainPushBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainPushBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainPushBpsData(request)
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

// DescribeLiveDomainPushBpsDataRequest is the request struct for api DescribeLiveDomainPushBpsData
type DescribeLiveDomainPushBpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeLiveDomainPushBpsDataResponse is the response struct for api DescribeLiveDomainPushBpsData
type DescribeLiveDomainPushBpsDataResponse struct {
	*responses.BaseResponse
	EndTime            string                                            `json:"EndTime" xml:"EndTime"`
	StartTime          string                                            `json:"StartTime" xml:"StartTime"`
	RequestId          string                                            `json:"RequestId" xml:"RequestId"`
	DomainName         string                                            `json:"DomainName" xml:"DomainName"`
	DataInterval       string                                            `json:"DataInterval" xml:"DataInterval"`
	BpsDataPerInterval BpsDataPerIntervalInDescribeLiveDomainPushBpsData `json:"BpsDataPerInterval" xml:"BpsDataPerInterval"`
}

// CreateDescribeLiveDomainPushBpsDataRequest creates a request to invoke DescribeLiveDomainPushBpsData API
func CreateDescribeLiveDomainPushBpsDataRequest() (request *DescribeLiveDomainPushBpsDataRequest) {
	request = &DescribeLiveDomainPushBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainPushBpsData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainPushBpsDataResponse creates a response to parse from DescribeLiveDomainPushBpsData response
func CreateDescribeLiveDomainPushBpsDataResponse() (response *DescribeLiveDomainPushBpsDataResponse) {
	response = &DescribeLiveDomainPushBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
