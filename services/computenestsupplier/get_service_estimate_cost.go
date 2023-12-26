package computenestsupplier

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

// GetServiceEstimateCost invokes the computenestsupplier.GetServiceEstimateCost API synchronously
func (client *Client) GetServiceEstimateCost(request *GetServiceEstimateCostRequest) (response *GetServiceEstimateCostResponse, err error) {
	response = CreateGetServiceEstimateCostResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceEstimateCostWithChan invokes the computenestsupplier.GetServiceEstimateCost API asynchronously
func (client *Client) GetServiceEstimateCostWithChan(request *GetServiceEstimateCostRequest) (<-chan *GetServiceEstimateCostResponse, <-chan error) {
	responseChan := make(chan *GetServiceEstimateCostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceEstimateCost(request)
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

// GetServiceEstimateCostWithCallback invokes the computenestsupplier.GetServiceEstimateCost API asynchronously
func (client *Client) GetServiceEstimateCostWithCallback(request *GetServiceEstimateCostRequest, callback func(response *GetServiceEstimateCostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceEstimateCostResponse
		var err error
		defer close(result)
		response, err = client.GetServiceEstimateCost(request)
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

// GetServiceEstimateCostRequest is the request struct for api GetServiceEstimateCost
type GetServiceEstimateCostRequest struct {
	*requests.RpcRequest
	ClientToken       string `position:"Query" name:"ClientToken"`
	ServiceVersion    string `position:"Query" name:"ServiceVersion"`
	TemplateName      string `position:"Query" name:"TemplateName"`
	ServiceId         string `position:"Query" name:"ServiceId"`
	Parameters        string `position:"Query" name:"Parameters"`
	ServiceInstanceId string `position:"Query" name:"ServiceInstanceId"`
}

// GetServiceEstimateCostResponse is the response struct for api GetServiceEstimateCost
type GetServiceEstimateCostResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Resources map[string]interface{} `json:"Resources" xml:"Resources"`
}

// CreateGetServiceEstimateCostRequest creates a request to invoke GetServiceEstimateCost API
func CreateGetServiceEstimateCostRequest() (request *GetServiceEstimateCostRequest) {
	request = &GetServiceEstimateCostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "GetServiceEstimateCost", "", "")
	request.Method = requests.POST
	return
}

// CreateGetServiceEstimateCostResponse creates a response to parse from GetServiceEstimateCost response
func CreateGetServiceEstimateCostResponse() (response *GetServiceEstimateCostResponse) {
	response = &GetServiceEstimateCostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
