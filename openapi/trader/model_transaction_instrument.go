/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trader

import (
	"encoding/json"
	"fmt"
)

// TransactionInstrument - struct for TransactionInstrument
type TransactionInstrument struct {
	CollectiveInvestment *CollectiveInvestment
	Currency *Currency
	Forex *Forex
	Future *Future
	Index *Index
	Product *Product
	TransactionCashEquivalent *TransactionCashEquivalent
	TransactionEquity *TransactionEquity
	TransactionFixedIncome *TransactionFixedIncome
	TransactionMutualFund *TransactionMutualFund
	TransactionOption *TransactionOption
}

// CollectiveInvestmentAsTransactionInstrument is a convenience function that returns CollectiveInvestment wrapped in TransactionInstrument
func CollectiveInvestmentAsTransactionInstrument(v *CollectiveInvestment) TransactionInstrument {
	return TransactionInstrument{
		CollectiveInvestment: v,
	}
}

// CurrencyAsTransactionInstrument is a convenience function that returns Currency wrapped in TransactionInstrument
func CurrencyAsTransactionInstrument(v *Currency) TransactionInstrument {
	return TransactionInstrument{
		Currency: v,
	}
}

// ForexAsTransactionInstrument is a convenience function that returns Forex wrapped in TransactionInstrument
func ForexAsTransactionInstrument(v *Forex) TransactionInstrument {
	return TransactionInstrument{
		Forex: v,
	}
}

// FutureAsTransactionInstrument is a convenience function that returns Future wrapped in TransactionInstrument
func FutureAsTransactionInstrument(v *Future) TransactionInstrument {
	return TransactionInstrument{
		Future: v,
	}
}

// IndexAsTransactionInstrument is a convenience function that returns Index wrapped in TransactionInstrument
func IndexAsTransactionInstrument(v *Index) TransactionInstrument {
	return TransactionInstrument{
		Index: v,
	}
}

// ProductAsTransactionInstrument is a convenience function that returns Product wrapped in TransactionInstrument
func ProductAsTransactionInstrument(v *Product) TransactionInstrument {
	return TransactionInstrument{
		Product: v,
	}
}

// TransactionCashEquivalentAsTransactionInstrument is a convenience function that returns TransactionCashEquivalent wrapped in TransactionInstrument
func TransactionCashEquivalentAsTransactionInstrument(v *TransactionCashEquivalent) TransactionInstrument {
	return TransactionInstrument{
		TransactionCashEquivalent: v,
	}
}

// TransactionEquityAsTransactionInstrument is a convenience function that returns TransactionEquity wrapped in TransactionInstrument
func TransactionEquityAsTransactionInstrument(v *TransactionEquity) TransactionInstrument {
	return TransactionInstrument{
		TransactionEquity: v,
	}
}

// TransactionFixedIncomeAsTransactionInstrument is a convenience function that returns TransactionFixedIncome wrapped in TransactionInstrument
func TransactionFixedIncomeAsTransactionInstrument(v *TransactionFixedIncome) TransactionInstrument {
	return TransactionInstrument{
		TransactionFixedIncome: v,
	}
}

// TransactionMutualFundAsTransactionInstrument is a convenience function that returns TransactionMutualFund wrapped in TransactionInstrument
func TransactionMutualFundAsTransactionInstrument(v *TransactionMutualFund) TransactionInstrument {
	return TransactionInstrument{
		TransactionMutualFund: v,
	}
}

// TransactionOptionAsTransactionInstrument is a convenience function that returns TransactionOption wrapped in TransactionInstrument
func TransactionOptionAsTransactionInstrument(v *TransactionOption) TransactionInstrument {
	return TransactionInstrument{
		TransactionOption: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionInstrument) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CollectiveInvestment
	err = newStrictDecoder(data).Decode(&dst.CollectiveInvestment)
	if err == nil {
		jsonCollectiveInvestment, _ := json.Marshal(dst.CollectiveInvestment)
		if string(jsonCollectiveInvestment) == "{}" { // empty struct
			dst.CollectiveInvestment = nil
		} else {
			match++
		}
	} else {
		dst.CollectiveInvestment = nil
	}

	// try to unmarshal data into Currency
	err = newStrictDecoder(data).Decode(&dst.Currency)
	if err == nil {
		jsonCurrency, _ := json.Marshal(dst.Currency)
		if string(jsonCurrency) == "{}" { // empty struct
			dst.Currency = nil
		} else {
			match++
		}
	} else {
		dst.Currency = nil
	}

	// try to unmarshal data into Forex
	err = newStrictDecoder(data).Decode(&dst.Forex)
	if err == nil {
		jsonForex, _ := json.Marshal(dst.Forex)
		if string(jsonForex) == "{}" { // empty struct
			dst.Forex = nil
		} else {
			match++
		}
	} else {
		dst.Forex = nil
	}

	// try to unmarshal data into Future
	err = newStrictDecoder(data).Decode(&dst.Future)
	if err == nil {
		jsonFuture, _ := json.Marshal(dst.Future)
		if string(jsonFuture) == "{}" { // empty struct
			dst.Future = nil
		} else {
			match++
		}
	} else {
		dst.Future = nil
	}

	// try to unmarshal data into Index
	err = newStrictDecoder(data).Decode(&dst.Index)
	if err == nil {
		jsonIndex, _ := json.Marshal(dst.Index)
		if string(jsonIndex) == "{}" { // empty struct
			dst.Index = nil
		} else {
			match++
		}
	} else {
		dst.Index = nil
	}

	// try to unmarshal data into Product
	err = newStrictDecoder(data).Decode(&dst.Product)
	if err == nil {
		jsonProduct, _ := json.Marshal(dst.Product)
		if string(jsonProduct) == "{}" { // empty struct
			dst.Product = nil
		} else {
			match++
		}
	} else {
		dst.Product = nil
	}

	// try to unmarshal data into TransactionCashEquivalent
	err = newStrictDecoder(data).Decode(&dst.TransactionCashEquivalent)
	if err == nil {
		jsonTransactionCashEquivalent, _ := json.Marshal(dst.TransactionCashEquivalent)
		if string(jsonTransactionCashEquivalent) == "{}" { // empty struct
			dst.TransactionCashEquivalent = nil
		} else {
			match++
		}
	} else {
		dst.TransactionCashEquivalent = nil
	}

	// try to unmarshal data into TransactionEquity
	err = newStrictDecoder(data).Decode(&dst.TransactionEquity)
	if err == nil {
		jsonTransactionEquity, _ := json.Marshal(dst.TransactionEquity)
		if string(jsonTransactionEquity) == "{}" { // empty struct
			dst.TransactionEquity = nil
		} else {
			match++
		}
	} else {
		dst.TransactionEquity = nil
	}

	// try to unmarshal data into TransactionFixedIncome
	err = newStrictDecoder(data).Decode(&dst.TransactionFixedIncome)
	if err == nil {
		jsonTransactionFixedIncome, _ := json.Marshal(dst.TransactionFixedIncome)
		if string(jsonTransactionFixedIncome) == "{}" { // empty struct
			dst.TransactionFixedIncome = nil
		} else {
			match++
		}
	} else {
		dst.TransactionFixedIncome = nil
	}

	// try to unmarshal data into TransactionMutualFund
	err = newStrictDecoder(data).Decode(&dst.TransactionMutualFund)
	if err == nil {
		jsonTransactionMutualFund, _ := json.Marshal(dst.TransactionMutualFund)
		if string(jsonTransactionMutualFund) == "{}" { // empty struct
			dst.TransactionMutualFund = nil
		} else {
			match++
		}
	} else {
		dst.TransactionMutualFund = nil
	}

	// try to unmarshal data into TransactionOption
	err = newStrictDecoder(data).Decode(&dst.TransactionOption)
	if err == nil {
		jsonTransactionOption, _ := json.Marshal(dst.TransactionOption)
		if string(jsonTransactionOption) == "{}" { // empty struct
			dst.TransactionOption = nil
		} else {
			match++
		}
	} else {
		dst.TransactionOption = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CollectiveInvestment = nil
		dst.Currency = nil
		dst.Forex = nil
		dst.Future = nil
		dst.Index = nil
		dst.Product = nil
		dst.TransactionCashEquivalent = nil
		dst.TransactionEquity = nil
		dst.TransactionFixedIncome = nil
		dst.TransactionMutualFund = nil
		dst.TransactionOption = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransactionInstrument)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransactionInstrument)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionInstrument) MarshalJSON() ([]byte, error) {
	if src.CollectiveInvestment != nil {
		return json.Marshal(&src.CollectiveInvestment)
	}

	if src.Currency != nil {
		return json.Marshal(&src.Currency)
	}

	if src.Forex != nil {
		return json.Marshal(&src.Forex)
	}

	if src.Future != nil {
		return json.Marshal(&src.Future)
	}

	if src.Index != nil {
		return json.Marshal(&src.Index)
	}

	if src.Product != nil {
		return json.Marshal(&src.Product)
	}

	if src.TransactionCashEquivalent != nil {
		return json.Marshal(&src.TransactionCashEquivalent)
	}

	if src.TransactionEquity != nil {
		return json.Marshal(&src.TransactionEquity)
	}

	if src.TransactionFixedIncome != nil {
		return json.Marshal(&src.TransactionFixedIncome)
	}

	if src.TransactionMutualFund != nil {
		return json.Marshal(&src.TransactionMutualFund)
	}

	if src.TransactionOption != nil {
		return json.Marshal(&src.TransactionOption)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionInstrument) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CollectiveInvestment != nil {
		return obj.CollectiveInvestment
	}

	if obj.Currency != nil {
		return obj.Currency
	}

	if obj.Forex != nil {
		return obj.Forex
	}

	if obj.Future != nil {
		return obj.Future
	}

	if obj.Index != nil {
		return obj.Index
	}

	if obj.Product != nil {
		return obj.Product
	}

	if obj.TransactionCashEquivalent != nil {
		return obj.TransactionCashEquivalent
	}

	if obj.TransactionEquity != nil {
		return obj.TransactionEquity
	}

	if obj.TransactionFixedIncome != nil {
		return obj.TransactionFixedIncome
	}

	if obj.TransactionMutualFund != nil {
		return obj.TransactionMutualFund
	}

	if obj.TransactionOption != nil {
		return obj.TransactionOption
	}

	// all schemas are nil
	return nil
}

type NullableTransactionInstrument struct {
	value *TransactionInstrument
	isSet bool
}

func (v NullableTransactionInstrument) Get() *TransactionInstrument {
	return v.value
}

func (v *NullableTransactionInstrument) Set(val *TransactionInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInstrument(val *TransactionInstrument) *NullableTransactionInstrument {
	return &NullableTransactionInstrument{value: val, isSet: true}
}

func (v NullableTransactionInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


