package webplus

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

// CheckInstancesImportStatus invokes the webplus.CheckInstancesImportStatus API synchronously
// api document: https://help.aliyun.com/api/webplus/checkinstancesimportstatus.html
func (client *Client) CheckInstancesImportStatus(request *CheckInstancesImportStatusRequest) (response *CheckInstancesImportStatusResponse, err error) {
	response = CreateCheckInstancesImportStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckInstancesImportStatusWithChan invokes the webplus.CheckInstancesImportStatus API asynchronously
// api document: https://help.aliyun.com/api/webplus/checkinstancesimportstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckInstancesImportStatusWithChan(request *CheckInstancesImportStatusRequest) (<-chan *CheckInstancesImportStatusResponse, <-chan error) {
	responseChan := make(chan *CheckInstancesImportStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckInstancesImportStatus(request)
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

// CheckInstancesImportStatusWithCallback invokes the webplus.CheckInstancesImportStatus API asynchronously
// api document: https://help.aliyun.com/api/webplus/checkinstancesimportstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckInstancesImportStatusWithCallback(request *CheckInstancesImportStatusRequest, callback func(response *CheckInstancesImportStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckInstancesImportStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckInstancesImportStatus(request)
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

// CheckInstancesImportStatusRequest is the request struct for api CheckInstancesImportStatus
type CheckInstancesImportStatusRequest struct {
	*requests.RoaRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

// CheckInstancesImportStatusResponse is the response struct for api CheckInstancesImportStatus
type CheckInstancesImportStatusResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Code                  string                `json:"Code" xml:"Code"`
	Message               string                `json:"Message" xml:"Message"`
	InstancesImportStatus InstancesImportStatus `json:"InstancesImportStatus" xml:"InstancesImportStatus"`
}

// CreateCheckInstancesImportStatusRequest creates a request to invoke CheckInstancesImportStatus API
func CreateCheckInstancesImportStatusRequest() (request *CheckInstancesImportStatusRequest) {
	request = &CheckInstancesImportStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "CheckInstancesImportStatus", "/pop/v1/wam/instance/ecsImportStat", "", "")
	request.Method = requests.GET
	return
}

// CreateCheckInstancesImportStatusResponse creates a response to parse from CheckInstancesImportStatus response
func CreateCheckInstancesImportStatusResponse() (response *CheckInstancesImportStatusResponse) {
	response = &CheckInstancesImportStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
