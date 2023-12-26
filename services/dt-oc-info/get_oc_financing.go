package dt_oc_info

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

// GetOcFinancing invokes the dt_oc_info.GetOcFinancing API synchronously
func (client *Client) GetOcFinancing(request *GetOcFinancingRequest) (response *GetOcFinancingResponse, err error) {
	response = CreateGetOcFinancingResponse()
	err = client.DoAction(request, response)
	return
}

// GetOcFinancingWithChan invokes the dt_oc_info.GetOcFinancing API asynchronously
func (client *Client) GetOcFinancingWithChan(request *GetOcFinancingRequest) (<-chan *GetOcFinancingResponse, <-chan error) {
	responseChan := make(chan *GetOcFinancingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOcFinancing(request)
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

// GetOcFinancingWithCallback invokes the dt_oc_info.GetOcFinancing API asynchronously
func (client *Client) GetOcFinancingWithCallback(request *GetOcFinancingRequest, callback func(response *GetOcFinancingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOcFinancingResponse
		var err error
		defer close(result)
		response, err = client.GetOcFinancing(request)
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

// GetOcFinancingRequest is the request struct for api GetOcFinancing
type GetOcFinancingRequest struct {
	*requests.RpcRequest
	PageNo    requests.Integer `position:"Body" name:"PageNo"`
	PageSize  requests.Integer `position:"Body" name:"PageSize"`
	SearchKey string           `position:"Body" name:"SearchKey"`
}

// GetOcFinancingResponse is the response struct for api GetOcFinancing
type GetOcFinancingResponse struct {
	*responses.BaseResponse
	Code      string     `json:"Code" xml:"Code"`
	Success   bool       `json:"Success" xml:"Success"`
	Message   string     `json:"Message" xml:"Message"`
	TotalNum  int        `json:"TotalNum" xml:"TotalNum"`
	PageIndex int        `json:"PageIndex" xml:"PageIndex"`
	PageNum   int        `json:"PageNum" xml:"PageNum"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetOcFinancingRequest creates a request to invoke GetOcFinancing API
func CreateGetOcFinancingRequest() (request *GetOcFinancingRequest) {
	request = &GetOcFinancingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dt-oc-info", "2022-08-29", "GetOcFinancing", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOcFinancingResponse creates a response to parse from GetOcFinancing response
func CreateGetOcFinancingResponse() (response *GetOcFinancingResponse) {
	response = &GetOcFinancingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}