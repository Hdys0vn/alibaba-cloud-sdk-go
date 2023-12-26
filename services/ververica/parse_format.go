package ververica

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

// ParseFormat invokes the ververica.ParseFormat API synchronously
func (client *Client) ParseFormat(request *ParseFormatRequest) (response *ParseFormatResponse, err error) {
	response = CreateParseFormatResponse()
	err = client.DoAction(request, response)
	return
}

// ParseFormatWithChan invokes the ververica.ParseFormat API asynchronously
func (client *Client) ParseFormatWithChan(request *ParseFormatRequest) (<-chan *ParseFormatResponse, <-chan error) {
	responseChan := make(chan *ParseFormatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ParseFormat(request)
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

// ParseFormatWithCallback invokes the ververica.ParseFormat API asynchronously
func (client *Client) ParseFormatWithCallback(request *ParseFormatRequest, callback func(response *ParseFormatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ParseFormatResponse
		var err error
		defer close(result)
		response, err = client.ParseFormat(request)
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

// ParseFormatRequest is the request struct for api ParseFormat
type ParseFormatRequest struct {
	*requests.RoaRequest
	Workspace  string `position:"Path" name:"workspace"`
	ParamsJson string `position:"Body" name:"paramsJson"`
	Namespace  string `position:"Path" name:"namespace"`
}

// ParseFormatResponse is the response struct for api ParseFormat
type ParseFormatResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateParseFormatRequest creates a request to invoke ParseFormat API
func CreateParseFormatRequest() (request *ParseFormatRequest) {
	request = &ParseFormatRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ParseFormat", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/formats:analyze-jars", "", "")
	request.Method = requests.POST
	return
}

// CreateParseFormatResponse creates a response to parse from ParseFormat response
func CreateParseFormatResponse() (response *ParseFormatResponse) {
	response = &ParseFormatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
