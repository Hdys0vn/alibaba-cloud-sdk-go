package sas

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

// ExportVul invokes the sas.ExportVul API synchronously
func (client *Client) ExportVul(request *ExportVulRequest) (response *ExportVulResponse, err error) {
	response = CreateExportVulResponse()
	err = client.DoAction(request, response)
	return
}

// ExportVulWithChan invokes the sas.ExportVul API asynchronously
func (client *Client) ExportVulWithChan(request *ExportVulRequest) (<-chan *ExportVulResponse, <-chan error) {
	responseChan := make(chan *ExportVulResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportVul(request)
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

// ExportVulWithCallback invokes the sas.ExportVul API asynchronously
func (client *Client) ExportVulWithCallback(request *ExportVulRequest, callback func(response *ExportVulResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportVulResponse
		var err error
		defer close(result)
		response, err = client.ExportVul(request)
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

// ExportVulRequest is the request struct for api ExportVul
type ExportVulRequest struct {
	*requests.RpcRequest
	StatusList          string           `position:"Query" name:"StatusList"`
	TargetType          string           `position:"Query" name:"TargetType"`
	MinScore            requests.Integer `position:"Query" name:"MinScore"`
	Remark              string           `position:"Query" name:"Remark"`
	AttachTypes         string           `position:"Query" name:"AttachTypes"`
	Type                string           `position:"Query" name:"Type"`
	VpcInstanceIds      string           `position:"Query" name:"VpcInstanceIds"`
	ContainerFieldName  string           `position:"Query" name:"ContainerFieldName"`
	SourceIp            string           `position:"Query" name:"SourceIp"`
	ContainerFieldValue string           `position:"Query" name:"ContainerFieldValue"`
	Lang                string           `position:"Query" name:"Lang"`
	Level               string           `position:"Query" name:"Level"`
	Resource            string           `position:"Query" name:"Resource"`
	GroupId             string           `position:"Query" name:"GroupId"`
	Dealed              string           `position:"Query" name:"Dealed"`
	ClusterId           string           `position:"Query" name:"ClusterId"`
	BatchName           string           `position:"Query" name:"BatchName"`
	AliasName           string           `position:"Query" name:"AliasName"`
	SearchTags          string           `position:"Query" name:"SearchTags"`
	Name                string           `position:"Query" name:"Name"`
	Necessity           string           `position:"Query" name:"Necessity"`
	Uuids               string           `position:"Query" name:"Uuids"`
}

// ExportVulResponse is the response struct for api ExportVul
type ExportVulResponse struct {
	*responses.BaseResponse
	Link         string `json:"Link" xml:"Link"`
	Progress     int    `json:"Progress" xml:"Progress"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	CurrentCount int    `json:"CurrentCount" xml:"CurrentCount"`
	Message      string `json:"Message" xml:"Message"`
	FileName     string `json:"FileName" xml:"FileName"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	ExportStatus string `json:"ExportStatus" xml:"ExportStatus"`
	Id           int64  `json:"Id" xml:"Id"`
}

// CreateExportVulRequest creates a request to invoke ExportVul API
func CreateExportVulRequest() (request *ExportVulRequest) {
	request = &ExportVulRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ExportVul", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportVulResponse creates a response to parse from ExportVul response
func CreateExportVulResponse() (response *ExportVulResponse) {
	response = &ExportVulResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
