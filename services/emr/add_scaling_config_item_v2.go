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

// AddScalingConfigItemV2 invokes the emr.AddScalingConfigItemV2 API synchronously
func (client *Client) AddScalingConfigItemV2(request *AddScalingConfigItemV2Request) (response *AddScalingConfigItemV2Response, err error) {
	response = CreateAddScalingConfigItemV2Response()
	err = client.DoAction(request, response)
	return
}

// AddScalingConfigItemV2WithChan invokes the emr.AddScalingConfigItemV2 API asynchronously
func (client *Client) AddScalingConfigItemV2WithChan(request *AddScalingConfigItemV2Request) (<-chan *AddScalingConfigItemV2Response, <-chan error) {
	responseChan := make(chan *AddScalingConfigItemV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddScalingConfigItemV2(request)
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

// AddScalingConfigItemV2WithCallback invokes the emr.AddScalingConfigItemV2 API asynchronously
func (client *Client) AddScalingConfigItemV2WithCallback(request *AddScalingConfigItemV2Request, callback func(response *AddScalingConfigItemV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddScalingConfigItemV2Response
		var err error
		defer close(result)
		response, err = client.AddScalingConfigItemV2(request)
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

// AddScalingConfigItemV2Request is the request struct for api AddScalingConfigItemV2
type AddScalingConfigItemV2Request struct {
	*requests.RpcRequest
	ConfigItemType        string           `position:"Query" name:"ConfigItemType"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupBizId     string           `position:"Query" name:"ScalingGroupBizId"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	ConfigItemInformation string           `position:"Query" name:"ConfigItemInformation"`
}

// AddScalingConfigItemV2Response is the response struct for api AddScalingConfigItemV2
type AddScalingConfigItemV2Response struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateAddScalingConfigItemV2Request creates a request to invoke AddScalingConfigItemV2 API
func CreateAddScalingConfigItemV2Request() (request *AddScalingConfigItemV2Request) {
	request = &AddScalingConfigItemV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "AddScalingConfigItemV2", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddScalingConfigItemV2Response creates a response to parse from AddScalingConfigItemV2 response
func CreateAddScalingConfigItemV2Response() (response *AddScalingConfigItemV2Response) {
	response = &AddScalingConfigItemV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
