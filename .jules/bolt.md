## 2024-05-22 - [Optimizing String Concatenation]
**Learning:** In Go, when building a byte slice from strings (e.g., for serialization), using `bytes.Buffer` directly avoids the double allocation of `strings.Builder` -> `string` -> `[]byte`.
**Action:** Use `bytes.Buffer` when the final output is needed as `[]byte`, even if inputs are strings.
