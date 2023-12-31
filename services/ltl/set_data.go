package ltl

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

// SetData invokes the ltl.SetData API synchronously
func (client *Client) SetData(request *SetDataRequest) (response *SetDataResponse, err error) {
	response = CreateSetDataResponse()
	err = client.DoAction(request, response)
	return
}

// SetDataWithChan invokes the ltl.SetData API asynchronously
func (client *Client) SetDataWithChan(request *SetDataRequest) (<-chan *SetDataResponse, <-chan error) {
	responseChan := make(chan *SetDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetData(request)
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

// SetDataWithCallback invokes the ltl.SetData API asynchronously
func (client *Client) SetDataWithCallback(request *SetDataRequest, callback func(response *SetDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDataResponse
		var err error
		defer close(result)
		response, err = client.SetData(request)
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

// SetDataRequest is the request struct for api SetData
type SetDataRequest struct {
	*requests.RpcRequest
	ApiVersion string `position:"Query" name:"ApiVersion"`
	ProductKey string `position:"Query" name:"ProductKey"`
	Value      string `position:"Query" name:"Value"`
	Key        string `position:"Query" name:"Key"`
}

// SetDataResponse is the response struct for api SetData
type SetDataResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSetDataRequest creates a request to invoke SetData API
func CreateSetDataRequest() (request *SetDataRequest) {
	request = &SetDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "SetData", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDataResponse creates a response to parse from SetData response
func CreateSetDataResponse() (response *SetDataResponse) {
	response = &SetDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
