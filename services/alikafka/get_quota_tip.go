package alikafka

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

// GetQuotaTip invokes the alikafka.GetQuotaTip API synchronously
func (client *Client) GetQuotaTip(request *GetQuotaTipRequest) (response *GetQuotaTipResponse, err error) {
	response = CreateGetQuotaTipResponse()
	err = client.DoAction(request, response)
	return
}

// GetQuotaTipWithChan invokes the alikafka.GetQuotaTip API asynchronously
func (client *Client) GetQuotaTipWithChan(request *GetQuotaTipRequest) (<-chan *GetQuotaTipResponse, <-chan error) {
	responseChan := make(chan *GetQuotaTipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQuotaTip(request)
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

// GetQuotaTipWithCallback invokes the alikafka.GetQuotaTip API asynchronously
func (client *Client) GetQuotaTipWithCallback(request *GetQuotaTipRequest, callback func(response *GetQuotaTipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQuotaTipResponse
		var err error
		defer close(result)
		response, err = client.GetQuotaTip(request)
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

// GetQuotaTipRequest is the request struct for api GetQuotaTip
type GetQuotaTipRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetQuotaTipResponse is the response struct for api GetQuotaTip
type GetQuotaTipResponse struct {
	*responses.BaseResponse
	Code      int       `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Success   bool      `json:"Success" xml:"Success"`
	QuotaData QuotaData `json:"QuotaData" xml:"QuotaData"`
}

// CreateGetQuotaTipRequest creates a request to invoke GetQuotaTip API
func CreateGetQuotaTipRequest() (request *GetQuotaTipRequest) {
	request = &GetQuotaTipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetQuotaTip", "", "")
	request.Method = requests.POST
	return
}

// CreateGetQuotaTipResponse creates a response to parse from GetQuotaTip response
func CreateGetQuotaTipResponse() (response *GetQuotaTipResponse) {
	response = &GetQuotaTipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
