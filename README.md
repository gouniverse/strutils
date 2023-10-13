# String Utils <a href="https://gitpod.io/#https://github.com/gouniverse/strutils" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

- <b>Base64Decode(text []byte) (data []byte, err error)</b> - decodes base64 text to binary data.
- <b>Base64Encode(data []byte) (text []byte)</b> -  encodes binary data to base64 encoded text.
- <b>Base32ExtendedDecode(data []byte) string</b> -  decodes binary data to base32 extended (RFC 4648) encoded text.
- <b>Base32ExtendedEncode(data []byte) string</b> -  encodes binary data to base32 extended (RFC 4648) encoded text.
- <b>ContainsAnyChar(str string, charset string) bool</b> - returns true if the string contains any of the characters in the specified charset
- <b>ContainsOnly(str string, charset string) bool </b> - returns true is the string contains only charcters from the specified charset
- <b>LeftPad(s string, padStr string, overallLen int) string</b>
- <b>RightPad(s string, padStr string, overallLen int) string</b>
- <b>Slugify(str string) string</b>
- <b>ToSnake(str string) string</b>
- <b>UcFirst(str string) string</b>
