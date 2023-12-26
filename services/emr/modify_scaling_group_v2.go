package emr

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

// ModifyScalingGroupV2 invokes the emr.ModifyScalingGroupV2 API synchronously
func (client *Client) ModifyScalingGroupV2(request *ModifyScalingGroupV2Request) (response *ModifyScalingGroupV2Response, err error) {
	response = CreateModifyScalingGroupV2Response()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingGroupV2WithChan invokes the emr.ModifyScalingGroupV2 API asynchronously
func (client *Client) ModifyScalingGroupV2WithChan(request *ModifyScalingGroupV2Request) (<-chan *ModifyScalingGroupV2Response, <-chan error) {
	responseChan := make(chan *ModifyScalingGroupV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingGroupV2(request)
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

// ModifyScalingGroupV2WithCallback invokes the emr.ModifyScalingGroupV2 API asynchronously
func (client *Client) ModifyScalingGroupV2WithCallback(request *ModifyScalingGroupV2Request, callback func(response *ModifyScalingGroupV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingGroupV2Response
		var err error
		defer close(result)
		response, err = client.ModifyScalingGroupV2(request)
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

// ModifyScalingGroupV2Request is the request struct for api ModifyScalingGroupV2
type ModifyScalingGroupV2Request struct {
	*requests.RpcRequest
	ResourceOwnerId   requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description       string           `position:"Query" name:"Description"`
	ScalingGroupBizId string           `position:"Query" name:"ScalingGroupBizId"`
	ResourceGroupId   string           `position:"Query" name:"ResourceGroupId"`
	Name              string           `position:"Query" name:"Name"`
}

// ModifyScalingGroupV2Response is the response struct for api ModifyScalingGroupV2
type ModifyScalingGroupV2Response struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateModifyScalingGroupV2Request creates a request to invoke ModifyScalingGroupV2 API
func CreateModifyScalingGroupV2Request() (request *ModifyScalingGroupV2Request) {
	request = &ModifyScalingGroupV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyScalingGroupV2", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyScalingGroupV2Response creates a response to parse from ModifyScalingGroupV2 response
func CreateModifyScalingGroupV2Response() (response *ModifyScalingGroupV2Response) {
	response = &ModifyScalingGroupV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
