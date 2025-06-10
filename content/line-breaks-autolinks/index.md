# Line Breaks and Autolinks

## Hard Line Breaks

This line ends with two spaces  
So it should break here.

This line ends with a backslash\
So it should break here.

This line has multiple spaces at the end     
So it should break here.

This line has a mix of spaces and backslash \  
So it should break here.

### Automatic URL and Email links

Here is a link to a website: https://www.google.com

Here is another one: example.com

You can contact us by email at: test@example.com

To disable automatic URL/Email linking, simply enclose the link in a code span `https://www.google.com/`

## Enhanced Autolinks

### Basic URLs

<http://example.com>
<https://example.com>
<www.example.com>
<example.com>

### URLs with Paths and Parameters

<example.com/path>
<example.com/path?param=value>
<example.com/path#fragment>
<example.com/path?param=value#fragment>

### URLs with Subdomains

<sub.example.com>
<sub.sub.example.com>

### URLs with Ports

<example.com:8080>
<example.com:8080/path>

### URLs with Authentication

<user:pass@example.com>
<user@example.com>

**Note:** For autolink emails with authentication (e.g., `user:pass@example.com`), only the username is shown as the link text (`user@example.com`), but the full value (including the password) is used in the link's `href` attribute. This is intentional for privacy and security, so that passwords are not displayed in the rendered HTML.

### URLs with Special Characters

<example.com/path/with/special/chars>
<example.com/path/with/underscores_and-hyphens>

### Email Addresses

<user@example.com>
<user.name@example.com>
<user+tag@example.com>
<user@sub.example.com>

### Invalid URLs

These should not be autolinks:
<not.a.url>
<http://>
<https://>
<www.>
<.com>
<user@>
<@example.com>

[Back to Index](../index.md) 