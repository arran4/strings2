# Smart Title Uppercase and Acronym Algorithms

This document outlines the design and evaluation of algorithms for handling uppercase words during Smart Title formatting. It compares methods for deciding whether a word in uppercase should be normalized to standard title case or preserved in uppercase.

## The Core Ambiguity: Screaming-Case vs. Acronyms

Screaming-case detection and acronym recognition are fundamentally different problems:

- **Screaming-case detection** aims to identify if the *entire input source* was written in an uppercase format (like `SCREAMING_SNAKE_CASE` or `SCREAMING-KEBAB-CASE`). If it was, the uppercase words within it are just ordinary words that happen to be uppercase due to the format constraint. They should usually be normalized to standard title case (`A New Hope`).
- **Acronym recognition** aims to identify if a *specific word* is an intentional acronym (like `HTTP`, `API`, `JSON`). These should be preserved in uppercase regardless of the surrounding format, unless explicitly overridden.

### Why Casing Alone Is Insufficient

Because both screaming-case and acronyms use uppercase letters, casing alone cannot always determine the semantic intent, leading to unavoidable ambiguities.

Consider `NASA_API_CLIENT`:
- Without external domain knowledge, we cannot reliably determine whether this is screaming snake case for "Nasa Api Client", or a format containing three acronyms "NASA API CLIENT". Both are technically valid interpretations.
- Ambiguities like `COVID_19_RESPONSE` or `A_B_TESTING` similarly suffer from the overlap between ordinary words, acronyms, and formatting styles.

## Parser Provenance and Architectural Changes

In the previous architecture, `ParserConfig.SmartAcronyms` simply categorized any all-uppercase word with more than one character as an `AcronymWord`.
This conflated the intent of an acronym with the physical casing of the word in the source, destroying the provenance of "why" a word is uppercase.

To resolve this without breaking backward compatibility, the parsing step continues to emit `AcronymWord` or `UpperCaseWord`, but the formatting step now looks at the *entire slice of words*. If the entire slice consists only of uppercase types (`AcronymWord`, `UpperCaseWord`, and separators), it's identified as "screaming case". In a confident screaming context, the distinction between `AcronymWord` and `UpperCaseWord` is overridden by structural and domain-specific checks.

Furthermore, `ToTitle` has access to the raw input string (and thus its overall casing shape), while `FromWordsToTitle` operates on pre-parsed `[]Word` slices. This means `FromWordsToTitle` must infer the original shape solely from the sequence of `Word` types. By scanning the entire slice, we can reconstruct the global screaming state even from parsed tokens.

## Algorithm Families Evaluated

We explored multiple independent algorithms for resolving uppercase tokens during formatting:

### Algorithm A: Ratio Heuristic (Baseline)
Uses the proportion of uppercase/acronym-looking words to total words. If the ratio exceeds a threshold, it assumes the text is screaming case and normalizes all caps.

### Algorithm B: Whole-Source Screaming-Case Detection
Analyzes the global casing of the source string or the full sequence of parsed words. If the entire sequence lacks lowercase letters, it classifies the source as screaming case and normalizes standard uppercase words.

### Algorithm C: Word-Type/Provenance Algorithm
Relies entirely on the parsed word types: intentional `AcronymWord` vs. incidental `UpperCaseWord`. Since the current parser heuristic conflates the two based on length, this requires structural modifications or manual overrides to work accurately.

### Algorithm D: Structural Pattern Algorithm
Uses local transitions. An isolated uppercase word (e.g., `parse_HTTP_request`) is highly likely to be an acronym. Contiguous runs of uppercase words are less clear, but surrounding context (mixed-case neighbours) can infer intent without global thresholds.

### Algorithm E: Lexical-Shape Algorithm
Examines the characters inside a token. Tokens with digits (`COVID19`), single letters (`A`), or missing vowels (`XML`) might be scored differently. This is lightweight but linguistically unreliable across domains.

### Algorithm F: Caller-Provided Acronym Recognition
Allows the caller to supply a domain-specific dictionary or predicate (`func(string) bool`) to definitively identify true acronyms (e.g., `NASA`, `API`).

### Algorithm G: Explicit Source-Mode Algorithm
The caller explicitly dictates the behavior: force normalize all uppercase words, force preserve them, or use automatic detection.

### Algorithm H: Multi-Signal Scoring Algorithm
Assigns numerical weights to global casing, isolation, lexical shape, and dictionary matches to compute a final action score for each token.

### Algorithm I: Staged Hybrid (Selected Default)
A deterministic decision tree that combines explicit overrides, domain predicates, global source shape, and structural isolation into a comprehensible, prioritized flow.

## Comparison Table

| Algorithm | Evidence used | Time | Space | Strength | Main failure |
|-----------|---------------|------|-------|----------|--------------|
| Ratio | Aggregate word types | O(n) | O(1) | Very small | Semantic ambiguity |
| Source shape | Original casing and delimiter | O(n) | O(1) | Detects screaming formats | Needs original source |
| Provenance | Word classifications | O(n) | O(1) | Preserves known acronyms | Depends on parser accuracy |
| Structural | Neighbouring token patterns | O(n) | O(1) | Handles isolated acronyms | Complex mixed runs |
| Lexical shape | Token characters | O(n) | O(1) | No dictionary | Linguistically unreliable |
| Acronym predicate | Caller knowledge | O(n) | caller-dependent | High domain accuracy | Requires configuration |
| Explicit mode | Caller declaration | O(n) | O(1) | Deterministic | No automatic inference |
| Scoring | Multiple signals | O(n) | O(1) | Flexible | Harder to reason about |
| Staged hybrid | Ordered decisions | O(n) | O(1) | Explainable and robust | More branches |

## Selected Default: Staged Hybrid Algorithm

The **Staged Hybrid** algorithm was selected as the default because it combines the strengths of deterministic rules, structural context, and domain knowledge while avoiding floating-point heuristics or opaque scoring systems.

### Configuration Options

The algorithm behavior can be controlled using new options:

- **Force Normalization**: `OptionSmartTitleUpperMode(SmartTitleUpperNormalize)` forces all ambiguous uppercase words to be normalized to Title Case.
- **Force Preservation**: `OptionSmartTitleUpperMode(SmartTitleUpperPreserve)` forces all uppercase words and acronyms to remain uppercase.
- **Domain Acronyms**: `OptionSmartTitleAcronymPredicate(func(word string) bool)` allows callers to define a custom function returning `true` for known acronyms. If matched, the word is preserved regardless of surrounding context.

### Reproducing PR #49 Fixed-Ratio Behavior

To maintain backward compatibility or explicitly use the ratio-based heuristic, callers can use the deprecated `OptionSmartTitleThreshold`:
```go
strings2.OptionSmartTitleThreshold(func(wc int) float64 { return 0.5 })
```
If this option is provided, the formatting logic will bypass the Staged Hybrid detection and fall back to the older ratio-based normalization logic.
