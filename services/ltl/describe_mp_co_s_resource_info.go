package ltl

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

// DescribeMPCoSResourceInfo invokes the ltl.DescribeMPCoSResourceInfo API synchronously
func (client *Client) DescribeMPCoSResourceInfo(request *DescribeMPCoSResourceInfoRequest) (response *DescribeMPCoSResourceInfoResponse, err error) {
	response = CreateDescribeMPCoSResourceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMPCoSResourceInfoWithChan invokes the ltl.DescribeMPCoSResourceInfo API asynchronously
func (client *Client) DescribeMPCoSResourceInfoWithChan(request *DescribeMPCoSResourceInfoRequest) (<-chan *DescribeMPCoSResourceInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeMPCoSResourceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMPCoSResourceInfo(request)
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

// DescribeMPCoSResourceInfoWithCallback invokes the ltl.DescribeMPCoSResourceInfo API asynchronously
func (client *Client) DescribeMPCoSResourceInfoWithCallback(request *DescribeMPCoSResourceInfoRequest, callback func(response *DescribeMPCoSResourceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMPCoSResourceInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeMPCoSResourceInfo(request)
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

// DescribeMPCoSResourceInfoRequest is the request struct for api DescribeMPCoSResourceInfo
type DescribeMPCoSResourceInfoRequest struct {
	*requests.RpcRequest
	ApiVersion string `position:"Query" name:"ApiVersion"`
	BizChainId string `position:"Query" name:"BizChainId"`
}

// DescribeMPCoSResourceInfoResponse is the response struct for api DescribeMPCoSResourceInfo
type DescribeMPCoSResourceInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeMPCoSResourceInfoRequest creates a request to invoke DescribeMPCoSResourceInfo API
func CreateDescribeMPCoSResourceInfoRequest() (request *DescribeMPCoSResourceInfoRequest) {
	request = &DescribeMPCoSResourceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "DescribeMPCoSResourceInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeMPCoSResourceInfoResponse creates a response to parse from DescribeMPCoSResourceInfo response
func CreateDescribeMPCoSResourceInfoResponse() (response *DescribeMPCoSResourceInfoResponse) {
	response = &DescribeMPCoSResourceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
