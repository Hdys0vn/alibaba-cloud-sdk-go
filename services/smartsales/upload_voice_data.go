package smartsales

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

// UploadVoiceData invokes the smartsales.UploadVoiceData API synchronously
func (client *Client) UploadVoiceData(request *UploadVoiceDataRequest) (response *UploadVoiceDataResponse, err error) {
	response = CreateUploadVoiceDataResponse()
	err = client.DoAction(request, response)
	return
}

// UploadVoiceDataWithChan invokes the smartsales.UploadVoiceData API asynchronously
func (client *Client) UploadVoiceDataWithChan(request *UploadVoiceDataRequest) (<-chan *UploadVoiceDataResponse, <-chan error) {
	responseChan := make(chan *UploadVoiceDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadVoiceData(request)
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

// UploadVoiceDataWithCallback invokes the smartsales.UploadVoiceData API asynchronously
func (client *Client) UploadVoiceDataWithCallback(request *UploadVoiceDataRequest, callback func(response *UploadVoiceDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadVoiceDataResponse
		var err error
		defer close(result)
		response, err = client.UploadVoiceData(request)
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

// UploadVoiceDataRequest is the request struct for api UploadVoiceData
type UploadVoiceDataRequest struct {
	*requests.RpcRequest
	AgentKey      string                          `position:"Query" name:"AgentKey"`
	VoiceDataList *[]UploadVoiceDataVoiceDataList `position:"Query" name:"VoiceDataList"  type:"Json"`
}

// UploadVoiceDataVoiceDataList is a repeated param struct in UploadVoiceDataRequest
type UploadVoiceDataVoiceDataList struct {
	CustomParamJson   string `name:"CustomParamJson"`
	ClueStatusRemark  string `name:"ClueStatusRemark"`
	BeginTime         string `name:"BeginTime"`
	CallChannel       string `name:"CallChannel"`
	BusinessResult    string `name:"BusinessResult"`
	CustomCallId      string `name:"CustomCallId"`
	SalesPersonId     string `name:"SalesPersonId"`
	TeamId            string `name:"TeamId"`
	CustomerId        string `name:"CustomerId"`
	FileUrl           string `name:"FileUrl"`
	ClientTrackNumber string `name:"ClientTrackNumber"`
	CustomerName      string `name:"CustomerName"`
	CallType          string `name:"CallType"`
	CalleeNumber      string `name:"CalleeNumber"`
	SalesPersonName   string `name:"SalesPersonName"`
	CallerNumber      string `name:"CallerNumber"`
}

// UploadVoiceDataResponse is the response struct for api UploadVoiceData
type UploadVoiceDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateUploadVoiceDataRequest creates a request to invoke UploadVoiceData API
func CreateUploadVoiceDataRequest() (request *UploadVoiceDataRequest) {
	request = &UploadVoiceDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SmartSales", "2022-08-18", "UploadVoiceData", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadVoiceDataResponse creates a response to parse from UploadVoiceData response
func CreateUploadVoiceDataResponse() (response *UploadVoiceDataResponse) {
	response = &UploadVoiceDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
