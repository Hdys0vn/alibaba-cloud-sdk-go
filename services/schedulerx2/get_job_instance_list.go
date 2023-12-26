package schedulerx2

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

// GetJobInstanceList invokes the schedulerx2.GetJobInstanceList API synchronously
func (client *Client) GetJobInstanceList(request *GetJobInstanceListRequest) (response *GetJobInstanceListResponse, err error) {
	response = CreateGetJobInstanceListResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobInstanceListWithChan invokes the schedulerx2.GetJobInstanceList API asynchronously
func (client *Client) GetJobInstanceListWithChan(request *GetJobInstanceListRequest) (<-chan *GetJobInstanceListResponse, <-chan error) {
	responseChan := make(chan *GetJobInstanceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobInstanceList(request)
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

// GetJobInstanceListWithCallback invokes the schedulerx2.GetJobInstanceList API asynchronously
func (client *Client) GetJobInstanceListWithCallback(request *GetJobInstanceListRequest, callback func(response *GetJobInstanceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobInstanceListResponse
		var err error
		defer close(result)
		response, err = client.GetJobInstanceList(request)
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

// GetJobInstanceListRequest is the request struct for api GetJobInstanceList
type GetJobInstanceListRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	StartTimestamp  requests.Integer `position:"Query" name:"StartTimestamp"`
	EndTimestamp    requests.Integer `position:"Query" name:"EndTimestamp"`
	JobId           requests.Integer `position:"Query" name:"JobId"`
	Namespace       string           `position:"Query" name:"Namespace"`
	Status          requests.Integer `position:"Query" name:"Status"`
}

// GetJobInstanceListResponse is the response struct for api GetJobInstanceList
type GetJobInstanceListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetJobInstanceListRequest creates a request to invoke GetJobInstanceList API
func CreateGetJobInstanceListRequest() (request *GetJobInstanceListRequest) {
	request = &GetJobInstanceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "GetJobInstanceList", "", "")
	request.Method = requests.GET
	return
}

// CreateGetJobInstanceListResponse creates a response to parse from GetJobInstanceList response
func CreateGetJobInstanceListResponse() (response *GetJobInstanceListResponse) {
	response = &GetJobInstanceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
