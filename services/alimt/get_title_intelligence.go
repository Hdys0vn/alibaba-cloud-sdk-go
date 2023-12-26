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

// GetTitleIntelligence invokes the alimt.GetTitleIntelligence API synchronously
func (client *Client) GetTitleIntelligence(request *GetTitleIntelligenceRequest) (response *GetTitleIntelligenceResponse, err error) {
	response = CreateGetTitleIntelligenceResponse()
	err = client.DoAction(request, response)
	return
}

// GetTitleIntelligenceWithChan invokes the alimt.GetTitleIntelligence API asynchronously
func (client *Client) GetTitleIntelligenceWithChan(request *GetTitleIntelligenceRequest) (<-chan *GetTitleIntelligenceResponse, <-chan error) {
	responseChan := make(chan *GetTitleIntelligenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTitleIntelligence(request)
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

// GetTitleIntelligenceWithCallback invokes the alimt.GetTitleIntelligence API asynchronously
func (client *Client) GetTitleIntelligenceWithCallback(request *GetTitleIntelligenceRequest, callback func(response *GetTitleIntelligenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTitleIntelligenceResponse
		var err error
		defer close(result)
		response, err = client.GetTitleIntelligence(request)
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

// GetTitleIntelligenceRequest is the request struct for api GetTitleIntelligence
type GetTitleIntelligenceRequest struct {
	*requests.RpcRequest
	CatLevelThreeId requests.Integer `position:"Body" name:"CatLevelThreeId"`
	CatLevelTwoId   requests.Integer `position:"Body" name:"CatLevelTwoId"`
	Keywords        string           `position:"Body" name:"Keywords"`
	Platform        string           `position:"Body" name:"Platform"`
	Extra           string           `position:"Body" name:"Extra"`
}

// GetTitleIntelligenceResponse is the response struct for api GetTitleIntelligence
type GetTitleIntelligenceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetTitleIntelligenceRequest creates a request to invoke GetTitleIntelligence API
func CreateGetTitleIntelligenceRequest() (request *GetTitleIntelligenceRequest) {
	request = &GetTitleIntelligenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetTitleIntelligence", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTitleIntelligenceResponse creates a response to parse from GetTitleIntelligence response
func CreateGetTitleIntelligenceResponse() (response *GetTitleIntelligenceResponse) {
	response = &GetTitleIntelligenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
