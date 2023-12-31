package apds

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

// CreateSurveyJob invokes the apds.CreateSurveyJob API synchronously
func (client *Client) CreateSurveyJob(request *CreateSurveyJobRequest) (response *CreateSurveyJobResponse, err error) {
	response = CreateCreateSurveyJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSurveyJobWithChan invokes the apds.CreateSurveyJob API asynchronously
func (client *Client) CreateSurveyJobWithChan(request *CreateSurveyJobRequest) (<-chan *CreateSurveyJobResponse, <-chan error) {
	responseChan := make(chan *CreateSurveyJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSurveyJob(request)
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

// CreateSurveyJobWithCallback invokes the apds.CreateSurveyJob API asynchronously
func (client *Client) CreateSurveyJobWithCallback(request *CreateSurveyJobRequest, callback func(response *CreateSurveyJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSurveyJobResponse
		var err error
		defer close(result)
		response, err = client.CreateSurveyJob(request)
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

// CreateSurveyJobRequest is the request struct for api CreateSurveyJob
type CreateSurveyJobRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateSurveyJobResponse is the response struct for api CreateSurveyJob
type CreateSurveyJobResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateCreateSurveyJobRequest creates a request to invoke CreateSurveyJob API
func CreateCreateSurveyJobRequest() (request *CreateSurveyJobRequest) {
	request = &CreateSurveyJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "CreateSurveyJob", "/okss-services/winback/add-survey-job", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSurveyJobResponse creates a response to parse from CreateSurveyJob response
func CreateCreateSurveyJobResponse() (response *CreateSurveyJobResponse) {
	response = &CreateSurveyJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
