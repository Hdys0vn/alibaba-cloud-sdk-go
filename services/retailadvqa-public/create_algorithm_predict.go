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

// CreateAlgorithmPredict invokes the retailadvqa_public.CreateAlgorithmPredict API synchronously
func (client *Client) CreateAlgorithmPredict(request *CreateAlgorithmPredictRequest) (response *CreateAlgorithmPredictResponse, err error) {
	response = CreateCreateAlgorithmPredictResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAlgorithmPredictWithChan invokes the retailadvqa_public.CreateAlgorithmPredict API asynchronously
func (client *Client) CreateAlgorithmPredictWithChan(request *CreateAlgorithmPredictRequest) (<-chan *CreateAlgorithmPredictResponse, <-chan error) {
	responseChan := make(chan *CreateAlgorithmPredictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlgorithmPredict(request)
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

// CreateAlgorithmPredictWithCallback invokes the retailadvqa_public.CreateAlgorithmPredict API asynchronously
func (client *Client) CreateAlgorithmPredictWithCallback(request *CreateAlgorithmPredictRequest, callback func(response *CreateAlgorithmPredictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlgorithmPredictResponse
		var err error
		defer close(result)
		response, err = client.CreateAlgorithmPredict(request)
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

// CreateAlgorithmPredictRequest is the request struct for api CreateAlgorithmPredict
type CreateAlgorithmPredictRequest struct {
	*requests.RpcRequest
	AccessId           string           `position:"Query" name:"AccessId"`
	TaskType           string           `position:"Body" name:"TaskType"`
	GoodsPoolTableName string           `position:"Body" name:"GoodsPoolTableName"`
	TenantId           string           `position:"Query" name:"TenantId"`
	TaskName           string           `position:"Body" name:"TaskName"`
	ItemSize           requests.Integer `position:"Body" name:"ItemSize"`
	AudienceTableName  string           `position:"Body" name:"AudienceTableName"`
	WorkspaceId        string           `position:"Query" name:"WorkspaceId"`
}

// CreateAlgorithmPredictResponse is the response struct for api CreateAlgorithmPredict
type CreateAlgorithmPredictResponse struct {
	*responses.BaseResponse
	Success   bool                         `json:"Success" xml:"Success"`
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	TraceId   string                       `json:"TraceId" xml:"TraceId"`
	ErrorCode string                       `json:"ErrorCode" xml:"ErrorCode"`
	ErrorDesc string                       `json:"ErrorDesc" xml:"ErrorDesc"`
	Data      DataInCreateAlgorithmPredict `json:"Data" xml:"Data"`
}

// CreateCreateAlgorithmPredictRequest creates a request to invoke CreateAlgorithmPredict API
func CreateCreateAlgorithmPredictRequest() (request *CreateAlgorithmPredictRequest) {
	request = &CreateAlgorithmPredictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CreateAlgorithmPredict", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAlgorithmPredictResponse creates a response to parse from CreateAlgorithmPredict response
func CreateCreateAlgorithmPredictResponse() (response *CreateAlgorithmPredictResponse) {
	response = &CreateAlgorithmPredictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
