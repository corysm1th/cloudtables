package cloudtables_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	mock "github.com/corysm1th/cloudtables/mock"
	. "github.com/corysm1th/cloudtables/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	err      error
	listener net.Listener
	URL      string
	store    Storage
)

var _ = Describe("Cloudtables", func() {
	debug := log.Logger{}
	debug.SetOutput(ioutil.Discard)
	if os.Getenv("DEBUG") == "true" {
		debug.Println("Debug On")
		debug.SetOutput(os.Stdout)
	}
	config := Config{
		Addr:       "127.0.0.1:9000",
		MutualAuth: false,
		Storage:    "memory",
	}

	URL = fmt.Sprintf("http://%s/api/v1", config.Addr)

	// TODO: Run a sync against the mock AWS API to populate objects.
	Describe("API /api/v1/", func() {
		BeforeEach(func() {
			// New server instance
			store = NewStorageMem()
			listener, err = net.Listen("tcp", config.Addr)
			Expect(err).To(BeNil())
			state := NewState()
			Run(&config, store, listener, state)

		})
		AfterEach(func() {
			listener.Close()
		})
		Context("With a properly formed request", func() {
			Describe("GET objects", func() {
				It("Should return a json array of objects", func() {
					URI := fmt.Sprintf("%s/objects", URL)
					request, err := http.NewRequest("GET", URI, nil)
					Expect(err).To(BeNil())
					client := &http.Client{Timeout: 10 * time.Second}
					resp, err := client.Do(request)
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
				})
			})

			Describe("GET ping", func() {
				It("Should return 204", func() {
					URI := fmt.Sprintf("%s/ping", URL)
					request, err := http.NewRequest("GET", URI, nil)
					Expect(err).To(BeNil())
					client := &http.Client{Timeout: 10 * time.Second}
					resp, err := client.Do(request)
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
				})
			})

			Describe("GET /api/v1/sync", func() {
				It("Should return a status 202", func() {
					// r := httptest.NewRequest(http.MethodGet, "/api/v1/sync", nil)
					// w := httptest.NewRecorder()
					// resp := w.Result()
					// Expect(resp.StatusCode).To(Equal(http.StatusAccepted))
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
					e, count, err := GetAWSInstances(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(1))
					for _, obj := range e {
						debug.Printf("Account: %s  Region: %s  ID: %s  ", obj.Account, obj.Region, obj.Name)
					}
				})

				It("Should parse Elastic IP addresses", func() {
					debug.Println("Elastic IPs")
					account, region := "Test_Account", "us-west-2"
					mockSvc := &mock.EC2Client{}
					e, count, err := GetAWSAddresses(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(1))
					for _, obj := range e {
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
					d, count, err := GetDynamoDB(mockSvc, account, region)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(2))
					debug.Println("DynamoDB:")
					for _, table := range d {
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
					b, count, err := GetAWSBuckets(mockSvc, account)
					Expect(err).ToNot(HaveOccurred())
					Expect(count).To(Equal(3))
					for _, obj := range b {
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

	Describe("State", func() {
		var err error
		var state *State

		BeforeEach(func() {
			state = NewState()
			names := []string{"prod", "non-prod", "ops"}
			for _, n := range names {
				err = state.AddAccount("AWS", n)
				Expect(err).To(BeNil())
			}
		})

		It("Should Add and Account", func() {
			state.AddAccount("AWS", "foo")
			counter := 0
			for _, a := range state.Accounts {
				if a.Name == "foo" {
					counter++
				}
			}
			Expect(counter).To(Equal(1))
		})

		It("Should set the state of an account", func() {
			err = state.SetState("AWS", "non-prod", SyncComplete)
			Expect(err).To(BeNil())
			for _, a := range state.Accounts {
				switch a.Name {
				case "non-prod":
					Expect(a.State).To(Equal(SyncComplete))
				}
			}
		})

		It("Should return account states", func() {
			s := state.GetState()
			Expect(len(s)).To(Equal(3))
		})

	})

	Describe("Storage", func() {
		Context("With a working storage backend", func() {
			BeforeEach(func() {
				store = NewStorageMem()
			})

			It("Should store account information", func() {
				Expect(nil).To(BeNil())
			})

			It("Should handle EC2 objects", func() {
				//Save and Retrieve
				objs := mock.CreateEC2Instances()
				err := store.StoreEC2InstObj(objs)
				Expect(err).To(BeNil())
				ec2s, err := store.SelectEC2InstObj()
				Expect(err).To(BeNil())
				Expect(len(ec2s)).To(Equal(4))
			})

			It("Should handle Elastic IP addresses", func() {
				//Save and Retrieve
				objs := mock.CreateEIPs()
				err := store.StoreEC2EIPObj(objs)
				Expect(err).To(BeNil())
				eips, err := store.SelectEC2EIPObj()
				Expect(err).To(BeNil())
				Expect(len(eips)).To(Equal(4))
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
				objs := mock.CreateDynamoDBs()
				err := store.StoreDynamoDBObj(objs)
				Expect(err).To(BeNil())
				ddbs, err := store.SelectDynamoDBObj()
				Expect(err).To(BeNil())
				Expect(len(ddbs)).To(Equal(4))

			})

			It("Should handle Elastic Container Service Instances", func() {
				//Save and Retrieve
				Expect(nil).To(BeNil())
			})

			It("Should handle S3 Buckets", func() {
				// Save and Retrieve
				objs := mock.CreateBuckets()
				err := store.StoreS3BucketObj(objs)
				Expect(err).To(BeNil())
				buckets, err := store.SelectS3BucketObj()
				Expect(len(buckets)).To(Equal(4))
			})

			It("Should handle metrics", func() {
				//Save and Retrieve
				Expect(nil).To(BeNil())
			})
		})
	})
})
