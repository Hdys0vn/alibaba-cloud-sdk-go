package antiddos_public

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

// ModifyIpDefenseThreshold invokes the antiddos_public.ModifyIpDefenseThreshold API synchronously
func (client *Client) ModifyIpDefenseThreshold(request *ModifyIpDefenseThresholdRequest) (response *ModifyIpDefenseThresholdResponse, err error) {
	response = CreateModifyIpDefenseThresholdResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIpDefenseThresholdWithChan invokes the antiddos_public.ModifyIpDefenseThreshold API asynchronously
func (client *Client) ModifyIpDefenseThresholdWithChan(request *ModifyIpDefenseThresholdRequest) (<-chan *ModifyIpDefenseThresholdResponse, <-chan error) {
	responseChan := make(chan *ModifyIpDefenseThresholdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIpDefenseThreshold(request)
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

// ModifyIpDefenseThresholdWithCallback invokes the antiddos_public.ModifyIpDefenseThreshold API asynchronously
func (client *Client) ModifyIpDefenseThresholdWithCallback(request *ModifyIpDefenseThresholdRequest, callback func(response *ModifyIpDefenseThresholdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIpDefenseThresholdResponse
		var err error
		defer close(result)
		response, err = client.ModifyIpDefenseThreshold(request)
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

// ModifyIpDefenseThresholdRequest is the request struct for api ModifyIpDefenseThreshold
type ModifyIpDefenseThresholdRequest struct {
	*requests.RpcRequest
	InternetIp   string           `position:"Query" name:"InternetIp"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	DdosRegionId string           `position:"Query" name:"DdosRegionId"`
	InstanceType string           `position:"Query" name:"InstanceType"`
	Lang         string           `position:"Query" name:"Lang"`
	Bps          requests.Integer `position:"Query" name:"Bps"`
	Pps          requests.Integer `position:"Query" name:"Pps"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	IsAuto       requests.Boolean `position:"Query" name:"IsAuto"`
}

// ModifyIpDefenseThresholdResponse is the response struct for api ModifyIpDefenseThreshold
type ModifyIpDefenseThresholdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIpDefenseThresholdRequest creates a request to invoke ModifyIpDefenseThreshold API
func CreateModifyIpDefenseThresholdRequest() (request *ModifyIpDefenseThresholdRequest) {
	request = &ModifyIpDefenseThresholdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "ModifyIpDefenseThreshold", "ddosbasic", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyIpDefenseThresholdResponse creates a response to parse from ModifyIpDefenseThreshold response
func CreateModifyIpDefenseThresholdResponse() (response *ModifyIpDefenseThresholdResponse) {
	response = &ModifyIpDefenseThresholdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}