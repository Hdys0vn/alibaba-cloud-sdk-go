package imageprocess

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

// RunMedQA invokes the imageprocess.RunMedQA API synchronously
func (client *Client) RunMedQA(request *RunMedQARequest) (response *RunMedQAResponse, err error) {
	response = CreateRunMedQAResponse()
	err = client.DoAction(request, response)
	return
}

// RunMedQAWithChan invokes the imageprocess.RunMedQA API asynchronously
func (client *Client) RunMedQAWithChan(request *RunMedQARequest) (<-chan *RunMedQAResponse, <-chan error) {
	responseChan := make(chan *RunMedQAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunMedQA(request)
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

// RunMedQAWithCallback invokes the imageprocess.RunMedQA API asynchronously
func (client *Client) RunMedQAWithCallback(request *RunMedQARequest, callback func(response *RunMedQAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunMedQAResponse
		var err error
		defer close(result)
		response, err = client.RunMedQA(request)
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

// RunMedQARequest is the request struct for api RunMedQA
type RunMedQARequest struct {
	*requests.RpcRequest
	SessionId           string                         `position:"Body" name:"SessionId"`
	OrgName             string                         `position:"Body" name:"OrgName"`
	AnswerImageDataList *[]RunMedQAAnswerImageDataList `position:"Body" name:"AnswerImageDataList"  type:"Repeated"`
	AnswerTextList      *[]RunMedQAAnswerTextList      `position:"Body" name:"AnswerTextList"  type:"Repeated"`
	Department          string                         `position:"Body" name:"Department"`
	AnswerImageURLList  *[]RunMedQAAnswerImageURLList  `position:"Body" name:"AnswerImageURLList"  type:"Repeated"`
	QuestionType        string                         `position:"Body" name:"QuestionType"`
	OrgId               string                         `position:"Body" name:"OrgId"`
}

// RunMedQAAnswerImageDataList is a repeated param struct in RunMedQARequest
type RunMedQAAnswerImageDataList struct {
	AnswerImageData string `name:"AnswerImageData"`
}

// RunMedQAAnswerTextList is a repeated param struct in RunMedQARequest
type RunMedQAAnswerTextList struct {
	AnswerText string `name:"AnswerText"`
}

// RunMedQAAnswerImageURLList is a repeated param struct in RunMedQARequest
type RunMedQAAnswerImageURLList struct {
	AnswerImageURL string `name:"AnswerImageURL"`
}

// RunMedQAResponse is the response struct for api RunMedQA
type RunMedQAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRunMedQARequest creates a request to invoke RunMedQA API
func CreateRunMedQARequest() (request *RunMedQARequest) {
	request = &RunMedQARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "RunMedQA", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunMedQAResponse creates a response to parse from RunMedQA response
func CreateRunMedQAResponse() (response *RunMedQAResponse) {
	response = &RunMedQAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
