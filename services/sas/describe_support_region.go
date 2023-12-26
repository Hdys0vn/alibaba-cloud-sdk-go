package sas

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

// DescribeSupportRegion invokes the sas.DescribeSupportRegion API synchronously
func (client *Client) DescribeSupportRegion(request *DescribeSupportRegionRequest) (response *DescribeSupportRegionResponse, err error) {
	response = CreateDescribeSupportRegionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSupportRegionWithChan invokes the sas.DescribeSupportRegion API asynchronously
func (client *Client) DescribeSupportRegionWithChan(request *DescribeSupportRegionRequest) (<-chan *DescribeSupportRegionResponse, <-chan error) {
	responseChan := make(chan *DescribeSupportRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSupportRegion(request)
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

// DescribeSupportRegionWithCallback invokes the sas.DescribeSupportRegion API asynchronously
func (client *Client) DescribeSupportRegionWithCallback(request *DescribeSupportRegionRequest, callback func(response *DescribeSupportRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSupportRegionResponse
		var err error
		defer close(result)
		response, err = client.DescribeSupportRegion(request)
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

// DescribeSupportRegionRequest is the request struct for api DescribeSupportRegion
type DescribeSupportRegionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
}

// DescribeSupportRegionResponse is the response struct for api DescribeSupportRegion
type DescribeSupportRegionResponse struct {
	*responses.BaseResponse
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	SupportRegion []string `json:"SupportRegion" xml:"SupportRegion"`
}

// CreateDescribeSupportRegionRequest creates a request to invoke DescribeSupportRegion API
func CreateDescribeSupportRegionRequest() (request *DescribeSupportRegionRequest) {
	request = &DescribeSupportRegionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeSupportRegion", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSupportRegionResponse creates a response to parse from DescribeSupportRegion response
func CreateDescribeSupportRegionResponse() (response *DescribeSupportRegionResponse) {
	response = &DescribeSupportRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}