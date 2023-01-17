## Effective Go

### Names

- The visibility of a name outside a package is determined by whether its first character is upper case.

### Package Names

- When a package is imported, the package name becomes an accessor for the contents. After
  import "bytes"
- Packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. In the rare
  case of a collision the importing package can choose a different name to use locally.
- Another convention is that the package name is the base name of its source directory; the package in
  src/encoding/base64 is imported as "encoding/base64" but has name base64, not encoding_base64 and not encodingBase64.
- Another short example is once.Do; once.Do(setup) reads well and would not be improved by writing
  once.DoOrWaitUntilDone(setup). Long names don't automatically make things more readable. A helpful doc comment can
  often be more valuable than an extra long name.

### Getters

- Go doesn't provide automatic support for getters and setters. There's nothing wrong with providing getters and setters
  yourself, and it's often appropriate to do so, but it's neither idiomatic nor necessary to put Get into the getter's
  name. If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case,
  exported), not GetOwner. The use of upper-case names for export provides the hook to discriminate the field from the
  method. A setter function, if needed, will likely be called SetOwner. Both names read well in practice:
  ```
    owner := obj.Owner()
    if owner != user {
      obj.SetOwner(user)
    }
  ```

### Interface names

By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to
construct
an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

### MixedCaps

The convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names.

### Semicolons
Like C, Go's formal grammar uses semicolons to terminate statements, but unlike in C, those semicolons do not appear
in the source. Instead the lexer uses a simple rule to insert semicolons automatically as it scans, so the input text is
mostly free of them.

### Control structures

The control structures of Go are related to those of C but differ in important ways. There is no do or while loop, only
a slightly generalized for; switch is more flexible; if and switch accept an optional initialization statement like that
of for; break and continue statements take an optional label to identify what to break or continue; and there are new
control structures including a type switch and a multiway communications multiplexer, select. The syntax is also
slightly different: there are no parentheses and the bodies must always be brace-delimited.