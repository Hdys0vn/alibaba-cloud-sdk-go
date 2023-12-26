package cloud_siem

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

// GetHistograms invokes the cloud_siem.GetHistograms API synchronously
func (client *Client) GetHistograms(request *GetHistogramsRequest) (response *GetHistogramsResponse, err error) {
	response = CreateGetHistogramsResponse()
	err = client.DoAction(request, response)
	return
}

// GetHistogramsWithChan invokes the cloud_siem.GetHistograms API asynchronously
func (client *Client) GetHistogramsWithChan(request *GetHistogramsRequest) (<-chan *GetHistogramsResponse, <-chan error) {
	responseChan := make(chan *GetHistogramsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetHistograms(request)
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

// GetHistogramsWithCallback invokes the cloud_siem.GetHistograms API asynchronously
func (client *Client) GetHistogramsWithCallback(request *GetHistogramsRequest, callback func(response *GetHistogramsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetHistogramsResponse
		var err error
		defer close(result)
		response, err = client.GetHistograms(request)
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

// GetHistogramsRequest is the request struct for api GetHistograms
type GetHistogramsRequest struct {
	*requests.RpcRequest
	From  requests.Integer `position:"Body" name:"From"`
	Query string           `position:"Body" name:"Query"`
	To    requests.Integer `position:"Body" name:"To"`
}

// GetHistogramsResponse is the response struct for api GetHistograms
type GetHistogramsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetHistogramsRequest creates a request to invoke GetHistograms API
func CreateGetHistogramsRequest() (request *GetHistogramsRequest) {
	request = &GetHistogramsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "GetHistograms", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetHistogramsResponse creates a response to parse from GetHistograms response
func CreateGetHistogramsResponse() (response *GetHistogramsResponse) {
	response = &GetHistogramsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
