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

// AddCustomLine invokes the alidns.AddCustomLine API synchronously
func (client *Client) AddCustomLine(request *AddCustomLineRequest) (response *AddCustomLineResponse, err error) {
	response = CreateAddCustomLineResponse()
	err = client.DoAction(request, response)
	return
}

// AddCustomLineWithChan invokes the alidns.AddCustomLine API asynchronously
func (client *Client) AddCustomLineWithChan(request *AddCustomLineRequest) (<-chan *AddCustomLineResponse, <-chan error) {
	responseChan := make(chan *AddCustomLineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCustomLine(request)
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

// AddCustomLineWithCallback invokes the alidns.AddCustomLine API asynchronously
func (client *Client) AddCustomLineWithCallback(request *AddCustomLineRequest, callback func(response *AddCustomLineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCustomLineResponse
		var err error
		defer close(result)
		response, err = client.AddCustomLine(request)
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

// AddCustomLineRequest is the request struct for api AddCustomLine
type AddCustomLineRequest struct {
	*requests.RpcRequest
	DomainName   string                    `position:"Query" name:"DomainName"`
	IpSegment    *[]AddCustomLineIpSegment `position:"Query" name:"IpSegment"  type:"Repeated"`
	UserClientIp string                    `position:"Query" name:"UserClientIp"`
	LineName     string                    `position:"Query" name:"LineName"`
	Lang         string                    `position:"Query" name:"Lang"`
}

// AddCustomLineIpSegment is a repeated param struct in AddCustomLineRequest
type AddCustomLineIpSegment struct {
	EndIp   string `name:"EndIp"`
	StartIp string `name:"StartIp"`
}

// AddCustomLineResponse is the response struct for api AddCustomLine
type AddCustomLineResponse struct {
	*responses.BaseResponse
	LineId    int64  `json:"LineId" xml:"LineId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	LineCode  string `json:"LineCode" xml:"LineCode"`
}

// CreateAddCustomLineRequest creates a request to invoke AddCustomLine API
func CreateAddCustomLineRequest() (request *AddCustomLineRequest) {
	request = &AddCustomLineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "AddCustomLine", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddCustomLineResponse creates a response to parse from AddCustomLine response
func CreateAddCustomLineResponse() (response *AddCustomLineResponse) {
	response = &AddCustomLineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
