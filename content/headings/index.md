# Basic Headings

SSGo supports standard Markdown headings (H1-H6) and has custom handling for extra # characters.

## Standard Headings

# H1 Heading (one #)

## H2 Heading (two #)

### H3 Heading (three #)

#### H4 Heading (four #)

##### H5 Heading (five #)

###### H6 Heading (six #)

## Custom Behavior

SSGo treats any heading with more than 6 # characters as an H6 heading, assuming it's a user error. This is different from standard Markdown which would treat these as paragraphs.

####### Seven # characters -> H6

######## Eight # characters -> H6

######### Nine # characters -> H6

This behavior helps prevent accidental paragraph formatting when users add too many # characters.

[Back to Index](../index.md) 