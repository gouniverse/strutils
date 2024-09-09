# String Utils <a href="https://gitpod.io/#https://github.com/gouniverse/strutils" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/strutils/workflows/tests/badge.svg)

- <b>Base64Decode(text []byte) (data []byte, err error)</b> - decodes base64 text to binary data.
- <b>Base64Encode(data []byte) (text []byte)</b> -  encodes binary data to base64 encoded text.
- <b>Base32ExtendedDecode(data []byte) string</b> -  decodes binary data to base32 extended (RFC 4648) encoded text.
- <b>Base32ExtendedEncode(data []byte) string</b> -  encodes binary data to base32 extended (RFC 4648) encoded text.
- <b>BcryptHashCompare(str string) bool</b> - compares the string to a bcrypt hash
- <b>>Between(str string, startNeedle string, endNeedle string) (result string, found bool)</b> - returns the substring between two needles
- <b>ContainsAnyChar(str string, charset string) bool</b> - returns true if the string contains any of the characters in the specified charset
- <b>ContainsOnly(str string, charset string) bool </b> - returns true is the string contains only charcters from the specified charset
- <b>IntToBase32(num int) string</b>
- <b>IntToBase36(num int) string</b>
- <b>LeftFrom(s string, needle, string) string</b> - returns the substring on the left side of the needle
- <b>LeftPad(s string, padStr string, overallLen int) string</b>
- <b>StrRightFrom(s string, needle, string) string</b> - returns the substring on the right side of the needle
- <b>RightPad(s string, padStr string, overallLen int) string</b>
- <b>Slugify(str string) string</b>
- <b>ToBcryptHash(str string) string</b> - converts the string to BCrypt hash, use BcryptHashCompare to check
- <b>ToBytes(s string) []byte</b> - converts string to bytes
- <b>ToCamel(str string) string</b> - converts the given string to camel case
- <b>ToSnake(str string) string</b> - converts the given string to snake case
- <b>UcFirst(str string) string</b> - convert first letter into upper case
