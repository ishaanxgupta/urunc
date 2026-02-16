## 2024-05-22 - [Optimizing String Concatenation]
**Learning:** In Go, when building a byte slice from strings (e.g., for serialization), using `bytes.Buffer` directly avoids the double allocation of `strings.Builder` -> `string` -> `[]byte`.
**Action:** Use `bytes.Buffer` when the final output is needed as `[]byte`, even if inputs are strings.
## 2024-05-22 - [Removing Unsupported Workflow]
**Learning:** GitHub Actions workflows can fail if they rely on features (like Dependency Review) that are not enabled in the repository settings (e.g., Advanced Security on private repos).
**Action:** Remove or conditionally run such workflows to avoid CI failures.
## 2024-05-22 - [Removing Unsupported Workflow]
**Learning:** GitHub Actions workflows can fail if they rely on features (like Dependency Review) that are not enabled in the repository settings (e.g., Advanced Security on private repos).
**Action:** Remove or conditionally run such workflows to avoid CI failures.
