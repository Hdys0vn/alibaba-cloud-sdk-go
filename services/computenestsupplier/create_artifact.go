package computenestsupplier

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

// CreateArtifact invokes the computenestsupplier.CreateArtifact API synchronously
func (client *Client) CreateArtifact(request *CreateArtifactRequest) (response *CreateArtifactResponse, err error) {
	response = CreateCreateArtifactResponse()
	err = client.DoAction(request, response)
	return
}

// CreateArtifactWithChan invokes the computenestsupplier.CreateArtifact API asynchronously
func (client *Client) CreateArtifactWithChan(request *CreateArtifactRequest) (<-chan *CreateArtifactResponse, <-chan error) {
	responseChan := make(chan *CreateArtifactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateArtifact(request)
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

// CreateArtifactWithCallback invokes the computenestsupplier.CreateArtifact API asynchronously
func (client *Client) CreateArtifactWithCallback(request *CreateArtifactRequest, callback func(response *CreateArtifactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateArtifactResponse
		var err error
		defer close(result)
		response, err = client.CreateArtifact(request)
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

// CreateArtifactRequest is the request struct for api CreateArtifact
type CreateArtifactRequest struct {
	*requests.RpcRequest
	Description      string    `position:"Query" name:"Description"`
	SupportRegionIds *[]string `position:"Query" name:"SupportRegionIds"  type:"Repeated"`
	ArtifactType     string    `position:"Query" name:"ArtifactType"`
	Name             string    `position:"Query" name:"Name"`
	ArtifactId       string    `position:"Query" name:"ArtifactId"`
	ArtifactProperty string    `position:"Query" name:"ArtifactProperty"`
	VersionName      string    `position:"Query" name:"VersionName"`
}

// CreateArtifactResponse is the response struct for api CreateArtifact
type CreateArtifactResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ArtifactId       string `json:"ArtifactId" xml:"ArtifactId"`
	ArtifactType     string `json:"ArtifactType" xml:"ArtifactType"`
	Name             string `json:"Name" xml:"Name"`
	VersionName      string `json:"VersionName" xml:"VersionName"`
	ArtifactVersion  string `json:"ArtifactVersion" xml:"ArtifactVersion"`
	Description      string `json:"Description" xml:"Description"`
	GmtModified      string `json:"GmtModified" xml:"GmtModified"`
	Status           string `json:"Status" xml:"Status"`
	MaxVersion       int64  `json:"MaxVersion" xml:"MaxVersion"`
	ArtifactProperty string `json:"ArtifactProperty" xml:"ArtifactProperty"`
	SupportRegionIds string `json:"SupportRegionIds" xml:"SupportRegionIds"`
}

// CreateCreateArtifactRequest creates a request to invoke CreateArtifact API
func CreateCreateArtifactRequest() (request *CreateArtifactRequest) {
	request = &CreateArtifactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "CreateArtifact", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateArtifactResponse creates a response to parse from CreateArtifact response
func CreateCreateArtifactResponse() (response *CreateArtifactResponse) {
	response = &CreateArtifactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
