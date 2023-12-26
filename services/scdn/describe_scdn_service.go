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

// DescribeScdnService invokes the scdn.DescribeScdnService API synchronously
func (client *Client) DescribeScdnService(request *DescribeScdnServiceRequest) (response *DescribeScdnServiceResponse, err error) {
	response = CreateDescribeScdnServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnServiceWithChan invokes the scdn.DescribeScdnService API asynchronously
func (client *Client) DescribeScdnServiceWithChan(request *DescribeScdnServiceRequest) (<-chan *DescribeScdnServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnService(request)
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

// DescribeScdnServiceWithCallback invokes the scdn.DescribeScdnService API asynchronously
func (client *Client) DescribeScdnServiceWithCallback(request *DescribeScdnServiceRequest, callback func(response *DescribeScdnServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnService(request)
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

// DescribeScdnServiceRequest is the request struct for api DescribeScdnService
type DescribeScdnServiceRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeScdnServiceResponse is the response struct for api DescribeScdnService
type DescribeScdnServiceResponse struct {
	*responses.BaseResponse
	DomainCountValue              string         `json:"DomainCountValue" xml:"DomainCountValue"`
	DDoSBasicValue                string         `json:"DDoSBasicValue" xml:"DDoSBasicValue"`
	ChangingAffectTime            string         `json:"ChangingAffectTime" xml:"ChangingAffectTime"`
	CcProtection                  string         `json:"CcProtection" xml:"CcProtection"`
	CurrentDomainCount            string         `json:"CurrentDomainCount" xml:"CurrentDomainCount"`
	PricingCycle                  string         `json:"PricingCycle" xml:"PricingCycle"`
	OpenTime                      string         `json:"OpenTime" xml:"OpenTime"`
	ChangingChargeType            string         `json:"ChangingChargeType" xml:"ChangingChargeType"`
	RequestId                     string         `json:"RequestId" xml:"RequestId"`
	CurrentBandwidthValue         string         `json:"CurrentBandwidthValue" xml:"CurrentBandwidthValue"`
	Bandwidth                     string         `json:"Bandwidth" xml:"Bandwidth"`
	DomainCount                   string         `json:"DomainCount" xml:"DomainCount"`
	ProtectTypeValue              string         `json:"ProtectTypeValue" xml:"ProtectTypeValue"`
	CurrentBandwidth              string         `json:"CurrentBandwidth" xml:"CurrentBandwidth"`
	PriceType                     string         `json:"PriceType" xml:"PriceType"`
	CcProtectionValue             string         `json:"CcProtectionValue" xml:"CcProtectionValue"`
	CurrentDDoSBasic              string         `json:"CurrentDDoSBasic" xml:"CurrentDDoSBasic"`
	ProtectType                   string         `json:"ProtectType" xml:"ProtectType"`
	CurrentElasticProtection      string         `json:"CurrentElasticProtection" xml:"CurrentElasticProtection"`
	CurrentElasticProtectionValue string         `json:"CurrentElasticProtectionValue" xml:"CurrentElasticProtectionValue"`
	InstanceId                    string         `json:"InstanceId" xml:"InstanceId"`
	CurrentProtectType            string         `json:"CurrentProtectType" xml:"CurrentProtectType"`
	ElasticProtection             string         `json:"ElasticProtection" xml:"ElasticProtection"`
	EndTime                       string         `json:"EndTime" xml:"EndTime"`
	CurrentDDoSBasicValue         string         `json:"CurrentDDoSBasicValue" xml:"CurrentDDoSBasicValue"`
	BandwidthValue                string         `json:"BandwidthValue" xml:"BandwidthValue"`
	DDoSBasic                     string         `json:"DDoSBasic" xml:"DDoSBasic"`
	CurrentDomainCountValue       string         `json:"CurrentDomainCountValue" xml:"CurrentDomainCountValue"`
	ElasticProtectionValue        string         `json:"ElasticProtectionValue" xml:"ElasticProtectionValue"`
	CurrentCcProtection           string         `json:"CurrentCcProtection" xml:"CurrentCcProtection"`
	InternetChargeType            string         `json:"InternetChargeType" xml:"InternetChargeType"`
	CurrentProtectTypeValue       string         `json:"CurrentProtectTypeValue" xml:"CurrentProtectTypeValue"`
	CurrentCcProtectionValue      string         `json:"CurrentCcProtectionValue" xml:"CurrentCcProtectionValue"`
	OperationLocks                OperationLocks `json:"OperationLocks" xml:"OperationLocks"`
}

// CreateDescribeScdnServiceRequest creates a request to invoke DescribeScdnService API
func CreateDescribeScdnServiceRequest() (request *DescribeScdnServiceRequest) {
	request = &DescribeScdnServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnService", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnServiceResponse creates a response to parse from DescribeScdnService response
func CreateDescribeScdnServiceResponse() (response *DescribeScdnServiceResponse) {
	response = &DescribeScdnServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}