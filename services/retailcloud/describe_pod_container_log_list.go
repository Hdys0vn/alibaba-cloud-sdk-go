package retailcloud

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

// DescribePodContainerLogList invokes the retailcloud.DescribePodContainerLogList API synchronously
func (client *Client) DescribePodContainerLogList(request *DescribePodContainerLogListRequest) (response *DescribePodContainerLogListResponse, err error) {
	response = CreateDescribePodContainerLogListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePodContainerLogListWithChan invokes the retailcloud.DescribePodContainerLogList API asynchronously
func (client *Client) DescribePodContainerLogListWithChan(request *DescribePodContainerLogListRequest) (<-chan *DescribePodContainerLogListResponse, <-chan error) {
	responseChan := make(chan *DescribePodContainerLogListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePodContainerLogList(request)
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

// DescribePodContainerLogListWithCallback invokes the retailcloud.DescribePodContainerLogList API asynchronously
func (client *Client) DescribePodContainerLogListWithCallback(request *DescribePodContainerLogListRequest, callback func(response *DescribePodContainerLogListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePodContainerLogListResponse
		var err error
		defer close(result)
		response, err = client.DescribePodContainerLogList(request)
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

// DescribePodContainerLogListRequest is the request struct for api DescribePodContainerLogList
type DescribePodContainerLogListRequest struct {
	*requests.RpcRequest
	Line    requests.Integer `position:"Query" name:"Line"`
	AppId   requests.Integer `position:"Query" name:"AppId"`
	PodName string           `position:"Query" name:"PodName"`
	EnvId   requests.Integer `position:"Query" name:"EnvId"`
}

// DescribePodContainerLogListResponse is the response struct for api DescribePodContainerLogList
type DescribePodContainerLogListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribePodContainerLogListRequest creates a request to invoke DescribePodContainerLogList API
func CreateDescribePodContainerLogListRequest() (request *DescribePodContainerLogListRequest) {
	request = &DescribePodContainerLogListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DescribePodContainerLogList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePodContainerLogListResponse creates a response to parse from DescribePodContainerLogList response
func CreateDescribePodContainerLogListResponse() (response *DescribePodContainerLogListResponse) {
	response = &DescribePodContainerLogListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
