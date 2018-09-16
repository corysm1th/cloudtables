package cloudtables_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cloudtables", func() {
	Describe("UI", func() {
		Context("Without UI certs on disk", func() {
			It("Should listen on HTTP", func() {
				Expect(nil).To(BeNil())
			})
		})

		Context("With UI certs on disk", func() {
			It("Should listen with TLS mutual auth", func() {
				Expect(nil).To(BeNil())
			})

			Context("Without a valid client certificate", func() {
				It("Should refuse connections", func() {
					Expect(nil).To(BeNil())
				})
			})

			Context("With a valid client certificate", func() {
				It("Should accept UI requests", func() {
					Expect(nil).To(BeNil())
				})
			})
		})
	})

	Describe("API", func() {
		Context("Without API certs on disk", func() {
			It("Should listen on HTTP", func() {
				Expect(nil).To(BeNil())
			})
		})

		Context("With API certs on disk", func() {
			It("Should listen on TLS", func() {
				Expect(nil).To(BeNil())
			})

			Context("Without a valid client certificate", func() {
				It("Should refuse connections", func() {
					Expect(nil).To(BeNil())
				})
			})

			Context("With a valid client certificate", func() {
				It("Should accept API requests", func() {
					Expect(nil).To(BeNil())
				})

				Describe("GET /api/v1/objects", func() {
					It("Should return a json array of objects", func() {
						Expect(nil).To(BeNil())
					})
				})

				Describe("GET /api/v1/sync", func() {
					It("Should return a status 201", func() {
						Expect(nil).To(BeNil())
					})
				})

				Describe("GET /api/v1/metrics", func() {
					It("Should return a json array of metrics", func() {
						Expect(nil).To(BeNil())
					})
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
					Expect(nil).To(BeNil())
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
					Expect(nil).To(BeNil())
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
