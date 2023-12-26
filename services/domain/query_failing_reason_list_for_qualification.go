package domain

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

// QueryFailingReasonListForQualification invokes the domain.QueryFailingReasonListForQualification API synchronously
func (client *Client) QueryFailingReasonListForQualification(request *QueryFailingReasonListForQualificationRequest) (response *QueryFailingReasonListForQualificationResponse, err error) {
	response = CreateQueryFailingReasonListForQualificationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFailingReasonListForQualificationWithChan invokes the domain.QueryFailingReasonListForQualification API asynchronously
func (client *Client) QueryFailingReasonListForQualificationWithChan(request *QueryFailingReasonListForQualificationRequest) (<-chan *QueryFailingReasonListForQualificationResponse, <-chan error) {
	responseChan := make(chan *QueryFailingReasonListForQualificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFailingReasonListForQualification(request)
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

// QueryFailingReasonListForQualificationWithCallback invokes the domain.QueryFailingReasonListForQualification API asynchronously
func (client *Client) QueryFailingReasonListForQualificationWithCallback(request *QueryFailingReasonListForQualificationRequest, callback func(response *QueryFailingReasonListForQualificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFailingReasonListForQualificationResponse
		var err error
		defer close(result)
		response, err = client.QueryFailingReasonListForQualification(request)
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

// QueryFailingReasonListForQualificationRequest is the request struct for api QueryFailingReasonListForQualification
type QueryFailingReasonListForQualificationRequest struct {
	*requests.RpcRequest
	QualificationType string           `position:"Query" name:"QualificationType"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	UserClientIp      string           `position:"Query" name:"UserClientIp"`
	Limit             requests.Integer `position:"Query" name:"Limit"`
	Lang              string           `position:"Query" name:"Lang"`
}

// QueryFailingReasonListForQualificationResponse is the response struct for api QueryFailingReasonListForQualification
type QueryFailingReasonListForQualificationResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Data      []FailRecord `json:"Data" xml:"Data"`
}

// CreateQueryFailingReasonListForQualificationRequest creates a request to invoke QueryFailingReasonListForQualification API
func CreateQueryFailingReasonListForQualificationRequest() (request *QueryFailingReasonListForQualificationRequest) {
	request = &QueryFailingReasonListForQualificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryFailingReasonListForQualification", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryFailingReasonListForQualificationResponse creates a response to parse from QueryFailingReasonListForQualification response
func CreateQueryFailingReasonListForQualificationResponse() (response *QueryFailingReasonListForQualificationResponse) {
	response = &QueryFailingReasonListForQualificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}