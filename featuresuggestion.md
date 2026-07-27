# Feature Suggestion: Unified Delimiter Detector CLI Argument

## Background

Currently, the `strings2` CLI relies on multiple individual boolean flags to control parsing behavior around boundaries and delimiters (e.g., `--number-splitting`, `--non-alphanumeric`). As more rules are added, this approach bloats the CLI arguments and prevents complex, composable logic from being expressed easily by the user.

## Proposal

Instead of individual flags, introduce a unified `--delimiter-detector` argument that accepts a flexible syntax for defining boundaries dynamically.

## Syntax Examples

**Before:**
```bash
strings2 camel "hello_world 123 foo" --number-splitting --non-alphanumeric
```

**After:**
```bash
strings2 camel "hello_world 123 foo" --delimiter-detector "any(numeric, whitespace, nonalphanumeric)"
```

Or for more complex requirements:
```bash
strings2 camel "..." --delimiter-detector "union(whitespace, not(tab))"
```

## Impact

This allows users to pass sophisticated detection rules directly to the `PartitionerConfig.DelimiterDetector` function dynamically at runtime, reducing the need for hardcoded boolean options in both the `strings2` library API and the CLI application itself.
