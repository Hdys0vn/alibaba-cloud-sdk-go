package qualitycheck

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

// InsertScoreForApi invokes the qualitycheck.InsertScoreForApi API synchronously
func (client *Client) InsertScoreForApi(request *InsertScoreForApiRequest) (response *InsertScoreForApiResponse, err error) {
	response = CreateInsertScoreForApiResponse()
	err = client.DoAction(request, response)
	return
}

// InsertScoreForApiWithChan invokes the qualitycheck.InsertScoreForApi API asynchronously
func (client *Client) InsertScoreForApiWithChan(request *InsertScoreForApiRequest) (<-chan *InsertScoreForApiResponse, <-chan error) {
	responseChan := make(chan *InsertScoreForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertScoreForApi(request)
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

// InsertScoreForApiWithCallback invokes the qualitycheck.InsertScoreForApi API asynchronously
func (client *Client) InsertScoreForApiWithCallback(request *InsertScoreForApiRequest, callback func(response *InsertScoreForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertScoreForApiResponse
		var err error
		defer close(result)
		response, err = client.InsertScoreForApi(request)
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

// InsertScoreForApiRequest is the request struct for api InsertScoreForApi
type InsertScoreForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// InsertScoreForApiResponse is the response struct for api InsertScoreForApi
type InsertScoreForApiResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateInsertScoreForApiRequest creates a request to invoke InsertScoreForApi API
func CreateInsertScoreForApiRequest() (request *InsertScoreForApiRequest) {
	request = &InsertScoreForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "InsertScoreForApi", "", "")
	request.Method = requests.POST
	return
}

// CreateInsertScoreForApiResponse creates a response to parse from InsertScoreForApi response
func CreateInsertScoreForApiResponse() (response *InsertScoreForApiResponse) {
	response = &InsertScoreForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
