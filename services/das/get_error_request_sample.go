package das

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

// GetErrorRequestSample invokes the das.GetErrorRequestSample API synchronously
func (client *Client) GetErrorRequestSample(request *GetErrorRequestSampleRequest) (response *GetErrorRequestSampleResponse, err error) {
	response = CreateGetErrorRequestSampleResponse()
	err = client.DoAction(request, response)
	return
}

// GetErrorRequestSampleWithChan invokes the das.GetErrorRequestSample API asynchronously
func (client *Client) GetErrorRequestSampleWithChan(request *GetErrorRequestSampleRequest) (<-chan *GetErrorRequestSampleResponse, <-chan error) {
	responseChan := make(chan *GetErrorRequestSampleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetErrorRequestSample(request)
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

// GetErrorRequestSampleWithCallback invokes the das.GetErrorRequestSample API asynchronously
func (client *Client) GetErrorRequestSampleWithCallback(request *GetErrorRequestSampleRequest, callback func(response *GetErrorRequestSampleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetErrorRequestSampleResponse
		var err error
		defer close(result)
		response, err = client.GetErrorRequestSample(request)
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

// GetErrorRequestSampleRequest is the request struct for api GetErrorRequestSample
type GetErrorRequestSampleRequest struct {
	*requests.RpcRequest
	SqlId      string           `position:"Query" name:"SqlId"`
	Start      requests.Integer `position:"Query" name:"Start"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	DbName     string           `position:"Query" name:"DbName"`
	End        requests.Integer `position:"Query" name:"End"`
	NodeId     string           `position:"Query" name:"NodeId"`
}

// GetErrorRequestSampleResponse is the response struct for api GetErrorRequestSample
type GetErrorRequestSampleResponse struct {
	*responses.BaseResponse
	Code      int64    `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	Data      []Sample `json:"Data" xml:"Data"`
}

// CreateGetErrorRequestSampleRequest creates a request to invoke GetErrorRequestSample API
func CreateGetErrorRequestSampleRequest() (request *GetErrorRequestSampleRequest) {
	request = &GetErrorRequestSampleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetErrorRequestSample", "", "")
	request.Method = requests.POST
	return
}

// CreateGetErrorRequestSampleResponse creates a response to parse from GetErrorRequestSample response
func CreateGetErrorRequestSampleResponse() (response *GetErrorRequestSampleResponse) {
	response = &GetErrorRequestSampleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
