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

// AlgorithmTrainDetail invokes the retailadvqa_public.AlgorithmTrainDetail API synchronously
func (client *Client) AlgorithmTrainDetail(request *AlgorithmTrainDetailRequest) (response *AlgorithmTrainDetailResponse, err error) {
	response = CreateAlgorithmTrainDetailResponse()
	err = client.DoAction(request, response)
	return
}

// AlgorithmTrainDetailWithChan invokes the retailadvqa_public.AlgorithmTrainDetail API asynchronously
func (client *Client) AlgorithmTrainDetailWithChan(request *AlgorithmTrainDetailRequest) (<-chan *AlgorithmTrainDetailResponse, <-chan error) {
	responseChan := make(chan *AlgorithmTrainDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AlgorithmTrainDetail(request)
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

// AlgorithmTrainDetailWithCallback invokes the retailadvqa_public.AlgorithmTrainDetail API asynchronously
func (client *Client) AlgorithmTrainDetailWithCallback(request *AlgorithmTrainDetailRequest, callback func(response *AlgorithmTrainDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AlgorithmTrainDetailResponse
		var err error
		defer close(result)
		response, err = client.AlgorithmTrainDetail(request)
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

// AlgorithmTrainDetailRequest is the request struct for api AlgorithmTrainDetail
type AlgorithmTrainDetailRequest struct {
	*requests.RpcRequest
	TrainId     requests.Integer `position:"Query" name:"TrainId"`
	AccessId    string           `position:"Query" name:"AccessId"`
	TenantId    string           `position:"Query" name:"TenantId"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// AlgorithmTrainDetailResponse is the response struct for api AlgorithmTrainDetail
type AlgorithmTrainDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAlgorithmTrainDetailRequest creates a request to invoke AlgorithmTrainDetail API
func CreateAlgorithmTrainDetailRequest() (request *AlgorithmTrainDetailRequest) {
	request = &AlgorithmTrainDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "AlgorithmTrainDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateAlgorithmTrainDetailResponse creates a response to parse from AlgorithmTrainDetail response
func CreateAlgorithmTrainDetailResponse() (response *AlgorithmTrainDetailResponse) {
	response = &AlgorithmTrainDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
