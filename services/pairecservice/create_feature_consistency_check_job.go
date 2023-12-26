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

// CreateFeatureConsistencyCheckJob invokes the pairecservice.CreateFeatureConsistencyCheckJob API synchronously
func (client *Client) CreateFeatureConsistencyCheckJob(request *CreateFeatureConsistencyCheckJobRequest) (response *CreateFeatureConsistencyCheckJobResponse, err error) {
	response = CreateCreateFeatureConsistencyCheckJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFeatureConsistencyCheckJobWithChan invokes the pairecservice.CreateFeatureConsistencyCheckJob API asynchronously
func (client *Client) CreateFeatureConsistencyCheckJobWithChan(request *CreateFeatureConsistencyCheckJobRequest) (<-chan *CreateFeatureConsistencyCheckJobResponse, <-chan error) {
	responseChan := make(chan *CreateFeatureConsistencyCheckJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFeatureConsistencyCheckJob(request)
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

// CreateFeatureConsistencyCheckJobWithCallback invokes the pairecservice.CreateFeatureConsistencyCheckJob API asynchronously
func (client *Client) CreateFeatureConsistencyCheckJobWithCallback(request *CreateFeatureConsistencyCheckJobRequest, callback func(response *CreateFeatureConsistencyCheckJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFeatureConsistencyCheckJobResponse
		var err error
		defer close(result)
		response, err = client.CreateFeatureConsistencyCheckJob(request)
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

// CreateFeatureConsistencyCheckJobRequest is the request struct for api CreateFeatureConsistencyCheckJob
type CreateFeatureConsistencyCheckJobRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateFeatureConsistencyCheckJobResponse is the response struct for api CreateFeatureConsistencyCheckJob
type CreateFeatureConsistencyCheckJobResponse struct {
	*responses.BaseResponse
	RequestId                    string `json:"RequestId" xml:"RequestId"`
	FeatureConsistencyCheckJobId string `json:"FeatureConsistencyCheckJobId" xml:"FeatureConsistencyCheckJobId"`
}

// CreateCreateFeatureConsistencyCheckJobRequest creates a request to invoke CreateFeatureConsistencyCheckJob API
func CreateCreateFeatureConsistencyCheckJobRequest() (request *CreateFeatureConsistencyCheckJobRequest) {
	request = &CreateFeatureConsistencyCheckJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateFeatureConsistencyCheckJob", "/api/v1/featureconsistencycheck/jobs", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFeatureConsistencyCheckJobResponse creates a response to parse from CreateFeatureConsistencyCheckJob response
func CreateCreateFeatureConsistencyCheckJobResponse() (response *CreateFeatureConsistencyCheckJobResponse) {
	response = &CreateFeatureConsistencyCheckJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
