package vs

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

// DescribeVsDomainDetail invokes the vs.DescribeVsDomainDetail API synchronously
func (client *Client) DescribeVsDomainDetail(request *DescribeVsDomainDetailRequest) (response *DescribeVsDomainDetailResponse, err error) {
	response = CreateDescribeVsDomainDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsDomainDetailWithChan invokes the vs.DescribeVsDomainDetail API asynchronously
func (client *Client) DescribeVsDomainDetailWithChan(request *DescribeVsDomainDetailRequest) (<-chan *DescribeVsDomainDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeVsDomainDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsDomainDetail(request)
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

// DescribeVsDomainDetailWithCallback invokes the vs.DescribeVsDomainDetail API asynchronously
func (client *Client) DescribeVsDomainDetailWithCallback(request *DescribeVsDomainDetailRequest, callback func(response *DescribeVsDomainDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsDomainDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsDomainDetail(request)
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

// DescribeVsDomainDetailRequest is the request struct for api DescribeVsDomainDetail
type DescribeVsDomainDetailRequest struct {
	*requests.RpcRequest
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsDomainDetailResponse is the response struct for api DescribeVsDomainDetail
type DescribeVsDomainDetailResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	DomainConfig DomainConfig `json:"DomainConfig" xml:"DomainConfig"`
}

// CreateDescribeVsDomainDetailRequest creates a request to invoke DescribeVsDomainDetail API
func CreateDescribeVsDomainDetailRequest() (request *DescribeVsDomainDetailRequest) {
	request = &DescribeVsDomainDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsDomainDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVsDomainDetailResponse creates a response to parse from DescribeVsDomainDetail response
func CreateDescribeVsDomainDetailResponse() (response *DescribeVsDomainDetailResponse) {
	response = &DescribeVsDomainDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
