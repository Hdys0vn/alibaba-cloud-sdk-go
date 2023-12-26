package sgw

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

// DescribeSharesBucketInfoForExpressSync invokes the sgw.DescribeSharesBucketInfoForExpressSync API synchronously
func (client *Client) DescribeSharesBucketInfoForExpressSync(request *DescribeSharesBucketInfoForExpressSyncRequest) (response *DescribeSharesBucketInfoForExpressSyncResponse, err error) {
	response = CreateDescribeSharesBucketInfoForExpressSyncResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSharesBucketInfoForExpressSyncWithChan invokes the sgw.DescribeSharesBucketInfoForExpressSync API asynchronously
func (client *Client) DescribeSharesBucketInfoForExpressSyncWithChan(request *DescribeSharesBucketInfoForExpressSyncRequest) (<-chan *DescribeSharesBucketInfoForExpressSyncResponse, <-chan error) {
	responseChan := make(chan *DescribeSharesBucketInfoForExpressSyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSharesBucketInfoForExpressSync(request)
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

// DescribeSharesBucketInfoForExpressSyncWithCallback invokes the sgw.DescribeSharesBucketInfoForExpressSync API asynchronously
func (client *Client) DescribeSharesBucketInfoForExpressSyncWithCallback(request *DescribeSharesBucketInfoForExpressSyncRequest, callback func(response *DescribeSharesBucketInfoForExpressSyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSharesBucketInfoForExpressSyncResponse
		var err error
		defer close(result)
		response, err = client.DescribeSharesBucketInfoForExpressSync(request)
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

// DescribeSharesBucketInfoForExpressSyncRequest is the request struct for api DescribeSharesBucketInfoForExpressSync
type DescribeSharesBucketInfoForExpressSyncRequest struct {
	*requests.RpcRequest
	BucketRegion  string `position:"Query" name:"BucketRegion"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	BucketName    string `position:"Query" name:"BucketName"`
}

// DescribeSharesBucketInfoForExpressSyncResponse is the response struct for api DescribeSharesBucketInfoForExpressSync
type DescribeSharesBucketInfoForExpressSyncResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Success     bool        `json:"Success" xml:"Success"`
	Code        string      `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	BucketInfos BucketInfos `json:"BucketInfos" xml:"BucketInfos"`
}

// CreateDescribeSharesBucketInfoForExpressSyncRequest creates a request to invoke DescribeSharesBucketInfoForExpressSync API
func CreateDescribeSharesBucketInfoForExpressSyncRequest() (request *DescribeSharesBucketInfoForExpressSyncRequest) {
	request = &DescribeSharesBucketInfoForExpressSyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeSharesBucketInfoForExpressSync", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSharesBucketInfoForExpressSyncResponse creates a response to parse from DescribeSharesBucketInfoForExpressSync response
func CreateDescribeSharesBucketInfoForExpressSyncResponse() (response *DescribeSharesBucketInfoForExpressSyncResponse) {
	response = &DescribeSharesBucketInfoForExpressSyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}