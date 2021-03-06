// Code generated by "stringer -type=dataType"; DO NOT EDIT.

package tds

import "strconv"

const _dataType_name = "UInt4ImageTextExtendedTypeVarbinaryIntNVarcharBinaryCharTinyintDateBitTimeSmallintDecimalIntSmalldatetimeRealMoneyDatetimeFloatNumericUTinyintUSmallintUIntUBigintUIntNBigDateTimeNDecimalNNumericNFloatNMoneyNDatetimeNSmallmoneyDateNUnicharTimeNUnivarcharTextLocatorImageLocatorUnitextLocatorUnitextLongCharBigdatetimeNBigtimeNBigdatetimeBigtimeBigintLongBinary"

var _dataType_map = map[dataType]string{
	29:  _dataType_name[0:5],
	34:  _dataType_name[5:10],
	35:  _dataType_name[10:14],
	36:  _dataType_name[14:26],
	37:  _dataType_name[26:35],
	38:  _dataType_name[35:39],
	39:  _dataType_name[39:46],
	45:  _dataType_name[46:52],
	47:  _dataType_name[52:56],
	48:  _dataType_name[56:63],
	49:  _dataType_name[63:67],
	50:  _dataType_name[67:70],
	51:  _dataType_name[70:74],
	52:  _dataType_name[74:82],
	55:  _dataType_name[82:89],
	56:  _dataType_name[89:92],
	58:  _dataType_name[92:105],
	59:  _dataType_name[105:109],
	60:  _dataType_name[109:114],
	61:  _dataType_name[114:122],
	62:  _dataType_name[122:127],
	63:  _dataType_name[127:134],
	64:  _dataType_name[134:142],
	65:  _dataType_name[142:151],
	66:  _dataType_name[151:155],
	67:  _dataType_name[155:162],
	68:  _dataType_name[162:167],
	80:  _dataType_name[167:179],
	106: _dataType_name[179:187],
	108: _dataType_name[187:195],
	109: _dataType_name[195:201],
	110: _dataType_name[201:207],
	111: _dataType_name[207:216],
	122: _dataType_name[216:226],
	123: _dataType_name[226:231],
	135: _dataType_name[231:238],
	147: _dataType_name[238:243],
	155: _dataType_name[243:253],
	169: _dataType_name[253:264],
	170: _dataType_name[264:276],
	171: _dataType_name[276:290],
	174: _dataType_name[290:297],
	175: _dataType_name[297:305],
	187: _dataType_name[305:317],
	188: _dataType_name[317:325],
	189: _dataType_name[325:336],
	190: _dataType_name[336:343],
	191: _dataType_name[343:349],
	225: _dataType_name[349:359],
}

func (i dataType) String() string {
	if str, ok := _dataType_map[i]; ok {
		return str
	}
	return "dataType(" + strconv.FormatInt(int64(i), 10) + ")"
}
