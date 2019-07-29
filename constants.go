package goharu

/*----------------------------------------------------------------------------*/
const TRUE int = 1
const FALSE int = 0

const OK int = 0
const NOERROR int = 0

/* default value of Graphic State */
const DEF_FONT string = "Helvetica"
const DEF_WORDSPACE int = 0
const DEF_CHARSPACE int = 0
const DEF_FONTSIZE int = 10
const DEF_HSCALING int = 100
const DEF_LEADING int = 0
const DEF_RISE int = 0
const DEF_LINEWIDTH int = 1
const DEF_MITERLIMIT int = 10
const DEF_FLATNESS int = 1
const DEF_PAGE_NUM int = 1

const BS_DEF_WIDTH int = 1

/* defalt page-size */
const DEF_PAGE_WIDTH float32 = 595.276
const DEF_PAGE_HEIGHT float32 = 841.89

const COMP_NONE int = 0x00
const COMP_TEXT int = 0x01
const COMP_IMAGE int = 0x02
const COMP_METADATA int = 0x04
const COMP_ALL int = 0x0F
const COMP_MASK int = 0xFF

/*----------------------------------------------------------------------------*/
/*----- permission flags (only Revision 2 is supported)-----------------------*/

const ENABLE_READ int = 0
const ENABLE_PRINT int = 4
const ENABLE_EDIT_ALL int = 8
const ENABLE_COPY int = 16
const ENABLE_EDIT int = 32

/*----------------------------------------------------------------------------*/
/*------ viewer preferences definitions --------------------------------------*/

const HIDE_TOOLBAR int = 1
const HIDE_MENUBAR int = 2
const HIDE_WINDOW_UI int = 4
const FIT_WINDOW int = 8
const CENTER_WINDOW int = 16
const PRINT_SCALING_NONE int = 32

/*---------------------------------------------------------------------------*/
/*------ limitation of object implementation (PDF1.4) -----------------------*/

const LIMIT_MAX_INT int = 2147483647
const LIMIT_MIN_INT int = -2147483647

const LIMIT_MAX_REAL float32 = 3.4e38  // per PDF 1.7 spec, Annex C, old value  32767
const LIMIT_MIN_REAL float32 = -3.4e38 // per PDF 1.7 spec, Annex C, old value -32767

const LIMIT_MAX_STRING_LEN int = 2147483646 // per PDF 1.7 spec, limit 32767 is for strings in content stream and no limit in other cases => open the limit to max Integer, old value 65535
const LIMIT_MAX_NAME_LEN int = 127

const LIMIT_MAX_ARRAY int = 8388607        // per PDF 1.7 spec, "Maximum number of indirect objects in a PDF file" is 8388607, old value 8191
const LIMIT_MAX_DICT_ELEMENT int = 8388607 // per PDF 1.7 spec, "Maximum number of indirect objects in a PDF file" is 8388607, old value 4095
const LIMIT_MAX_XREF_ELEMENT int = 8388607
const LIMIT_MAX_GSTATE int = 28
const LIMIT_MAX_DEVICE_N int = 8
const LIMIT_MAX_DEVICE_N_V15 int = 32
const LIMIT_MAX_CID int = 65535
const MAX_GENERATION_NUM int = 65535

const MIN_PAGE_HEIGHT int = 3
const MIN_PAGE_WIDTH int = 3
const MAX_PAGE_HEIGHT int = 14400
const MAX_PAGE_WIDTH int = 14400
const MIN_MAGNIFICATION_FACTOR int = 8
const MAX_MAGNIFICATION_FACTOR int = 3200

/*---------------------------------------------------------------------------*/
/*------ limitation of various properties -----------------------------------*/

const MIN_PAGE_SIZE int = 3
const MAX_PAGE_SIZE int = 14400
const MIN_HORIZONTALSCALING int = 10
const MAX_HORIZONTALSCALING int = 300
const MIN_WORDSPACE int = -30
const MAX_WORDSPACE int = 300
const MIN_CHARSPACE int = -30
const MAX_CHARSPACE int = 300
const MAX_FONTSIZE int = 600
const MAX_ZOOMSIZE int = 10
const MAX_LEADING int = 300
const MAX_LINEWIDTH int = 100
const MAX_DASH_PATTERN int = 100
const MAX_JWW_NUM int = 128

/*----------------------------------------------------------------------------*/
/*----- country code definition ----------------------------------------------*/

const COUNTRY_AF string = "AF" /* AFGHANISTAN */
const COUNTRY_AL string = "AL" /* ALBANIA */
const COUNTRY_DZ string = "DZ" /* ALGERIA */
const COUNTRY_AS string = "AS" /* AMERICAN SAMOA */
const COUNTRY_AD string = "AD" /* ANDORRA */
const COUNTRY_AO string = "AO" /* ANGOLA */
const COUNTRY_AI string = "AI" /* ANGUILLA */
const COUNTRY_AQ string = "AQ" /* ANTARCTICA */
const COUNTRY_AG string = "AG" /* ANTIGUA AND BARBUDA */
const COUNTRY_AR string = "AR" /* ARGENTINA */
const COUNTRY_AM string = "AM" /* ARMENIA */
const COUNTRY_AW string = "AW" /* ARUBA */
const COUNTRY_AU string = "AU" /* AUSTRALIA */
const COUNTRY_AT string = "AT" /* AUSTRIA */
const COUNTRY_AZ string = "AZ" /* AZERBAIJAN */
const COUNTRY_BS string = "BS" /* BAHAMAS */
const COUNTRY_BH string = "BH" /* BAHRAIN */
const COUNTRY_BD string = "BD" /* BANGLADESH */
const COUNTRY_BB string = "BB" /* BARBADOS */
const COUNTRY_BY string = "BY" /* BELARUS */
const COUNTRY_BE string = "BE" /* BELGIUM */
const COUNTRY_BZ string = "BZ" /* BELIZE */
const COUNTRY_BJ string = "BJ" /* BENIN */
const COUNTRY_BM string = "BM" /* BERMUDA */
const COUNTRY_BT string = "BT" /* BHUTAN */
const COUNTRY_BO string = "BO" /* BOLIVIA */
const COUNTRY_BA string = "BA" /* BOSNIA AND HERZEGOVINA */
const COUNTRY_BW string = "BW" /* BOTSWANA */
const COUNTRY_BV string = "BV" /* BOUVET ISLAND */
const COUNTRY_BR string = "BR" /* BRAZIL */
const COUNTRY_IO string = "IO" /* BRITISH INDIAN OCEAN TERRITORY */
const COUNTRY_BN string = "BN" /* BRUNEI DARUSSALAM */
const COUNTRY_BG string = "BG" /* BULGARIA */
const COUNTRY_BF string = "BF" /* BURKINA FASO */
const COUNTRY_BI string = "BI" /* BURUNDI */
const COUNTRY_KH string = "KH" /* CAMBODIA */
const COUNTRY_CM string = "CM" /* CAMEROON */
const COUNTRY_CA string = "CA" /* CANADA */
const COUNTRY_CV string = "CV" /* CAPE VERDE */
const COUNTRY_KY string = "KY" /* CAYMAN ISLANDS */
const COUNTRY_CF string = "CF" /* CENTRAL AFRICAN REPUBLIC */
const COUNTRY_TD string = "TD" /* CHAD */
const COUNTRY_CL string = "CL" /* CHILE */
const COUNTRY_CN string = "CN" /* CHINA */
const COUNTRY_CX string = "CX" /* CHRISTMAS ISLAND */
const COUNTRY_CC string = "CC" /* COCOS (KEELING) ISLANDS */
const COUNTRY_CO string = "CO" /* COLOMBIA */
const COUNTRY_KM string = "KM" /* COMOROS */
const COUNTRY_CG string = "CG" /* CONGO */
const COUNTRY_CK string = "CK" /* COOK ISLANDS */
const COUNTRY_CR string = "CR" /* COSTA RICA */
const COUNTRY_CI string = "CI" /* COTE D'IVOIRE */
const COUNTRY_HR string = "HR" /* CROATIA (local name: Hrvatska) */
const COUNTRY_CU string = "CU" /* CUBA */
const COUNTRY_CY string = "CY" /* CYPRUS */
const COUNTRY_CZ string = "CZ" /* CZECH REPUBLIC */
const COUNTRY_DK string = "DK" /* DENMARK */
const COUNTRY_DJ string = "DJ" /* DJIBOUTI */
const COUNTRY_DM string = "DM" /* DOMINICA */
const COUNTRY_DO string = "DO" /* DOMINICAN REPUBLIC */
const COUNTRY_TP string = "TP" /* EAST TIMOR */
const COUNTRY_EC string = "EC" /* ECUADOR */
const COUNTRY_EG string = "EG" /* EGYPT */
const COUNTRY_SV string = "SV" /* EL SALVADOR */
const COUNTRY_GQ string = "GQ" /* EQUATORIAL GUINEA */
const COUNTRY_ER string = "ER" /* ERITREA */
const COUNTRY_EE string = "EE" /* ESTONIA */
const COUNTRY_ET string = "ET" /* ETHIOPIA */
const COUNTRY_FK string = "FK" /* FALKLAND ISLANDS (MALVINAS) */
const COUNTRY_FO string = "FO" /* FAROE ISLANDS */
const COUNTRY_FJ string = "FJ" /* FIJI */
const COUNTRY_FI string = "FI" /* FINLAND */
const COUNTRY_FR string = "FR" /* FRANCE */
const COUNTRY_FX string = "FX" /* FRANCE, METROPOLITAN */
const COUNTRY_GF string = "GF" /* FRENCH GUIANA */
const COUNTRY_PF string = "PF" /* FRENCH POLYNESIA */
const COUNTRY_TF string = "TF" /* FRENCH SOUTHERN TERRITORIES */
const COUNTRY_GA string = "GA" /* GABON */
const COUNTRY_GM string = "GM" /* GAMBIA */
const COUNTRY_GE string = "GE" /* GEORGIA */
const COUNTRY_DE string = "DE" /* GERMANY */
const COUNTRY_GH string = "GH" /* GHANA */
const COUNTRY_GI string = "GI" /* GIBRALTAR */
const COUNTRY_GR string = "GR" /* GREECE */
const COUNTRY_GL string = "GL" /* GREENLAND */
const COUNTRY_GD string = "GD" /* GRENADA */
const COUNTRY_GP string = "GP" /* GUADELOUPE */
const COUNTRY_GU string = "GU" /* GUAM */
const COUNTRY_GT string = "GT" /* GUATEMALA */
const COUNTRY_GN string = "GN" /* GUINEA */
const COUNTRY_GW string = "GW" /* GUINEA-BISSAU */
const COUNTRY_GY string = "GY" /* GUYANA */
const COUNTRY_HT string = "HT" /* HAITI */
const COUNTRY_HM string = "HM" /* HEARD AND MC DONALD ISLANDS */
const COUNTRY_HN string = "HN" /* HONDURAS */
const COUNTRY_HK string = "HK" /* HONG KONG */
const COUNTRY_HU string = "HU" /* HUNGARY */
const COUNTRY_IS string = "IS" /* ICELAND */
const COUNTRY_IN string = "IN" /* INDIA */
const COUNTRY_ID string = "ID" /* INDONESIA */
const COUNTRY_IR string = "IR" /* IRAN (ISLAMIC REPUBLIC OF) */
const COUNTRY_IQ string = "IQ" /* IRAQ */
const COUNTRY_IE string = "IE" /* IRELAND */
const COUNTRY_IL string = "IL" /* ISRAEL */
const COUNTRY_IT string = "IT" /* ITALY */
const COUNTRY_JM string = "JM" /* JAMAICA */
const COUNTRY_JP string = "JP" /* JAPAN */
const COUNTRY_JO string = "JO" /* JORDAN */
const COUNTRY_KZ string = "KZ" /* KAZAKHSTAN */
const COUNTRY_KE string = "KE" /* KENYA */
const COUNTRY_KI string = "KI" /* KIRIBATI */
const COUNTRY_KP string = "KP" /* KOREA, DEMOCRATIC PEOPLE'S REPUBLIC OF */
const COUNTRY_KR string = "KR" /* KOREA, REPUBLIC OF */
const COUNTRY_KW string = "KW" /* KUWAIT */
const COUNTRY_KG string = "KG" /* KYRGYZSTAN */
const COUNTRY_LA string = "LA" /* LAO PEOPLE'S DEMOCRATIC REPUBLIC */
const COUNTRY_LV string = "LV" /* LATVIA */
const COUNTRY_LB string = "LB" /* LEBANON */
const COUNTRY_LS string = "LS" /* LESOTHO */
const COUNTRY_LR string = "LR" /* LIBERIA */
const COUNTRY_LY string = "LY" /* LIBYAN ARAB JAMAHIRIYA */
const COUNTRY_LI string = "LI" /* LIECHTENSTEIN */
const COUNTRY_LT string = "LT" /* LITHUANIA */
const COUNTRY_LU string = "LU" /* LUXEMBOURG */
const COUNTRY_MO string = "MO" /* MACAU */
const COUNTRY_MK string = "MK" /* MACEDONIA, THE FORMER YUGOSLAV REPUBLIC OF */
const COUNTRY_MG string = "MG" /* MADAGASCAR */
const COUNTRY_MW string = "MW" /* MALAWI */
const COUNTRY_MY string = "MY" /* MALAYSIA */
const COUNTRY_MV string = "MV" /* MALDIVES */
const COUNTRY_ML string = "ML" /* MALI */
const COUNTRY_MT string = "MT" /* MALTA */
const COUNTRY_MH string = "MH" /* MARSHALL ISLANDS */
const COUNTRY_MQ string = "MQ" /* MARTINIQUE */
const COUNTRY_MR string = "MR" /* MAURITANIA */
const COUNTRY_MU string = "MU" /* MAURITIUS */
const COUNTRY_YT string = "YT" /* MAYOTTE */
const COUNTRY_MX string = "MX" /* MEXICO */
const COUNTRY_FM string = "FM" /* MICRONESIA, FEDERATED STATES OF */
const COUNTRY_MD string = "MD" /* MOLDOVA, REPUBLIC OF */
const COUNTRY_MC string = "MC" /* MONACO */
const COUNTRY_MN string = "MN" /* MONGOLIA */
const COUNTRY_MS string = "MS" /* MONTSERRAT */
const COUNTRY_MA string = "MA" /* MOROCCO */
const COUNTRY_MZ string = "MZ" /* MOZAMBIQUE */
const COUNTRY_MM string = "MM" /* MYANMAR */
const COUNTRY_NA string = "NA" /* NAMIBIA */
const COUNTRY_NR string = "NR" /* NAURU */
const COUNTRY_NP string = "NP" /* NEPAL */
const COUNTRY_NL string = "NL" /* NETHERLANDS */
const COUNTRY_AN string = "AN" /* NETHERLANDS ANTILLES */
const COUNTRY_NC string = "NC" /* NEW CALEDONIA */
const COUNTRY_NZ string = "NZ" /* NEW ZEALAND */
const COUNTRY_NI string = "NI" /* NICARAGUA */
const COUNTRY_NE string = "NE" /* NIGER */
const COUNTRY_NG string = "NG" /* NIGERIA */
const COUNTRY_NU string = "NU" /* NIUE */
const COUNTRY_NF string = "NF" /* NORFOLK ISLAND */
const COUNTRY_MP string = "MP" /* NORTHERN MARIANA ISLANDS */
const COUNTRY_NO string = "NO" /* NORWAY */
const COUNTRY_OM string = "OM" /* OMAN */
const COUNTRY_PK string = "PK" /* PAKISTAN */
const COUNTRY_PW string = "PW" /* PALAU */
const COUNTRY_PA string = "PA" /* PANAMA */
const COUNTRY_PG string = "PG" /* PAPUA NEW GUINEA */
const COUNTRY_PY string = "PY" /* PARAGUAY */
const COUNTRY_PE string = "PE" /* PERU */
const COUNTRY_PH string = "PH" /* PHILIPPINES */
const COUNTRY_PN string = "PN" /* PITCAIRN */
const COUNTRY_PL string = "PL" /* POLAND */
const COUNTRY_PT string = "PT" /* PORTUGAL */
const COUNTRY_PR string = "PR" /* PUERTO RICO */
const COUNTRY_QA string = "QA" /* QATAR */
const COUNTRY_RE string = "RE" /* REUNION */
const COUNTRY_RO string = "RO" /* ROMANIA */
const COUNTRY_RU string = "RU" /* RUSSIAN FEDERATION */
const COUNTRY_RW string = "RW" /* RWANDA */
const COUNTRY_KN string = "KN" /* SAINT KITTS AND NEVIS */
const COUNTRY_LC string = "LC" /* SAINT LUCIA */
const COUNTRY_VC string = "VC" /* SAINT VINCENT AND THE GRENADINES */
const COUNTRY_WS string = "WS" /* SAMOA */
const COUNTRY_SM string = "SM" /* SAN MARINO */
const COUNTRY_ST string = "ST" /* SAO TOME AND PRINCIPE */
const COUNTRY_SA string = "SA" /* SAUDI ARABIA */
const COUNTRY_SN string = "SN" /* SENEGAL */
const COUNTRY_SC string = "SC" /* SEYCHELLES */
const COUNTRY_SL string = "SL" /* SIERRA LEONE */
const COUNTRY_SG string = "SG" /* SINGAPORE */
const COUNTRY_SK string = "SK" /* SLOVAKIA (Slovak Republic) */
const COUNTRY_SI string = "SI" /* SLOVENIA */
const COUNTRY_SB string = "SB" /* SOLOMON ISLANDS */
const COUNTRY_SO string = "SO" /* SOMALIA */
const COUNTRY_ZA string = "ZA" /* SOUTH AFRICA */
const COUNTRY_ES string = "ES" /* SPAIN */
const COUNTRY_LK string = "LK" /* SRI LANKA */
const COUNTRY_SH string = "SH" /* ST. HELENA */
const COUNTRY_PM string = "PM" /* ST. PIERRE AND MIQUELON */
const COUNTRY_SD string = "SD" /* SUDAN */
const COUNTRY_SR string = "SR" /* SURINAME */
const COUNTRY_SJ string = "SJ" /* SVALBARD AND JAN MAYEN ISLANDS */
const COUNTRY_SZ string = "SZ" /* SWAZILAND */
const COUNTRY_SE string = "SE" /* SWEDEN */
const COUNTRY_CH string = "CH" /* SWITZERLAND */
const COUNTRY_SY string = "SY" /* SYRIAN ARAB REPUBLIC */
const COUNTRY_TW string = "TW" /* TAIWAN, PROVINCE OF CHINA */
const COUNTRY_TJ string = "TJ" /* TAJIKISTAN */
const COUNTRY_TZ string = "TZ" /* TANZANIA, UNITED REPUBLIC OF */
const COUNTRY_TH string = "TH" /* THAILAND */
const COUNTRY_TG string = "TG" /* TOGO */
const COUNTRY_TK string = "TK" /* TOKELAU */
const COUNTRY_TO string = "TO" /* TONGA */
const COUNTRY_TT string = "TT" /* TRINIDAD AND TOBAGO */
const COUNTRY_TN string = "TN" /* TUNISIA */
const COUNTRY_TR string = "TR" /* TURKEY */
const COUNTRY_TM string = "TM" /* TURKMENISTAN */
const COUNTRY_TC string = "TC" /* TURKS AND CAICOS ISLANDS */
const COUNTRY_TV string = "TV" /* TUVALU */
const COUNTRY_UG string = "UG" /* UGANDA */
const COUNTRY_UA string = "UA" /* UKRAINE */
const COUNTRY_AE string = "AE" /* UNITED ARAB EMIRATES */
const COUNTRY_GB string = "GB" /* UNITED KINGDOM */
const COUNTRY_US string = "US" /* UNITED STATES */
const COUNTRY_UM string = "UM" /* UNITED STATES MINOR OUTLYING ISLANDS */
const COUNTRY_UY string = "UY" /* URUGUAY */
const COUNTRY_UZ string = "UZ" /* UZBEKISTAN */
const COUNTRY_VU string = "VU" /* VANUATU */
const COUNTRY_VA string = "VA" /* VATICAN CITY STATE (HOLY SEE) */
const COUNTRY_VE string = "VE" /* VENEZUELA */
const COUNTRY_VN string = "VN" /* VIET NAM */
const COUNTRY_VG string = "VG" /* VIRGIN ISLANDS (BRITISH) */
const COUNTRY_VI string = "VI" /* VIRGIN ISLANDS (U.S.) */
const COUNTRY_WF string = "WF" /* WALLIS AND FUTUNA ISLANDS */
const COUNTRY_EH string = "EH" /* WESTERN SAHARA */
const COUNTRY_YE string = "YE" /* YEMEN */
const COUNTRY_YU string = "YU" /* YUGOSLAVIA */
const COUNTRY_ZR string = "ZR" /* ZAIRE */
const COUNTRY_ZM string = "ZM" /* ZAMBIA */
const COUNTRY_ZW string = "ZW" /* ZIMBABWE */

/*----------------------------------------------------------------------------*/
/*----- lang code definition -------------------------------------------------*/

const LANG_AA string = "aa" /* Afar */
const LANG_AB string = "ab" /* Abkhazian */
const LANG_AF string = "af" /* Afrikaans */
const LANG_AM string = "am" /* Amharic */
const LANG_AR string = "ar" /* Arabic */
const LANG_AS string = "as" /* Assamese */
const LANG_AY string = "ay" /* Aymara */
const LANG_AZ string = "az" /* Azerbaijani */
const LANG_BA string = "ba" /* Bashkir */
const LANG_BE string = "be" /* Byelorussian */
const LANG_BG string = "bg" /* Bulgarian */
const LANG_BH string = "bh" /* Bihari */
const LANG_BI string = "bi" /* Bislama */
const LANG_BN string = "bn" /* Bengali Bangla */
const LANG_BO string = "bo" /* Tibetan */
const LANG_BR string = "br" /* Breton */
const LANG_CA string = "ca" /* Catalan */
const LANG_CO string = "co" /* Corsican */
const LANG_CS string = "cs" /* Czech */
const LANG_CY string = "cy" /* Welsh */
const LANG_DA string = "da" /* Danish */
const LANG_DE string = "de" /* German */
const LANG_DZ string = "dz" /* Bhutani */
const LANG_EL string = "el" /* Greek */
const LANG_EN string = "en" /* English */
const LANG_EO string = "eo" /* Esperanto */
const LANG_ES string = "es" /* Spanish */
const LANG_ET string = "et" /* Estonian */
const LANG_EU string = "eu" /* Basque */
const LANG_FA string = "fa" /* Persian */
const LANG_FI string = "fi" /* Finnish */
const LANG_FJ string = "fj" /* Fiji */
const LANG_FO string = "fo" /* Faeroese */
const LANG_FR string = "fr" /* French */
const LANG_FY string = "fy" /* Frisian */
const LANG_GA string = "ga" /* Irish */
const LANG_GD string = "gd" /* Scots Gaelic */
const LANG_GL string = "gl" /* Galician */
const LANG_GN string = "gn" /* Guarani */
const LANG_GU string = "gu" /* Gujarati */
const LANG_HA string = "ha" /* Hausa */
const LANG_HI string = "hi" /* Hindi */
const LANG_HR string = "hr" /* Croatian */
const LANG_HU string = "hu" /* Hungarian */
const LANG_HY string = "hy" /* Armenian */
const LANG_IA string = "ia" /* Interlingua */
const LANG_IE string = "ie" /* Interlingue */
const LANG_IK string = "ik" /* Inupiak */
const LANG_IN string = "in" /* Indonesian */
const LANG_IS string = "is" /* Icelandic */
const LANG_IT string = "it" /* Italian */
const LANG_IW string = "iw" /* Hebrew */
const LANG_JA string = "ja" /* Japanese */
const LANG_JI string = "ji" /* Yiddish */
const LANG_JW string = "jw" /* Javanese */
const LANG_KA string = "ka" /* Georgian */
const LANG_KK string = "kk" /* Kazakh */
const LANG_KL string = "kl" /* Greenlandic */
const LANG_KM string = "km" /* Cambodian */
const LANG_KN string = "kn" /* Kannada */
const LANG_KO string = "ko" /* Korean */
const LANG_KS string = "ks" /* Kashmiri */
const LANG_KU string = "ku" /* Kurdish */
const LANG_KY string = "ky" /* Kirghiz */
const LANG_LA string = "la" /* Latin */
const LANG_LN string = "ln" /* Lingala */
const LANG_LO string = "lo" /* Laothian */
const LANG_LT string = "lt" /* Lithuanian */
const LANG_LV string = "lv" /* Latvian,Lettish */
const LANG_MG string = "mg" /* Malagasy */
const LANG_MI string = "mi" /* Maori */
const LANG_MK string = "mk" /* Macedonian */
const LANG_ML string = "ml" /* Malayalam */
const LANG_MN string = "mn" /* Mongolian */
const LANG_MO string = "mo" /* Moldavian */
const LANG_MR string = "mr" /* Marathi */
const LANG_MS string = "ms" /* Malay */
const LANG_MT string = "mt" /* Maltese */
const LANG_MY string = "my" /* Burmese */
const LANG_NA string = "na" /* Nauru */
const LANG_NE string = "ne" /* Nepali */
const LANG_NL string = "nl" /* Dutch */
const LANG_NO string = "no" /* Norwegian */
const LANG_OC string = "oc" /* Occitan */
const LANG_OM string = "om" /* (Afan)Oromo */
const LANG_OR string = "or" /* Oriya */
const LANG_PA string = "pa" /* Punjabi */
const LANG_PL string = "pl" /* Polish */
const LANG_PS string = "ps" /* Pashto,Pushto */
const LANG_PT string = "pt" /* Portuguese  */
const LANG_QU string = "qu" /* Quechua */
const LANG_RM string = "rm" /* Rhaeto-Romance */
const LANG_RN string = "rn" /* Kirundi */
const LANG_RO string = "ro" /* Romanian */
const LANG_RU string = "ru" /* Russian */
const LANG_RW string = "rw" /* Kinyarwanda */
const LANG_SA string = "sa" /* Sanskrit */
const LANG_SD string = "sd" /* Sindhi */
const LANG_SG string = "sg" /* Sangro */
const LANG_SH string = "sh" /* Serbo-Croatian */
const LANG_SI string = "si" /* Singhalese */
const LANG_SK string = "sk" /* Slovak */
const LANG_SL string = "sl" /* Slovenian */
const LANG_SM string = "sm" /* Samoan */
const LANG_SN string = "sn" /* Shona */
const LANG_SO string = "so" /* Somali */
const LANG_SQ string = "sq" /* Albanian */
const LANG_SR string = "sr" /* Serbian */
const LANG_SS string = "ss" /* Siswati */
const LANG_ST string = "st" /* Sesotho */
const LANG_SU string = "su" /* Sundanese */
const LANG_SV string = "sv" /* Swedish */
const LANG_SW string = "sw" /* Swahili */
const LANG_TA string = "ta" /* Tamil */
const LANG_TE string = "te" /* Tegulu */
const LANG_TG string = "tg" /* Tajik */
const LANG_TH string = "th" /* Thai */
const LANG_TI string = "ti" /* Tigrinya */
const LANG_TK string = "tk" /* Turkmen */
const LANG_TL string = "tl" /* Tagalog */
const LANG_TN string = "tn" /* Setswanato Tonga */
const LANG_TR string = "tr" /* Turkish */
const LANG_TS string = "ts" /* Tsonga */
const LANG_TT string = "tt" /* Tatar */
const LANG_TW string = "tw" /* Twi */
const LANG_UK string = "uk" /* Ukrainian */
const LANG_UR string = "ur" /* Urdu */
const LANG_UZ string = "uz" /* Uzbek */
const LANG_VI string = "vi" /* Vietnamese */
const LANG_VO string = "vo" /* Volapuk */
const LANG_WO string = "wo" /* Wolof */
const LANG_XH string = "xh" /* Xhosa */
const LANG_YO string = "yo" /* Yoruba */
const LANG_ZH string = "zh" /* Chinese */
const LANG_ZU string = "zu" /* Zulu */

/*----------------------------------------------------------------------------*/
/*----- Graphics mode ---------------------------------------------------------*/

const GMODE_PAGE_DESCRIPTION int = 0x0001
const GMODE_PATH_OBJECT int = 0x0002
const GMODE_TEXT_OBJECT int = 0x0004
const GMODE_CLIPPING_PATH int = 0x0008
const GMODE_SHADING int = 0x0010
const GMODE_INLINE_IMAGE int = 0x0020
const GMODE_EXTERNAL_OBJECT int = 0x0040

/*----------------------------------------------------------------------------*/
const ENCODING_FONT_SPECIFIC string = "FontSpecific"
const ENCODING_STANDARD string = "StandardEncoding"
const ENCODING_MAC_ROMAN string = "MacRomanEncoding"
const ENCODING_WIN_ANSI string = "WinAnsiEncoding"
const ENCODING_ISO8859_2 string = "ISO8859-2"
const ENCODING_ISO8859_3 string = "ISO8859-3"
const ENCODING_ISO8859_4 string = "ISO8859-4"
const ENCODING_ISO8859_5 string = "ISO8859-5"
const ENCODING_ISO8859_6 string = "ISO8859-6"
const ENCODING_ISO8859_7 string = "ISO8859-7"
const ENCODING_ISO8859_8 string = "ISO8859-8"
const ENCODING_ISO8859_9 string = "ISO8859-9"
const ENCODING_ISO8859_10 string = "ISO8859-10"
const ENCODING_ISO8859_11 string = "ISO8859-11"
const ENCODING_ISO8859_13 string = "ISO8859-13"
const ENCODING_ISO8859_14 string = "ISO8859-14"
const ENCODING_ISO8859_15 string = "ISO8859-15"
const ENCODING_ISO8859_16 string = "ISO8859-16"
const ENCODING_CP1250 string = "CP1250"
const ENCODING_CP1251 string = "CP1251"
const ENCODING_CP1252 string = "CP1252"
const ENCODING_CP1253 string = "CP1253"
const ENCODING_CP1254 string = "CP1254"
const ENCODING_CP1255 string = "CP1255"
const ENCODING_CP1256 string = "CP1256"
const ENCODING_CP1257 string = "CP1257"
const ENCODING_CP1258 string = "CP1258"
const ENCODING_KOI8_R string = "KOI8-R"

/* error-code */
const ERR_ARRAY_COUNT_ERR int = 0x1001
const ERR_ARRAY_ITEM_NOT_FOUND int = 0x1002
const ERR_ARRAY_ITEM_UNEXPECTED_TYPE int = 0x1003
const ERR_BINARY_LENGTH_ERR int = 0x1004
const ERR_CANNOT_GET_PALLET int = 0x1005
const ERR_DICT_COUNT_ERR int = 0x1007
const ERR_DICT_ITEM_NOT_FOUND int = 0x1008
const ERR_DICT_ITEM_UNEXPECTED_TYPE int = 0x1009
const ERR_DICT_STREAM_LENGTH_NOT_FOUND int = 0x100A
const ERR_DOC_ENCRYPTDICT_NOT_FOUND int = 0x100B
const ERR_DOC_INVALID_OBJECT int = 0x100C

/*                                                0x100D */
const ERR_DUPLICATE_REGISTRATION int = 0x100E
const ERR_EXCEED_JWW_CODE_NUM_LIMIT int = 0x100F

/*                                                0x1010 */
const ERR_ENCRYPT_INVALID_PASSWORD int = 0x1011

/*                                                0x1012 */
const ERR_ERR_UNKNOWN_CLASS int = 0x1013
const ERR_EXCEED_GSTATE_LIMIT int = 0x1014
const ERR_FAILED_TO_ALLOC_MEM int = 0x1015
const ERR_FILE_IO_ERROR int = 0x1016
const ERR_FILE_OPEN_ERROR int = 0x1017

/*                                                0x1018 */
const ERR_FONT_EXISTS int = 0x1019
const ERR_FONT_INVALID_WIDTHS_TABLE int = 0x101A
const ERR_INVALID_AFM_HEADER int = 0x101B
const ERR_INVALID_ANNOTATION int = 0x101C

/*                                                0x101D */
const ERR_INVALID_BIT_PER_COMPONENT int = 0x101E
const ERR_INVALID_CHAR_MATRIX_DATA int = 0x101F
const ERR_INVALID_COLOR_SPACE int = 0x1020
const ERR_INVALID_COMPRESSION_MODE int = 0x1021
const ERR_INVALID_DATE_TIME int = 0x1022
const ERR_INVALID_DESTINATION int = 0x1023

/*                                                0x1024 */
const ERR_INVALID_DOCUMENT int = 0x1025
const ERR_INVALID_DOCUMENT_STATE int = 0x1026
const ERR_INVALID_ENCODER int = 0x1027
const ERR_INVALID_ENCODER_TYPE int = 0x1028

/*                                                0x1029 */
/*                                                0x102A */
const ERR_INVALID_ENCODING_NAME int = 0x102B
const ERR_INVALID_ENCRYPT_KEY_LEN int = 0x102C
const ERR_INVALID_FONTDEF_DATA int = 0x102D
const ERR_INVALID_FONTDEF_TYPE int = 0x102E
const ERR_INVALID_FONT_NAME int = 0x102F
const ERR_INVALID_IMAGE int = 0x1030
const ERR_INVALID_JPEG_DATA int = 0x1031
const ERR_INVALID_N_DATA int = 0x1032
const ERR_INVALID_OBJECT int = 0x1033
const ERR_INVALID_OBJ_ID int = 0x1034
const ERR_INVALID_OPERATION int = 0x1035
const ERR_INVALID_OUTLINE int = 0x1036
const ERR_INVALID_PAGE int = 0x1037
const ERR_INVALID_PAGES int = 0x1038
const ERR_INVALID_PARAMETER int = 0x1039

/*                                                0x103A */
const ERR_INVALID_PNG_IMAGE int = 0x103B
const ERR_INVALID_STREAM int = 0x103C
const ERR_MISSING_FILE_NAME_ENTRY int = 0x103D

/*                                                0x103E */
const ERR_INVALID_TTC_FILE int = 0x103F
const ERR_INVALID_TTC_INDEX int = 0x1040
const ERR_INVALID_WX_DATA int = 0x1041
const ERR_ITEM_NOT_FOUND int = 0x1042
const ERR_LIBPNG_ERROR int = 0x1043
const ERR_NAME_INVALID_VALUE int = 0x1044
const ERR_NAME_OUT_OF_RANGE int = 0x1045

/*                                                0x1046 */
/*                                                0x1047 */
const ERR_PAGE_INVALID_PARAM_COUNT int = 0x1048
const ERR_PAGES_MISSING_KIDS_ENTRY int = 0x1049
const ERR_PAGE_CANNOT_FIND_OBJECT int = 0x104A
const ERR_PAGE_CANNOT_GET_ROOT_PAGES int = 0x104B
const ERR_PAGE_CANNOT_RESTORE_GSTATE int = 0x104C
const ERR_PAGE_CANNOT_SET_PARENT int = 0x104D
const ERR_PAGE_FONT_NOT_FOUND int = 0x104E
const ERR_PAGE_INVALID_FONT int = 0x104F
const ERR_PAGE_INVALID_FONT_SIZE int = 0x1050
const ERR_PAGE_INVALID_GMODE int = 0x1051
const ERR_PAGE_INVALID_INDEX int = 0x1052
const ERR_PAGE_INVALID_ROTATE_VALUE int = 0x1053
const ERR_PAGE_INVALID_SIZE int = 0x1054
const ERR_PAGE_INVALID_X_OBJECT int = 0x1055
const ERR_PAGE_OUT_OF_RANGE int = 0x1056
const ERR_REAL_OUT_OF_RANGE int = 0x1057
const ERR_STREAM_EOF int = 0x1058
const ERR_STREAM_READ_CONTINUE int = 0x1059

/*                                                0x105A */
const ERR_STRING_OUT_OF_RANGE int = 0x105B
const ERR_THIS_FUNC_WAS_SKIPPED int = 0x105C
const ERR_TTF_CANNOT_EMBEDDING_FONT int = 0x105D
const ERR_TTF_INVALID_CHAR_MAP int = 0x105E
const ERR_TTF_INVALID_FORMAT int = 0x105F
const ERR_TTF_MISSING_TABLE int = 0x1060
const ERR_UNSUPPORTED_FONT_TYPE int = 0x1061
const ERR_UNSUPPORTED_FUNC int = 0x1062
const ERR_UNSUPPORTED_JPEG_FORMAT int = 0x1063
const ERR_UNSUPPORTED_TYPE1_FONT int = 0x1064
const ERR_XREF_COUNT_ERR int = 0x1065
const ERR_ZLIB_ERROR int = 0x1066
const ERR_INVALID_PAGE_INDEX int = 0x1067
const ERR_INVALID_URI int = 0x1068
const ERR_PAGE_LAYOUT_OUT_OF_RANGE int = 0x1069
const ERR_PAGE_MODE_OUT_OF_RANGE int = 0x1070
const ERR_PAGE_NUM_STYLE_OUT_OF_RANGE int = 0x1071
const ERR_ANNOT_INVALID_ICON int = 0x1072
const ERR_ANNOT_INVALID_BORDER_STYLE int = 0x1073
const ERR_PAGE_INVALID_DIRECTION int = 0x1074
const ERR_INVALID_FONT int = 0x1075
const ERR_PAGE_INSUFFICIENT_SPACE int = 0x1076
const ERR_PAGE_INVALID_DISPLAY_TIME int = 0x1077
const ERR_PAGE_INVALID_TRANSITION_TIME int = 0x1078
const ERR_INVALID_PAGE_SLIDESHOW_TYPE int = 0x1079
const ERR_EXT_GSTATE_OUT_OF_RANGE int = 0x1080
const ERR_INVALID_EXT_GSTATE int = 0x1081
const ERR_EXT_GSTATE_READ_ONLY int = 0x1082
const ERR_INVALID_U3D_DATA int = 0x1083
const ERR_NAME_CANNOT_GET_NAMES int = 0x1084
const ERR_INVALID_ICC_COMPONENT_NUM int = 0x1085

const FONT_FIXED_WIDTH int = 1
const FONT_SERIF int = 2
const FONT_SYMBOLIC int = 4
const FONT_SCRIPT int = 8

/* Reserved                    16 */
const FONT_STD_CHARSET int = 32
const FONT_ITALIC int = 64

/* Reserved                    128
   Reserved                    256
   Reserved                    512
   Reserved                    1024
   Reserved                    2048
   Reserved                    4096
   Reserved                    8192
   Reserved                    16384
   Reserved                    32768 */
const FONT_ALL_CAP int = 65536
const FONT_SMALL_CAP int = 131072
const FONT_FORCE_BOLD int = 262144

const CID_W_TYPE_FROM_TO int = 0
const CID_W_TYPE_FROM_ARRAY int = 1

const OTYPE_NONE int = 0x00000000
const OTYPE_DIRECT int = 0x80000000
const OTYPE_INDIRECT int = 0x40000000
const OTYPE_HIDDEN int = 0x10000000

const OCLASS_UNKNOWN int = 0x0001
const OCLASS_NULL int = 0x0002
const OCLASS_BOOLEAN int = 0x0003
const OCLASS_NUMBER int = 0x0004
const OCLASS_REAL int = 0x0005
const OCLASS_NAME int = 0x0006
const OCLASS_STRING int = 0x0007
const OCLASS_BINARY int = 0x0008
const OCLASS_ARRAY int = 0x0010
const OCLASS_DICT int = 0x0011
const OCLASS_PROXY int = 0x0012
const OCLASS_ANY int = 0x00FF

const OSUBCLASS_FONT int = 0x0100
const OSUBCLASS_CATALOG int = 0x0200
const OSUBCLASS_PAGES int = 0x0300
const OSUBCLASS_PAGE int = 0x0400
const OSUBCLASS_X_OBJECT int = 0x0500
const OSUBCLASS_OUTLINE int = 0x0600
const OSUBCLASS_DESTINATION int = 0x0700
const OSUBCLASS_ANNOTATION int = 0x0800
const OSUBCLASS_ENCRYPT int = 0x0900
const OSUBCLASS_EXT_GSTATE int = 0x0A00
const OSUBCLASS_EXT_GSTATE_R int = 0x0B00 /* read only object */
const OSUBCLASS_NAMEDICT int = 0x0C00
const OSUBCLASS_NAMETREE int = 0x0D00

const BS_SOLID int = 0
const BS_DASHED int = 1
const BS_BEVELED int = 2
const BS_INSET int = 3
const BS_UNDERLINED int = 4
