package ddoscoo

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

// DescribeDomainSecurityProfile invokes the ddoscoo.DescribeDomainSecurityProfile API synchronously
func (client *Client) DescribeDomainSecurityProfile(request *DescribeDomainSecurityProfileRequest) (response *DescribeDomainSecurityProfileResponse, err error) {
	response = CreateDescribeDomainSecurityProfileResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainSecurityProfileWithChan invokes the ddoscoo.DescribeDomainSecurityProfile API asynchronously
func (client *Client) DescribeDomainSecurityProfileWithChan(request *DescribeDomainSecurityProfileRequest) (<-chan *DescribeDomainSecurityProfileResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSecurityProfileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSecurityProfile(request)
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

// DescribeDomainSecurityProfileWithCallback invokes the ddoscoo.DescribeDomainSecurityProfile API asynchronously
func (client *Client) DescribeDomainSecurityProfileWithCallback(request *DescribeDomainSecurityProfileRequest, callback func(response *DescribeDomainSecurityProfileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSecurityProfileResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSecurityProfile(request)
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

// DescribeDomainSecurityProfileRequest is the request struct for api DescribeDomainSecurityProfile
type DescribeDomainSecurityProfileRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Domain   string `position:"Query" name:"Domain"`
}

// DescribeDomainSecurityProfileResponse is the response struct for api DescribeDomainSecurityProfile
type DescribeDomainSecurityProfileResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateDescribeDomainSecurityProfileRequest creates a request to invoke DescribeDomainSecurityProfile API
func CreateDescribeDomainSecurityProfileRequest() (request *DescribeDomainSecurityProfileRequest) {
	request = &DescribeDomainSecurityProfileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeDomainSecurityProfile", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainSecurityProfileResponse creates a response to parse from DescribeDomainSecurityProfile response
func CreateDescribeDomainSecurityProfileResponse() (response *DescribeDomainSecurityProfileResponse) {
	response = &DescribeDomainSecurityProfileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
