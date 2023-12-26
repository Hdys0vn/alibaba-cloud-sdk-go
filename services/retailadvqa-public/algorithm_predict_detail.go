package retailadvqa_public

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

// AlgorithmPredictDetail invokes the retailadvqa_public.AlgorithmPredictDetail API synchronously
func (client *Client) AlgorithmPredictDetail(request *AlgorithmPredictDetailRequest) (response *AlgorithmPredictDetailResponse, err error) {
	response = CreateAlgorithmPredictDetailResponse()
	err = client.DoAction(request, response)
	return
}

// AlgorithmPredictDetailWithChan invokes the retailadvqa_public.AlgorithmPredictDetail API asynchronously
func (client *Client) AlgorithmPredictDetailWithChan(request *AlgorithmPredictDetailRequest) (<-chan *AlgorithmPredictDetailResponse, <-chan error) {
	responseChan := make(chan *AlgorithmPredictDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AlgorithmPredictDetail(request)
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

// AlgorithmPredictDetailWithCallback invokes the retailadvqa_public.AlgorithmPredictDetail API asynchronously
func (client *Client) AlgorithmPredictDetailWithCallback(request *AlgorithmPredictDetailRequest, callback func(response *AlgorithmPredictDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AlgorithmPredictDetailResponse
		var err error
		defer close(result)
		response, err = client.AlgorithmPredictDetail(request)
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

// AlgorithmPredictDetailRequest is the request struct for api AlgorithmPredictDetail
type AlgorithmPredictDetailRequest struct {
	*requests.RpcRequest
	AccessId    string           `position:"Query" name:"AccessId"`
	PredictId   requests.Integer `position:"Query" name:"PredictId"`
	TenantId    string           `position:"Query" name:"TenantId"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// AlgorithmPredictDetailResponse is the response struct for api AlgorithmPredictDetail
type AlgorithmPredictDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAlgorithmPredictDetailRequest creates a request to invoke AlgorithmPredictDetail API
func CreateAlgorithmPredictDetailRequest() (request *AlgorithmPredictDetailRequest) {
	request = &AlgorithmPredictDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "AlgorithmPredictDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateAlgorithmPredictDetailResponse creates a response to parse from AlgorithmPredictDetail response
func CreateAlgorithmPredictDetailResponse() (response *AlgorithmPredictDetailResponse) {
	response = &AlgorithmPredictDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
