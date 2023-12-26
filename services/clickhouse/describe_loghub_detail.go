package clickhouse

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

// DescribeLoghubDetail invokes the clickhouse.DescribeLoghubDetail API synchronously
func (client *Client) DescribeLoghubDetail(request *DescribeLoghubDetailRequest) (response *DescribeLoghubDetailResponse, err error) {
	response = CreateDescribeLoghubDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoghubDetailWithChan invokes the clickhouse.DescribeLoghubDetail API asynchronously
func (client *Client) DescribeLoghubDetailWithChan(request *DescribeLoghubDetailRequest) (<-chan *DescribeLoghubDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeLoghubDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoghubDetail(request)
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

// DescribeLoghubDetailWithCallback invokes the clickhouse.DescribeLoghubDetail API asynchronously
func (client *Client) DescribeLoghubDetailWithCallback(request *DescribeLoghubDetailRequest, callback func(response *DescribeLoghubDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoghubDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoghubDetail(request)
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

// DescribeLoghubDetailRequest is the request struct for api DescribeLoghubDetail
type DescribeLoghubDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ProjectName          string           `position:"Query" name:"ProjectName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ExportName           string           `position:"Query" name:"ExportName"`
}

// DescribeLoghubDetailResponse is the response struct for api DescribeLoghubDetail
type DescribeLoghubDetailResponse struct {
	*responses.BaseResponse
	RequestId  string                           `json:"RequestId" xml:"RequestId"`
	LoghubInfo LoghubInfoInDescribeLoghubDetail `json:"LoghubInfo" xml:"LoghubInfo"`
}

// CreateDescribeLoghubDetailRequest creates a request to invoke DescribeLoghubDetail API
func CreateDescribeLoghubDetailRequest() (request *DescribeLoghubDetailRequest) {
	request = &DescribeLoghubDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeLoghubDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLoghubDetailResponse creates a response to parse from DescribeLoghubDetail response
func CreateDescribeLoghubDetailResponse() (response *DescribeLoghubDetailResponse) {
	response = &DescribeLoghubDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
