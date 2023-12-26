package arms

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

// GetOnCallSchedulesDetail invokes the arms.GetOnCallSchedulesDetail API synchronously
func (client *Client) GetOnCallSchedulesDetail(request *GetOnCallSchedulesDetailRequest) (response *GetOnCallSchedulesDetailResponse, err error) {
	response = CreateGetOnCallSchedulesDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetOnCallSchedulesDetailWithChan invokes the arms.GetOnCallSchedulesDetail API asynchronously
func (client *Client) GetOnCallSchedulesDetailWithChan(request *GetOnCallSchedulesDetailRequest) (<-chan *GetOnCallSchedulesDetailResponse, <-chan error) {
	responseChan := make(chan *GetOnCallSchedulesDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOnCallSchedulesDetail(request)
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

// GetOnCallSchedulesDetailWithCallback invokes the arms.GetOnCallSchedulesDetail API asynchronously
func (client *Client) GetOnCallSchedulesDetailWithCallback(request *GetOnCallSchedulesDetailRequest, callback func(response *GetOnCallSchedulesDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOnCallSchedulesDetailResponse
		var err error
		defer close(result)
		response, err = client.GetOnCallSchedulesDetail(request)
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

// GetOnCallSchedulesDetailRequest is the request struct for api GetOnCallSchedulesDetail
type GetOnCallSchedulesDetailRequest struct {
	*requests.RpcRequest
	EndTime   string           `position:"Query" name:"EndTime"`
	Id        requests.Integer `position:"Query" name:"Id"`
	StartTime string           `position:"Query" name:"StartTime"`
}

// GetOnCallSchedulesDetailResponse is the response struct for api GetOnCallSchedulesDetail
type GetOnCallSchedulesDetailResponse struct {
	*responses.BaseResponse
}

// CreateGetOnCallSchedulesDetailRequest creates a request to invoke GetOnCallSchedulesDetail API
func CreateGetOnCallSchedulesDetailRequest() (request *GetOnCallSchedulesDetailRequest) {
	request = &GetOnCallSchedulesDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetOnCallSchedulesDetail", "arms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetOnCallSchedulesDetailResponse creates a response to parse from GetOnCallSchedulesDetail response
func CreateGetOnCallSchedulesDetailResponse() (response *GetOnCallSchedulesDetailResponse) {
	response = &GetOnCallSchedulesDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
