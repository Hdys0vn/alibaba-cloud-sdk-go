package domain

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

// QueryDomainRealNameVerificationInfo invokes the domain.QueryDomainRealNameVerificationInfo API synchronously
func (client *Client) QueryDomainRealNameVerificationInfo(request *QueryDomainRealNameVerificationInfoRequest) (response *QueryDomainRealNameVerificationInfoResponse, err error) {
	response = CreateQueryDomainRealNameVerificationInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainRealNameVerificationInfoWithChan invokes the domain.QueryDomainRealNameVerificationInfo API asynchronously
func (client *Client) QueryDomainRealNameVerificationInfoWithChan(request *QueryDomainRealNameVerificationInfoRequest) (<-chan *QueryDomainRealNameVerificationInfoResponse, <-chan error) {
	responseChan := make(chan *QueryDomainRealNameVerificationInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainRealNameVerificationInfo(request)
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

// QueryDomainRealNameVerificationInfoWithCallback invokes the domain.QueryDomainRealNameVerificationInfo API asynchronously
func (client *Client) QueryDomainRealNameVerificationInfoWithCallback(request *QueryDomainRealNameVerificationInfoRequest, callback func(response *QueryDomainRealNameVerificationInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainRealNameVerificationInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainRealNameVerificationInfo(request)
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

// QueryDomainRealNameVerificationInfoRequest is the request struct for api QueryDomainRealNameVerificationInfo
type QueryDomainRealNameVerificationInfoRequest struct {
	*requests.RpcRequest
	FetchImage   requests.Boolean `position:"Query" name:"FetchImage"`
	DomainName   string           `position:"Query" name:"DomainName"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

// QueryDomainRealNameVerificationInfoResponse is the response struct for api QueryDomainRealNameVerificationInfo
type QueryDomainRealNameVerificationInfoResponse struct {
	*responses.BaseResponse
	IdentityCredentialType string `json:"IdentityCredentialType" xml:"IdentityCredentialType"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
	InstanceId             string `json:"InstanceId" xml:"InstanceId"`
	DomainName             string `json:"DomainName" xml:"DomainName"`
	IdentityCredential     string `json:"IdentityCredential" xml:"IdentityCredential"`
	SubmissionDate         string `json:"SubmissionDate" xml:"SubmissionDate"`
	IdentityCredentialNo   string `json:"IdentityCredentialNo" xml:"IdentityCredentialNo"`
	IdentityCredentialUrl  string `json:"IdentityCredentialUrl" xml:"IdentityCredentialUrl"`
}

// CreateQueryDomainRealNameVerificationInfoRequest creates a request to invoke QueryDomainRealNameVerificationInfo API
func CreateQueryDomainRealNameVerificationInfoRequest() (request *QueryDomainRealNameVerificationInfoRequest) {
	request = &QueryDomainRealNameVerificationInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainRealNameVerificationInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDomainRealNameVerificationInfoResponse creates a response to parse from QueryDomainRealNameVerificationInfo response
func CreateQueryDomainRealNameVerificationInfoResponse() (response *QueryDomainRealNameVerificationInfoResponse) {
	response = &QueryDomainRealNameVerificationInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
