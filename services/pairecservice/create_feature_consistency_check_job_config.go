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

// CreateFeatureConsistencyCheckJobConfig invokes the pairecservice.CreateFeatureConsistencyCheckJobConfig API synchronously
func (client *Client) CreateFeatureConsistencyCheckJobConfig(request *CreateFeatureConsistencyCheckJobConfigRequest) (response *CreateFeatureConsistencyCheckJobConfigResponse, err error) {
	response = CreateCreateFeatureConsistencyCheckJobConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFeatureConsistencyCheckJobConfigWithChan invokes the pairecservice.CreateFeatureConsistencyCheckJobConfig API asynchronously
func (client *Client) CreateFeatureConsistencyCheckJobConfigWithChan(request *CreateFeatureConsistencyCheckJobConfigRequest) (<-chan *CreateFeatureConsistencyCheckJobConfigResponse, <-chan error) {
	responseChan := make(chan *CreateFeatureConsistencyCheckJobConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFeatureConsistencyCheckJobConfig(request)
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

// CreateFeatureConsistencyCheckJobConfigWithCallback invokes the pairecservice.CreateFeatureConsistencyCheckJobConfig API asynchronously
func (client *Client) CreateFeatureConsistencyCheckJobConfigWithCallback(request *CreateFeatureConsistencyCheckJobConfigRequest, callback func(response *CreateFeatureConsistencyCheckJobConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFeatureConsistencyCheckJobConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateFeatureConsistencyCheckJobConfig(request)
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

// CreateFeatureConsistencyCheckJobConfigRequest is the request struct for api CreateFeatureConsistencyCheckJobConfig
type CreateFeatureConsistencyCheckJobConfigRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateFeatureConsistencyCheckJobConfigResponse is the response struct for api CreateFeatureConsistencyCheckJobConfig
type CreateFeatureConsistencyCheckJobConfigResponse struct {
	*responses.BaseResponse
	RequestId                          string `json:"RequestId" xml:"RequestId"`
	FeatureConsistencyCheckJobConfigId string `json:"FeatureConsistencyCheckJobConfigId" xml:"FeatureConsistencyCheckJobConfigId"`
}

// CreateCreateFeatureConsistencyCheckJobConfigRequest creates a request to invoke CreateFeatureConsistencyCheckJobConfig API
func CreateCreateFeatureConsistencyCheckJobConfigRequest() (request *CreateFeatureConsistencyCheckJobConfigRequest) {
	request = &CreateFeatureConsistencyCheckJobConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateFeatureConsistencyCheckJobConfig", "/api/v1/featureconsistencycheck/jobconfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFeatureConsistencyCheckJobConfigResponse creates a response to parse from CreateFeatureConsistencyCheckJobConfig response
func CreateCreateFeatureConsistencyCheckJobConfigResponse() (response *CreateFeatureConsistencyCheckJobConfigResponse) {
	response = &CreateFeatureConsistencyCheckJobConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
