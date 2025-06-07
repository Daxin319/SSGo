# Invalid HTML Entities

Invalid HTML entities should be treated as literal text. Here are some examples of invalid entities:

&invalid; &#invalid; &#xinvalid;

## Examples of Invalid Entities

- `&invalid;` - Invalid named entity
- `&#invalid;` - Invalid decimal reference
- `&#xinvalid;` - Invalid hexadecimal reference

These invalid entities will be displayed as-is in the browser, rather than being converted to special characters.

[Back to Index](../../index.md) 