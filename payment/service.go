package payment

import (
	"bwa-golang/user"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/iris"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	//Initiate client for Midtrans CoreAPI
	var clientVar = coreapi.Client{}
	clientVar.New("YOUR-SERVER-KEY", midtrans.Sandbox)

	//Initiate client for Midtrans Snap
	var snapClient = snap.Client{}
	snapClient.New("YOUR-SERVER-KEY", midtrans.Sandbox)

	//Initiate client for Iris disbursement
	var irisClient = iris.Client{}
	irisClient.New("IRIS-API-KEY", midtrans.Sandbox)

	custAddress := &midtrans.CustomerDetails{
		FName: user.Name,
		Email: user.Email,
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: custAddress,
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		panic(errSnap)
	}

	return response.RedirectURL, nil

}
