package cbn

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

// ModifyCenBandwidthPackageSpec invokes the cbn.ModifyCenBandwidthPackageSpec API synchronously
func (client *Client) ModifyCenBandwidthPackageSpec(request *ModifyCenBandwidthPackageSpecRequest) (response *ModifyCenBandwidthPackageSpecResponse, err error) {
	response = CreateModifyCenBandwidthPackageSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCenBandwidthPackageSpecWithChan invokes the cbn.ModifyCenBandwidthPackageSpec API asynchronously
func (client *Client) ModifyCenBandwidthPackageSpecWithChan(request *ModifyCenBandwidthPackageSpecRequest) (<-chan *ModifyCenBandwidthPackageSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyCenBandwidthPackageSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCenBandwidthPackageSpec(request)
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

// ModifyCenBandwidthPackageSpecWithCallback invokes the cbn.ModifyCenBandwidthPackageSpec API asynchronously
func (client *Client) ModifyCenBandwidthPackageSpecWithCallback(request *ModifyCenBandwidthPackageSpecRequest, callback func(response *ModifyCenBandwidthPackageSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCenBandwidthPackageSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyCenBandwidthPackageSpec(request)
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

// ModifyCenBandwidthPackageSpecRequest is the request struct for api ModifyCenBandwidthPackageSpec
type ModifyCenBandwidthPackageSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth             requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	Version               string           `position:"Query" name:"Version"`
	ServiceType           string           `position:"Query" name:"ServiceType"`
	CenBandwidthPackageId string           `position:"Query" name:"CenBandwidthPackageId"`
}

// ModifyCenBandwidthPackageSpecResponse is the response struct for api ModifyCenBandwidthPackageSpec
type ModifyCenBandwidthPackageSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCenBandwidthPackageSpecRequest creates a request to invoke ModifyCenBandwidthPackageSpec API
func CreateModifyCenBandwidthPackageSpecRequest() (request *ModifyCenBandwidthPackageSpecRequest) {
	request = &ModifyCenBandwidthPackageSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ModifyCenBandwidthPackageSpec", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyCenBandwidthPackageSpecResponse creates a response to parse from ModifyCenBandwidthPackageSpec response
func CreateModifyCenBandwidthPackageSpecResponse() (response *ModifyCenBandwidthPackageSpecResponse) {
	response = &ModifyCenBandwidthPackageSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
