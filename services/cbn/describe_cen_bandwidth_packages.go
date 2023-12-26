package cbn

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

// DescribeCenBandwidthPackages invokes the cbn.DescribeCenBandwidthPackages API synchronously
func (client *Client) DescribeCenBandwidthPackages(request *DescribeCenBandwidthPackagesRequest) (response *DescribeCenBandwidthPackagesResponse, err error) {
	response = CreateDescribeCenBandwidthPackagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCenBandwidthPackagesWithChan invokes the cbn.DescribeCenBandwidthPackages API asynchronously
func (client *Client) DescribeCenBandwidthPackagesWithChan(request *DescribeCenBandwidthPackagesRequest) (<-chan *DescribeCenBandwidthPackagesResponse, <-chan error) {
	responseChan := make(chan *DescribeCenBandwidthPackagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCenBandwidthPackages(request)
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

// DescribeCenBandwidthPackagesWithCallback invokes the cbn.DescribeCenBandwidthPackages API asynchronously
func (client *Client) DescribeCenBandwidthPackagesWithCallback(request *DescribeCenBandwidthPackagesRequest, callback func(response *DescribeCenBandwidthPackagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCenBandwidthPackagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCenBandwidthPackages(request)
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

// DescribeCenBandwidthPackagesRequest is the request struct for api DescribeCenBandwidthPackages
type DescribeCenBandwidthPackagesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer                      `position:"Query" name:"ResourceOwnerId"`
	IncludeReservationData requests.Boolean                      `position:"Query" name:"IncludeReservationData"`
	PageNumber             requests.Integer                      `position:"Query" name:"PageNumber"`
	IsOrKey                requests.Boolean                      `position:"Query" name:"IsOrKey"`
	ResourceGroupId        string                                `position:"Query" name:"ResourceGroupId"`
	PageSize               requests.Integer                      `position:"Query" name:"PageSize"`
	Tag                    *[]DescribeCenBandwidthPackagesTag    `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount   string                                `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string                                `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer                      `position:"Query" name:"OwnerId"`
	Version                string                                `position:"Query" name:"Version"`
	Filter                 *[]DescribeCenBandwidthPackagesFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

// DescribeCenBandwidthPackagesTag is a repeated param struct in DescribeCenBandwidthPackagesRequest
type DescribeCenBandwidthPackagesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeCenBandwidthPackagesFilter is a repeated param struct in DescribeCenBandwidthPackagesRequest
type DescribeCenBandwidthPackagesFilter struct {
	Value *[]string `name:"Value" type:"Repeated"`
	Key   string    `name:"Key"`
}

// DescribeCenBandwidthPackagesResponse is the response struct for api DescribeCenBandwidthPackages
type DescribeCenBandwidthPackagesResponse struct {
	*responses.BaseResponse
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	CenBandwidthPackages CenBandwidthPackages `json:"CenBandwidthPackages" xml:"CenBandwidthPackages"`
}

// CreateDescribeCenBandwidthPackagesRequest creates a request to invoke DescribeCenBandwidthPackages API
func CreateDescribeCenBandwidthPackagesRequest() (request *DescribeCenBandwidthPackagesRequest) {
	request = &DescribeCenBandwidthPackagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenBandwidthPackages", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCenBandwidthPackagesResponse creates a response to parse from DescribeCenBandwidthPackages response
func CreateDescribeCenBandwidthPackagesResponse() (response *DescribeCenBandwidthPackagesResponse) {
	response = &DescribeCenBandwidthPackagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
