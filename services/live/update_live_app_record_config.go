package live

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

// UpdateLiveAppRecordConfig invokes the live.UpdateLiveAppRecordConfig API synchronously
func (client *Client) UpdateLiveAppRecordConfig(request *UpdateLiveAppRecordConfigRequest) (response *UpdateLiveAppRecordConfigResponse, err error) {
	response = CreateUpdateLiveAppRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveAppRecordConfigWithChan invokes the live.UpdateLiveAppRecordConfig API asynchronously
func (client *Client) UpdateLiveAppRecordConfigWithChan(request *UpdateLiveAppRecordConfigRequest) (<-chan *UpdateLiveAppRecordConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveAppRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveAppRecordConfig(request)
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

// UpdateLiveAppRecordConfigWithCallback invokes the live.UpdateLiveAppRecordConfig API asynchronously
func (client *Client) UpdateLiveAppRecordConfigWithCallback(request *UpdateLiveAppRecordConfigRequest, callback func(response *UpdateLiveAppRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveAppRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveAppRecordConfig(request)
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

// UpdateLiveAppRecordConfigRequest is the request struct for api UpdateLiveAppRecordConfig
type UpdateLiveAppRecordConfigRequest struct {
	*requests.RpcRequest
	OssEndpoint           string                                            `position:"Query" name:"OssEndpoint"`
	DelayTime             requests.Integer                                  `position:"Query" name:"DelayTime"`
	TranscodeTemplates    *[]string                                         `position:"Query" name:"TranscodeTemplates"  type:"Repeated"`
	StartTime             string                                            `position:"Query" name:"StartTime"`
	AppName               string                                            `position:"Query" name:"AppName"`
	SecurityToken         string                                            `position:"Query" name:"SecurityToken"`
	TranscodeRecordFormat *[]UpdateLiveAppRecordConfigTranscodeRecordFormat `position:"Query" name:"TranscodeRecordFormat"  type:"Repeated"`
	OnDemand              requests.Integer                                  `position:"Query" name:"OnDemand"`
	StreamName            string                                            `position:"Query" name:"StreamName"`
	DomainName            string                                            `position:"Query" name:"DomainName"`
	EndTime               string                                            `position:"Query" name:"EndTime"`
	OwnerId               requests.Integer                                  `position:"Query" name:"OwnerId"`
	RecordFormat          *[]UpdateLiveAppRecordConfigRecordFormat          `position:"Query" name:"RecordFormat"  type:"Repeated"`
}

// UpdateLiveAppRecordConfigTranscodeRecordFormat is a repeated param struct in UpdateLiveAppRecordConfigRequest
type UpdateLiveAppRecordConfigTranscodeRecordFormat struct {
	SliceDuration string `name:"SliceDuration"`
	Format        string `name:"Format"`
	CycleDuration string `name:"CycleDuration"`
}

// UpdateLiveAppRecordConfigRecordFormat is a repeated param struct in UpdateLiveAppRecordConfigRequest
type UpdateLiveAppRecordConfigRecordFormat struct {
	SliceDuration string `name:"SliceDuration"`
	Format        string `name:"Format"`
	CycleDuration string `name:"CycleDuration"`
}

// UpdateLiveAppRecordConfigResponse is the response struct for api UpdateLiveAppRecordConfig
type UpdateLiveAppRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveAppRecordConfigRequest creates a request to invoke UpdateLiveAppRecordConfig API
func CreateUpdateLiveAppRecordConfigRequest() (request *UpdateLiveAppRecordConfigRequest) {
	request = &UpdateLiveAppRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLiveAppRecordConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLiveAppRecordConfigResponse creates a response to parse from UpdateLiveAppRecordConfig response
func CreateUpdateLiveAppRecordConfigResponse() (response *UpdateLiveAppRecordConfigResponse) {
	response = &UpdateLiveAppRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
