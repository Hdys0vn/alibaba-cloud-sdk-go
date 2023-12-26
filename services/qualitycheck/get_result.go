package qualitycheck

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

// GetResult invokes the qualitycheck.GetResult API synchronously
func (client *Client) GetResult(request *GetResultRequest) (response *GetResultResponse, err error) {
	response = CreateGetResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetResultWithChan invokes the qualitycheck.GetResult API asynchronously
func (client *Client) GetResultWithChan(request *GetResultRequest) (<-chan *GetResultResponse, <-chan error) {
	responseChan := make(chan *GetResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResult(request)
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

// GetResultWithCallback invokes the qualitycheck.GetResult API asynchronously
func (client *Client) GetResultWithCallback(request *GetResultRequest, callback func(response *GetResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResultResponse
		var err error
		defer close(result)
		response, err = client.GetResult(request)
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

// GetResultRequest is the request struct for api GetResult
type GetResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetResultResponse is the response struct for api GetResult
type GetResultResponse struct {
	*responses.BaseResponse
	RequestId     string          `json:"RequestId" xml:"RequestId"`
	Success       bool            `json:"Success" xml:"Success"`
	ResultCountId string          `json:"ResultCountId" xml:"ResultCountId"`
	Code          string          `json:"Code" xml:"Code"`
	Message       string          `json:"Message" xml:"Message"`
	PageNumber    int             `json:"PageNumber" xml:"PageNumber"`
	PageSize      int             `json:"PageSize" xml:"PageSize"`
	Count         int             `json:"Count" xml:"Count"`
	Data          DataInGetResult `json:"Data" xml:"Data"`
}

// CreateGetResultRequest creates a request to invoke GetResult API
func CreateGetResultRequest() (request *GetResultRequest) {
	request = &GetResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetResult", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResultResponse creates a response to parse from GetResult response
func CreateGetResultResponse() (response *GetResultResponse) {
	response = &GetResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}