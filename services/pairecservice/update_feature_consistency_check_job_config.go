package pairecservice

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

// UpdateFeatureConsistencyCheckJobConfig invokes the pairecservice.UpdateFeatureConsistencyCheckJobConfig API synchronously
func (client *Client) UpdateFeatureConsistencyCheckJobConfig(request *UpdateFeatureConsistencyCheckJobConfigRequest) (response *UpdateFeatureConsistencyCheckJobConfigResponse, err error) {
	response = CreateUpdateFeatureConsistencyCheckJobConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFeatureConsistencyCheckJobConfigWithChan invokes the pairecservice.UpdateFeatureConsistencyCheckJobConfig API asynchronously
func (client *Client) UpdateFeatureConsistencyCheckJobConfigWithChan(request *UpdateFeatureConsistencyCheckJobConfigRequest) (<-chan *UpdateFeatureConsistencyCheckJobConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateFeatureConsistencyCheckJobConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFeatureConsistencyCheckJobConfig(request)
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

// UpdateFeatureConsistencyCheckJobConfigWithCallback invokes the pairecservice.UpdateFeatureConsistencyCheckJobConfig API asynchronously
func (client *Client) UpdateFeatureConsistencyCheckJobConfigWithCallback(request *UpdateFeatureConsistencyCheckJobConfigRequest, callback func(response *UpdateFeatureConsistencyCheckJobConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFeatureConsistencyCheckJobConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateFeatureConsistencyCheckJobConfig(request)
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

// UpdateFeatureConsistencyCheckJobConfigRequest is the request struct for api UpdateFeatureConsistencyCheckJobConfig
type UpdateFeatureConsistencyCheckJobConfigRequest struct {
	*requests.RoaRequest
	Body                               string `position:"Body" name:"body"`
	FeatureConsistencyCheckJobConfigId string `position:"Path" name:"FeatureConsistencyCheckJobConfigId"`
}

// UpdateFeatureConsistencyCheckJobConfigResponse is the response struct for api UpdateFeatureConsistencyCheckJobConfig
type UpdateFeatureConsistencyCheckJobConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateFeatureConsistencyCheckJobConfigRequest creates a request to invoke UpdateFeatureConsistencyCheckJobConfig API
func CreateUpdateFeatureConsistencyCheckJobConfigRequest() (request *UpdateFeatureConsistencyCheckJobConfigRequest) {
	request = &UpdateFeatureConsistencyCheckJobConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateFeatureConsistencyCheckJobConfig", "/api/v1/featureconsistencycheck/jobconfigs/[FeatureConsistencyCheckJobConfigId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateFeatureConsistencyCheckJobConfigResponse creates a response to parse from UpdateFeatureConsistencyCheckJobConfig response
func CreateUpdateFeatureConsistencyCheckJobConfigResponse() (response *UpdateFeatureConsistencyCheckJobConfigResponse) {
	response = &UpdateFeatureConsistencyCheckJobConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
