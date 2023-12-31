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

// GetMainDomainName invokes the alidns.GetMainDomainName API synchronously
func (client *Client) GetMainDomainName(request *GetMainDomainNameRequest) (response *GetMainDomainNameResponse, err error) {
	response = CreateGetMainDomainNameResponse()
	err = client.DoAction(request, response)
	return
}

// GetMainDomainNameWithChan invokes the alidns.GetMainDomainName API asynchronously
func (client *Client) GetMainDomainNameWithChan(request *GetMainDomainNameRequest) (<-chan *GetMainDomainNameResponse, <-chan error) {
	responseChan := make(chan *GetMainDomainNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMainDomainName(request)
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

// GetMainDomainNameWithCallback invokes the alidns.GetMainDomainName API asynchronously
func (client *Client) GetMainDomainNameWithCallback(request *GetMainDomainNameRequest, callback func(response *GetMainDomainNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMainDomainNameResponse
		var err error
		defer close(result)
		response, err = client.GetMainDomainName(request)
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

// GetMainDomainNameRequest is the request struct for api GetMainDomainName
type GetMainDomainNameRequest struct {
	*requests.RpcRequest
	InputString  string `position:"Query" name:"InputString"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// GetMainDomainNameResponse is the response struct for api GetMainDomainName
type GetMainDomainNameResponse struct {
	*responses.BaseResponse
	RR          string `json:"RR" xml:"RR"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DomainName  string `json:"DomainName" xml:"DomainName"`
	DomainLevel int64  `json:"DomainLevel" xml:"DomainLevel"`
}

// CreateGetMainDomainNameRequest creates a request to invoke GetMainDomainName API
func CreateGetMainDomainNameRequest() (request *GetMainDomainNameRequest) {
	request = &GetMainDomainNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "GetMainDomainName", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMainDomainNameResponse creates a response to parse from GetMainDomainName response
func CreateGetMainDomainNameResponse() (response *GetMainDomainNameResponse) {
	response = &GetMainDomainNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
