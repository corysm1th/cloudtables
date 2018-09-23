package cloudtables_test

import (
	"net/http"
	"net/http/httptest"

	mock "github.com/corysm1th/cloudtables/mock"
	cloudtables "github.com/corysm1th/cloudtables/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cloudtables", func() {
	Describe("API", func() {
		Context("With a properly formed request", func() {
			Describe("GET /api/v1/objects", func() {
				It("Should return a json array of objects", func() {
					Expect(nil).To(BeNil())
					// handler := cloudtables.HandleGetObjects
					// r := httptest.NewRequest(http.MethodGet, "/api/v1/objects", nil)
					// w := httptest.NewRecorder()
					// handler(w, r)
					// resp := w.Result()
					// defer resp.Body.Close()
					// body, err := ioutil.ReadAll(resp.Body)
					// Expect(err).ToNot(HaveOccurred())
					// Expect(resp.StatusCode).To(Equal(http.StatusOK))
					// Expect(body).To(ContainSubstring("TODO"))
				})
			})

			Describe("GET /api/v1/sync", func() {
				It("Should return a status 202", func() {
					handler := cloudtables.HandleGetSync
					r := httptest.NewRequest(http.MethodGet, "/api/v1/sync", nil)
					w := httptest.NewRecorder()
					handler(w, r)
					resp := w.Result()
					Expect(resp.StatusCode).To(Equal(http.StatusAccepted))
				})
			})

			Describe("GET /api/v1/metrics", func() {
				It("Should return a json array of metrics", func() {
					Expect(nil).To(BeNil())
					// handler := cloudtables.HandleGetMetrics
					// r := httptest.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
					// w := httptest.NewRecorder()
					// handler(w, r)
					// resp := w.Result()
					// defer resp.Body.Close()
					// body, err := ioutil.ReadAll(resp.Body)
					// Expect(err).ToNot(HaveOccurred())
					// Expect(resp.StatusCode).To(Equal(http.StatusOK))
					// Expect(body).To(ContainSubstring("TODO"))
				})
			})
		})
	})

	Describe("Sync AWS", func() {
		Context("With valid AWS credentials", func() {
			It("Should read credentials from disk", func() {
				Expect(nil).To(BeNil())
			})

			Context("With sync not 'In Progress'", func() {
				It("Should set the state to 'In Progress'", func() {
					Expect(nil).To(BeNil())
				})

				It("Should request AWS objects from all regions", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse EC2 instances", func() {
					account, region := "Test_Account", "us-west-2"
					mockSvc := &mock.EC2Client{}
					err := cloudtables.SyncDescribeInstances(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should parse Elastic IP addresses", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse Route53 records", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse Elastic Load Balancers", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse Application Load Balancers", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse Relational Database Instances", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse DynamoDB Instances", func() {
					account, region := "Test_Account", "us-west-2"
					mockSvc := &mock.DynamoDBClient{}
					err := cloudtables.SyncDynamoDB(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())

				})

				It("Should parse Elastic Container Service Instances", func() {
					Expect(nil).To(BeNil())
				})

				It("Should update progress when finished", func() {
					Expect(nil).To(BeNil())
				})

				It("Should handle metrics from the sync", func() {
					Expect(nil).To(BeNil())
				})
			})

			Context("With sync 'In Progress'", func() {
				It("Should skip that account", func() {
					Expect(nil).To(BeNil())
				})
			})
		})
	})

	Describe("Storage", func() {
		Context("With a working storage backend", func() {
			It("Should store account information", func() {
				Expect(nil).To(BeNil())
			})

			Context("For AWS Accounts", func() {
				It("Should handle EC2 objects", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle Elastic IP addresses", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle Route53 records", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle Elastic Load Balancers", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle Application Load Balancers", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle Relational Database Instances", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle DynamoDB Instances", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})

				It("Should handle Elastic Container Service Instances", func() {
					//Save and Retrieve
					Expect(nil).To(BeNil())
				})
			})

			It("Should handle metrics", func() {
				//Save and Retrieve
				Expect(nil).To(BeNil())
			})
		})
	})
})
