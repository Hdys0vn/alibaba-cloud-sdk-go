package alidns

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

// DescribeDomainInfo invokes the alidns.DescribeDomainInfo API synchronously
func (client *Client) DescribeDomainInfo(request *DescribeDomainInfoRequest) (response *DescribeDomainInfoResponse, err error) {
	response = CreateDescribeDomainInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainInfoWithChan invokes the alidns.DescribeDomainInfo API asynchronously
func (client *Client) DescribeDomainInfoWithChan(request *DescribeDomainInfoRequest) (<-chan *DescribeDomainInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainInfo(request)
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

// DescribeDomainInfoWithCallback invokes the alidns.DescribeDomainInfo API asynchronously
func (client *Client) DescribeDomainInfoWithCallback(request *DescribeDomainInfoRequest, callback func(response *DescribeDomainInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainInfo(request)
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

// DescribeDomainInfoRequest is the request struct for api DescribeDomainInfo
type DescribeDomainInfoRequest struct {
	*requests.RpcRequest
	DomainName           string           `position:"Query" name:"DomainName"`
	UserClientIp         string           `position:"Query" name:"UserClientIp"`
	Lang                 string           `position:"Query" name:"Lang"`
	NeedDetailAttributes requests.Boolean `position:"Query" name:"NeedDetailAttributes"`
}

// DescribeDomainInfoResponse is the response struct for api DescribeDomainInfo
type DescribeDomainInfoResponse struct {
	*responses.BaseResponse
	RecordLineTreeJson string                          `json:"RecordLineTreeJson" xml:"RecordLineTreeJson"`
	GroupName          string                          `json:"GroupName" xml:"GroupName"`
	InBlackHole        bool                            `json:"InBlackHole" xml:"InBlackHole"`
	RegionLines        bool                            `json:"RegionLines" xml:"RegionLines"`
	SlaveDns           bool                            `json:"SlaveDns" xml:"SlaveDns"`
	AliDomain          bool                            `json:"AliDomain" xml:"AliDomain"`
	RequestId          string                          `json:"RequestId" xml:"RequestId"`
	ResourceGroupId    string                          `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceId         string                          `json:"InstanceId" xml:"InstanceId"`
	DomainName         string                          `json:"DomainName" xml:"DomainName"`
	CreateTime         string                          `json:"CreateTime" xml:"CreateTime"`
	PunyCode           string                          `json:"PunyCode" xml:"PunyCode"`
	Remark             string                          `json:"Remark" xml:"Remark"`
	GroupId            string                          `json:"GroupId" xml:"GroupId"`
	VersionCode        string                          `json:"VersionCode" xml:"VersionCode"`
	DomainId           string                          `json:"DomainId" xml:"DomainId"`
	MinTtl             int64                           `json:"MinTtl" xml:"MinTtl"`
	InClean            bool                            `json:"InClean" xml:"InClean"`
	VersionName        string                          `json:"VersionName" xml:"VersionName"`
	LineType           string                          `json:"LineType" xml:"LineType"`
	DnsServers         DnsServersInDescribeDomainInfo  `json:"DnsServers" xml:"DnsServers"`
	AvailableTtls      AvailableTtls                   `json:"AvailableTtls" xml:"AvailableTtls"`
	RecordLines        RecordLinesInDescribeDomainInfo `json:"RecordLines" xml:"RecordLines"`
}

// CreateDescribeDomainInfoRequest creates a request to invoke DescribeDomainInfo API
func CreateDescribeDomainInfoRequest() (request *DescribeDomainInfoRequest) {
	request = &DescribeDomainInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainInfo", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainInfoResponse creates a response to parse from DescribeDomainInfo response
func CreateDescribeDomainInfoResponse() (response *DescribeDomainInfoResponse) {
	response = &DescribeDomainInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
