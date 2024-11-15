package main_test

import (
	"import-testing/bank"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("bank", func() {
	Context("Testing Valid Amount Deposit", func() {
		account := bank.Account{}
		It("should return nil", func() {
			var amount float64 = 100.50
			err := account.Deposit(amount)
			Expect(err).To(BeNil())
			Expect(account.Balance).To(Equal(amount))
		})
	})

	Context("Testing Invalid Amount Deposit", func() {
		account := bank.Account{}
		It("should return Error: Invalid amount to deposit", func() {
			var amount float64 = -100
			err := account.Deposit(amount)
			Expect(err).To(HaveOccurred())
			Expect(account.Balance).To(Equal(float64(0)))
		})
	})

	Context("Testing Valid Amount Withdraw", func() {
		account := bank.Account{}
		It("should return nil", func() {
			var amount float64 = 100.5
			err := account.Deposit(amount)
			Expect(err).To(BeNil())
			Expect(account.Balance).To(Equal(amount))
			var withdraw float64 = 50.5
			err = account.Withdraw(withdraw)
			Expect(err).To(BeNil())
			Expect(account.Balance).To(Equal(amount - withdraw))
		})
	})

	Context("Testing Invalid Amount Withdraw", func() {
		account := bank.Account{}
		It("should return Error: Invalid amount to withdraw", func() {
			var amount float64 = 100.5
			err := account.Deposit(amount)
			Expect(err).To(BeNil())
			Expect(account.Balance).To(Equal(amount))
			var withdraw float64 = -50.5
			err = account.Withdraw(withdraw)
			Expect(err).To(HaveOccurred())
			Expect(account.Balance).To(Equal(amount))
		})
	})

	Context("Testing Insufficient Funds Withdraw", func() {
		account := bank.Account{}
		It("should return Error: insufficient funds", func() {
			var amount float64 = 100.5
			err := account.Deposit(amount)
			Expect(err).To(BeNil())
			Expect(account.Balance).To(Equal(amount))
			var withdraw float64 = 101
			err = account.Withdraw(withdraw)
			Expect(err).To(HaveOccurred())
			Expect(account.Balance).To(Equal(amount))
		})
	})

	Context("Testing Get Balance", func() {
		account := bank.Account{}
		It("should return deposit amount", func() {
			var amount float64 = 100.5
			err := account.Deposit(amount)
			Expect(err).To(BeNil())
			Expect(account.Balance).To(Equal(amount))
		})
	})
})
