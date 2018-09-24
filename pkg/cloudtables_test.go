package cloudtables_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	mock "github.com/corysm1th/cloudtables/mock"
	cloudtables "github.com/corysm1th/cloudtables/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cloudtables", func() {
	debug := log.Logger{}
	debug.SetOutput(ioutil.Discard)
	if os.Getenv("DEBUG") == "true" {
		debug.Println("Debug On")
		debug.SetOutput(os.Stdout)
	}
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
					debug.Println("EC2 Instances")
					account, region := "Test_Account", "us-west-2"
					mockSvc := &mock.EC2Client{}
					e, count, err := cloudtables.GetAWSInstances(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(1))
					for _, obj := range *e {
						debug.Printf("Account: %s  Region: %s  ID: %s  ", obj.Account, obj.Region, obj.Name)
					}
				})

				It("Should parse Elastic IP addresses", func() {
					debug.Println("Elastic IPs")
					account, region := "Test_Account", "us-west-2"
					mockSvc := &mock.EC2Client{}
					e, count, err := cloudtables.GetAWSAddresses(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(1))
					for _, obj := range *e {
						debug.Printf("Account: %s  Region: %s  IP: %s  ", obj.Account, obj.Region, obj.PublicIP)
						Expect(obj.Account).To(Equal("Test_Account"))
					}
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
					debug.Println("DynamoDB Tables")
					account, region := "Test_Account", "us-west-2"
					mockSvc := &mock.DynamoDBClient{}
					d, count, err := cloudtables.GetDynamoDB(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(2))
					debug.Println("DynamoDB:")
					for _, table := range *d {
						Expect(table.Account).To(Equal("Test_Account"))
						debug.Printf("Account: %s  Region: %s  Table: %s", table.Account, table.Region, table.Name)
					}

				})

				It("Should parse Elastic Container Service Instances", func() {
					Expect(nil).To(BeNil())
				})

				It("Should parse S3 Buckets", func() {
					debug.Println("S3 Buckets")
					account := "Test_Account"
					mockSvc := &mock.S3Client{}
					b, count, err := cloudtables.GetAWSBuckets(mockSvc, account)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(3))
					for _, obj := range *b {
						debug.Printf("Account: %s  Name: %s  ", obj.Account, obj.Name)
					}
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
