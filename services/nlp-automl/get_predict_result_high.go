package nlp_automl

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

// GetPredictResultHigh invokes the nlp_automl.GetPredictResultHigh API synchronously
func (client *Client) GetPredictResultHigh(request *GetPredictResultHighRequest) (response *GetPredictResultHighResponse, err error) {
	response = CreateGetPredictResultHighResponse()
	err = client.DoAction(request, response)
	return
}

// GetPredictResultHighWithChan invokes the nlp_automl.GetPredictResultHigh API asynchronously
func (client *Client) GetPredictResultHighWithChan(request *GetPredictResultHighRequest) (<-chan *GetPredictResultHighResponse, <-chan error) {
	responseChan := make(chan *GetPredictResultHighResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPredictResultHigh(request)
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

// GetPredictResultHighWithCallback invokes the nlp_automl.GetPredictResultHigh API asynchronously
func (client *Client) GetPredictResultHighWithCallback(request *GetPredictResultHighRequest, callback func(response *GetPredictResultHighResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPredictResultHighResponse
		var err error
		defer close(result)
		response, err = client.GetPredictResultHigh(request)
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

// GetPredictResultHighRequest is the request struct for api GetPredictResultHigh
type GetPredictResultHighRequest struct {
	*requests.RpcRequest
	TopK         requests.Integer `position:"Body" name:"TopK"`
	Product      string           `position:"Body" name:"Product"`
	ModelId      requests.Integer `position:"Body" name:"ModelId"`
	DetailTag    string           `position:"Body" name:"DetailTag"`
	Content      string           `position:"Body" name:"Content"`
	ModelVersion string           `position:"Body" name:"ModelVersion"`
}

// GetPredictResultHighResponse is the response struct for api GetPredictResultHigh
type GetPredictResultHighResponse struct {
	*responses.BaseResponse
	Content   string `json:"Content" xml:"Content"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetPredictResultHighRequest creates a request to invoke GetPredictResultHigh API
func CreateGetPredictResultHighRequest() (request *GetPredictResultHighRequest) {
	request = &GetPredictResultHighRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nlp-automl", "2019-11-11", "GetPredictResultHigh", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPredictResultHighResponse creates a response to parse from GetPredictResultHigh response
func CreateGetPredictResultHighResponse() (response *GetPredictResultHighResponse) {
	response = &GetPredictResultHighResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}