package lib

const (
	EmptyValue               = ""
	RefundFeeByParticipant   = "refund_fee_by_participant"
	SettlementTypeDefault    = "default"
	SettlementTypeCustom     = "custom"
	DivisionTypeProportional = "proportional"
)

type SettlementType string
type FeeDivisionType string
type RefundFeeCalculusType string
type ZeroFeeWhenChargebackNetAmountIsZero string

type Spec struct {
	Name         string
	WillValidate []any
}

type Dictionary map[string]Spec

var Specs = make(Dictionary, 4)

func init() {
	var specNames = map[string]string{
		"settlement_type":                             "SettlementType",
		"fee_division_type":                           "FeeDivisionType",
		"refund_fee_calculus_type":                    "RefundFeeCalculusType",
		"zero_fee_when_chargeback_net_amount_is_zero": "ZeroFeeWhenChargebackNetAmountIsZero",
	}

	var validateTypes = map[string][]any{
		"settlement_type":                             {SettlementTypeDefault, SettlementTypeCustom},
		"fee_division_type":                           {DivisionTypeProportional, EmptyValue},
		"refund_fee_calculus_type":                    {RefundFeeByParticipant, EmptyValue},
		"zero_fee_when_chargeback_net_amount_is_zero": {true, false},
	}

	for i, name := range specNames {
		Specs[i] = Spec{Name: name, WillValidate: validateTypes[i]}
	}
}
