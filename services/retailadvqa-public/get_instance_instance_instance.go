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

// GetInstanceInstanceInstance invokes the retailadvqa_public.GetInstanceInstanceInstance API synchronously
func (client *Client) GetInstanceInstanceInstance(request *GetInstanceInstanceInstanceRequest) (response *GetInstanceInstanceInstanceResponse, err error) {
	response = CreateGetInstanceInstanceInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceInstanceInstanceWithChan invokes the retailadvqa_public.GetInstanceInstanceInstance API asynchronously
func (client *Client) GetInstanceInstanceInstanceWithChan(request *GetInstanceInstanceInstanceRequest) (<-chan *GetInstanceInstanceInstanceResponse, <-chan error) {
	responseChan := make(chan *GetInstanceInstanceInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceInstanceInstance(request)
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

// GetInstanceInstanceInstanceWithCallback invokes the retailadvqa_public.GetInstanceInstanceInstance API asynchronously
func (client *Client) GetInstanceInstanceInstanceWithCallback(request *GetInstanceInstanceInstanceRequest, callback func(response *GetInstanceInstanceInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceInstanceInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceInstanceInstance(request)
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

// GetInstanceInstanceInstanceRequest is the request struct for api GetInstanceInstanceInstance
type GetInstanceInstanceInstanceRequest struct {
	*requests.RpcRequest
	AccessId    string `position:"Query" name:"AccessId"`
	AccountId   string `position:"Query" name:"AccountId"`
	AccountName string `position:"Query" name:"AccountName"`
}

// GetInstanceInstanceInstanceResponse is the response struct for api GetInstanceInstanceInstance
type GetInstanceInstanceInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetInstanceInstanceInstanceRequest creates a request to invoke GetInstanceInstanceInstance API
func CreateGetInstanceInstanceInstanceRequest() (request *GetInstanceInstanceInstanceRequest) {
	request = &GetInstanceInstanceInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "GetInstanceInstanceInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceInstanceInstanceResponse creates a response to parse from GetInstanceInstanceInstance response
func CreateGetInstanceInstanceInstanceResponse() (response *GetInstanceInstanceInstanceResponse) {
	response = &GetInstanceInstanceInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
