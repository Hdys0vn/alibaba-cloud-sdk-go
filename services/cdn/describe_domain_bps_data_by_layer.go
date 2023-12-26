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

// DescribeDomainBpsDataByLayer invokes the cdn.DescribeDomainBpsDataByLayer API synchronously
func (client *Client) DescribeDomainBpsDataByLayer(request *DescribeDomainBpsDataByLayerRequest) (response *DescribeDomainBpsDataByLayerResponse, err error) {
	response = CreateDescribeDomainBpsDataByLayerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainBpsDataByLayerWithChan invokes the cdn.DescribeDomainBpsDataByLayer API asynchronously
func (client *Client) DescribeDomainBpsDataByLayerWithChan(request *DescribeDomainBpsDataByLayerRequest) (<-chan *DescribeDomainBpsDataByLayerResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainBpsDataByLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainBpsDataByLayer(request)
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

// DescribeDomainBpsDataByLayerWithCallback invokes the cdn.DescribeDomainBpsDataByLayer API asynchronously
func (client *Client) DescribeDomainBpsDataByLayerWithCallback(request *DescribeDomainBpsDataByLayerRequest, callback func(response *DescribeDomainBpsDataByLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainBpsDataByLayerResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainBpsDataByLayer(request)
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

// DescribeDomainBpsDataByLayerRequest is the request struct for api DescribeDomainBpsDataByLayer
type DescribeDomainBpsDataByLayerRequest struct {
	*requests.RpcRequest
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	Interval       string `position:"Query" name:"Interval"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	Layer          string `position:"Query" name:"Layer"`
}

// DescribeDomainBpsDataByLayerResponse is the response struct for api DescribeDomainBpsDataByLayer
type DescribeDomainBpsDataByLayerResponse struct {
	*responses.BaseResponse
	DataInterval    string          `json:"DataInterval" xml:"DataInterval"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	BpsDataInterval BpsDataInterval `json:"BpsDataInterval" xml:"BpsDataInterval"`
}

// CreateDescribeDomainBpsDataByLayerRequest creates a request to invoke DescribeDomainBpsDataByLayer API
func CreateDescribeDomainBpsDataByLayerRequest() (request *DescribeDomainBpsDataByLayerRequest) {
	request = &DescribeDomainBpsDataByLayerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainBpsDataByLayer", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainBpsDataByLayerResponse creates a response to parse from DescribeDomainBpsDataByLayer response
func CreateDescribeDomainBpsDataByLayerResponse() (response *DescribeDomainBpsDataByLayerResponse) {
	response = &DescribeDomainBpsDataByLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
