package sgw

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

// ReportBlockVolumes invokes the sgw.ReportBlockVolumes API synchronously
func (client *Client) ReportBlockVolumes(request *ReportBlockVolumesRequest) (response *ReportBlockVolumesResponse, err error) {
	response = CreateReportBlockVolumesResponse()
	err = client.DoAction(request, response)
	return
}

// ReportBlockVolumesWithChan invokes the sgw.ReportBlockVolumes API asynchronously
func (client *Client) ReportBlockVolumesWithChan(request *ReportBlockVolumesRequest) (<-chan *ReportBlockVolumesResponse, <-chan error) {
	responseChan := make(chan *ReportBlockVolumesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportBlockVolumes(request)
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

// ReportBlockVolumesWithCallback invokes the sgw.ReportBlockVolumes API asynchronously
func (client *Client) ReportBlockVolumesWithCallback(request *ReportBlockVolumesRequest, callback func(response *ReportBlockVolumesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportBlockVolumesResponse
		var err error
		defer close(result)
		response, err = client.ReportBlockVolumes(request)
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

// ReportBlockVolumesRequest is the request struct for api ReportBlockVolumes
type ReportBlockVolumesRequest struct {
	*requests.RpcRequest
	ClientUUID    string `position:"Query" name:"ClientUUID"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
	Info          string `position:"Query" name:"Info"`
}

// ReportBlockVolumesResponse is the response struct for api ReportBlockVolumes
type ReportBlockVolumesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateReportBlockVolumesRequest creates a request to invoke ReportBlockVolumes API
func CreateReportBlockVolumesRequest() (request *ReportBlockVolumesRequest) {
	request = &ReportBlockVolumesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "ReportBlockVolumes", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportBlockVolumesResponse creates a response to parse from ReportBlockVolumes response
func CreateReportBlockVolumesResponse() (response *ReportBlockVolumesResponse) {
	response = &ReportBlockVolumesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
