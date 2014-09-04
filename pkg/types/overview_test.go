package types

import "testing"

var (
	owner = User{
		Id:       100,
		FullName: "Owner",
	}

	par1 = User{
		Id:       200,
		FullName: "Par1",
	}

	par2 = User{
		Id:       201,
		FullName: "Par2",
	}

	pay1 = Payment{
		FromId: 200,
		From:   par1,
		ToId:   201,
		To:     par2,
		Price: Price{
			Value:    10,
			Currency: "EUR",
		},
	}

	pay2 = Payment{
		FromId: 201,
		From:   par2,
		ToId:   200,
		To:     par1,
		Price: Price{
			Value:    5,
			Currency: "EUR",
		},
	}

	pay3 = Payment{
		FromId: 201,
		From:   par2,
		ToId:   200,
		To:     par1,
		Price: Price{
			Value:    2,
			Currency: "EUR",
		},
	}

	pay4 = Payment{
		FromId: 100,
		From:   owner,
		ToId:   200,
		To:     par1,
		Price: Price{
			Value:    4,
			Currency: "EUR",
		},
	}
)

func TestComputeUserAccounts(t *testing.T) {
	overview := Overview{
		OwnerId:      owner.Id,
		Owner:        owner,
		Participants: Users{par1, par2},
		Payments:     Payments{pay1, pay2, pay3, pay4},
	}

	overview.ComputeUserAccounts()

	for _, userAccount := range overview.UserAccounts {
		switch userAccount.User.Id {
		case 100:
			if userAccount.Balances[200].Value != 4 {
				t.Fatalf("UserAccount %d has wrong balance with 200! Expected 4 but got %d", userAccount.User.Id, userAccount.Balances[200].Value)
			}
		case 200:
			if userAccount.Balances[100].Value != -4 {
				t.Fatalf("UserAccount %d has wrong balance with 100! Expected -4 but got %d", userAccount.User.Id, userAccount.Balances[200].Value)
			}
			if userAccount.Balances[201].Value != 3 {
				t.Fatalf("UserAccount %d has wrong balance with 201! Expected 3 but got %d", userAccount.User.Id, userAccount.Balances[200].Value)
			}
		case 201:
			if userAccount.Balances[200].Value != -3 {
				t.Fatalf("UserAccount %d has wrong balance with 200! Expected -3 but got %d", userAccount.User.Id, userAccount.Balances[200].Value)
			}
		}

		/*
			log.Printf("UserAccount[%d]\n", userAccount.User.Id)
			for _, balance := range userAccount.Balances {
				log.Printf("Balance with [%d] %d %s\n", balance.User.Id, balance.Price.Value, balance.Price.Currency)

				for _, debit := range balance.Debits {
					log.Printf("Debit %d %s\n", debit.Value, debit.Currency)
				}

				for _, credit := range balance.Credits {
					log.Printf("Credit %d %s\n", credit.Value, credit.Currency)
				}
			}
		*/
	}
}
