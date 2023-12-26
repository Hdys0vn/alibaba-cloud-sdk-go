package sasti

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

// DescribeDomainReport invokes the sasti.DescribeDomainReport API synchronously
func (client *Client) DescribeDomainReport(request *DescribeDomainReportRequest) (response *DescribeDomainReportResponse, err error) {
	response = CreateDescribeDomainReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainReportWithChan invokes the sasti.DescribeDomainReport API asynchronously
func (client *Client) DescribeDomainReportWithChan(request *DescribeDomainReportRequest) (<-chan *DescribeDomainReportResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainReport(request)
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

// DescribeDomainReportWithCallback invokes the sasti.DescribeDomainReport API asynchronously
func (client *Client) DescribeDomainReportWithCallback(request *DescribeDomainReportRequest, callback func(response *DescribeDomainReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainReport(request)
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

// DescribeDomainReportRequest is the request struct for api DescribeDomainReport
type DescribeDomainReportRequest struct {
	*requests.RpcRequest
	SourceIp    string `position:"Query" name:"SourceIp"`
	Field       string `position:"Query" name:"Field"`
	Domain      string `position:"Query" name:"Domain"`
	ServiceLang string `position:"Query" name:"ServiceLang"`
}

// DescribeDomainReportResponse is the response struct for api DescribeDomainReport
type DescribeDomainReportResponse struct {
	*responses.BaseResponse
	Intelligences         string `json:"Intelligences" xml:"Intelligences"`
	Domain                string `json:"Domain" xml:"Domain"`
	SslCert               string `json:"SslCert" xml:"SslCert"`
	AttackPreferenceTop5  string `json:"AttackPreferenceTop5" xml:"AttackPreferenceTop5"`
	ThreatTypes           string `json:"ThreatTypes" xml:"ThreatTypes"`
	Confidence            string `json:"Confidence" xml:"Confidence"`
	ThreatLevel           string `json:"ThreatLevel" xml:"ThreatLevel"`
	AttackCntByThreatType string `json:"AttackCntByThreatType" xml:"AttackCntByThreatType"`
	Context               string `json:"Context" xml:"Context"`
	Whois                 string `json:"Whois" xml:"Whois"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	Scenario              string `json:"Scenario" xml:"Scenario"`
	Basic                 string `json:"Basic" xml:"Basic"`
	Group                 string `json:"Group" xml:"Group"`
}

// CreateDescribeDomainReportRequest creates a request to invoke DescribeDomainReport API
func CreateDescribeDomainReportRequest() (request *DescribeDomainReportRequest) {
	request = &DescribeDomainReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sasti", "2020-05-12", "DescribeDomainReport", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainReportResponse creates a response to parse from DescribeDomainReport response
func CreateDescribeDomainReportResponse() (response *DescribeDomainReportResponse) {
	response = &DescribeDomainReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
