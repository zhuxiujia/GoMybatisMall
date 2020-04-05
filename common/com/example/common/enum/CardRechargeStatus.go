package enum

type CardRechargeStatus int

const (
	CardRechargeStatus_Pay_Fail CardRechargeStatus = -1
	CardRechargeStatus_Created  CardRechargeStatus = 1
	CardRechargeStatus_Paying   CardRechargeStatus = 2
	CardRechargeStatus_Payed    CardRechargeStatus = 3
)
