package seedgen

import "errors"

//ErrConvertingDiceRolls dice rolls
var ErrConvertingDiceRolls = errors.New("error converting dice rolls to binary")

//ErrConvertingTrimmedString binary
var ErrConvertingTrimmedString = errors.New("error converting trimmed string to binary")
