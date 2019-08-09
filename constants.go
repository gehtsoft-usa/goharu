package goharu

//True is true value for libharu
const True int = 1

//False is false value for libharu
const False int = 0

//Ok is no error value for libharu
const Ok int = 0

//Noerror no error value for libharu
const Noerror int = 0

//DefFont is default font
const DefFont string = "Helvetica"

//DefWordSpace is default word space
const DefWordSpace int = 0

//DefCharSpace is default character space
const DefCharSpace int = 0

//DefFontSize is default font size
const DefFontSize int = 10

//DefHscaling is default horizontal scaling in percents
const DefHscaling int = 100

//DefLeading is default leading value
const DefLeading int = 0

//DefRise is default text rise value
const DefRise int = 0

//DefLineWidth default line width
const DefLineWidth int = 1

//DefMiterlimit is default miter limit
const DefMiterlimit int = 10

//DefFlatness is default flatness level
const DefFlatness int = 1

//DefPageNum is default page numeration style
const DefPageNum int = 1

//BsDefWidth is default line width
const BsDefWidth int = 1

//DefPageWidth is default page width
const DefPageWidth float32 = 595.276

//DefPageHeight is default page height
const DefPageHeight float32 = 841.89

//CompNone is flag indicating that document should be saved with no compression
const CompNone int = 0x00

//CompText is flag indicating that document should be saved with text only compression
const CompText int = 0x01

//CompImage is flag indicating that document should be saved with images only compression
const CompImage int = 0x02

//CompMetadata is flag indicating that document should be saved with metadata only compression
const CompMetadata int = 0x04

//CompAll is flag indicating that document should be saved with full compression
const CompAll int = 0x0F

/*----------------------------------------------------------------------------*/
/*----- permission flags (only Revision 2 is supported)-----------------------*/

//EnableRead is permission that enables document reading
const EnableRead int = 0

//EnablePrint is permission that enables document printing
const EnablePrint int = 4

//EnableEditAll is permission that enables document modification
const EnableEditAll int = 8

//EnableCopy is is permission that enables copying information from the document
const EnableCopy int = 16

//EnableEdit is is permission that enables document modification
const EnableEdit int = 32

/*----------------------------------------------------------------------------*/
/*------ viewer preferences definitions --------------------------------------*/

//HideToolbar is a viewer preference value
const HideToolbar int = 1

//HideMenubar is a viewer preference value
const HideMenubar int = 2

//HideWindowUI is a viewer preference value
const HideWindowUI int = 4

//FitWindow is a viewer preference value
const FitWindow int = 8

//CenterWindow is a viewer preference value
const CenterWindow int = 16

//PrintScalingNone is a viewer preference value
const PrintScalingNone int = 32

/*---------------------------------------------------------------------------*/
/*------ limitation of object implementation (PDF1.4) -----------------------*/

//LimitMaxInt is maximum integer value for PDF 1.4
const LimitMaxInt int = 2147483647

//LimitMinInt is minimum integer value for PDF 1.4
const LimitMinInt int = -2147483647

//LimitMaxReal is maximum real value for PDF 1.4
const LimitMaxReal float32 = 3.4e38 // per PDF 1.7 spec, Annex C, old value  32767

//LimitMinReal is minimum real value for PDF 1.4
const LimitMinReal float32 = -3.4e38 // per PDF 1.7 spec, Annex C, old value -32767

//LimitMaxStringLen is maximum text length for PDF 1.4
const LimitMaxStringLen int = 2147483646 // per PDF 1.7 spec, limit 32767 is for strings in content stream and no limit in other cases => open the limit to max Integer, old value 65535

//LimitMaxNameLen is maximum object name length for PDF 1.4
const LimitMaxNameLen int = 127

//LimitMaxArray the maximum number of object in PDF file
const LimitMaxArray int = 8388607 // per PDF 1.7 spec, "Maximum number of indirect objects in a PDF file" is 8388607, old value 8191
//LimitMaxDictElement the maxmum dictionary elements
const LimitMaxDictElement int = 8388607 // per PDF 1.7 spec, "Maximum number of indirect objects in a PDF file" is 8388607, old value 4095
//LimitMaxXrefElement the maximum XRef elements
const LimitMaxXrefElement int = 8388607

//LimitMaxGstate the maximum of states in stack
const LimitMaxGstate int = 28

//LimitMaxDeviceN ???
const LimitMaxDeviceN int = 8

//LimitMaxDeviceNV15 ???
const LimitMaxDeviceNV15 int = 32

//LimitMaxCid ???
const LimitMaxCid int = 65535

//MaxGenerationNum ???
const MaxGenerationNum int = 65535

//MinPageHeight is the mimimum page height
const MinPageHeight int = 3

//MinPageWidth the mimimum page width
const MinPageWidth int = 3

//MaxPageHeight is the maximum page height
const MaxPageHeight int = 14400

//MaxPageWidth is the maximum page width
const MaxPageWidth int = 14400

//MinMagnificationFactor is the minimum magnification factor
const MinMagnificationFactor int = 8

//MaxMagnificationFactor the maximum magnification factor
const MaxMagnificationFactor int = 3200

/*---------------------------------------------------------------------------*/
/*------ limitation of various properties -----------------------------------*/

//MinPageSize is the minimum page size
const MinPageSize int = 3

//MaxPageSize is the maximum page size
const MaxPageSize int = 14400

//MinHorizontalscaling is the minimum horizontal scaling
const MinHorizontalscaling int = 10

//MaxHorizontalscaling is the maximum horizontal scaling
const MaxHorizontalscaling int = 300

//MinWordspace is the minimum word space
const MinWordspace int = -30

//MaxWordspace is the maximum word space
const MaxWordspace int = 300

//MinCharspace is the minimum character space
const MinCharspace int = -30

//MaxCharspace is the maximum character space
const MaxCharspace int = 300

//MaxFontSize is the maximum font size
const MaxFontSize int = 600

//MaxZoomSize is the minimum font size
const MaxZoomSize int = 10

//MaxLeading is the maximum leading
const MaxLeading int = 300

//MaxLineWidth is maximum line width
const MaxLineWidth int = 100

//MaxDashPattern is maximum hash pattern length
const MaxDashPattern int = 100

//MaxJwwNum is ???
const MaxJwwNum int = 128

/*----------------------------------------------------------------------------*/
/*----- country code definition ----------------------------------------------*/

//CountryAf is country code for /* AFGHANISTAN */
const CountryAf string = "AF" /* AFGHANISTAN */
//CountryAl is country code for /* ALBANIA */
const CountryAl string = "AL" /* ALBANIA */
//CountryDz is country code for /* ALGERIA */
const CountryDz string = "DZ" /* ALGERIA */
//CountryAs is country code for /* AMERICAN SAMOA */
const CountryAs string = "AS" /* AMERICAN SAMOA */
//CountryAd is country code for /* ANDORRA */
const CountryAd string = "AD" /* ANDORRA */
//CountryAo is country code for /* ANGOLA */
const CountryAo string = "AO" /* ANGOLA */
//CountryAi is country code for /* ANGUILLA */
const CountryAi string = "AI" /* ANGUILLA */
//CountryAq is country code for /* ANTARCTICA */
const CountryAq string = "AQ" /* ANTARCTICA */
//CountryAg is country code for /* ANTIGUA AND BARBUDA */
const CountryAg string = "AG" /* ANTIGUA AND BARBUDA */
//CountryAr is country code for /* ARGENTINA */
const CountryAr string = "AR" /* ARGENTINA */
//CountryAm is country code for /* ARMENIA */
const CountryAm string = "AM" /* ARMENIA */
//CountryAw is country code for /* ARUBA */
const CountryAw string = "AW" /* ARUBA */
//CountryAu is country code for /* AUSTRALIA */
const CountryAu string = "AU" /* AUSTRALIA */
//CountryAt is country code for /* AUSTRIA */
const CountryAt string = "AT" /* AUSTRIA */
//CountryAz is country code for /* AZERBAIJAN */
const CountryAz string = "AZ" /* AZERBAIJAN */
//CountryBs is country code for /* BAHAMAS */
const CountryBs string = "BS" /* BAHAMAS */
//CountryBh is country code for /* BAHRAIN */
const CountryBh string = "BH" /* BAHRAIN */
//CountryBd is country code for /* BANGLADESH */
const CountryBd string = "BD" /* BANGLADESH */
//CountryBb is country code for /* BARBADOS */
const CountryBb string = "BB" /* BARBADOS */
//CountryBy is country code for /* BELARUS */
const CountryBy string = "BY" /* BELARUS */
//CountryBe is country code for /* BELGIUM */
const CountryBe string = "BE" /* BELGIUM */
//CountryBz is country code for /* BELIZE */
const CountryBz string = "BZ" /* BELIZE */
//CountryBj is country code for /* BENIN */
const CountryBj string = "BJ" /* BENIN */
//CountryBm is country code for /* BERMUDA */
const CountryBm string = "BM" /* BERMUDA */
//CountryBt is country code for /* BHUTAN */
const CountryBt string = "BT" /* BHUTAN */
//CountryBo is country code for /* BOLIVIA */
const CountryBo string = "BO" /* BOLIVIA */
//CountryBa is country code for /* BOSNIA AND HERZEGOVINA */
const CountryBa string = "BA" /* BOSNIA AND HERZEGOVINA */
//CountryBw is country code for /* BOTSWANA */
const CountryBw string = "BW" /* BOTSWANA */
//CountryBv is country code for /* BOUVET ISLAND */
const CountryBv string = "BV" /* BOUVET ISLAND */
//CountryBr is country code for /* BRAZIL */
const CountryBr string = "BR" /* BRAZIL */
//CountryIo is country code for /* BRITISH INDIAN OCEAN TERRITORY */
const CountryIo string = "IO" /* BRITISH INDIAN OCEAN TERRITORY */
//CountryBn is country code for /* BRUNEI DARUSSALAM */
const CountryBn string = "BN" /* BRUNEI DARUSSALAM */
//CountryBg is country code for /* BULGARIA */
const CountryBg string = "BG" /* BULGARIA */
//CountryBf is country code for /* BURKINA FASO */
const CountryBf string = "BF" /* BURKINA FASO */
//CountryBi is country code for /* BURUNDI */
const CountryBi string = "BI" /* BURUNDI */
//CountryKh is country code for /* CAMBODIA */
const CountryKh string = "KH" /* CAMBODIA */
//CountryCm is country code for /* CAMEROON */
const CountryCm string = "CM" /* CAMEROON */
//CountryCa is country code for /* CANADA */
const CountryCa string = "CA" /* CANADA */
//CountryCv is country code for /* CAPE VERDE */
const CountryCv string = "CV" /* CAPE VERDE */
//CountryKy is country code for /* CAYMAN ISLANDS */
const CountryKy string = "KY" /* CAYMAN ISLANDS */
//CountryCf is country code for /* CENTRAL AFRICAN REPUBLIC */
const CountryCf string = "CF" /* CENTRAL AFRICAN REPUBLIC */
//CountryTd is country code for /* CHAD */
const CountryTd string = "TD" /* CHAD */
//CountryCl is country code for /* CHILE */
const CountryCl string = "CL" /* CHILE */
//CountryCn is country code for /* CHINA */
const CountryCn string = "CN" /* CHINA */
//CountryCx is country code for /* CHRISTMAS ISLAND */
const CountryCx string = "CX" /* CHRISTMAS ISLAND */
//CountryCc is country code for /* COCOS (KEELING) ISLANDS */
const CountryCc string = "CC" /* COCOS (KEELING) ISLANDS */
//CountryCo is country code for /* COLOMBIA */
const CountryCo string = "CO" /* COLOMBIA */
//CountryKm is country code for /* COMOROS */
const CountryKm string = "KM" /* COMOROS */
//CountryCg is country code for /* CONGO */
const CountryCg string = "CG" /* CONGO */
//CountryCk is country code for /* COOK ISLANDS */
const CountryCk string = "CK" /* COOK ISLANDS */
//CountryCr is country code for /* COSTA RICA */
const CountryCr string = "CR" /* COSTA RICA */
//CountryCi is country code for /* COTE D'IVOIRE */
const CountryCi string = "CI" /* COTE D'IVOIRE */
//CountryHr is country code for /* CROATIA (local name: Hrvatska) */
const CountryHr string = "HR" /* CROATIA (local name: Hrvatska) */
//CountryCu is country code for /* CUBA */
const CountryCu string = "CU" /* CUBA */
//CountryCy is country code for /* CYPRUS */
const CountryCy string = "CY" /* CYPRUS */
//CountryCz is country code for /* CZECH REPUBLIC */
const CountryCz string = "CZ" /* CZECH REPUBLIC */
//CountryDk is country code for /* DENMARK */
const CountryDk string = "DK" /* DENMARK */
//CountryDj is country code for /* DJIBOUTI */
const CountryDj string = "DJ" /* DJIBOUTI */
//CountryDm is country code for /* DOMINICA */
const CountryDm string = "DM" /* DOMINICA */
//CountryDo is country code for /* DOMINICAN REPUBLIC */
const CountryDo string = "DO" /* DOMINICAN REPUBLIC */
//CountryTp is country code for /* EAST TIMOR */
const CountryTp string = "TP" /* EAST TIMOR */
//CountryEc is country code for /* ECUADOR */
const CountryEc string = "EC" /* ECUADOR */
//CountryEg is country code for /* EGYPT */
const CountryEg string = "EG" /* EGYPT */
//CountrySv is country code for /* EL SALVADOR */
const CountrySv string = "SV" /* EL SALVADOR */
//CountryGq is country code for /* EQUATORIAL GUINEA */
const CountryGq string = "GQ" /* EQUATORIAL GUINEA */
//CountryEr is country code for /* ERITREA */
const CountryEr string = "ER" /* ERITREA */
//CountryEe is country code for /* ESTONIA */
const CountryEe string = "EE" /* ESTONIA */
//CountryEt is country code for /* ETHIOPIA */
const CountryEt string = "ET" /* ETHIOPIA */
//CountryFk is country code for /* FALKLAND ISLANDS (MALVINAS) */
const CountryFk string = "FK" /* FALKLAND ISLANDS (MALVINAS) */
//CountryFo is country code for /* FAROE ISLANDS */
const CountryFo string = "FO" /* FAROE ISLANDS */
//CountryFj is country code for /* FIJI */
const CountryFj string = "FJ" /* FIJI */
//CountryFi is country code for /* FINLAND */
const CountryFi string = "FI" /* FINLAND */
//CountryFr is country code for /* FRANCE */
const CountryFr string = "FR" /* FRANCE */
//CountryFx is country code for /* FRANCE, METROPOLITAN */
const CountryFx string = "FX" /* FRANCE, METROPOLITAN */
//CountryGf is country code for /* FRENCH GUIANA */
const CountryGf string = "GF" /* FRENCH GUIANA */
//CountryPf is country code for /* FRENCH POLYNESIA */
const CountryPf string = "PF" /* FRENCH POLYNESIA */
//CountryTf is country code for /* FRENCH SOUTHERN TERRITORIES */
const CountryTf string = "TF" /* FRENCH SOUTHERN TERRITORIES */
//CountryGa is country code for /* GABON */
const CountryGa string = "GA" /* GABON */
//CountryGm is country code for /* GAMBIA */
const CountryGm string = "GM" /* GAMBIA */
//CountryGe is country code for /* GEORGIA */
const CountryGe string = "GE" /* GEORGIA */
//CountryDe is country code for /* GERMANY */
const CountryDe string = "DE" /* GERMANY */
//CountryGh is country code for /* GHANA */
const CountryGh string = "GH" /* GHANA */
//CountryGi is country code for /* GIBRALTAR */
const CountryGi string = "GI" /* GIBRALTAR */
//CountryGr is country code for /* GREECE */
const CountryGr string = "GR" /* GREECE */
//CountryGl is country code for /* GREENLAND */
const CountryGl string = "GL" /* GREENLAND */
//CountryGd is country code for /* GRENADA */
const CountryGd string = "GD" /* GRENADA */
//CountryGp is country code for /* GUADELOUPE */
const CountryGp string = "GP" /* GUADELOUPE */
//CountryGu is country code for /* GUAM */
const CountryGu string = "GU" /* GUAM */
//CountryGt is country code for /* GUATEMALA */
const CountryGt string = "GT" /* GUATEMALA */
//CountryGn is country code for /* GUINEA */
const CountryGn string = "GN" /* GUINEA */
//CountryGw is country code for /* GUINEA-BISSAU */
const CountryGw string = "GW" /* GUINEA-BISSAU */
//CountryGy is country code for /* GUYANA */
const CountryGy string = "GY" /* GUYANA */
//CountryHt is country code for /* HAITI */
const CountryHt string = "HT" /* HAITI */
//CountryHm is country code for /* HEARD AND MC DONALD ISLANDS */
const CountryHm string = "HM" /* HEARD AND MC DONALD ISLANDS */
//CountryHn is country code for /* HONDURAS */
const CountryHn string = "HN" /* HONDURAS */
//CountryHk is country code for /* HONG KONG */
const CountryHk string = "HK" /* HONG KONG */
//CountryHu is country code for /* HUNGARY */
const CountryHu string = "HU" /* HUNGARY */
//CountryIs is country code for /* ICELAND */
const CountryIs string = "IS" /* ICELAND */
//CountryIn is country code for /* INDIA */
const CountryIn string = "IN" /* INDIA */
//CountryInd is country code for /* INDONESIA */
const CountryInd string = "ID" /* INDONESIA */
//CountryIr is country code for /* IRAN (ISLAMIC REPUBLIC OF) */
const CountryIr string = "IR" /* IRAN (ISLAMIC REPUBLIC OF) */
//CountryIq is country code for /* IRAQ */
const CountryIq string = "IQ" /* IRAQ */
//CountryIe is country code for /* IRELAND */
const CountryIe string = "IE" /* IRELAND */
//CountryIl is country code for /* ISRAEL */
const CountryIl string = "IL" /* ISRAEL */
//CountryIt is country code for /* ITALY */
const CountryIt string = "IT" /* ITALY */
//CountryJm is country code for /* JAMAICA */
const CountryJm string = "JM" /* JAMAICA */
//CountryJp is country code for /* JAPAN */
const CountryJp string = "JP" /* JAPAN */
//CountryJo is country code for /* JORDAN */
const CountryJo string = "JO" /* JORDAN */
//CountryKz is country code for /* KAZAKHSTAN */
const CountryKz string = "KZ" /* KAZAKHSTAN */
//CountryKe is country code for /* KENYA */
const CountryKe string = "KE" /* KENYA */
//CountryKi is country code for /* KIRIBATI */
const CountryKi string = "KI" /* KIRIBATI */
//CountryKp is country code for /* KOREA, DEMOCRATIC PEOPLE'S REPUBLIC OF */
const CountryKp string = "KP" /* KOREA, DEMOCRATIC PEOPLE'S REPUBLIC OF */
//CountryKr is country code for /* KOREA, REPUBLIC OF */
const CountryKr string = "KR" /* KOREA, REPUBLIC OF */
//CountryKw is country code for /* KUWAIT */
const CountryKw string = "KW" /* KUWAIT */
//CountryKg is country code for /* KYRGYZSTAN */
const CountryKg string = "KG" /* KYRGYZSTAN */
//CountryLa is country code for /* LAO PEOPLE'S DEMOCRATIC REPUBLIC */
const CountryLa string = "LA" /* LAO PEOPLE'S DEMOCRATIC REPUBLIC */
//CountryLv is country code for /* LATVIA */
const CountryLv string = "LV" /* LATVIA */
//CountryLb is country code for /* LEBANON */
const CountryLb string = "LB" /* LEBANON */
//CountryLs is country code for /* LESOTHO */
const CountryLs string = "LS" /* LESOTHO */
//CountryLr is country code for /* LIBERIA */
const CountryLr string = "LR" /* LIBERIA */
//CountryLy is country code for /* LIBYAN ARAB JAMAHIRIYA */
const CountryLy string = "LY" /* LIBYAN ARAB JAMAHIRIYA */
//CountryLi is country code for /* LIECHTENSTEIN */
const CountryLi string = "LI" /* LIECHTENSTEIN */
//CountryLt is country code for /* LITHUANIA */
const CountryLt string = "LT" /* LITHUANIA */
//CountryLu is country code for /* LUXEMBOURG */
const CountryLu string = "LU" /* LUXEMBOURG */
//CountryMo is country code for /* MACAU */
const CountryMo string = "MO" /* MACAU */
//CountryMk is country code for /* MACEDONIA, THE FORMER YUGOSLAV REPUBLIC OF */
const CountryMk string = "MK" /* MACEDONIA, THE FORMER YUGOSLAV REPUBLIC OF */
//CountryMg is country code for /* MADAGASCAR */
const CountryMg string = "MG" /* MADAGASCAR */
//CountryMw is country code for /* MALAWI */
const CountryMw string = "MW" /* MALAWI */
//CountryMy is country code for /* MALAYSIA */
const CountryMy string = "MY" /* MALAYSIA */
//CountryMv is country code for /* MALDIVES */
const CountryMv string = "MV" /* MALDIVES */
//CountryMl is country code for /* MALI */
const CountryMl string = "ML" /* MALI */
//CountryMt is country code for /* MALTA */
const CountryMt string = "MT" /* MALTA */
//CountryMh is country code for /* MARSHALL ISLANDS */
const CountryMh string = "MH" /* MARSHALL ISLANDS */
//CountryMq is country code for /* MARTINIQUE */
const CountryMq string = "MQ" /* MARTINIQUE */
//CountryMr is country code for /* MAURITANIA */
const CountryMr string = "MR" /* MAURITANIA */
//CountryMu is country code for /* MAURITIUS */
const CountryMu string = "MU" /* MAURITIUS */
//CountryYt is country code for /* MAYOTTE */
const CountryYt string = "YT" /* MAYOTTE */
//CountryMx is country code for /* MEXICO */
const CountryMx string = "MX" /* MEXICO */
//CountryFm is country code for /* MICRONESIA, FEDERATED STATES OF */
const CountryFm string = "FM" /* MICRONESIA, FEDERATED STATES OF */
//CountryMd is country code for /* MOLDOVA, REPUBLIC OF */
const CountryMd string = "MD" /* MOLDOVA, REPUBLIC OF */
//CountryMc is country code for /* MONACO */
const CountryMc string = "MC" /* MONACO */
//CountryMn is country code for /* MONGOLIA */
const CountryMn string = "MN" /* MONGOLIA */
//CountryMs is country code for /* MONTSERRAT */
const CountryMs string = "MS" /* MONTSERRAT */
//CountryMa is country code for /* MOROCCO */
const CountryMa string = "MA" /* MOROCCO */
//CountryMz is country code for /* MOZAMBIQUE */
const CountryMz string = "MZ" /* MOZAMBIQUE */
//CountryMm is country code for /* MYANMAR */
const CountryMm string = "MM" /* MYANMAR */
//CountryNa is country code for /* NAMIBIA */
const CountryNa string = "NA" /* NAMIBIA */
//CountryNr is country code for /* NAURU */
const CountryNr string = "NR" /* NAURU */
//CountryNp is country code for /* NEPAL */
const CountryNp string = "NP" /* NEPAL */
//CountryNl is country code for /* NETHERLANDS */
const CountryNl string = "NL" /* NETHERLANDS */
//CountryAn is country code for /* NETHERLANDS ANTILLES */
const CountryAn string = "AN" /* NETHERLANDS ANTILLES */
//CountryNc is country code for /* NEW CALEDONIA */
const CountryNc string = "NC" /* NEW CALEDONIA */
//CountryNz is country code for /* NEW ZEALAND */
const CountryNz string = "NZ" /* NEW ZEALAND */
//CountryNi is country code for /* NICARAGUA */
const CountryNi string = "NI" /* NICARAGUA */
//CountryNe is country code for /* NIGER */
const CountryNe string = "NE" /* NIGER */
//CountryNg is country code for /* NIGERIA */
const CountryNg string = "NG" /* NIGERIA */
//CountryNu is country code for /* NIUE */
const CountryNu string = "NU" /* NIUE */
//CountryNf is country code for /* NORFOLK ISLAND */
const CountryNf string = "NF" /* NORFOLK ISLAND */
//CountryMp is country code for /* NORTHERN MARIANA ISLANDS */
const CountryMp string = "MP" /* NORTHERN MARIANA ISLANDS */
//CountryNo is country code for /* NORWAY */
const CountryNo string = "NO" /* NORWAY */
//CountryOm is country code for /* OMAN */
const CountryOm string = "OM" /* OMAN */
//CountryPk is country code for /* PAKISTAN */
const CountryPk string = "PK" /* PAKISTAN */
//CountryPw is country code for /* PALAU */
const CountryPw string = "PW" /* PALAU */
//CountryPa is country code for /* PANAMA */
const CountryPa string = "PA" /* PANAMA */
//CountryPg is country code for /* PAPUA NEW GUINEA */
const CountryPg string = "PG" /* PAPUA NEW GUINEA */
//CountryPy is country code for /* PARAGUAY */
const CountryPy string = "PY" /* PARAGUAY */
//CountryPe is country code for /* PERU */
const CountryPe string = "PE" /* PERU */
//CountryPh is country code for /* PHILIPPINES */
const CountryPh string = "PH" /* PHILIPPINES */
//CountryPn is country code for /* PITCAIRN */
const CountryPn string = "PN" /* PITCAIRN */
//CountryPl is country code for /* POLAND */
const CountryPl string = "PL" /* POLAND */
//CountryPt is country code for /* PORTUGAL */
const CountryPt string = "PT" /* PORTUGAL */
//CountryPr is country code for /* PUERTO RICO */
const CountryPr string = "PR" /* PUERTO RICO */
//CountryQa is country code for /* QATAR */
const CountryQa string = "QA" /* QATAR */
//CountryRe is country code for /* REUNION */
const CountryRe string = "RE" /* REUNION */
//CountryRo is country code for /* ROMANIA */
const CountryRo string = "RO" /* ROMANIA */
//CountryRu is country code for /* RUSSIAN FEDERATION */
const CountryRu string = "RU" /* RUSSIAN FEDERATION */
//CountryRw is country code for /* RWANDA */
const CountryRw string = "RW" /* RWANDA */
//CountryKn is country code for /* SAINT KITTS AND NEVIS */
const CountryKn string = "KN" /* SAINT KITTS AND NEVIS */
//CountryLc is country code for /* SAINT LUCIA */
const CountryLc string = "LC" /* SAINT LUCIA */
//CountryVc is country code for /* SAINT VINCENT AND THE GRENADINES */
const CountryVc string = "VC" /* SAINT VINCENT AND THE GRENADINES */
//CountryWs is country code for /* SAMOA */
const CountryWs string = "WS" /* SAMOA */
//CountrySm is country code for /* SAN MARINO */
const CountrySm string = "SM" /* SAN MARINO */
//CountrySt is country code for /* SAO TOME AND PRINCIPE */
const CountrySt string = "ST" /* SAO TOME AND PRINCIPE */
//CountrySa is country code for /* SAUDI ARABIA */
const CountrySa string = "SA" /* SAUDI ARABIA */
//CountrySn is country code for /* SENEGAL */
const CountrySn string = "SN" /* SENEGAL */
//CountrySc is country code for /* SEYCHELLES */
const CountrySc string = "SC" /* SEYCHELLES */
//CountrySl is country code for /* SIERRA LEONE */
const CountrySl string = "SL" /* SIERRA LEONE */
//CountrySg is country code for /* SINGAPORE */
const CountrySg string = "SG" /* SINGAPORE */
//CountrySk is country code for /* SLOVAKIA (Slovak Republic) */
const CountrySk string = "SK" /* SLOVAKIA (Slovak Republic) */
//CountrySi is country code for /* SLOVENIA */
const CountrySi string = "SI" /* SLOVENIA */
//CountrySb is country code for /* SOLOMON ISLANDS */
const CountrySb string = "SB" /* SOLOMON ISLANDS */
//CountrySo is country code for /* SOMALIA */
const CountrySo string = "SO" /* SOMALIA */
//CountryZa is country code for /* SOUTH AFRICA */
const CountryZa string = "ZA" /* SOUTH AFRICA */
//CountryEs is country code for /* SPAIN */
const CountryEs string = "ES" /* SPAIN */
//CountryLk is country code for /* SRI LANKA */
const CountryLk string = "LK" /* SRI LANKA */
//CountrySh is country code for /* ST. HELENA */
const CountrySh string = "SH" /* ST. HELENA */
//CountryPm is country code for /* ST. PIERRE AND MIQUELON */
const CountryPm string = "PM" /* ST. PIERRE AND MIQUELON */
//CountrySd is country code for /* SUDAN */
const CountrySd string = "SD" /* SUDAN */
//CountrySr is country code for /* SURINAME */
const CountrySr string = "SR" /* SURINAME */
//CountrySj is country code for /* SVALBARD AND JAN MAYEN ISLANDS */
const CountrySj string = "SJ" /* SVALBARD AND JAN MAYEN ISLANDS */
//CountrySz is country code for /* SWAZILAND */
const CountrySz string = "SZ" /* SWAZILAND */
//CountrySe is country code for /* SWEDEN */
const CountrySe string = "SE" /* SWEDEN */
//CountryCh is country code for /* SWITZERLAND */
const CountryCh string = "CH" /* SWITZERLAND */
//CountrySy is country code for /* SYRIAN ARAB REPUBLIC */
const CountrySy string = "SY" /* SYRIAN ARAB REPUBLIC */
//CountryTw is country code for /* TAIWAN, PROVINCE OF CHINA */
const CountryTw string = "TW" /* TAIWAN, PROVINCE OF CHINA */
//CountryTj is country code for /* TAJIKISTAN */
const CountryTj string = "TJ" /* TAJIKISTAN */
//CountryTz is country code for /* TANZANIA, UNITED REPUBLIC OF */
const CountryTz string = "TZ" /* TANZANIA, UNITED REPUBLIC OF */
//CountryTh is country code for /* THAILAND */
const CountryTh string = "TH" /* THAILAND */
//CountryTg is country code for /* TOGO */
const CountryTg string = "TG" /* TOGO */
//CountryTk is country code for /* TOKELAU */
const CountryTk string = "TK" /* TOKELAU */
//CountryTo is country code for /* TONGA */
const CountryTo string = "TO" /* TONGA */
//CountryTt is country code for /* TRINIDAD AND TOBAGO */
const CountryTt string = "TT" /* TRINIDAD AND TOBAGO */
//CountryTn is country code for /* TUNISIA */
const CountryTn string = "TN" /* TUNISIA */
//CountryTr is country code for /* TURKEY */
const CountryTr string = "TR" /* TURKEY */
//CountryTm is country code for /* TURKMENISTAN */
const CountryTm string = "TM" /* TURKMENISTAN */
//CountryTc is country code for /* TURKS AND CAICOS ISLANDS */
const CountryTc string = "TC" /* TURKS AND CAICOS ISLANDS */
//CountryTv is country code for /* TUVALU */
const CountryTv string = "TV" /* TUVALU */
//CountryUg is country code for /* UGANDA */
const CountryUg string = "UG" /* UGANDA */
//CountryUa is country code for /* UKRAINE */
const CountryUa string = "UA" /* UKRAINE */
//CountryAe is country code for /* UNITED ARAB EMIRATES */
const CountryAe string = "AE" /* UNITED ARAB EMIRATES */
//CountryGb is country code for /* UNITED KINGDOM */
const CountryGb string = "GB" /* UNITED KINGDOM */
//CountryUs is country code for /* UNITED STATES */
const CountryUs string = "US" /* UNITED STATES */
//CountryUm is country code for /* UNITED STATES MINOR OUTLYING ISLANDS */
const CountryUm string = "UM" /* UNITED STATES MINOR OUTLYING ISLANDS */
//CountryUy is country code for /* URUGUAY */
const CountryUy string = "UY" /* URUGUAY */
//CountryUz is country code for /* UZBEKISTAN */
const CountryUz string = "UZ" /* UZBEKISTAN */
//CountryVu is country code for /* VANUATU */
const CountryVu string = "VU" /* VANUATU */
//CountryVa is country code for /* VATICAN CITY STATE (HOLY SEE) */
const CountryVa string = "VA" /* VATICAN CITY STATE (HOLY SEE) */
//CountryVe is country code for /* VENEZUELA */
const CountryVe string = "VE" /* VENEZUELA */
//CountryVn is country code for /* VIET NAM */
const CountryVn string = "VN" /* VIET NAM */
//CountryVg is country code for /* VIRGIN ISLANDS (BRITISH) */
const CountryVg string = "VG" /* VIRGIN ISLANDS (BRITISH) */
//CountryVi is country code for /* VIRGIN ISLANDS (U.S.) */
const CountryVi string = "VI" /* VIRGIN ISLANDS (U.S.) */
//CountryWf is country code for /* WALLIS AND FUTUNA ISLANDS */
const CountryWf string = "WF" /* WALLIS AND FUTUNA ISLANDS */
//CountryEh is country code for /* WESTERN SAHARA */
const CountryEh string = "EH" /* WESTERN SAHARA */
//CountryYe is country code for /* YEMEN */
const CountryYe string = "YE" /* YEMEN */
//CountryYu is country code for /* YUGOSLAVIA */
const CountryYu string = "YU" /* YUGOSLAVIA */
//CountryZr is country code for /* ZAIRE */
const CountryZr string = "ZR" /* ZAIRE */
//CountryZm is country code for /* ZAMBIA */
const CountryZm string = "ZM" /* ZAMBIA */
//CountryZw is country code for /* ZIMBABWE */
const CountryZw string = "ZW" /* ZIMBABWE */

/*----------------------------------------------------------------------------*/
/*----- lang code definition -------------------------------------------------*/

//LangAa is language code for /* Afar */
const LangAa string = "aa" /* Afar */
//LangAb is language code for /* Abkhazian */
const LangAb string = "ab" /* Abkhazian */
//LangAf is language code for /* Afrikaans */
const LangAf string = "af" /* Afrikaans */
//LangAm is language code for /* Amharic */
const LangAm string = "am" /* Amharic */
//LangAr is language code for /* Arabic */
const LangAr string = "ar" /* Arabic */
//LangAs is language code for /* Assamese */
const LangAs string = "as" /* Assamese */
//LangAy is language code for /* Aymara */
const LangAy string = "ay" /* Aymara */
//LangAz is language code for /* Azerbaijani */
const LangAz string = "az" /* Azerbaijani */
//LangBa is language code for /* Bashkir */
const LangBa string = "ba" /* Bashkir */
//LangBe is language code for /* Byelorussian */
const LangBe string = "be" /* Byelorussian */
//LangBg is language code for /* Bulgarian */
const LangBg string = "bg" /* Bulgarian */
//LangBh is language code for /* Bihari */
const LangBh string = "bh" /* Bihari */
//LangBi is language code for /* Bislama */
const LangBi string = "bi" /* Bislama */
//LangBn is language code for /* Bengali Bangla */
const LangBn string = "bn" /* Bengali Bangla */
//LangBo is language code for /* Tibetan */
const LangBo string = "bo" /* Tibetan */
//LangBr is language code for /* Breton */
const LangBr string = "br" /* Breton */
//LangCa is language code for /* Catalan */
const LangCa string = "ca" /* Catalan */
//LangCo is language code for /* Corsican */
const LangCo string = "co" /* Corsican */
//LangCs is language code for /* Czech */
const LangCs string = "cs" /* Czech */
//LangCy is language code for /* Welsh */
const LangCy string = "cy" /* Welsh */
//LangDa is language code for /* Danish */
const LangDa string = "da" /* Danish */
//LangDe is language code for /* German */
const LangDe string = "de" /* German */
//LangDz is language code for /* Bhutani */
const LangDz string = "dz" /* Bhutani */
//LangEl is language code for /* Greek */
const LangEl string = "el" /* Greek */
//LangEn is language code for /* English */
const LangEn string = "en" /* English */
//LangEo is language code for /* Esperanto */
const LangEo string = "eo" /* Esperanto */
//LangEs is language code for /* Spanish */
const LangEs string = "es" /* Spanish */
//LangEt is language code for /* Estonian */
const LangEt string = "et" /* Estonian */
//LangEu is language code for /* Basque */
const LangEu string = "eu" /* Basque */
//LangFa is language code for /* Persian */
const LangFa string = "fa" /* Persian */
//LangFi is language code for /* Finnish */
const LangFi string = "fi" /* Finnish */
//LangFj is language code for /* Fiji */
const LangFj string = "fj" /* Fiji */
//LangFo is language code for /* Faeroese */
const LangFo string = "fo" /* Faeroese */
//LangFr is language code for /* French */
const LangFr string = "fr" /* French */
//LangFy is language code for /* Frisian */
const LangFy string = "fy" /* Frisian */
//LangGa is language code for /* Irish */
const LangGa string = "ga" /* Irish */
//LangGd is language code for /* Scots Gaelic */
const LangGd string = "gd" /* Scots Gaelic */
//LangGl is language code for /* Galician */
const LangGl string = "gl" /* Galician */
//LangGn is language code for /* Guarani */
const LangGn string = "gn" /* Guarani */
//LangGu is language code for /* Gujarati */
const LangGu string = "gu" /* Gujarati */
//LangHa is language code for /* Hausa */
const LangHa string = "ha" /* Hausa */
//LangHi is language code for /* Hindi */
const LangHi string = "hi" /* Hindi */
//LangHr is language code for /* Croatian */
const LangHr string = "hr" /* Croatian */
//LangHu is language code for /* Hungarian */
const LangHu string = "hu" /* Hungarian */
//LangHy is language code for /* Armenian */
const LangHy string = "hy" /* Armenian */
//LangIa is language code for /* Interlingua */
const LangIa string = "ia" /* Interlingua */
//LangIe is language code for /* Interlingue */
const LangIe string = "ie" /* Interlingue */
//LangIk is language code for /* Inupiak */
const LangIk string = "ik" /* Inupiak */
//LangIn is language code for /* Indonesian */
const LangIn string = "in" /* Indonesian */
//LangIs is language code for /* Icelandic */
const LangIs string = "is" /* Icelandic */
//LangIt is language code for /* Italian */
const LangIt string = "it" /* Italian */
//LangIw is language code for /* Hebrew */
const LangIw string = "iw" /* Hebrew */
//LangJa is language code for /* Japanese */
const LangJa string = "ja" /* Japanese */
//LangJi is language code for /* Yiddish */
const LangJi string = "ji" /* Yiddish */
//LangJw is language code for /* Javanese */
const LangJw string = "jw" /* Javanese */
//LangKa is language code for /* Georgian */
const LangKa string = "ka" /* Georgian */
//LangKk is language code for /* Kazakh */
const LangKk string = "kk" /* Kazakh */
//LangKl is language code for /* Greenlandic */
const LangKl string = "kl" /* Greenlandic */
//LangKm is language code for /* Cambodian */
const LangKm string = "km" /* Cambodian */
//LangKn is language code for /* Kannada */
const LangKn string = "kn" /* Kannada */
//LangKo is language code for /* Korean */
const LangKo string = "ko" /* Korean */
//LangKs is language code for /* Kashmiri */
const LangKs string = "ks" /* Kashmiri */
//LangKu is language code for /* Kurdish */
const LangKu string = "ku" /* Kurdish */
//LangKy is language code for /* Kirghiz */
const LangKy string = "ky" /* Kirghiz */
//LangLa is language code for /* Latin */
const LangLa string = "la" /* Latin */
//LangLn is language code for /* Lingala */
const LangLn string = "ln" /* Lingala */
//LangLo is language code for /* Laothian */
const LangLo string = "lo" /* Laothian */
//LangLt is language code for /* Lithuanian */
const LangLt string = "lt" /* Lithuanian */
//LangLv is language code for /* Latvian,Lettish */
const LangLv string = "lv" /* Latvian,Lettish */
//LangMg is language code for /* Malagasy */
const LangMg string = "mg" /* Malagasy */
//LangMi is language code for /* Maori */
const LangMi string = "mi" /* Maori */
//LangMk is language code for /* Macedonian */
const LangMk string = "mk" /* Macedonian */
//LangMl is language code for /* Malayalam */
const LangMl string = "ml" /* Malayalam */
//LangMn is language code for /* Mongolian */
const LangMn string = "mn" /* Mongolian */
//LangMo is language code for /* Moldavian */
const LangMo string = "mo" /* Moldavian */
//LangMr is language code for /* Marathi */
const LangMr string = "mr" /* Marathi */
//LangMs is language code for /* Malay */
const LangMs string = "ms" /* Malay */
//LangMt is language code for /* Maltese */
const LangMt string = "mt" /* Maltese */
//LangMy is language code for /* Burmese */
const LangMy string = "my" /* Burmese */
//LangNa is language code for /* Nauru */
const LangNa string = "na" /* Nauru */
//LangNe is language code for /* Nepali */
const LangNe string = "ne" /* Nepali */
//LangNl is language code for /* Dutch */
const LangNl string = "nl" /* Dutch */
//LangNo is language code for /* Norwegian */
const LangNo string = "no" /* Norwegian */
//LangOc is language code for /* Occitan */
const LangOc string = "oc" /* Occitan */
//LangOm is language code for /* (Afan)Oromo */
const LangOm string = "om" /* (Afan)Oromo */
//LangOr is language code for /* Oriya */
const LangOr string = "or" /* Oriya */
//LangPa is language code for /* Punjabi */
const LangPa string = "pa" /* Punjabi */
//LangPl is language code for /* Polish */
const LangPl string = "pl" /* Polish */
//LangPs is language code for /* Pashto,Pushto */
const LangPs string = "ps" /* Pashto,Pushto */
//LangPt is language code for /* Portuguese  */
const LangPt string = "pt" /* Portuguese  */
//LangQu is language code for /* Quechua */
const LangQu string = "qu" /* Quechua */
//LangRm is language code for /* Rhaeto-Romance */
const LangRm string = "rm" /* Rhaeto-Romance */
//LangRn is language code for /* Kirundi */
const LangRn string = "rn" /* Kirundi */
//LangRo is language code for /* Romanian */
const LangRo string = "ro" /* Romanian */
//LangRu is language code for /* Russian */
const LangRu string = "ru" /* Russian */
//LangRw is language code for /* Kinyarwanda */
const LangRw string = "rw" /* Kinyarwanda */
//LangSa is language code for /* Sanskrit */
const LangSa string = "sa" /* Sanskrit */
//LangSd is language code for /* Sindhi */
const LangSd string = "sd" /* Sindhi */
//LangSg is language code for /* Sangro */
const LangSg string = "sg" /* Sangro */
//LangSh is language code for /* Serbo-Croatian */
const LangSh string = "sh" /* Serbo-Croatian */
//LangSi is language code for /* Singhalese */
const LangSi string = "si" /* Singhalese */
//LangSk is language code for /* Slovak */
const LangSk string = "sk" /* Slovak */
//LangSl is language code for /* Slovenian */
const LangSl string = "sl" /* Slovenian */
//LangSm is language code for /* Samoan */
const LangSm string = "sm" /* Samoan */
//LangSn is language code for /* Shona */
const LangSn string = "sn" /* Shona */
//LangSo is language code for /* Somali */
const LangSo string = "so" /* Somali */
//LangSq is language code for /* Albanian */
const LangSq string = "sq" /* Albanian */
//LangSr is language code for /* Serbian */
const LangSr string = "sr" /* Serbian */
//LangSs is language code for /* Siswati */
const LangSs string = "ss" /* Siswati */
//LangSt is language code for /* Sesotho */
const LangSt string = "st" /* Sesotho */
//LangSu is language code for /* Sundanese */
const LangSu string = "su" /* Sundanese */
//LangSv is language code for /* Swedish */
const LangSv string = "sv" /* Swedish */
//LangSw is language code for /* Swahili */
const LangSw string = "sw" /* Swahili */
//LangTa is language code for /* Tamil */
const LangTa string = "ta" /* Tamil */
//LangTe is language code for /* Tegulu */
const LangTe string = "te" /* Tegulu */
//LangTg is language code for /* Tajik */
const LangTg string = "tg" /* Tajik */
//LangTh is language code for /* Thai */
const LangTh string = "th" /* Thai */
//LangTi is language code for /* Tigrinya */
const LangTi string = "ti" /* Tigrinya */
//LangTk is language code for /* Turkmen */
const LangTk string = "tk" /* Turkmen */
//LangTl is language code for /* Tagalog */
const LangTl string = "tl" /* Tagalog */
//LangTn is language code for /* Setswanato Tonga */
const LangTn string = "tn" /* Setswanato Tonga */
//LangTr is language code for /* Turkish */
const LangTr string = "tr" /* Turkish */
//LangTs is language code for /* Tsonga */
const LangTs string = "ts" /* Tsonga */
//LangTt is language code for /* Tatar */
const LangTt string = "tt" /* Tatar */
//LangTw is language code for /* Twi */
const LangTw string = "tw" /* Twi */
//LangUk is language code for /* Ukrainian */
const LangUk string = "uk" /* Ukrainian */
//LangUr is language code for /* Urdu */
const LangUr string = "ur" /* Urdu */
//LangUz is language code for /* Uzbek */
const LangUz string = "uz" /* Uzbek */
//LangVi is language code for /* Vietnamese */
const LangVi string = "vi" /* Vietnamese */
//LangVo is language code for /* Volapuk */
const LangVo string = "vo" /* Volapuk */
//LangWo is language code for /* Wolof */
const LangWo string = "wo" /* Wolof */
//LangXh is language code for /* Xhosa */
const LangXh string = "xh" /* Xhosa */
//LangYo is language code for /* Yoruba */
const LangYo string = "yo" /* Yoruba */
//LangZh is language code for /* Chinese */
const LangZh string = "zh" /* Chinese */
//LangZu is language code for /* Zulu */
const LangZu string = "zu" /* Zulu */

/*----------------------------------------------------------------------------*/
/*----- Graphics mode ---------------------------------------------------------*/

//GmodePageDescription is page graphics mode
const GmodePageDescription int = 0x0001

//GmodePathObject is page graphics mode
const GmodePathObject int = 0x0002

//GmodeTextObject is page graphics mode
const GmodeTextObject int = 0x0004

//GmodeClippingPath is page graphics mode
const GmodeClippingPath int = 0x0008

//GmodeShading is page graphics mode
const GmodeShading int = 0x0010

//GmodeInlineImage is page graphics mode
const GmodeInlineImage int = 0x0020

//GmodeExternalObject is page graphics mode
const GmodeExternalObject int = 0x0040

//EncodingFontSpecific is the name of font specific encoding
const EncodingFontSpecific string = "FontSpecific"

//EncodingStandard is the name of the standard encoding
const EncodingStandard string = "StandardEncoding"

//EncodingMacRoman is the name of mac roman encoding
const EncodingMacRoman string = "MacRomanEncoding"

//EncodingWinAnsi is the name of windows ansi encoding
const EncodingWinAnsi string = "WinAnsiEncoding"

//EncodingIso88592 is the name of ISO8859-2 encoding
const EncodingIso88592 string = "ISO8859-2"

//EncodingIso88593 is the name of ISO8859-3 encoding
const EncodingIso88593 string = "ISO8859-3"

//EncodingIso88594 is the name of ISO8859-4 encoding
const EncodingIso88594 string = "ISO8859-4"

//EncodingIso88595 is the name of ISO8859-5 encoding
const EncodingIso88595 string = "ISO8859-5"

//EncodingIso88596 is the name of ISO8859-6 encoding
const EncodingIso88596 string = "ISO8859-6"

//EncodingIso88597 is the name of ISO8859-7 encoding
const EncodingIso88597 string = "ISO8859-7"

//EncodingIso88598 is the name of ISO8859-8 encoding
const EncodingIso88598 string = "ISO8859-8"

//EncodingIso88599 is the name of ISO8859-9 encoding
const EncodingIso88599 string = "ISO8859-9"

//EncodingIso885910 is the name of ISO8859-10 encoding
const EncodingIso885910 string = "ISO8859-10"

//EncodingIso885911 is the name of ISO8859-11 encoding
const EncodingIso885911 string = "ISO8859-11"

//EncodingIso885913 is the name of ISO8859-13 encoding
const EncodingIso885913 string = "ISO8859-13"

//EncodingIso885914 is the name of ISO8859-14 encoding
const EncodingIso885914 string = "ISO8859-14"

//EncodingIso885915 is the name of ISO8859-15 encoding
const EncodingIso885915 string = "ISO8859-15"

//EncodingIso885916 is the name of ISO8859-16 encoding
const EncodingIso885916 string = "ISO8859-16"

//EncodingCp1250 is the name of codepage 1250 encoding
const EncodingCp1250 string = "CP1250"

//EncodingCp1251 is the name of codepage 1251 encoding
const EncodingCp1251 string = "CP1251"

//EncodingCp1252 is the name of codepage 1252 encoding
const EncodingCp1252 string = "CP1252"

//EncodingCp1253 is the name of codepage 1253 encoding
const EncodingCp1253 string = "CP1253"

//EncodingCp1254 is the name of codepage 1254 encoding
const EncodingCp1254 string = "CP1254"

//EncodingCp1255 is the name of codepage 1255 encoding
const EncodingCp1255 string = "CP1255"

//EncodingCp1256 is the name of codepage 1256 encoding
const EncodingCp1256 string = "CP1256"

//EncodingCp1257 is the name of codepage 1257 encoding
const EncodingCp1257 string = "CP1257"

//EncodingCp1258 is the name of codepage 1258 encoding
const EncodingCp1258 string = "CP1258"

//EncodingKoi8R is the name of codepage KOI-8(R) encoding
const EncodingKoi8R string = "KOI8-R"

//ErrArrayCountErr is a error code
const ErrArrayCountErr int = 0x1001

//ErrArrayItemNotFound is a error code
const ErrArrayItemNotFound int = 0x1002

//ErrArrayItemUnexpectedType is a error code
const ErrArrayItemUnexpectedType int = 0x1003

//ErrBinaryLengthErr is a error code
const ErrBinaryLengthErr int = 0x1004

//ErrCannotGetPallet is a error code
const ErrCannotGetPallet int = 0x1005

//ErrDictCountErr is a error code
const ErrDictCountErr int = 0x1007

//ErrDictItemNotFound is a error code
const ErrDictItemNotFound int = 0x1008

//ErrDictItemUnexpectedType is a error code
const ErrDictItemUnexpectedType int = 0x1009

//ErrDictStreamLengthNotFound is a error code
const ErrDictStreamLengthNotFound int = 0x100A

//ErrDocEncryptdictNotFound is a error code
const ErrDocEncryptdictNotFound int = 0x100B

//ErrDocInvalidObject is a error code
const ErrDocInvalidObject int = 0x100C

//ErrDuplicateRegistration is a error code
const ErrDuplicateRegistration int = 0x100E

//ErrExceedJwwCodeNumLimit is a error code
const ErrExceedJwwCodeNumLimit int = 0x100F

//ErrEncryptInvalidPassword is a error code
const ErrEncryptInvalidPassword int = 0x1011

//ErrErrUnknownClass is a error code
const ErrErrUnknownClass int = 0x1013

//ErrExceedGstateLimit is a error code
const ErrExceedGstateLimit int = 0x1014

//ErrFailedToAllocMem is a error code
const ErrFailedToAllocMem int = 0x1015

//ErrFileIoError is a error code
const ErrFileIoError int = 0x1016

//ErrFileOpenError is a error code
const ErrFileOpenError int = 0x1017

//ErrFontExists is a error code
const ErrFontExists int = 0x1019

//ErrFontInvalidWidthsTable is a error code
const ErrFontInvalidWidthsTable int = 0x101A

//ErrInvalidAfmHeader is a error code
const ErrInvalidAfmHeader int = 0x101B

//ErrInvalidAnnotation is a error code
const ErrInvalidAnnotation int = 0x101C

//ErrInvalidBitPerComponent is a error code
const ErrInvalidBitPerComponent int = 0x101E

//ErrInvalidCharMatrixData is a error code
const ErrInvalidCharMatrixData int = 0x101F

//ErrInvalidColorSpace is a error code
const ErrInvalidColorSpace int = 0x1020

//ErrInvalidCompressionMode is a error code
const ErrInvalidCompressionMode int = 0x1021

//ErrInvalidDateTime is a error code
const ErrInvalidDateTime int = 0x1022

//ErrInvalidDestination is a error code
const ErrInvalidDestination int = 0x1023

//ErrInvalidDocument is a error code
const ErrInvalidDocument int = 0x1025

//ErrInvalidDocumentState is a error code
const ErrInvalidDocumentState int = 0x1026

//ErrInvalidEncoder is a error code
const ErrInvalidEncoder int = 0x1027

//ErrInvalidEncoderType is a error code
const ErrInvalidEncoderType int = 0x1028

//ErrInvalidEncodingName is a error code
const ErrInvalidEncodingName int = 0x102B

//ErrInvalidEncryptKeyLen is a error code
const ErrInvalidEncryptKeyLen int = 0x102C

//ErrInvalidFontdefData is a error code
const ErrInvalidFontdefData int = 0x102D

//ErrInvalidFontdefType is a error code
const ErrInvalidFontdefType int = 0x102E

//ErrInvalidFontName is a error code
const ErrInvalidFontName int = 0x102F

//ErrInvalidImage is a error code
const ErrInvalidImage int = 0x1030

//ErrInvalidJpegData is a error code
const ErrInvalidJpegData int = 0x1031

//ErrInvalidNData is a error code
const ErrInvalidNData int = 0x1032

//ErrInvalidObject is a error code
const ErrInvalidObject int = 0x1033

//ErrInvalidObjID is a error code
const ErrInvalidObjID int = 0x1034

//ErrInvalidOperation is a error code
const ErrInvalidOperation int = 0x1035

//ErrInvalidOutline is a error code
const ErrInvalidOutline int = 0x1036

//ErrInvalidPage is a error code
const ErrInvalidPage int = 0x1037

//ErrInvalidPages is a error code
const ErrInvalidPages int = 0x1038

//ErrInvalidParameter is a error code
const ErrInvalidParameter int = 0x1039

//ErrInvalidPngImage is a error code
const ErrInvalidPngImage int = 0x103B

//ErrInvalidStream is a error code
const ErrInvalidStream int = 0x103C

//ErrMissingFileNameEntry is a error code
const ErrMissingFileNameEntry int = 0x103D

//ErrInvalidTtcFile is a error code
const ErrInvalidTtcFile int = 0x103F

//ErrInvalidTtcIndex is a error code
const ErrInvalidTtcIndex int = 0x1040

//ErrInvalidWxData is a error code
const ErrInvalidWxData int = 0x1041

//ErrItemNotFound is a error code
const ErrItemNotFound int = 0x1042

//ErrLibpngError is a error code
const ErrLibpngError int = 0x1043

//ErrNameInvalidValue is a error code
const ErrNameInvalidValue int = 0x1044

//ErrNameOutOfRange is a error code
const ErrNameOutOfRange int = 0x1045

//ErrPageInvalidParamCount is a error code
const ErrPageInvalidParamCount int = 0x1048

//ErrPagesMissingKidsEntry is a error code
const ErrPagesMissingKidsEntry int = 0x1049

//ErrPageCannotFindObject is a error code
const ErrPageCannotFindObject int = 0x104A

//ErrPageCannotGetRootPages is a error code
const ErrPageCannotGetRootPages int = 0x104B

//ErrPageCannotRestoreGstate is a error code
const ErrPageCannotRestoreGstate int = 0x104C

//ErrPageCannotSetParent is a error code
const ErrPageCannotSetParent int = 0x104D

//ErrPageFontNotFound is a error code
const ErrPageFontNotFound int = 0x104E

//ErrPageInvalidFont is a error code
const ErrPageInvalidFont int = 0x104F

//ErrPageInvalidFontSize is a error code
const ErrPageInvalidFontSize int = 0x1050

//ErrPageInvalidGmode is a error code
const ErrPageInvalidGmode int = 0x1051

//ErrPageInvalidIndex is a error code
const ErrPageInvalidIndex int = 0x1052

//ErrPageInvalidRotateValue is a error code
const ErrPageInvalidRotateValue int = 0x1053

//ErrPageInvalidSize is a error code
const ErrPageInvalidSize int = 0x1054

//ErrPageInvalidXObject is a error code
const ErrPageInvalidXObject int = 0x1055

//ErrPageOutOfRange is a error code
const ErrPageOutOfRange int = 0x1056

//ErrRealOutOfRange is a error code
const ErrRealOutOfRange int = 0x1057

//ErrStreamEOF is a error code
const ErrStreamEOF int = 0x1058

//ErrStreamReadContinue is a error code
const ErrStreamReadContinue int = 0x1059

//ErrStringOutOfRange is a error code
const ErrStringOutOfRange int = 0x105B

//ErrThisFuncWasSkipped is a error code
const ErrThisFuncWasSkipped int = 0x105C

//ErrTtfCannotEmbeddingFont is a error code
const ErrTtfCannotEmbeddingFont int = 0x105D

//ErrTtfInvalidCharMap is a error code
const ErrTtfInvalidCharMap int = 0x105E

//ErrTtfInvalidFormat is a error code
const ErrTtfInvalidFormat int = 0x105F

//ErrTtfMissingTable is a error code
const ErrTtfMissingTable int = 0x1060

//ErrUnsupportedFontType is a error code
const ErrUnsupportedFontType int = 0x1061

//ErrUnsupportedFunc is a error code
const ErrUnsupportedFunc int = 0x1062

//ErrUnsupportedJpegFormat is a error code
const ErrUnsupportedJpegFormat int = 0x1063

//ErrUnsupportedType1Font is a error code
const ErrUnsupportedType1Font int = 0x1064

//ErrXrefCountErr is a error code
const ErrXrefCountErr int = 0x1065

//ErrZlibError is a error code
const ErrZlibError int = 0x1066

//ErrInvalidPageIndex is a error code
const ErrInvalidPageIndex int = 0x1067

//ErrInvalidURI is a error code
const ErrInvalidURI int = 0x1068

//ErrPageLayoutOutOfRange is a error code
const ErrPageLayoutOutOfRange int = 0x1069

//ErrPageModeOutOfRange is a error code
const ErrPageModeOutOfRange int = 0x1070

//ErrPageNumStyleOutOfRange is a error code
const ErrPageNumStyleOutOfRange int = 0x1071

//ErrAnnotInvalidIcon is a error code
const ErrAnnotInvalidIcon int = 0x1072

//ErrAnnotInvalidBorderStyle is a error code
const ErrAnnotInvalidBorderStyle int = 0x1073

//ErrPageInvalidDirection is a error code
const ErrPageInvalidDirection int = 0x1074

//ErrInvalidFont is a error code
const ErrInvalidFont int = 0x1075

//ErrPageInsufficientSpace is a error code
const ErrPageInsufficientSpace int = 0x1076

//ErrPageInvalidDisplayTime is a error code
const ErrPageInvalidDisplayTime int = 0x1077

//ErrPageInvalidTransitionTime is a error code
const ErrPageInvalidTransitionTime int = 0x1078

//ErrInvalidPageSlideshowType is a error code
const ErrInvalidPageSlideshowType int = 0x1079

//ErrExtGstateOutOfRange is a error code
const ErrExtGstateOutOfRange int = 0x1080

//ErrInvalidExtGstate is a error code
const ErrInvalidExtGstate int = 0x1081

//ErrExtGstateReadOnly is a error code
const ErrExtGstateReadOnly int = 0x1082

//ErrInvalidU3dData is a error code
const ErrInvalidU3dData int = 0x1083

//ErrNameCannotGetNames is a error code
const ErrNameCannotGetNames int = 0x1084

//ErrInvalidIccComponentNum is a error code
const ErrInvalidIccComponentNum int = 0x1085

//FontFixedWidth is font type
const FontFixedWidth int = 1

//FontSerif is font type
const FontSerif int = 2

//FontSymbolic is font type
const FontSymbolic int = 4

//FontScript is font type
const FontScript int = 8

//FontStdCharset is font type
const FontStdCharset int = 32

//FontItalic is font type
const FontItalic int = 64

//FontAllCap is font type
const FontAllCap int = 65536

//FontSmallCap is font type
const FontSmallCap int = 131072

//FontForceBold is font type
const FontForceBold int = 262144

//BsSolid is a border style subtype
const BsSolid int = 0

//BsDashed is a border style subtype
const BsDashed int = 1

//BsBeveled is a border style subtype
const BsBeveled int = 2

//BsInset is a border style subtype
const BsInset int = 3

//BsUnderlined is a border style subtype
const BsUnderlined int = 4
