package types

type PaymentMethodKind int8

func ParsePaymentMethodKind(value int) PaymentMethodKind {
	switch value {
	case 1:
		return PaymentMethodKindCash
	case 2:
		return PaymentMethodKindCredit
	case 4:
		return PaymentMethodKindDebit
	case 8:
		return PaymentMethodKindCheque
	case 16:
		return PaymentMethodKindBankTransfer
	case 32:
		return PaymentMethodKindOther
	default:
		return PaymentMethodKindUnknown
	}
}

func (self PaymentMethodKind) String() string {
	switch self {
	case PaymentMethodKindCash:
		return "cash"
	case PaymentMethodKindCredit:
		return "credit"
	case PaymentMethodKindDebit:
		return "debit"
	case PaymentMethodKindCheque:
		return "cheque"
	case PaymentMethodKindBankTransfer:
		return "bank_transfer"
	case PaymentMethodKindOther:
		return "other"
	default:
		return "unknown"
	}
}

var (
	PaymentMethodKindUnknown      PaymentMethodKind = 0
	PaymentMethodKindCash         PaymentMethodKind = 1
	PaymentMethodKindCredit       PaymentMethodKind = 2
	PaymentMethodKindDebit        PaymentMethodKind = 4
	PaymentMethodKindCheque       PaymentMethodKind = 8
	PaymentMethodKindBankTransfer PaymentMethodKind = 16
	PaymentMethodKindOther        PaymentMethodKind = 32
)

type PaymentMethod struct {
	Kind          PaymentMethodKind `json:"kind" datastore:"kind"`
	KindFormatted string            `json:"kind_formatted" datastore:"-"`
	Comment       string            `json:"comment" datastore:"comment"`
}

func (self *PaymentMethod) SetFormattedValues() {
	self.KindFormatted = self.Kind.String()
}
