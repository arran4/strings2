
- When adding new features or types, always ensure you provide an example in the form of a testable 'ExampleXxx' function (like those in 'example_*_test.go') following standard Go conventions.
- If a feature modifies parsing or formatting behavior, reflect the usage in the README.md to ensure users are aware of it.

## Parsing and option contracts

- Preserve public case-conversion, partitioning and formatting semantics when adding features. If an API already accepts typed options through `...any`, keep accepted types discoverable, avoid ambiguous raw primitive options, and check/document repeated-option precedence before changing it.
- Prefer a small interface or typed function only when it matches the operation's actual capabilities. Do not rewrite existing public signatures simply to make them mockable; the type-switched variadic pattern is optional, not mandatory for internal helpers.
- For parsing/formatting changes, add ExampleXxx plus behaviour tests for Unicode, empty input, punctuation, acronyms, separators and loss/canonicalisation behaviour **where those cases are supported**. Do not promise byte-for-byte round trips for deliberate normalisers.

## Engineering guidance and agent audit

Use https://arran4.github.io/blog/post/2026/055-suggested-coding-standards-for-repositories-and-agents/ selectively, without superseding the more specific rules above. A coding agent must inspect implementation/tests, distinguish verified gaps from hypotheses, and record results and existing issue overlap at https://github.com/arran4/strings2/issues/72. Do not equate adding instructions with completing the implementation audit; do not merge automatically.
