package accounts

import (
  "bank/holders"
)

type SavingsAccount struct {
  AccountHolder holders.Holder
  AccountCode, AccountNumber, Operations int
  accountBalance float64
}

func (a *SavingsAccount) Withdraw(withdrawalAmount float64) (string, float64){
  checkAmount := withdrawalAmount > 0 && withdrawalAmount <= a.accountBalance
  if checkAmount {
    a.accountBalance -= withdrawalAmount
    return "Success.", a.accountBalance
  } else {
    return "Withdraw failed.", a.accountBalance
  }
}

func (a *SavingsAccount) Deposit(depositAmount float64) (string, float64) {
  checkAmount := depositAmount > 0
  if checkAmount {
    a.accountBalance += depositAmount
    return "Success.", a.accountBalance
  } else {
    return "Deposit failed.", a.accountBalance
  }
}

func (a *SavingsAccount) Transfer(transferAmount float64, destinationAccount *SavingsAccount) (string, float64){
  if transferAmount > 0 && transferAmount < a.accountBalance {
    a.accountBalance -= transferAmount
    destinationAccount.accountBalance += transferAmount
    return "Sucess.", destinationAccount.accountBalance
  } else {
    return "Failed.", destinationAccount.accountBalance
  }
}

func (a *SavingsAccount) GetBalance() float64{
  return a.accountBalance
}
