# PROTEIN LANGUAGE DESIGN DOC

## Lexing

### 1 Lexical analysis

Source code -1-> Tokens -2-> AST (Abstract Syntax Tree)

Step 1 is lexical analysis.

For a protein language `var x int = 100 + 5;`, the lexer looks like:

```
[
    VAR,
    IDENTIFIER("x"),
    INT,
    EQUAL_SIGN,
    INTEGER(100),
    PLUSH_SIGN,
    INTEGER(100),
    SEMICOLON
]
```

### 2 Defining Tokens

```
var five int = 5;
var ten int = 10;

var add = fn(a int, b int) int {
    return a + b;
};

var result int = add(five, ten);
```

### 3 The Lexer

We are going to writ our own lexer. It'll tak source code as input and output the tokens that represent the source
code.

TODO 1: Attach filenames and line numbers to tokens.
TODO 2: Fully support Unicode and emojis.
TODO 3: Generalize `readNumber` and `readLetter`
TODO 4: Support int8, int16, int32, int64, float and double
