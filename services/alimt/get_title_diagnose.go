package alimt

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

// GetTitleDiagnose invokes the alimt.GetTitleDiagnose API synchronously
func (client *Client) GetTitleDiagnose(request *GetTitleDiagnoseRequest) (response *GetTitleDiagnoseResponse, err error) {
	response = CreateGetTitleDiagnoseResponse()
	err = client.DoAction(request, response)
	return
}

// GetTitleDiagnoseWithChan invokes the alimt.GetTitleDiagnose API asynchronously
func (client *Client) GetTitleDiagnoseWithChan(request *GetTitleDiagnoseRequest) (<-chan *GetTitleDiagnoseResponse, <-chan error) {
	responseChan := make(chan *GetTitleDiagnoseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTitleDiagnose(request)
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

// GetTitleDiagnoseWithCallback invokes the alimt.GetTitleDiagnose API asynchronously
func (client *Client) GetTitleDiagnoseWithCallback(request *GetTitleDiagnoseRequest, callback func(response *GetTitleDiagnoseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTitleDiagnoseResponse
		var err error
		defer close(result)
		response, err = client.GetTitleDiagnose(request)
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

// GetTitleDiagnoseRequest is the request struct for api GetTitleDiagnose
type GetTitleDiagnoseRequest struct {
	*requests.RpcRequest
	Language   string `position:"Body" name:"Language"`
	Title      string `position:"Body" name:"Title"`
	Platform   string `position:"Body" name:"Platform"`
	Extra      string `position:"Body" name:"Extra"`
	CategoryId string `position:"Body" name:"CategoryId"`
}

// GetTitleDiagnoseResponse is the response struct for api GetTitleDiagnose
type GetTitleDiagnoseResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetTitleDiagnoseRequest creates a request to invoke GetTitleDiagnose API
func CreateGetTitleDiagnoseRequest() (request *GetTitleDiagnoseRequest) {
	request = &GetTitleDiagnoseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetTitleDiagnose", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTitleDiagnoseResponse creates a response to parse from GetTitleDiagnose response
func CreateGetTitleDiagnoseResponse() (response *GetTitleDiagnoseResponse) {
	response = &GetTitleDiagnoseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
