package aliyuncvc

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

// CreateEvaluation invokes the aliyuncvc.CreateEvaluation API synchronously
func (client *Client) CreateEvaluation(request *CreateEvaluationRequest) (response *CreateEvaluationResponse, err error) {
	response = CreateCreateEvaluationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEvaluationWithChan invokes the aliyuncvc.CreateEvaluation API asynchronously
func (client *Client) CreateEvaluationWithChan(request *CreateEvaluationRequest) (<-chan *CreateEvaluationResponse, <-chan error) {
	responseChan := make(chan *CreateEvaluationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEvaluation(request)
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

// CreateEvaluationWithCallback invokes the aliyuncvc.CreateEvaluation API asynchronously
func (client *Client) CreateEvaluationWithCallback(request *CreateEvaluationRequest, callback func(response *CreateEvaluationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEvaluationResponse
		var err error
		defer close(result)
		response, err = client.CreateEvaluation(request)
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

// CreateEvaluationRequest is the request struct for api CreateEvaluation
type CreateEvaluationRequest struct {
	*requests.RpcRequest
	CreateTime  requests.Integer `position:"Query" name:"CreateTime"`
	Memo        string           `position:"Query" name:"Memo"`
	Description string           `position:"Query" name:"Description"`
	MemberUUID  string           `position:"Query" name:"MemberUUID"`
	UserId      string           `position:"Query" name:"UserId"`
	Evaluation  string           `position:"Query" name:"Evaluation"`
	Score       string           `position:"Query" name:"Score"`
	MeetingUUID string           `position:"Query" name:"MeetingUUID"`
	AppId       string           `position:"Query" name:"AppId"`
}

// CreateEvaluationResponse is the response struct for api CreateEvaluation
type CreateEvaluationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateEvaluationRequest creates a request to invoke CreateEvaluation API
func CreateCreateEvaluationRequest() (request *CreateEvaluationRequest) {
	request = &CreateEvaluationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "CreateEvaluation", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEvaluationResponse creates a response to parse from CreateEvaluation response
func CreateCreateEvaluationResponse() (response *CreateEvaluationResponse) {
	response = &CreateEvaluationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
