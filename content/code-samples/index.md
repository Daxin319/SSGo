# Code Samples

SSGo supports both inline code and code blocks, with some custom behaviors for escaping and formatting.

## Inline Code

Basic inline code: `let result = compute(42)`

Nested inline code: ``print(`Hello World`)``

### Custom Escape Behavior

**Escapes are not processed inside inline code spans or code blocks in SSGo.**  
This means that backslashes and backticks are treated as literal characters, making code examples more natural to write and read.

For example, this markdown:

```
Inline code with escapes (escapes shouldn't work in inline code spans/blocks): `Use \\\`backticks\\\` and a literal \\\\ here`
```

**SSGo renders this as:**

- Inline code with escapes (escapes shouldn't work in inline code spans/blocks): `Use \`backticks\` and a literal \\ here`

**CommonMark would render this as:**

- Inline code with escapes (escapes shouldn't work in inline code spans/blocks): `Use \ `backticks\` and a literal \ here`

Notice that in CommonMark, the backslashes are processed and disappear, while in SSGo, they are preserved as literal characters.

## Code Blocks

SSGo supports both triple backtick and triple tilde code blocks:

~~~
function greet(name) {
  return "Hello, " + name;
}
~~~

```
# Bash style
echo "Static sites are cool!"
```

### Custom Behaviors

1. Code blocks preserve exact whitespace and newlines
2. No need to escape backticks inside code blocks
3. Both ``` and ~~~ delimiters are supported
4. Language specification is not supported and will be ignored

[Back to Index](../index.md) 