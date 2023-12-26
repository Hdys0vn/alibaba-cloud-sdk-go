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

// DescribeCerts invokes the ddoscoo.DescribeCerts API synchronously
func (client *Client) DescribeCerts(request *DescribeCertsRequest) (response *DescribeCertsResponse, err error) {
	response = CreateDescribeCertsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCertsWithChan invokes the ddoscoo.DescribeCerts API asynchronously
func (client *Client) DescribeCertsWithChan(request *DescribeCertsRequest) (<-chan *DescribeCertsResponse, <-chan error) {
	responseChan := make(chan *DescribeCertsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCerts(request)
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

// DescribeCertsWithCallback invokes the ddoscoo.DescribeCerts API asynchronously
func (client *Client) DescribeCertsWithCallback(request *DescribeCertsRequest, callback func(response *DescribeCertsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCertsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCerts(request)
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

// DescribeCertsRequest is the request struct for api DescribeCerts
type DescribeCertsRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// DescribeCertsResponse is the response struct for api DescribeCerts
type DescribeCertsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Certs     []CertItem `json:"Certs" xml:"Certs"`
}

// CreateDescribeCertsRequest creates a request to invoke DescribeCerts API
func CreateDescribeCertsRequest() (request *DescribeCertsRequest) {
	request = &DescribeCertsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeCerts", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCertsResponse creates a response to parse from DescribeCerts response
func CreateDescribeCertsResponse() (response *DescribeCertsResponse) {
	response = &DescribeCertsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
